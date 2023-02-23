package pkg

import (
	"flag"
	"log"
	"strings"

	"github.com/spf13/viper"
	"github.com/sunbursthou/Go_blog/configs"
)

func InitViper() {
	var configFile string
	// 优先从命令行读取配置文件
	flag.StringVar(&configFile, "c", "", "choose config file")
	flag.Parse()
	if configFile != "" { // 命令行读取到自定义配置文件
		log.Printf("读取配置文件路径为:%s\n", configFile)
	} else {
		log.Printf("默认加载配置文件:configs/configs.toml")
		configFile = "configs/configs.toml"
	}

	v_config := viper.New()
	v_config.SetConfigFile(configFile)
	v_config.AutomaticEnv()
	v_config.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // _替换为.

	// 读取配置文件
	if err := v_config.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败: ", err)
	}

	// 反序列化到go结构体上
	if err := v_config.Unmarshal(&configs.Cfg); err != nil {
		log.Panic("配置文件加载失败", err)
	}

	log.Println("配置文件加载成功！")

}
