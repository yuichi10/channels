package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuichi10/channels/config"
	"github.com/yuichi10/channels/lircd"
)

func statusRoute(r *mux.Router) {
	status := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	}
	r.HandleFunc("/status", status)
}

func getGlobalIP() ([]byte, error) {
	req, err := http.NewRequest("GET", "http://ifconfig.io/ip", bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		fmt.Println("error occoer")
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("Can not get Status Code 2xx")
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, err
}

func globalIPRoute(r *mux.Router) {
	globalIP := func(w http.ResponseWriter, r *http.Request) {
		ip, err := getGlobalIP()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(ip)
		w.WriteHeader(http.StatusOK)
	}
	r.HandleFunc("/global-ip", globalIP)
}

func main() {
	config.LoadOption()
	lircds, err := lircd.LoadLircdConf(config.LircdDir)
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	statusRoute(r)
	globalIPRoute(r)
	lircd.StartServer(lircds, r)
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r)

}
