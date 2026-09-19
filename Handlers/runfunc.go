package handlers

import "net/http"

func GetFuncRunHandler(f func(), handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f()
		handler(w, r)
	}
}
