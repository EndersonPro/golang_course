package helpers


import (
	"encoding/json"
	"net/http"
)

type responseapi struct {
	Status 	int 	`json:"status"`
	Message string	`json:"message"`
	Route 	string	`json:"route"`
	Method 	string	`json:"method"`
}

// HandleResponse structura para todas las respuestas de la api
func HandleResponse(w http.ResponseWriter, r *http.Request, status int, m string) {
	route := r.RequestURI
	method := r.Method
	response := responseapi{ status, m, route, method }

	js, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}