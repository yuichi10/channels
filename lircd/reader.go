package lircd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	beginRemote  = "begin remote"
	endRemote    = "end remote"
	beginRawCode = "begin raw_codes"
	endRawCode   = "end raw_codes"
	beginCodes   = "begin codes"
	endCodes     = "end codes"
	codeName     = "name"
)

type Lircd struct {
	name   string
	action []string
}

func analyzeLircd(f *os.File) *Lircd {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return &Lircd{}
}

func loadFile(filepath string, lircd chan *Lircd) {
	f, err := os.Open(filepath)
	// when filed to load file cause panic
	if err != nil {
		panic(err)
	}
	lircd <- analyzeLircd(f)
}

func LoadLircdConf(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to load lircd file: %s", err)
	}
	// var wg sync.WaitGroup

	lircds := make([]*Lircd, 0, 5)
	lircd := make(chan *Lircd, 4)
	for _, file := range files {
		if !file.IsDir() {
			go loadFile(filepath.Join(dirPath, file.Name()), lircd)
		}
	}

	for range files {
		if l := <-lircd; l != nil {
			lircds = append(lircds, l)
		}
	}
	fmt.Println(lircds)
	return nil
}
