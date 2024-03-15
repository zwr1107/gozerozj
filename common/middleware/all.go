package middleware

import (
	"fmt"
	"net/http"
)

type AllMiddleware struct {
}

func NewAllMiddleware() *AllMiddleware {
	return &AllMiddleware{}
}

func (m *AllMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		fmt.Println("全局中间件AllMiddleware")
		// Passthrough to next handler if need
		next(w, r)
	}
}
