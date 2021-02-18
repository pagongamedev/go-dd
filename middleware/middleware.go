package middleware

import (
	godd "github.com/pagongamedev/go-dd"
)

// Middleware struct
type Middleware struct {
	LifeCycle        *godd.APILifeCycle
	HandlerStartList []godd.HandlerCycle
	HandlerEndList   []godd.HandlerCycle
}
