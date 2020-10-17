package controller

import "net/http"

// Index view index
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`shuwen-go-web`))
}
