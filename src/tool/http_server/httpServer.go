package http_server

import "net/http"
//import "fmt"
//跨域策略
func Cross(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
}