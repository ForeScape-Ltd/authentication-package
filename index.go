package authenticationpkg

import (
	"time"

	"github.com/felixgiftinfo/authentication-pkg/token"
)

// type AuthenticationContext struct {
// 	TokenMaker token.Maker
// }

func TODO(tokenSecretKey, refreshTokenSecretKey string, tokenDuration, refreshTokenDuration time.Duration) (token.Maker, error) {

	jwt, err := token.NewJWTMaker(tokenSecretKey, refreshTokenSecretKey, tokenDuration)

	return jwt, err
}
