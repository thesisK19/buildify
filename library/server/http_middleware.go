package server

import (
	"net/http"
	"strings"
	"sync"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// HTTPServerMiddleware is an interface of http server middleware
type HTTPServerMiddleware func(http.Handler) http.Handler

// PassedHeaderDeciderFunc returns true if given header should be passed to gRPC server metadata.
type PassedHeaderDeciderFunc func(string) bool

func createPassingHeaderMiddleware(decide PassedHeaderDeciderFunc) HTTPServerMiddleware {
	return func(next http.Handler) http.Handler {
		cache := new(sync.Map)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//nolint:gomnd
			newHeader := make(http.Header, 2*len(r.Header))

			for k := range r.Header {
				v := r.Header.Get(k)
				if newKey, ok := cache.Load(k); ok {
					newHeader.Set(newKey.(string), v)
				} else if decide(k) {
					newKey := runtime.MetadataHeaderPrefix + k
					cache.Store(k, newKey)
					newHeader.Set(newKey, v)
				}
				newHeader.Set(k, v)
			}

			r.Header = newHeader

			next.ServeHTTP(w, r)
		})
	}
}

// MiddlewareAuthentication used to handle the current user & roles by JWT/IAM service.
func MiddlewareAuthAttachUser(pattern string, roles ...string) HTTPServerMiddleware {
	return func(handler http.Handler) http.Handler {
		next := handler.ServeHTTP
		return HandlerAuthAttachUser(next, pattern, roles...)
	}
}

// HandlerAuthAttachUser used to handle the current user & roles by JWT/IAM service.
func HandlerAuthAttachUser(next http.HandlerFunc, pattern string, roles ...string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.RequestURI, pattern) {
			next.ServeHTTP(w, r)
			return
		}
		// if clientIAM == nil {
		// 	http.Error(w, "Missing IAM client", http.StatusInternalServerError)
		// 	return
		// }
		// user, err := clientIAM.GetUser(r)
		// if err != nil {
		// 	/**
		// 	|-------------------------------------------------------------------------
		// 	| @TODO This code does not work, could not convert embedded struct like
		// 	| below
		// 	|-----------------------------------------------------------------------*/
		// 	if _, ok := err.(*net.DNSError); ok {
		// 		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 		return
		// 	}
		// 	if _, ok := err.(*h.ConnectionError); ok {
		// 		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 		return
		// 	}
		// 	http.Error(w, err.Error(), http.StatusUnauthorized)
		// 	return
		// }
		// if !user.HasAllRoles(roles...) {
		// 	http.Error(w, ForbiddenMesage, http.StatusForbidden)
		// 	return
		// }
		// ctxNew := context.WithValue(r.Context(), Identity, user)
		// req := r.WithContext(ctxNew)
		// next.ServeHTTP(w, req)
		next.ServeHTTP(w, r)
	})
}

// MiddlewareAuthentication used to handle the current user & roles by JWT/IAM service.
func MiddlewareAuthentication(pattern string) HTTPServerMiddleware {
	return func(handler http.Handler) http.Handler {
		next := handler.ServeHTTP
		return HandlerAuthentication(next, pattern)
	}
}

// HandlerAuthentication used to handle the current user & roles by JWT/IAM service.
func HandlerAuthentication(next http.HandlerFunc, pattern string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.RequestURI, pattern) {
			next(w, r)
			return
		}
		// if clientIAM == nil {
		// 	http.Error(w, "Missing IAM client", http.StatusInternalServerError)
		// 	return
		// }
		// claims, err := clientIAM.GetStandardClaims(r)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusUnauthorized)
		// 	return
		// }
		// ctxNew := context.WithValue(r.Context(), StandardClaims, claims)
		// req := r.WithContext(ctxNew)
		// next(w, req)
		// TODO: authennnn
		next(w, r)
	})
}
