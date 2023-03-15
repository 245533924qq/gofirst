package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"log"
	"os"
)

// 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml" // 配置常量,对应配置文件名
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile) // 读取文件内容
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c) // 解析 yaml内容写入到指针里面，返回报错
	if err != nil {
		log.Fatalf("config Init Unmarshal:%v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}
