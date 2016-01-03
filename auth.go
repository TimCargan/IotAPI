package main

import(
	"github.com/gin-gonic/gin"
	"crypto/rand"
	"encoding/hex"
	"github.com/plimble/sessions"

	)

func get_session(c *gin.Context) *sessions.Sessions {
	session, err := c.Get("sessions")
	if err != false {
		return nil
	}
	ses_cat := session.(*sessions.Sessions)
	return ses_cat
}
func get_sessions(c *gin.Context) *sessions.Sessions {
	session, err := c.Get("sessions")
	if err != false {
		
	}
	ses_cat := session.(*sessions.Sessions)
	return ses_cat
}
/*
login
see docs for differnt auth levels
*/
func auth(c *gin.Context, level int64) bool{

	ses_cat := get_sessions(c)
	if ses_cat == nil {
		return false
	}
	//session := get_session(c)
	hold := ses_cat.Get("user")
	login := hold.GetInt("login", 0)
	/*
	session := c.Sessions("user")
	login := session.GetInt("login", 0) 
	*/

	switch level {
	//No login needed. Just Always return true
	case 0:
		return false

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
