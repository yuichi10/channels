package lircd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

func analyzeLircd(f *os.File) (lircd *Lircd) {
	lircd = &Lircd{}
	lircd.action = make([]string, 0, 10)
	scanner := bufio.NewScanner(f)
	flag := ""

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, beginRemote) {
			flag = beginRemote
			continue
		} else if strings.Contains(line, beginCodes) && flag == beginRemote {
			flag = beginCodes
			continue
		} else if strings.Contains(line, endCodes) && flag == beginCodes {
			flag = endCodes
			continue
		} else if strings.Contains(line, beginRawCode) && flag == beginRemote {
			flag = beginRawCode
			continue
		} else if strings.Contains(line, endRawCode) && flag == beginRawCode {
			flag = endRawCode
			continue
		} else if strings.Contains(line, endRemote) {
			break
		}

		splitSpace := func(c rune) bool {
			return c == ' '
		}

		// get lircd name
		if strings.Contains(line, codeName) && flag == beginRemote {
			names := strings.FieldsFunc(line, splitSpace)
			if len(names) < 1 {
				continue
			}
			lircd.name = strings.Trim(names[1], " ")
		}

		// gat names when the file is begin code
		if flag == beginCodes {
			names := strings.FieldsFunc(line, splitSpace)
			if len(names) < 2 {
				continue
			}
			lircd.action = append(lircd.action, names[0])
		}

		if flag == beginRawCode {
			if strings.Contains(line, codeName) {
				names := strings.FieldsFunc(line, splitSpace)
				if len(names) < 2 {
					continue
				}
				lircd.action = append(lircd.action, strings.Trim(names[1], " "))
			}
		}
	}
	return
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
			fmt.Println(l.name)
			fmt.Println(l.action)
			lircds = append(lircds, l)
		}
	}
	fmt.Println(lircds)
	return nil
}
