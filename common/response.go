package common

import (
	"encoding/json"
	"log"
)

type DefaultResponse struct {
	Message string `json:"message"`
}

func (dr DefaultResponse) ToBytes() []byte {
	response, err := json.Marshal(dr)
	if err != nil {
		log.Fatalln("Failed to serialize response")
	}

	return response
}

type ResponseSuccessWithData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (rs ResponseSuccessWithData) ToBytes() []byte {
	response, err := json.Marshal(rs)
	if err != nil {
		log.Fatalln("Failed to serialize response")
	}

	return response
}
