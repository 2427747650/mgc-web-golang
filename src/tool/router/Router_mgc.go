package router

import "net/http"

func Router(routerName string,serverFunc func(http.ResponseWriter, *http.Request)){
	http.HandleFunc(routerName, serverFunc)
}