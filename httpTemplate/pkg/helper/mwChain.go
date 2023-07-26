package helper

import "net/http"

//Middleware template
type Middleware func(http.HandlerFunc) http.HandlerFunc

//HandleFuncMW fn
//s.router.HandleFuncMW("/chain", s.chainMW(s.authHandler(), s.mwChain1, s.mwChain2, s.mwChain3))
//same as
//s.router.HandleFuncMW("/chain", s.mwChain1(s.mwChain2(s.mwChain3(s.authHandler()))))
func HandleFuncMW(h http.HandlerFunc, mw ...Middleware) http.HandlerFunc {
	if len(mw) < 1 {
		return h
	}
	wrapped := h
	// loop in reverse to preserve middleware order
	for i := len(mw) - 1; i >= 0; i-- {
		wrapped = mw[i](wrapped)
	}
	return wrapped
}

// ChainMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func ChainMiddleware(mw ...Middleware) Middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		last := final
		for i := len(mw) - 1; i >= 0; i-- {
			last = mw[i](last)
		}

		return func(w http.ResponseWriter, r *http.Request) {
			last(w, r)
		}
	}
}
