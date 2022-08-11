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
	router.Use(cors.New(cors.Config{
           AllowOrigins:     []string{"https://foo.com"},
           AllowMethods:     []string{"PUT", "PATCH"},
           AllowHeaders:     []string{"Origin"},
           ExposeHeaders:    []string{"Content-Length"},
           AllowCredentials: true,
           AllowOriginFunc: func(origin string) bool {
            return origin == "https://github.com"
          },
          MaxAge: 12 * time.Hour,
        }))

	h := db.Init(dbUrl)


	controller.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
