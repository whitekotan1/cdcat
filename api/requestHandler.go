package api

import (
	"cdcat/types"
	"encoding/json"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post", http.StatusMethodNotAllowed)
		return
	}

	var request types.Request

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := types.Response{
		Status:  "cat",
		Message: "cat received your repos" + request.RepoUrl,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
