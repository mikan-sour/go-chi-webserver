package domains

import (
	"fmt"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionId int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}

	if u.SessionId == 0 {
		return fmt.Errorf("token has no sessionId")
	}

	return nil
}
