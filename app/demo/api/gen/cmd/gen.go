package main

import (
	"flag"
	"log"
	"path/filepath"
	"runtime"

	"github.com/zeromicro/go-zero/core/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	"zero-fusion/app/demo/api/gen/annotations"
	"zero-fusion/app/demo/api/internal/config"
	"zero-fusion/share/utils"
)

var configFile = flag.String("f", "etc/dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	parentDir := utils.NLevelUp(currentDir, 2)

	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       filepath.Join(parentDir, "/internal/dao/query"), // output directory
		ModelPkgPath:  filepath.Join(parentDir, "/internal/model"),     // model package path
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// Initialize a *gorm.DB instance
	db, err := gorm.Open(mysql.Open(c.Gorm.Conn.MasterDefaultDSN))
	if err != nil {
		panic(err)
	}

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyInterface(func(annotations.Querier) {}, g.GenerateAllTable()...)

	//g.ApplyBasic(
	//	g.GenerateModel("user_subscribe"),
	//)

	g.ApplyBasic(g.GenerateAllTable()...)
	g.WithOpts(
		gen.FieldTrimSuffix("_"),
	)

	// Execute the generator
	g.Execute()

	log.Printf("generate done.")
}
