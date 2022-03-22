package config

import (
	"github.com/Unknwon/goconfig"
	"log"
	"os"
	"path/filepath"
	"sanguo_game/utils"
)

const configFile = "/conf/config.ini"

var File *goconfig.ConfigFile

func init() {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(currentPath, configFile)
	if !utils.FileExists(filePath) {
		panic("File does not exist")
	}
	File, err = goconfig.LoadConfigFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
