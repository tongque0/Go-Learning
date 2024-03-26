package main

import (
	"cache/pkg/cache"
	"fmt"
	"time"
)

func main() {
	c := cache.NewMemCache()
	c.SetMaxMemory("100KB")
	c.Set("int", 1, 1*time.Second)
	c.Set("bool", false, 1*time.Second)
	c.Set("string", "你好，灰太狼！凄凄切切凄凄切切！", 1*time.Second)
	c.Set("struct", "你好，灰太狼！", 1*time.Second)
	fmt.Print(c.Get("int"))
	c.Flush()
	fmt.Print(c.Get("int"))
}
