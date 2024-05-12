package chat

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// 定义一个结构体来映射 YAML 文件中的数据结构
type Config struct {
	Xfapi struct {
		Appid     string `yaml:"appid"`
		ApiKey    string `yaml:"apiKey"`
		ApiSecret string `yaml:"apiSecret"`
		HostUrl   string `yaml:"hostUrl"`
	} `yaml:"xfapi"`
}

// 读取并解析 YAML 配置文件
func readConfig() Config {
	var config Config

	// 读取文件内容
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 解析 YAML
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}

// 修改 EnvString 函数，使其从 YAML 文件中获取配置信息
func EnvString(key string) string {
	config := readConfig()

	switch key {
	case "APP_ID":
		return config.Xfapi.Appid
	case "API_KEY":
		return config.Xfapi.ApiKey
	case "API_SECRET":
		return config.Xfapi.ApiSecret
	case "HOST_URL":
		return config.Xfapi.HostUrl
	default:
		fmt.Printf("Key %s not found\n", key)
		return ""
	}
}
