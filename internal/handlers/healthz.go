package handlers

import "net/http"

//healthz handler
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte(`{"status": "ok"}`))
}
