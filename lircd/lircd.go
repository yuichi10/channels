package lircd

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

const irsend = "irsend"
const sendOnce = "SEND_ONCE"

// Lircd is type which have name and action
type Lircd struct {
	name   string
	action []string
}

// Lircds is list of lircd
type Lircds []*Lircd

func (l Lircd) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := func(c rune) bool {
		return c == '/'
	}
	paths := strings.FieldsFunc(r.URL.Path, f)
	if len(paths) != 4 {
		fmt.Println("came from invalid url:", r.URL.Path, "len:", len(paths))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(irsend, sendOnce, paths[2], paths[3])
	err := exec.Command(irsend, sendOnce, paths[2], paths[3]).Run()
	if err != nil {
		fmt.Println("failed to exec command:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
