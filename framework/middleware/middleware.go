package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	framework "campfire-software.com/simplifi-framework/framework/models"
	"github.com/gin-gonic/gin"
)

// LogProperties logs out all of the application properties
func LogProperties(props *framework.Properties) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("Properties")
		if !exists {
			c.Set("Properties", props)
			value, _ = c.Get("Properties")
		}
		log.Printf("Properties: %+v", value)
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
