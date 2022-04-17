package helper

import (
	"strings"
)

// Berisi modul untuk memberikan respon yang berhasil maupun gagal


// struct Response dengan pengembalian nilai json 
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}


// struct EmptyObj ketik data tersebut tidak dibolehkan null pada json
type EmptyObj struct{}

// method BuildResponse untuk memberikan nilai dari data secara dynamic dengan respon success 
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

// method BuildResponse untuk memberikan nilai dari data secara dynamic dengan respon failed
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splitedError,
		Data:    data,
	}
	return res
}
