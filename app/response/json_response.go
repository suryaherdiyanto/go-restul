package response

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, res Response) {
	enc := json.NewEncoder(w)
	w.WriteHeader(res.GetStatus())
	enc.Encode(res)
}
