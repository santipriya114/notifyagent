package middleware

func ValidateToken(token string) bool {

	return token == "test123"

}
