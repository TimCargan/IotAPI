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
	
	session_t := get_session(c)
    if session_t == nil {
    	return
    }
	session := get_session(c).Get("user")
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
    
    session_t := get_session(c)
    if session_t != nil {
    	session := session_t.Get("user")
    	attemts := session.GetInt("attemts", 0)
    	if attemts > MAX_LOGIN_ATTEMPTS {
    		c.String(500, "stop hacking")
    		return
    	}else {
    		//Increase attempt counter
    		session.Set("attemts", attemts + 1)
    	}
	}
    
    con := dial_db(c)
		
    //Load needed db
	db := con.DB("user").C("users")
	pw_db := con.DB("pass").C("pass")

	login := Login{}
	c.Bind(&login)

	//Get User
	user := User{}.getByEmail(login.Email, db)
	if user == nil {
		abort_login(c)
		return
	}
	//Get Password
	hash := Pass{}.getById(user.Id, pw_db)
	if hash == nil {
		abort_login(c)
		return
	}

	if hash.password_valid(c, login.Pass){
		set_login_user(c, user)
		c.JSON(200, gin.H{"Hi " : user.Name.Nickname,})
		return
	}

	abort_login(c)
}

func abort_login(c *gin.Context) {
	c.JSON(401, gin.H{"stat": "Incorect Username or password",})
	c.Next()
}
func set_login_user(c *gin.Context, user *User){
	session_t := get_session(c)
    if session_t == nil {
    	abort_login(c)
    	return
    }
    session := session_t.Get("user")
	session.Set("uid", user.Id.Hex())
	session.Set("login", 1)
}

func (p Pass) password_valid(c *gin.Context, pass string) bool{
	pass_valid := bcrypt.CompareHashAndPassword(p.Hash, []byte(pass))
	if pass_valid == nil {
		return true
	}
    return false
}