package middleware

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func Getcors() []rest.RunOption {
	var opts []rest.RunOption
	opt := rest.WithCustomCors(func(header http.Header) {
		header.Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id,OS,Platform, Version")
		header.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS,PATCH")
		header.Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, token")
	}, nil, "*")
	opts = append(opts, opt)

	domains := []string{"*", "http://127.0.0.1", "http://localhost"}
	opt = rest.WithCors(domains...)
	opts = append(opts, opt)

	return opts
}
