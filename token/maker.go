package token

import "time"


type Maker interface {
	CreateToken(userID int32, email string, duration time.Duration)(string,error)
	VerifyToken(token string) (*Payload, error)
}