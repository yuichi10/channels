package lircd

import (
	"fmt"
	"io/ioutil"
)

type Lircd struct {
	name   string
	action []string
}

func loadFile(filepath string) error {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

}

func LoadLircdConf(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to load lircd file: %s", err)
	}

}
