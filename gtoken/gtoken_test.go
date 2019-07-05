package gtoken

import (
	"testing"
)

func TestEncryptDecryptToken(t *testing.T) {
	t.Log("encrypt and decrypt token test ")
	gfToken := GfToken{}
	gfToken.Init()

	userKey := "123123"
	token := gfToken.EncryptToken(userKey)
	if !token.Success() {
		t.Error(token.Json())
	}
	t.Log(token.DataString())

	token2 := gfToken.DecryptToken(token.GetString("token"))
	if !token2.Success() {
		t.Error(token2.Json())
	}
	t.Log(token2.DataString())
	if userKey != token2.GetString("userKey") {
		t.Error("token decrypt userKey error")
	}
	if token.GetString("uuid") != token2.GetString("uuid") {
		t.Error("token decrypt uuid error")
	}

}
