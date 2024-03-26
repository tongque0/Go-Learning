package cache_test

import (
	"cache/pkg/cache"
	"testing"
	"time"
)

func TestMemCache(t *testing.T) {
	// 测试初始化和设置最大内存
	mc := cache.NewMemCache()
	if !mc.SetMaxMemory("100KB") {
		t.Fatalf("SetMaxMemory failed")
	}

	// 测试添加缓存项
	key, val := "key1", "value1"
	if !mc.Set(key, val, 10*time.Second) {
		t.Fatalf("Set failed for key %s", key)
	}

	// 测试获取缓存项
	if !mc.Exists(key) {
		t.Fatalf("exists failed for key %s", key)
	}
	if gotVal, found := mc.Get(key); !found || gotVal != val {
		t.Fatalf("Get failed for key %s, expected %v, got %v", key, val, gotVal)
	}

	if !mc.Del(key) {
		t.Fatalf("Del failed for key %s", key)
	}
	// 等待过期
	time.Sleep(11 * time.Second)
	if _, found := mc.Get(key); found {
		t.Fatalf("Get found expired key %s", key)
	}

	// 测试Flush方法
	if !mc.Flush() {
		t.Fatalf("Flush failed")
	}
	if mc.Keys() != 0 {
		t.Fatalf("Flush failed, expected 0 keys, got %d", mc.Keys())
	}
}
