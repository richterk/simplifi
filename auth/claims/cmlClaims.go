package claims

import (
	"fmt"
	"strings"
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
	if !(time.Now().Unix() <= int64(claims.ExpiresAt)) {
		return AuthError{When: time.Now(), What: "token has expired"}
	}
	fmt.Printf("Claims: %+v \n\n", claims)
	for i := 0; i < len(claims.Roles); i++ {
		realRole := strings.Replace(claims.Roles[i], "\"", "", 2) // So stupid that this has to be done.
		fmt.Printf("Role: %+v, %v - Username: %s \n\n", claims.Roles[i], realRole == "ROLE_ADMIN", claims.Username)
		if realRole == "ROLE_ADMIN" {
			fmt.Printf("Admin...\n")
			return nil
		}

		if realRole == "ROLE_USER" {
			fmt.Printf("User...\n")
			return nil
		}
	}

	fmt.Printf("Error...\n")
	return AuthError{When: time.Now(), What: "You don't have the required role"}
}

//AuthError - Error type for Auth
type AuthError struct {
	When time.Time
	What string
}

func (e AuthError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}
