package response

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, res Response) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(res.GetStatus())

	enc := json.NewEncoder(w)
	enc.Encode(res)
}
