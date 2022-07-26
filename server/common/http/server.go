package http

import (
	"net/http"
	"server/common/errors"
	"server/common/http/codec"
)

type ResponseErr struct {
	Code    int    `json:"code"`
	SubCode int    `json:"subcode"`
	Message string `json:"message"`
}

type Response struct {
	Success bool         `json:"success"`
	Payload interface{}  `json:"payload"`
	Error   *ResponseErr `json:"error"`
}

func DecodeRequest(req *http.Request, v interface{}) error {
	return codec.DecodeRequest(req, v)
}

func EncodeResponse(res http.ResponseWriter, req *http.Request, v interface{}) error {
	return codec.EncodeResponse(res, req, v, func(v interface{}) interface{} {
		return Response{
			Success: true,
			Payload: v,
		}
	})
}

func EncodeError(res http.ResponseWriter, req *http.Request, err error) {
	codec.EncodeError(res, req, err, func(e *errors.OctopusError) interface{} {
		return Response{
			Success: false,
			Error: &ResponseErr{
				Code:    e.HTTPCode(),
				SubCode: e.HTTPSubCode(),
				Message: e.HTTPErrMsg(),
			},
		}
	})
}
