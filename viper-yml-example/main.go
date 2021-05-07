package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	App      App
	Database Database
}

type App struct {
	Foo  string
	Foo2 string
}

type Database struct {
	Password string
}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("PREFIX")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	var C Configuration
	// Uncomment this to check whether viper has read the config from yaml correctly
	// fmt.Println(viper.GetString("app"))
	if err := viper.Unmarshal(&C); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	C.App.Foo2 = viper.GetString("PORT")

	fmt.Println(C.Database)
	fmt.Println(C.App.Foo2)
}
