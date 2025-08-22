package core

import "gorm.io/gorm"

// IDatabaseProvider 数据库提供者接口
type IDatabaseProvider interface {
	GetMainDB() *gorm.DB
	GetOrderDB() *gorm.DB
	GetDatabase(name string) (*gorm.DB, bool)
}

// 实现接口
type DatabaseProvider struct{}

func (p *DatabaseProvider) GetMainDB() *gorm.DB {
	return GetDB()
}

func (p *DatabaseProvider) GetOrderDB() *gorm.DB {
	return GetOrderDB()
}

func (p *DatabaseProvider) GetDatabase(name string) (*gorm.DB, bool) {
	return GetDatabase(name)
}
