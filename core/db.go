package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"mwhtpay/config"
	"os"
	"sync"
	"time"
)

// 数据库标识符常量
const (
	DBMain  = "main"
	DBOrder = "order"
)

var (
	db        = initMysql()
	orderDb   = initOrderMysql()
	dbMutex   sync.RWMutex
	databases = make(map[string]*gorm.DB)
)

// 初始化时注册所有数据库
func init() {
	RegisterDatabase(DBMain, db)
	RegisterDatabase(DBOrder, orderDb)
}

// RegisterDatabase 注册数据库到管理器
func RegisterDatabase(name string, dbInstance *gorm.DB) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	databases[name] = dbInstance
	log.Printf("Database registered: %s", name)
}

// GetDB 获取主数据库（保持向后兼容）
func GetDB() *gorm.DB {
	return db
}

// GetOrderDB 获取订单数据库（保持向后兼容）
func GetOrderDB() *gorm.DB {
	return orderDb
}

// GetDatabase 通过标识符获取数据库
func GetDatabase(name string) (*gorm.DB, bool) {
	dbMutex.RLock()
	defer dbMutex.RUnlock()
	dbInstance, exists := databases[name]
	return dbInstance, exists
}

// GetAllDatabases 获取所有数据库
func GetAllDatabases() map[string]*gorm.DB {
	dbMutex.RLock()
	defer dbMutex.RUnlock()
	return databases
}

// initMysql 初始化mysql会话
func initMysql() *gorm.DB {
	// 日志配置
	slowThreshold := time.Second
	ignoreRecordNotFoundError := true
	logLevel := logger.Warn
	if config.Config.GinMode == "debug" {
		logLevel = logger.Info
		slowThreshold = 200 * time.Millisecond
		ignoreRecordNotFoundError = false
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             slowThreshold,             // 慢 SQL 阈值
			LogLevel:                  logLevel,                  // 日志级别
			IgnoreRecordNotFoundError: ignoreRecordNotFoundError, // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,                      // 彩色打印
		},
	)
	// 初始化会话
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Config.DatabaseUrl,         // DSN data source name
		DefaultStringSize:         config.Config.DbDefaultStringSize, // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                             // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Config.DbTablePrefix, // 表名前缀
			SingularTable: true,                        // 使用单一表名, eg. `User` => `user`
		},
		DisableForeignKeyConstraintWhenMigrating: true,      // 禁用自动创建外键约束
		Logger:                                   newLogger, // 自定义Logger
	})
	if err != nil {
		log.Fatal("initMysql gorm.Open err:", err)
	}
	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("initMysql db.DB err:", err)
	}
	// 数据库空闲连接池最大值
	sqlDB.SetMaxIdleConns(config.Config.DbMaxIdleConns)
	// 数据库连接池最大值
	sqlDB.SetMaxOpenConns(config.Config.DbMaxOpenConns)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.Config.DbConnMaxLifetimeHours) * time.Hour)
	return db
}

// initOrderMysql 初始化订单库mysql会话
func initOrderMysql() *gorm.DB {
	// 日志配置
	slowThreshold := time.Second
	ignoreRecordNotFoundError := true
	logLevel := logger.Warn
	if config.Config.GinMode == "debug" {
		logLevel = logger.Info
		slowThreshold = 200 * time.Millisecond
		ignoreRecordNotFoundError = false
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             slowThreshold,             // 慢 SQL 阈值
			LogLevel:                  logLevel,                  // 日志级别
			IgnoreRecordNotFoundError: ignoreRecordNotFoundError, // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,                      // 彩色打印
		},
	)
	// 初始化会话
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Config.OrderDatabaseUrl,    // DSN data source name
		DefaultStringSize:         config.Config.DbDefaultStringSize, // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                             // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Config.DbTablePrefix, // 表名前缀
			SingularTable: true,                        // 使用单一表名, eg. `User` => `user`
		},
		DisableForeignKeyConstraintWhenMigrating: true,      // 禁用自动创建外键约束
		Logger:                                   newLogger, // 自定义Logger
	})
	if err != nil {
		log.Fatal("initOrderMysql gorm.Open err:", err)
	}
	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("initOrderMysql db.DB err:", err)
	}
	// 数据库空闲连接池最大值
	sqlDB.SetMaxIdleConns(config.Config.DbMaxIdleConns)
	// 数据库连接池最大值
	sqlDB.SetMaxOpenConns(config.Config.DbMaxOpenConns)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.Config.DbConnMaxLifetimeHours) * time.Hour)
	return db
}

func DBTableName(model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(model); err != nil {
		log.Printf("parse model failed: %v", err)
		return ""
	}
	return stmt.Schema.Table
}

// DBTableName 获取表名（使用指定数据库）
func DBTableNameByDb(dbInstance *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: dbInstance}
	if err := stmt.Parse(model); err != nil {
		log.Printf("parse model failed: %v", err)
		return ""
	}
	return stmt.Schema.Table
}

// DBTableNameWithDB 通过数据库标识符获取表名
func DBTableNameWithDB(dbName string, model interface{}) string {
	if dbInstance, exists := GetDatabase(dbName); exists {
		return DBTableNameByDb(dbInstance, model)
	}
	log.Printf("database not found: %s", dbName)
	return ""
}
