package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	SEEKING string
	FOO     string
}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("app.env")
	viper.SetConfigType("env")
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error when reading config: %s", err)
	}

	var C Configuration
	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalf("Unable to decode to struct: %s", err)
	}

	fmt.Println("Hello world")
	// If you run with APP_SEEKING=Calling, this will print Calling
	fmt.Println(C.SEEKING)
	// Print Bar
	fmt.Println(C.FOO)
}
