package cache

import "time"

// Cache 定义了缓存系统的接口，包括设置和获取缓存项，删除缓存项，以及检查缓存项是否存在等操作。
type Cache interface {
	// setMaxMemory 设置缓存可以使用的最大内存大小。
	// size是一个字符串，表示内存大小，例如"64MB"或"1GB"。
	// 返回值表示是否设置成功。
	SetMaxMemory(size string) bool

	// Set 向缓存中添加一个项，如果键已存在，则更新其值。
	// key为要设置的项的键，val为值，expire是过期时间。
	// 返回值表示是否设置成功。
	Set(key string, val interface{}, expire time.Duration) bool

	// Get 从缓存中获取一个项的值。
	// key是要获取的项的键。
	// 返回值是找到的值（如果键存在）和一个布尔值，表示是否找到该键。
	Get(key string) (interface{}, bool)

	// Del 从缓存中删除一个项。
	// key是要删除的项的键。
	// 返回值表示是否删除成功。
	Del(key string) bool

	// Exiists 检查缓存中是否存在指定的键。
	// key是要检查的项的键。
	// 返回值表示该键是否存在。
	Exists(key string) bool

	// Flush 清除缓存中的所有项。
	// 返回值表示操作是否成功。
	Flush() bool

	// Keys 返回缓存中项的数量。
	// 返回值是缓存中项的总数。
	Keys() int64
}
