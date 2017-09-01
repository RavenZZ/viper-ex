package main

import (
	"fmt"

	"ravenzz/viper-ex/pkg1"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	// for k, v := range viper.AllSettings() {
	// 	fmt.Printf("key:%v   val: %v \n", k, v)
	// }

	pkg1.GetConfig()
}
