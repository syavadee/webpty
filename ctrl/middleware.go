package ctrl

import (
	"net/http"
)

func Middleware(fn func(res http.ResponseWriter, req *http.Request)) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		fn(res, req)
	}
}
