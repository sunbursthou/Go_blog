package routers

import (
	"github.com/sunbursthou/Go_blog/internal/pkg"
)

func InitGlobaleVariable() {
	// 初始化Viper
	pkg.InitViper()
	// 初始化数据库
	pkg.InitMySQL()

}
