package gormx

import (
	"fmt"

	sqlmysql "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"

	gormlogx "zero-fusion/share/orm/gormx/logging/logx"
)

type GormDB struct {
	*gorm.DB
}

// NewGormDB 新建数据库链接
func NewGormDB(config Config, registerCallbacks ...func(resolver *dbresolver.DBResolver)) *GormDB {
	var (
		db  *gorm.DB
		err error
	)

	dialect := config.Conn.Dialect
	switch dialect {
	case "mysql":
		gormConfig := &gorm.Config{}

		logLevel := logger.Silent
		if config.Logger.LogLevel == "info" {
			logLevel = logger.Info
		} else if config.Logger.LogLevel == "warn" {
			logLevel = logger.Warn
		} else if config.Logger.LogLevel == "error" {
			logLevel = logger.Error
		}

		loggerType := config.Logger.LoggerType
		switch loggerType {
		case "logx":
			gormConfig = &gorm.Config{
				Logger: gormlogx.NewGormLogger(
					gormlogx.Config{
						SlowThreshold:             config.Logger.SlowThreshold,
						LogLevel:                  logLevel,
						IgnoreRecordNotFoundError: config.Logger.IgnoreRecordNotFoundError,
					}),
			}
		}

		var defaultMasterDB *gorm.DB
		// 主节点默认节点
		masterDefaultDSNS := config.Conn.MasterDefaultDSN
		defaultGormDialect, parsedDSN := newGormDialect(dialect, masterDefaultDSNS)
		defaultMasterDB, err = gorm.Open(defaultGormDialect, gormConfig)
		if err != nil {
			err = fmt.Errorf("%s-%s new gorm db error: %+v", dialect, parsedDSN.Addr, err)
			logx.Error(err)
			panic(err)
		}

		// 主节点
		masterDSNS := config.Conn.MasterDSNS
		var gormDialectMasterList []gorm.Dialector
		for _, masterDSN := range masterDSNS {
			masterGormDialect, _ := newGormDialect(dialect, masterDSN)
			gormDialectMasterList = append(gormDialectMasterList, masterGormDialect)
		}

		// 从节点
		var gormDialectSlaveList []gorm.Dialector
		for _, slaveDSN := range config.Conn.SlaveDSNS {
			slaveGormDialect, _ := newGormDialect(dialect, slaveDSN)
			gormDialectSlaveList = append(gormDialectSlaveList, slaveGormDialect)
		}

		resolver := dbresolver.Register(dbresolver.Config{
			Sources:           gormDialectMasterList,
			Replicas:          gormDialectSlaveList,
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: true,
		})

		// 自定义主从切换规则
		for _, registerCallback := range registerCallbacks {
			registerCallback(resolver)
		}

		resolver.
			SetMaxIdleConns(config.Conn.MaxIdleConns).
			SetMaxOpenConns(config.Conn.MaxOpenConns).
			SetConnMaxLifetime(config.Conn.ConnMaxLifetime)

		err = defaultMasterDB.Use(resolver)
		if err != nil {
			errMsg := fmt.Sprintf("use plugin error %s error: %+v", dialect, err)
			logx.Error(errMsg)
			panic(errMsg)
		}

		db = defaultMasterDB
	default:
		errMsg := fmt.Sprintf("gorm unsupported dialects %s", dialect)
		logx.Error(errMsg)
		panic(err)
	}

	logx.Info(fmt.Sprintf("new db %s success", dialect))

	return &GormDB{DB: db}
}

func newGormDialect(dialect string, dsn string) (gorm.Dialector, *sqlmysql.Config) {
	parsedDSN, err := sqlmysql.ParseDSN(dsn)
	if err != nil {
		errMsg := fmt.Sprintf("%s failed to parse dsn err: %v", dialect, err)
		logx.Error(errMsg)
		panic(errMsg)
	}

	return mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), parsedDSN
}

// Debug debug模式
func (g *GormDB) Debug() {
	g.DB = g.DB.Debug()
}
