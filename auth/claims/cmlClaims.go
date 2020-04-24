package claims

import (
	"fmt"
	"time"
)

// CmlClaims are the claims for JWT tokens
type CmlClaims struct {
	ID        string
	Username  string
	Roles     []string
	ExpiresAt string
}

// Valid checks whether the claims the user has are valid.  Both ROLE_ADMIN and ROLE_USER are valid.
func (claims CmlClaims) Valid() error {
	expireTime, err := time.Parse(AuthDateFormat, claims.ExpiresAt)

	if err != nil {
		return AuthError{When: time.Now(), What: "Failed to parse the expirationDate"}
	}

	if !(time.Now().Unix() <= expireTime.Unix()) {
		return AuthError{When: time.Now(), What: "Token has expired"}
	}

	fmt.Printf("Claims: %+v \n\n", claims)

	return nil
}

//AuthError - Error type for Auth
type AuthError struct {
	When time.Time
	What string
}

func (e AuthError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}
