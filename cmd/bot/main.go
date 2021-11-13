package main

import (
	"github.com/Luismorlan/newsmux/server"
	"github.com/Luismorlan/newsmux/utils"
	"github.com/Luismorlan/newsmux/utils/dotenv"
	. "github.com/Luismorlan/newsmux/utils/flag"
	. "github.com/Luismorlan/newsmux/utils/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

func cleanup() {
	Log.Info("bot server shutdown")
}

func main() {
	defer cleanup()
	ParseFlags()
	InitLogger()

	if err := dotenv.LoadDotEnvs(); err != nil {
		panic(err)
	}

	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()

	router.Use(cors.Default())
	router.Use(gintrace.Middleware(*ServiceName))

	db, err := utils.GetDBConnection()
	utils.BotDBSetupAndMigration(db)
	if err != nil {
		panic("failed to connect to database")
	}

	slashCmdHandler := server.BotCommandHandler(db)
	router.POST("/cmd", slashCmdHandler)

	subscribeHandler := server.BotCommandHandler(db)
	router.POST("/sub", subscribeHandler)

	authHandler := server.AuthHandler(db)
	router.GET("/auth", authHandler)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Newsfeed server - API not found"})
	})

	Log.Info("bot server starts up")
	router.Run(":8080")
}
