package handlers

import "net/http"

func Listlisting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Kushagra!"))
}
