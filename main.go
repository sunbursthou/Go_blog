package main

import (
	"github.com/sunbursthou/Go_blog/internal/routers"
)

func main() {
	// 初始化全局，包括visper，MySQL，Redis等
	routers.InitGlobaleVariable()
	routers.InitRouter()

}
