package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/richterk/simplifi-framework/framework"
)

// LogProperties logs out all of the application properties
func LogProperties(props *framework.Properties) gin.HandlerFunc {
	fmt.Println("Props...")
	return func(c *gin.Context) {
		fmt.Println("Props handler...")
		_, exists := c.Get("Properties")
		if !exists {
			c.Set("Properties", props)
		}
		c.Next()
	}
}

// LoadPropertiesByEnv loads the properties file based on the environment
func LoadPropertiesByEnv(env string) *framework.Properties {
	propertiesReader := fmt.Sprintf("./config/config.%s.json", env)
	bytes, err := ioutil.ReadFile(propertiesReader)

	if err != nil {
		panic(err)
	}
	var props framework.Properties

	err = json.Unmarshal(bytes, &props)
	if err != nil {
		panic(err)
	}

	return &props

}

func CorsConfig() gin.HandlerFunc {

	fmt.Println("Creating config...")
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin, Access-Control-Allow-Header, Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
