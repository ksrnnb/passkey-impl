package jwt

type JwtKvs struct {
	invalidatedTokens []string
}

var jwtKvs JwtKvs

func init() {
	jwtKvs = JwtKvs{}
}

func isInvalidated(tokenString string) bool {
	for _, invalidatedTokenString := range jwtKvs.invalidatedTokens {
		if tokenString == invalidatedTokenString {
			return true
		}
	}
	return false
}

func Invalidate(tokenString string) {
	jwtKvs.invalidatedTokens = append(jwtKvs.invalidatedTokens, tokenString)
}
