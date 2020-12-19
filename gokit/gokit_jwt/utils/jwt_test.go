package utils

import (
	"testing"
)

func TestCreateJwtToken(t *testing.T) {
	jwtToken, err := CreateJwtToken("xmr", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(jwtToken)
	jwtInfo, err := ParseToken(jwtToken)
	if err != nil {
		t.Error(err)
	}
	t.Log(jwtInfo)
}
