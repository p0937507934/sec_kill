package main

import (
	"log"
	"os"
	"os/signal"
	"sec_kill/api"
	"sec_kill/driver"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	driver.GetService()
	go func() {
		engine.POST("/activity", api.SetActivity)
		engine.POST("/sec_kill", api.SetKill)
		err := engine.Run(":8000")
		if err != nil {
			panic(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit)
	<-quit
	log.Println("shutdown...")
}
