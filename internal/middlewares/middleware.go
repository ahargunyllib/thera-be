package middlewares

import "github.com/ahargunyllib/thera-be/pkg/jwt"

type Middleware struct {
	jwt jwt.CustomJwtInterface
}

func NewMiddleware(
	jwt jwt.CustomJwtInterface,
) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}
