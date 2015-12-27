package main

import(
	"github.com/gin-gonic/gin"
	"crypto/rand"
	"encoding/hex"
	)

/*
login
see docs for differnt auth levels
*/
func auth(c *gin.Context, level int64) bool{
	session := c.Sessions("user")
	login := session.GetInt("login", 0)

	switch level {
	//No login needed. Just Always return true
	case 0:
		return true

	case 1:
		
	case 2:
		
	case 3:
		if login != 0 {
			return true
		}
	case 4:
		if login != 0 {
			return true
		}
	}

	return false
}

func isuser (){

}

//Standard token of lenght 10
func genToken() string {
	return genToken_len(10)
}
//Copy from go doc (https://golang.org/pkg/crypto/rand/#example_Read)
//but tweeked so it can have a spesifided len
func genToken_len(len int) string {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
