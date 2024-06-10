package configs

import (
	"time"

	"github.com/gin-contrib/cors"
)

var CorsConfiguration = cors.New(cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"*"},
	AllowHeaders:     []string{"*"},
	ExposeHeaders:    []string{"*"},
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
		return origin == "https://github.com"
	},
	MaxAge: 12 * time.Hour,
})
