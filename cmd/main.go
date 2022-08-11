package main

import (
	"github.com/gin-gonic/gin"
	"pln/jatim/pkg/controller"
	"pln/jatim/pkg/db"
	"github.com/spf13/viper"
	"github.com/gin-contrib/cors"
)

func main() {
	viper.SetConfigFile("./pkg/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	r.Use(CORSMiddleware())
	h := db.Init(dbUrl)


	controller.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
