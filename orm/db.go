package orm

import (
	"log"
	"os"
	"time"

	"github.com/Cheng1622/web-short-video/config"
	"golang.org/x/exp/slog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitSqlite() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.SQLITE_DB), &gorm.Config{
		// 执行任何 SQL 时都会创建一个 prepared statement 并将其缓存
		PrepareStmt: false,
		// 禁用默认事务
		SkipDefaultTransaction: true,
		// 禁用嵌套事务
		DisableNestedTransaction: true,
		// 日志
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
				LogLevel:                  logger.Info,            // Log level
				IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,                  // Don't include params in the SQL log
				Colorful:                  true,                   // Disable color
			}),
	})
	if err != nil {
		slog.Error("connect server failed, err:", err)
		os.Exit(1)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := DB.DB()
	if err != nil {
		slog.Info("connect server failed, err:", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
