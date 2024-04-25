package model

import (
	g "gin-blog/internal/global"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func GetDB(c *gin.Context) *gorm.DB {
	// gin MustGet返回一个 any (interface{})
	// 通过类型断言的方式将它转为想要的 *gorm.DB 类型
	// 例子：
	// var a interface{} = 10
	// t, ok := a.(int) 转为整数类型
	// t1, ok1 := a.(float32) 转为浮点数类型
	return c.MustGet(g.CTX_DB).(*gorm.DB)
}

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

	if conf.Server.DbAutoMigrate {
		if err := makeMigrateDb(db); err != nil {
			log.Fatal("数据库迁移失败", err)
		}
		log.Println("数据库自动迁移成功")
	}

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

// 分页器校验
func Paginate(pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum < 1 {
			pageNum = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 1:
			pageSize = 10
		}

		limit := pageSize
		offset := (pageSize - 1) * pageNum

		return db.Limit(limit).Offset(offset)
	}
}
