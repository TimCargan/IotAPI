package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
    session := c.Sessions("user")
    //Validate the user hasnt tried to login more than the set number of times
    //TODO: this is a crude way to do it. Need to find a better way
    attemts := session.GetInt("attemts", 0)
    if attemts > MAX_LOGIN_ATTEMPTS {
    	c.String(500, "stop hacking")
    	return
    }else {
    	//Increase attempt counter
    	session.Set("attemts", attemts + 1)
    }


    //Connet to the DB
    con, err := mgo.Dial(mongo_address)
    if err != nil {
    	panic(err)
    }
    defer con.Close()

    //Load the user db and the pass db
	db := con.DB("test").C("foo")
	pass := con.DB("pass").C("pass")

	//TODO: might be worth it passing a token to ensure its the same user.
    //Help with the max login attempts.
	login := Login{}
	c.Bind(&login)


	//User struct to store the db respnce
	user := User{}
	//Query to send to the db
	q := bson.M{"email": login.Email}
	//Query DB
	dberr := db.Find(q).One(&user)
	if dberr != nil {
		c.String(200, "User not found")
	}

	hash := Pass{}
	q = bson.M{"_id": user.Id}
	dberr = pass.Find(q).One(&hash)
	if dberr != nil {
		c.String(200, "pass not found")

	}

	//Check to see if password is valid
	pass_valid := bcrypt.CompareHashAndPassword(hash.Hash, []byte(login.Pass))
	if pass_valid == nil {
		//TODO: refactor, move set auth into auth
		session.Set("uid", user.Id.Hex())
		session.Set("login", 1)
		c.String(200, ":")
	}else{
    	c.String(200, "nope " + login.Email)
	}
}