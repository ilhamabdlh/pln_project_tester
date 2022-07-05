package main

import (
	"github.com/gin-gonic/gin"
	"pln/jatim/pkg/controller"
	"pln/jatim/pkg/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)


	controller.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
