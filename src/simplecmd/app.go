package main

import (
	"simplecmd/cmd"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// import logrus
	log.WithFields(log.Fields{
		"function": "main",
	}).Info("First run main")
	// import gorm
	type DB struct {
		*gorm.DB
	}
	// import viper
	viper.SetConfigType("JSON")

	cmd.Execute()
}
