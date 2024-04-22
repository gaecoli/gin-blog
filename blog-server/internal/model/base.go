package model

import (
	g "gin-blog/internal/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func makeMigrateDb(db *gorm.DB) error {

	return db.AutoMigrate(
		&Article{},  // 文章
		&Category{}, // 分类
		&Tag{},      // 标签
	)
}

type Model struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// InitDB 连接数据库，目前暂时只支持 MySQL
func InitDB(conf *g.Config) *gorm.DB {
	dsn := conf.DbDSN()

	var level logger.LogLevel

	switch conf.Server.DbLogMode {
	case "silent":
		level = logger.Silent
	case "info":
		level = logger.Info
	case "warn":
		level = logger.Warn
	case "error":
		level = logger.Error
	default:
		level = logger.Error
	}

	config := &gorm.Config{
		Logger:                                   logger.Default.LogMode(level),
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
		},
	}

	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}

	log.Println("数据库连接成功: ", dsn)

	// TODO: 数据库迁移
	// if conf.Server.DbAutoMigrate
	//if conf.Server.DbAutoMigrate {
	//	if err := makeMigrateDb(db); err != nil {
	//		log.Fatal("数据库迁移失败", err)
	//	}
	//	log.Println("数据库自动迁移成功")
	//}

	return db
}

func Count[T any](db *gorm.DB, data *T, where ...any) (int, error) {
	var total int64
	db = db.Model(data)
	if len(where) > 0 {
		db = db.Where(where[0], where[1:]...)
	}
	result := db.Count(&total)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(total), nil
}
