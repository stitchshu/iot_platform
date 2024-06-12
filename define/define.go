package define

import "github.com/golang-jwt/jwt/v5"

type M map[string]interface{}

type UserClaim struct {
	Id       uint   `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

var (
	JwtKey = "iot"
)
