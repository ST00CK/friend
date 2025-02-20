package config

import "github.com/gin-contrib/cors"

func SetCors() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // 모든 origin 허용 (개발 환경에서 유용)
	return config
}
