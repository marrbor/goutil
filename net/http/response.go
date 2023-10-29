package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/marrbor/goutil/closer"
)

// Send response (for server side program)

// ResponseOK returns 200 ok.
func ResponseOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// errResponse returns error response.
func errResponse(w http.ResponseWriter, code int, err error) {
	msg := http.StatusText(code)
	if err != nil {
		m := err.Error()
		if m != "" {
			msg = m
		}
	}
	http.Error(w, msg, code)
}

// BadRequest returns http 400
func BadRequest(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusBadRequest, err)
}

// Unauthorized returns http 401
func Unauthorized(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusUnauthorized, err)
}

// Forbidden returns http 403
func Forbidden(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusForbidden, err)
}

// NotFound returns http 404
func NotFound(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusNotFound, err)
}

// MethodNotAllowed returns http 405
func MethodNotAllowed(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusMethodNotAllowed, err)
}

// InternalServerError returns http 500
func InternalServerError(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusInternalServerError, err)
}

// NotImplementedError returns http 501
func NotImplementedError(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusNotImplemented, err)
}

// BadGatewayError returns http 502
func BadGatewayError(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusBadGateway, err)
}

// ServiceUnavailableError returns http 503
func ServiceUnavailableError(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusServiceUnavailable, err)
}

// GatewayTimeoutError returns http 504
func GatewayTimeoutError(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusGatewayTimeout, err)
}

// JSONResponse returns JSON object
func JSONResponse(w http.ResponseWriter, data interface{}) error {
	if data == nil {
		ResponseOK(w)
		return nil
	}

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(j); err != nil {
		return err
	}
	return nil
}

// //// Receive response (for client side program)

// ResponseJSONToParams convert JSON body in response to given structure.
func ResponseJSONToParams(r *http.Response, params interface{}) error {
	defer closer.Close(r.Body)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, params); err != nil {
		return err
	}
	return nil
}

// ResponseToError convert status and body into an error if given function (default: IsSuccessful) returns false
func ResponseToError(r *http.Response, jf func(*http.Response) bool, ff func(int, string, string) string) error {
	if jf == nil {
		jf = IsSuccessful
	}
	if jf(r) {
		return nil
	}

	defer closer.Close(r.Body)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	bs := string(b)
	if ff == nil {
		ff = func(code int, status, body string) string {
			return fmt.Sprintf("%d %s:\n %s", code, status, body)
		}
	}
	es := ff(r.StatusCode, r.Status, bs)
	return fmt.Errorf(es)
}

// //// Response class checker https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html

// IsInformational returns whether response is informational class (1xx) or not
func IsInformational(r *http.Response) bool {
	return http.StatusContinue <= r.StatusCode && r.StatusCode < http.StatusOK
}

// IsSuccessful returns whether response is Successful class (2xx) or not
func IsSuccessful(r *http.Response) bool {
	return http.StatusOK <= r.StatusCode && r.StatusCode < http.StatusMultipleChoices
}

// IsRedirection returns whether response is Redirection class (3xx) or not
func IsRedirection(r *http.Response) bool {
	return http.StatusMultipleChoices <= r.StatusCode && r.StatusCode < http.StatusBadRequest
}

// IsClientError returns whether response is client error class (4xx) or not
func IsClientError(r *http.Response) bool {
	return http.StatusBadRequest <= r.StatusCode && r.StatusCode < http.StatusInternalServerError
}

// IsNotFound returns whether response is 404 or not
func IsNotFound(r *http.Response) bool {
	return http.StatusNotFound == r.StatusCode
}

// IsServerError returns whether response is server error class (5xx) or not
func IsServerError(r *http.Response) bool {
	return http.StatusInternalServerError <= r.StatusCode && r.StatusCode < 600
}
