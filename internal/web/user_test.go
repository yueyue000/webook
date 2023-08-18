package web

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := "hello#world123"
	// 加密
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	// 比较
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	require.NoError(t, err)
}
