package main

import (
	"fmt"
	"time"

	"github.com/RavenZZ/viper"
	_ "github.com/RavenZZ/viper/remote"
	"github.com/lunny/log"
)

func main() {
	viper.SetConfigType("yml")
	viper.AddRemoteProvider("etcd", "http://10.9.13.6:2379", "/config/zhuyingjun/test.yml")
	if err := viper.ReadRemoteConfig(); err == nil {
		fmt.Println("Using config file:", viper.AllKeys())
		go func() {
			for {
				time.Sleep(time.Millisecond * 10) // delay after each request
				viper.OnRemoteConfigChange(func() {
					fmt.Println("Using config file:", viper.AllKeys())
				})
				// currently, only tested with etcd support
				err := viper.WatchRemoteConfigOnChannel()
				if err != nil {
					log.Errorf("unable to read remote config: %v", err)
					continue
				}

				// unmarshal new config into our runtime config struct. you can also use channel
				// to implement a signal to notify the system of the changes
				// runtime_viper.Unmarshal(&runtime_conf)
			}
		}()
	} else {
		panic(err)
	}
	c := make(chan int)
	<-c
}
