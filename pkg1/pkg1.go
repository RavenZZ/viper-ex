package pkg1

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	config = map[string]interface{}{}
)

func init() {
	checkConfig()
}

func checkConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	loadConfig()
	if config["p1"] == nil {
		panic(fmt.Errorf("config p1 not found:  "))
	}
	if config["p2"] == nil {
		panic(fmt.Errorf("config p2 not found:  "))
	}
	if config["p3"] == nil {
		panic(fmt.Errorf("config p3 not found:  "))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("Config file changed:", e.Op)
		fmt.Println("Config file changed:", e.String())
		loadConfig()
		PrintConfig()
	})
}

func loadConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	config = viper.GetStringMap("pkg1")
}

func PrintConfig() {
	for k, v := range config {
		fmt.Printf("key:%v   val: %v \n", k, v)
	}
}
