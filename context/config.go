package context

import (
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

// GlobalConfig 配置信息
var GlobalConfig Config

// Config 配置信息结构
type Config struct {
	Bind struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	LogConf struct {
		Logdir   string `yaml:"logdir"`
		Filename string `yaml:"filename"`
		LogLevel string `yaml:"loglevel"`
	}
}

func ParseConfig(cfg string) {
	data, err := ioutil.ReadFile(cfg)
	if err != nil {
		log.Println("read configuration error:", err)
		os.Exit(-1)
	}
	err = yaml.Unmarshal([]byte(data), &GlobalConfig)
	if err != nil {
		log.Println("configuration format error:", err)
		os.Exit(-2)
	}
	log.Println("bind.host", GlobalConfig.Bind.Host)
	log.Println("bind.port", GlobalConfig.Bind.Port)
	log.Println("logconf.logdir", GlobalConfig.LogConf.Logdir)
	log.Println("logconf.LogLevel", GlobalConfig.LogConf.LogLevel)
	// check configuration.
	err = GlobalConfig.Check()
	if err != nil {
		log.Println("configuration format error:", err)
		os.Exit(-3)
	}
}

func (c *Config) Check() error {
	return nil
}
