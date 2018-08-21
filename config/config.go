package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	LircdDir string
	Port     string
)

func arguments() error {
	currentPath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current path: %s", err)
	}
	lircdDir := filepath.Join(currentPath, "lircd.conf.d")
	flag.StringVar(&LircdDir, "lircd", lircdDir, "lircd directory")
	flag.StringVar(&Port, "port", "9999", "web api port")
	flag.Parse()
	return nil
}

func loadConfig() {
	os.Getenv("LIRCD_DIR")
	os.Getenv("PORT")
}

func LoadOption() {
	loadConfig()
	arguments()
}
