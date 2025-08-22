// di.go
package core

import (
	"sync"
)

// DI容器
var (
	diContainer = make(map[string]interface{})
	diMutex     sync.RWMutex
)

// ProvideForDIWithName 使用指定名称注册依赖
func ProvideForDIWithName(name string, instance interface{}) error {
	diMutex.Lock()
	defer diMutex.Unlock()
	diContainer[name] = instance
	return nil
}

// ResolveForDI 从DI容器解析依赖
func ResolveForDI(instance interface{}) error {
	// 这里需要根据具体的DI实现来编写
	// 简化示例：通过类型断言设置值
	return nil
}

// ResolveForDIByName 通过名称从DI容器解析依赖
func ResolveForDIByName(name string) (interface{}, bool) {
	diMutex.RLock()
	defer diMutex.RUnlock()
	instance, exists := diContainer[name]
	return instance, exists
}
