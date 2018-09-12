package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/yuichi10/channels/config"
	"github.com/yuichi10/channels/lircd"
)

func main() {
	config.LoadOption()
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to get working directory")
	}
	dir := filepath.Join(wd, "lircd.conf.d")
	lircds, err := lircd.LoadLircdConf(dir)
	if err != nil {
		panic(err)
	}
	lircd.StartServer(lircds)
}
