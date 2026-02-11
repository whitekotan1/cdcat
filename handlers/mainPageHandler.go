package api

import "net/http"

func HandlePage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "this is GET methd!", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")

}
