package handlers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Author", "Bhavesh")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Welcome"}`))
}
