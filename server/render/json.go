package render

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	StatusCode int                    `json:"-"`
	Data       interface{}            `json:"data"`
	Meta       map[string]interface{} `json:"meta"`
}

func (resp Response) Write(w http.ResponseWriter, r *http.Request) {

	bs, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		// TODO handle error
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if resp.StatusCode == 0 {
		resp.StatusCode = http.StatusOK
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(bs)

}

func render(statusCode int, result interface{}, w http.ResponseWriter, r *http.Request) {
	res := Response{
		StatusCode: statusCode,
		Meta:       map[string]interface{}{},
		Data:       result,
	}
	res.Write(w, r)
}

func RenderOk(result interface{}, w http.ResponseWriter, r *http.Request) {
	render(http.StatusOK, result, w, r)
}


