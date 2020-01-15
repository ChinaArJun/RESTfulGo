package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context is the context of the JSON web token.
type Context struct {
	ID       int
	Username string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err

		// Read the token if it's valid.
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = int(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// 自动解密
func ParseRequest(g *gin.Context) (*Context, error)  {
	header := g.Request.Header.Get("Authorization")
	// load the jwt secret from config
	secret := viper.GetString("jwt_secret")
	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}
	var t string
	fmt.Sscanf(header, "Bearer %s", &t)
	Context, error := Parse(t, secret)
	return Context, error
}

// 生成令牌
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error)  {
	// Load the jwt secret from the gin config
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	// the token content 使用id username 生成token令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": c.ID,
		"username": c.Username,
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
	})
	tokenString , err = token.SignedString([]byte(secret))
	return
}