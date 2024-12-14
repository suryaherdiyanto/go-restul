package response

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, res Response) {
	enc := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json")
	enc.Encode(res)
}
