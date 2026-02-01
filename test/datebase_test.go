// 测试数据库连接以及 Article 模型的基本操作
package test

import (
	"Server/conf"
	"Server/core"
	"Server/global"
	"Server/model"
	"log"
	"testing"

	"gorm.io/gorm"
)

func TestDatabaseConnection(t *testing.T) {
	// 不加载配置文件，直接创建 Config 实例
	//
	cfg := &conf.Config{
		Log: conf.LogConfig{
			Level:   "debug",
			Outputs: "stdout",
		},

		Database: conf.DatabaseConfig{
			Host:     "192.168.3.14",
			Port:     25432,
			Name:     "ginblog",
			User:     "ginadmin",
			Password: "ginpassword",
			Schema:   "public",
		},
	}
	global.GlobalConfig = cfg

	// 日志初始化
	logger, cleanup, err := core.InitLogger(global.GlobalConfig.Log)
	if err != nil {
		log.Fatalf("init logger failed: %v", err)
	}
	global.GlobalLogger = logger
	defer cleanup()

	// 初始化数据库
	db, err := core.InitDatabase(cfg.Database)
	if err != nil {
		t.Fatalf("failed to initialize database: %v", err)
	}
	global.GlobalDB = db
	// 自动迁移 Article 模型
	if err := global.GlobalDB.AutoMigrate(&model.Article{}); err != nil {
		t.Fatalf("auto migrate failed: %v", err)
	}

	// 创建一个新的 Article 实例
	article := model.Article{
		Title:   "Test Article",
		Content: "This is a test article.",
		Tags:    []string{"test", "article"},
	}
	if err := global.GlobalDB.Create(&article).Error; err != nil {
		t.Fatalf("failed to create article: %v", err)
	}

	// 读取刚刚创建的 Article 实例
	var readArticle model.Article
	if err := global.GlobalDB.First(&readArticle, article.ID).Error; err != nil {
		t.Fatalf("failed to read article: %v", err)
	}
	if readArticle.Title != article.Title || readArticle.Content != article.Content {
		t.Fatalf("read article does not match created article")
	}

	// 更新 Article 实例
	updatedContent := "This is an updated test article."
	if err := global.GlobalDB.Model(&readArticle).Update("Content", updatedContent).Error; err != nil {
		t.Fatalf("failed to update article: %v", err)
	}

	// 验证更新
	var updatedArticle model.Article
	if err := global.GlobalDB.First(&updatedArticle, article.ID).Error; err != nil {
		t.Fatalf("failed to read updated article: %v", err)
	}

	if updatedArticle.Content != updatedContent {
		t.Fatalf("article content was not updated")
	}

	// 删除 Article 实例
	if err := global.GlobalDB.Delete(&updatedArticle).Error; err != nil {
		t.Fatalf("failed to delete article: %v", err)
	}

	// 验证删除
	var deletedArticle model.Article
	if err := global.GlobalDB.First(&deletedArticle, article.ID).Error; err != gorm.ErrRecordNotFound {
		t.Fatalf("article was not deleted")
	}

	t.Logf("database connection and Article model operations succeeded")
}
