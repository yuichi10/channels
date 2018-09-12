package main

import (
	"github.com/yuichi10/channels/config"
	"github.com/yuichi10/channels/lircd"
)

func main() {
	config.LoadOption()
	lircds, err := lircd.LoadLircdConf(config.LircdDir)
	if err != nil {
		panic(err)
	}
	lircd.StartServer(lircds)
}
