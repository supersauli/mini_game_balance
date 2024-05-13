package configs

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type Server struct {
	Version int     `json:"version"` // 配置文件版版本号
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Sign    Sign    `mapstructure:"sign" json:"sign" yaml:"sign"`
	Balance Balance `mapstructure:"balance" json:"balance" yaml:"balance"`
}

var ServerConfig = &Server{}

func Init() {
	var removeAddr string
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.StringVar(&removeAddr, "r", "", "remove addr 远程配置文件地址")
	flag.Parse()
	if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
		config = "config.yaml"
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType(strings.ReplaceAll(path.Ext(config), ".", ""))
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(ServerConfig); err != nil {
		panic(fmt.Errorf("Fatal Unmarshal config file: %s \n", err))
	}

	info, _ := json.Marshal(*ServerConfig)
	log.Printf("config file = %s", string(info))

}
