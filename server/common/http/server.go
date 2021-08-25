package http

import (
	"io/ioutil"
	"net/http"
	commctx "server/common/context"
	"server/common/errors"
	"server/common/utils"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
)

const (
	requestIdHeader = "Octopus-Request-Id"
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

type HandleOptions struct {
	log *log.Helper
}

type UserId string
type RequestId string

const baseContentType = "application"

var (
	// acceptHeader      = http.CanonicalHeaderKey("Accept")
	contentTypeHeader = http.CanonicalHeaderKey("Content-Type")
)

func NewHandleOptions() *HandleOptions {
	return &HandleOptions{
		log: log.NewHelper("Server", log.DefaultLogger),
	}
}

func contentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

func contentSubtype(contentType string) string {
	if contentType == baseContentType {
		return ""
	}
	if !strings.HasPrefix(contentType, baseContentType) {
		return ""
	}
	// guaranteed since != baseContentType and has baseContentType prefix
	switch contentType[len(baseContentType)] {
	case '/', ';':
		if i := strings.Index(contentType, ";"); i != -1 {
			return contentType[len(baseContentType)+1 : i]
		}
		return contentType[len(baseContentType)+1:]
	default:
		return ""
	}
}

func (s *HandleOptions) DecodeRequest(req *http.Request, v interface{}) error {
	*req = *req.WithContext(commctx.RequestIdToContext(req.Context(), utils.GetUUIDWithoutSeparator()))
	subtype := contentSubtype(req.Header.Get(contentTypeHeader))
	if codec := encoding.GetCodec(subtype); codec != nil {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return errors.Errorf(err, errors.ErrorHttpReadBody)
		}
		err = codec.Unmarshal(data, v)
		if err != nil {
			return errors.Errorf(err, errors.ErrorJsonUnmarshal)
		}
		return nil
	}
	err := binding.BindForm(req, v)
	if err != nil {
		return errors.Errorf(err, errors.ErrorHttpBindFormFailed)
	}
	return nil
}

func (s *HandleOptions) EncodeResponse(res http.ResponseWriter, req *http.Request, v interface{}) error {
	subtype := contentSubtype(req.Header.Get("accept"))
	codec := encoding.GetCodec(subtype)
	if codec == nil {
		codec = encoding.GetCodec("json")
	}

	v = Response{
		Success: true,
		Payload: v,
	}
	data, err := codec.Marshal(v)
	if err != nil {
		return errors.Errorf(err, errors.ErrorJsonMarshal)
	}
	res.Header().Set("content-type", contentType(codec.Name()))
	res.Header().Set(requestIdHeader, commctx.RequestIdFromContext(req.Context()))
	_, err = res.Write(data)
	if err != nil {
		return errors.Errorf(err, errors.ErrorHttpWriteFailed)
	}
	return nil
}

func (s *HandleOptions) EncodeError(res http.ResponseWriter, req *http.Request, err error) {
	se, ok := errors.FromError(err)
	if !ok {
		se, _ = errors.FromError(errors.Errorf(nil, errors.ErrorUnknown))
	}

	subtype := contentSubtype(req.Header.Get("accept"))
	codec := encoding.GetCodec(subtype)
	if codec == nil {
		codec = encoding.GetCodec("json")
	}
	data, err := codec.Marshal(Response{
		Success: false,
		Error: &ResponseErr{
			Code:    se.HTTPCode(),
			SubCode: se.HTTPSubCode(),
			Message: se.HTTPErrMsg(),
		},
	})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("content-type", contentType(codec.Name()))
	res.Header().Set(requestIdHeader, commctx.RequestIdFromContext(req.Context()))
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(data)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
