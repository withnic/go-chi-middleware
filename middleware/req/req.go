package req

import (
	"context"
	"errors"
	"net/http"
)

type errKey int

const (
	requestURI errKey = iota
	method
	path
	referer
	userAgent
	addr
)

// New returns http.Handler.
func New(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, requestURI, r.RequestURI)
		ctx = context.WithValue(ctx, method, r.Method)
		ctx = context.WithValue(ctx, path, r.URL.Path)
		ctx = context.WithValue(ctx, referer, r.Referer())
		ctx = context.WithValue(ctx, userAgent, r.UserAgent())
		ctx = context.WithValue(ctx, addr, r.RemoteAddr)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// RemoteAddr returns remoteAddr.
func RemoteAddr(ctx context.Context, fallback string) string {
	a, err := GetRemoteAddr(ctx)
	if a == "" || err != nil {
		return fallback
	}

	return a
}

// GetRemoteAddr returns remoteAddr.
func GetRemoteAddr(ctx context.Context) (string, error) {
	addr, ok := ctx.Value(addr).(string)
	if !ok {
		return "", errors.New("failed to get remote address")
	}
	return addr, nil
}

// UserAgent returns user-agent or fallback
func UserAgent(ctx context.Context, fallback string) string {
	u, err := GetUserAgent(ctx)
	if u == "" || err != nil {
		return fallback
	}
	return u
}

// GetUserAgent returns user agent.
func GetUserAgent(ctx context.Context) (string, error) {
	ua, ok := ctx.Value(userAgent).(string)
	if !ok {
		return "", errors.New("failed to get user-agent")
	}
	return ua, nil
}

// Referer returns referer or fallback
func Referer(ctx context.Context, fallback string) string {
	r, err := GetReferer(ctx)
	if r == "" || err != nil {
		return fallback
	}
	return r
}

// GetReferer returns referer.
func GetReferer(ctx context.Context) (string, error) {
	r, ok := ctx.Value(referer).(string)
	if !ok {
		return "", errors.New("failed to get referer")
	}
	return r, nil
}

// Path returns path or fallback
func Path(ctx context.Context, fallback string) string {
	p, err := GetPath(ctx)
	if p == "" || err != nil {
		return fallback
	}
	return p
}

// GetPath returns path.
func GetPath(ctx context.Context) (string, error) {
	p, ok := ctx.Value(path).(string)
	if !ok {
		return "", errors.New("failed to get path")
	}
	return p, nil
}

// Method returns method or fallback.
func Method(ctx context.Context, fallback string) string {
	m, err := GetMethod(ctx)
	if m == "" || err != nil {
		return fallback
	}
	return m
}

// GetMethod returns method.
func GetMethod(ctx context.Context) (string, error) {
	m, ok := ctx.Value(method).(string)
	if !ok {
		return "", errors.New("failed to get method")
	}
	return m, nil
}

// RequestURI returns requestURI or fallback.
func RequestURI(ctx context.Context, fallback string) string {
	u, err := GetRequestURI(ctx)
	if u == "" || err != nil {
		return fallback
	}
	return u
}

// GetRequestURI returns requestURI.
func GetRequestURI(ctx context.Context) (string, error) {
	u, ok := ctx.Value(requestURI).(string)
	if !ok {
		return "", errors.New("failed to get requestURI")
	}
	return u, nil
}
