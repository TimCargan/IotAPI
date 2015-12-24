package main

import(
	"github.com/gin-gonic/gin"
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

func isuser

