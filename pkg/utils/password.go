package utils

import "golang.org/x/crypto/bcrypt"

const (
	DefaultSaltCost = 7
)

func HashPassword(password string, saltCost ...int) (string, error) {
	cost := DefaultSaltCost
	if len(saltCost) > 0 {
		cost = saltCost[0]
	}
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bcryptPassword), nil
}

func ComparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
