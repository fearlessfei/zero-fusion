package common

import (
	"context"
	"flag"
	"fmt"
	"reflect"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"

	"zero-fusion/app/demo/api/scripts/internal/config"
	"zero-fusion/app/demo/api/scripts/internal/svc"
	"zero-fusion/share/log"
	"zero-fusion/share/scriptrunner"
	"zero-fusion/share/utils"
)

var (
	configFile = flag.String("f", "etc/dev.yaml", "the config file")
	c          config.Config
)

func init() {
	flag.Parse()

	conf.MustLoad(*configFile, &c)
	c.MustSetUp()
	log.OutStdout(c.Mode)
}

// run start run script
func run(scriptName string, fn func(c config.Config, spanCtx context.Context, logger logx.Logger) error) {
	spanCtx, endSpan := StartSpan(context.Background(), scriptName, WithStopAgent())
	logger := logx.WithContext(spanCtx).WithFields(logx.Field("script", scriptName))
	logger.Infof("start run script: %s", scriptName)

	defer func() {
		err := logx.Close()
		if err != nil {
			logger.Errorf("close logx error: %v", err)
		}
	}()

	proc.AddShutdownListener(func() {
		ExitClean(logger)
	})

	startTime := time.Now()
	err := fn(c, spanCtx, logger)
	if err != nil {
		logx.WithContext(spanCtx).Errorf("run script error: %v", err)
	}
	logger.Infof("end run script: %s, elapsed: %s", scriptName, time.Since(startTime).String())

	endSpan(err)
}

func RunScript(scriptLogicRunner scriptrunner.ScriptLogicRunner) {
	scriptName := utils.GetPathFileNameNoExt(utils.RuntimeCallerSkipFile(2))
	run(scriptName, func(c config.Config, spanCtx context.Context, logger logx.Logger) error {
		val := reflect.ValueOf(scriptLogicRunner)
		if val.Kind() != reflect.Ptr {
			panic("scriptLogicRunner must be pointer")
		}

		injectFields(val, map[string]interface{}{
			"Logger": logger,
			"Ctx":    spanCtx,
			"SvcCtx": svc.NewServiceContext(c),
		})

		method := val.MethodByName("Run")
		if !method.IsValid() {
			panic("method is invalid")
		}

		ret := method.Call([]reflect.Value{})
		retInterface := ret[0].Interface()
		if retInterface == nil {
			return nil
		}

		return retInterface.(error)
	})
}

func setField(obj reflect.Value, fieldName string, value any) {
	// can only set field values on struct kind
	val := obj.Elem()
	field := val.FieldByName(fieldName)

	if !field.IsValid() || !field.CanSet() {
		panic(fmt.Sprintf("field %s is invalid", fieldName))
	}

	v := reflect.ValueOf(value)
	if v.Type().ConvertibleTo(field.Type()) {
		field.Set(v.Convert(field.Type()))
	}
}

func injectFields(val reflect.Value, fields map[string]interface{}) {
	for name, value := range fields {
		setField(val, name, value)
	}
}
