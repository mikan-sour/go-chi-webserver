package marco

import "net/http"

func Marco(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("polo"))
}
