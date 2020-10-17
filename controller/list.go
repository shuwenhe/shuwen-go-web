package controller

import (
	"io/ioutil"
	"net/http"
)

// List list
func List(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/pages/list.html")
	w.Write(buf)
}
