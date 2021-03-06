package helpers

import (
	"encoding/json"
	"net/http"

	instances "github.com/burubur/microservice/instances/infrastructures"
	structures "github.com/burubur/microservice/structures/https"
	infraStruct "github.com/burubur/microservice/structures/infrastructures"
)

// Response doc
type Response struct{}

// SendOK doc
func (r *Response) SendOK(w http.ResponseWriter, data structures.Response) {
	r.send(w, http.StatusOK, data)
}

// SendAccepted doc
func (r *Response) SendAccepted(w http.ResponseWriter, data structures.Response) {
	r.send(w, http.StatusAccepted, data)
}

// SendBadRequest doc
func (r *Response) SendBadRequest(w http.ResponseWriter, data structures.Response) {
	r.send(w, http.StatusBadRequest, data)
}

// SendError doc
func (r *Response) SendError(w http.ResponseWriter, httpStatus int, data structures.Response) {
	r.send(w, httpStatus, data)
}

func (r *Response) send(w http.ResponseWriter, httpStatus int, data structures.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	responseData := structures.Response{
		Message: data.Message,
		Data:    data.Data,
		Errors:  data.Errors,
	}
	customLog := &instances.Logger{}
	logContent := infraStruct.Log{Entry: responseData}
	customLog.Debug(logContent)
	json.NewEncoder(w).Encode(responseData)
}
