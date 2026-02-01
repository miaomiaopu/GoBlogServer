package core

import (
	"Server/conf"
	"Server/global"
	"Server/model"
	"Server/util"

	"gorm.io/gorm"
)

func InitDatabase(dbConf conf.DatabaseConfig) (*gorm.DB, error) {
	global.GlobalLogger.Info("Initializing database connection...")

	// 连接 PG 数据库
	db, err := util.InitPostgresDB(dbConf)
	if err != nil {
		global.GlobalLogger.Errorw("Failed to connect to database", "error", err)
		return nil, err
	}

	// 迁移数据库模型
	if err := autoMigrate(db, &model.Article{}); err != nil {
		return nil, err
	}

	return db, nil
}

// 自动迁移数据库模型
func autoMigrate(db *gorm.DB, models ...any) error {
	global.GlobalLogger.Info("Starting database auto-migration...")
	err := util.AutoMigrate(db, models...)
	if err != nil {
		global.GlobalLogger.Errorw("Database auto-migration failed", "error", err)
		return err
	}
	global.GlobalLogger.Info("Database auto-migration completed successfully.")
	return nil
}
