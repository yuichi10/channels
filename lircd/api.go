package lircd

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/yuichi10/channels/config"
)

const (
	apiPath = "/api"
	version = "/v1"
)

// StartServer stand up api
func StartServer(ls []*Lircd) {
	r := mux.NewRouter()
	base := path.Join(apiPath, version)
	fmt.Println("There is under paths")
	for _, l := range ls {
		for _, a := range l.action {
			path := path.Join(base, l.name, a)
			fmt.Println("  " + path)
			r.Handle(path, l)
		}
	}
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r)
}
