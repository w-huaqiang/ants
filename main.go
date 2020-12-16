package main

import (
	"log"
	"os"

	"bjzdgt.com/ants/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString(`server.port`)
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run()
}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: ", err)
	}
}
