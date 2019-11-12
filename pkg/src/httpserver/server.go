package httpserver

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"loyalty_go/pkg/src/domains/userProfile"
	"loyalty_go/pkg/src/rabbit"
	"loyalty_go/pkg/src/repositories/userProfileRepo"
)

func StartServer(){
	go rabbit.InitRabbit()
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(static.Serve("/v1/loyalty", static.LocalFile("www", true)))
	v1 := r.Group("/v1/loyalty")
	profileRepo := userProfileRepo.NewRepo()
	profileService := userProfile.NewService(profileRepo)
	profileHandler := NewProfileHandler(profileService)
	{
		v1.POST("/userprofile", profileHandler.newProfile)
		v1.GET("/userprofile", profileHandler.getProfile)
	}
	r.Run(":4100")
	fmt.Println("Documentation on http://localhost:4100/v1/loyalty/")
	return
}