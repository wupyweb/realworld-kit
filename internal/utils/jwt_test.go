package utils

import (
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	token, _ := GenerateToken(9527)
	t.Log(token)
	
	time.Sleep(time.Second * 30)
	payload, err := ParseToken(token)
	if err != nil {
	    t.Error(err)
	}
	t.Log(payload.UserID)
}