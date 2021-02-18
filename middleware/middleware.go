package middleware

import (
	godd "github.com/pagongamedev/go-dd"
)

// Middleware struct
type Middleware struct {
	LifeCycle        *godd.APILifeCycle
	handlerStartList []godd.Handler
	handlerEndList   []godd.Handler
}
