package main

import "net/http"

func returnHTTPError(w http.ResponseWriter, code int, errText string) {
	w.WriteHeader(code)
	w.Write([]byte(errText))
}
