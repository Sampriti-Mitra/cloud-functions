package cloud_functions

import "net/http"

func SimpleHelloWorldFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("Hello World!"))
}
