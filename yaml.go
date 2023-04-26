package main

import (
	"fmt"
	"packagetest"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("prometheus.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath("/home/go/go-learn/learn-code/class-code/20230410")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// global := viper.Get("global")
	//alerting := viper.Get("alerting.alertmanagers.targets")
	value := "127.0.0.1:9098"
	//err2 := viper.Set()
	// rule_files := viper.Get("rule_files")
	//scrape_configs := viper.Get("scrape_configs")
	// fmt.Println(global)
	// fmt.Println(alerting)
	// fmt.Println(rule_files)
	// fmt.Println(scrape_configs)
	fmt.Println(value)
	packagetest.Packagetest()
    packagetest.Packagetest()
	// 导入本地包步骤
	// 1)在自建包目录下执行“go mod init <包名>”
	// 2)在项目根目录下执行“go mod init <项目名>”
	// 3)在项目根目录下执行“go mod edit -replace <import导入路径>=<真实路径>”
	// 4)在项目根目录下执行“go get <包名>”，完成导入

}
