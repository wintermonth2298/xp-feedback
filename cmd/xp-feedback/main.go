package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wintermonth2298/xp-feedback/database"
	"github.com/wintermonth2298/xp-feedback/feedback"
	"github.com/wintermonth2298/xp-feedback/server"
)

var (
	configFile = "config/config.toml"
)

type config struct {
	Port  string `toml:"port"`
	Mongo struct {
		Name               string `toml:"name"`
		Uri                string `toml:"uri"`
		FeedbackCollection string `toml:"feedback_collection"`
	} `toml:"mongo"`
}

func main() {
	c := &config{}
	if _, err := toml.DecodeFile(configFile, &c); err != nil {
		log.Fatal(err)
	}

	db := database.Init(c.Mongo.Uri, c.Mongo.Name)
	feedbackRepo := feedback.Repository{DB: db.Collection(c.Mongo.FeedbackCollection)}
	feedbackService := feedback.Service{Repo: &feedbackRepo}
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		cors.Default(),
	)
	api := router.Group("/api")
	feedback.InitRoutes(&feedbackService, api)

	s := server.NewServer(router, c.Port)
	s.Run()
}
