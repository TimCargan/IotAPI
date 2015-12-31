package main

import (
	"github.com/gin-gonic/gin"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	)

/*
Handler for /
*/
func login_get(c *gin.Context) {
	session := c.Sessions("user")
	uid := session.GetString("uid", "none")
	c.HTML(200, "login.html", gin.H{
            "title": uid,
     	})
}

/*
Post handler for /login
*/
func login_post(c *gin.Context) {
	//Validate the user hasnt tried to login more than the set number of times
    //TODO: this is a crude way to do it. Need to find a better way
    session := c.Sessions("user")
    attemts := session.GetInt("attemts", 0)
    if attemts > MAX_LOGIN_ATTEMPTS {
    	c.String(500, "stop hacking")
    	return
    }else {
    	//Increase attempt counter
    	session.Set("attemts", attemts + 1)
    }


    con := dial_db(c)
    defer con.Close()
    //Load needed db
	db := con.DB("user").C("users")
	pw_db := con.DB("pass").C("pass")

	login := Login{}
	c.Bind(&login)

	//
	user := User{}.getByEmail(login.Email, db)
	if user == nil {
		//abort_login(c)
	}

	hash := Pass{}.getById(user.Id, pw_db)
	if hash == nil {
		//abort_login(c)
	}

	if hash.password_valid(c, login.Pass){
		//set_loggedin_user(c, user)
		return
	}


	
}

func set_login_user(c *gin.Context, user User){
	session := c.Sessions("user")
	session.Set("uid", user.Id.Hex())
	session.Set("login", 1)
	c.String(200, "Hi " + user.Name.Nickname)
}
func (p Pass) password_valid(c *gin.Context, pass string) bool{
	pass_valid := bcrypt.CompareHashAndPassword(p.Hash, []byte(pass))
	if pass_valid == nil {
		return true
	}
    return false
}