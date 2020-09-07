package gorilla

import (
	"encoding/json"
	"net/http"
)

// Response data format for HTTP
type Response struct {
	Status  string      `json:"status" bson:"status"`                       // Status code (error|fail|success)
	Code    int         `json:"code"  bson:"code"`                          // HTTP status code
	Message string      `json:"message,omitempty" bson:"message,omitempty"` // Error or status message
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"`       // Data payload
}

func sendResponse(hw http.ResponseWriter, status string, code int, message string, data interface{}) {

	response := Response{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}

	hw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(hw).Encode(response)
}
