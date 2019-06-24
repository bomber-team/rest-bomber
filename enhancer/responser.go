package enhancer

import (
	"encoding/json"
	"net/http"
)

type Responser struct {
	CurrentContentTypeResponse string
}

/*ResponseWithError - ответ с ошибкой*/
func (resp *Responser) ResponseWithError(w http.ResponseWriter, r *http.Request, httpStatus int, payload interface{}, contentType string) {
	resp.ResponseWithJSON(w, r, httpStatus, payload, contentType)
}

/*ResponseWithJSON - ответ в формате json*/
func (resp *Responser) ResponseWithJSON(w http.ResponseWriter, r *http.Request, httpStatus int, payload interface{}, contentType string) {

	if r.Header.Get("Content-Type") != contentType {

		w.Header().Set("Content-Type", resp.CurrentContentTypeResponse)
		w.WriteHeader(http.StatusConflict)
		response, _ := json.Marshal(map[string]string{"error": "you packet in non json format!"})

		w.Write(response)

	} else {

		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", resp.CurrentContentTypeResponse)
		w.WriteHeader(httpStatus)

		w.Write(response)

	}
}
