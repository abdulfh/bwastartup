package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
}

var SecretKey = []byte("BWASTARTUP")

func NewService() *jwtService {
	return &jwtService{}
}

func (service *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SecretKey)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
