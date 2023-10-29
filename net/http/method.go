// Package http wrap net/http.Method
package http

import "net/http"

type Method struct{ method string }

func (h Method) String() string {
	return h.method
}

var Methods = struct {
	GET     Method
	HEAD    Method
	POST    Method
	PUT     Method
	PATCH   Method
	DELETE  Method
	CONNECT Method
	OPTIONS Method
	TRACE   Method
}{
	GET:     Method{http.MethodGet},
	HEAD:    Method{http.MethodHead},
	POST:    Method{http.MethodPost},
	PUT:     Method{http.MethodPut},
	PATCH:   Method{http.MethodPatch},
	DELETE:  Method{http.MethodDelete},
	CONNECT: Method{http.MethodConnect},
	OPTIONS: Method{http.MethodOptions},
	TRACE:   Method{http.MethodTrace},
}
