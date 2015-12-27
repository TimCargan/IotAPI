package main

import(
	 "github.com/gin-gonic/gin"
	 "gopkg.in/mgo.v2"
	 "gopkg.in/mgo.v2/bson"
	 "github.com/fatih/structs"
	 "net/http"
	 "strings"
	 //"golang.org/x/crypto/bcrypt"
	 //"strconv"
	 )

/*
Handler for POST /user/
Handler for creating a user
Takes a json object (of type user) and will add it to the database

*/
func user_post(c *gin.Context) {
	//DB connection
	mon, err := mgo.Dial(mongo_address)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	defer mon.Close()
	db := mon.DB("user").C("users")

	//Make sure the index is there
	err = db.EnsureIndex(user_new_index)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	err = db.EnsureIndex(user_username_index)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	err = db.EnsureIndex(user_email_index)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	//sess.SetSafe(&mgo.Safe{})
	//collection := sess.DB("test").C("foo")
	//pw := sess.DB("pass").C("pass")
	user := User{}
	pass := Pass{}
	login := Login{}
	c.Bind(&user)
	c.Bind(&login)

	//Santize user

	//Set user info
	user.V = USER_V
	user.Id = bson.NewObjectId()
	user.Path = strings.ToLower(user.Username)
	user.Istemp = true
	user.EmailToken = genToken()

	dberr := db.Insert(user)
	switch {
	case mgo.IsDup(dberr):
		if strings.Contains(dberr.Error(), " email_1 ") {
			c.String(http.StatusOK, "email")
		}else{
			c.String(http.StatusOK, "user")
		}
		return
	case dberr != nil:
		c.AbortWithError(http.StatusInternalServerError, dberr)
		return
	}

	//pw, err := bcrypt.GenerateFromPassword([]byte(login.Pass), 12)
	//pass.Hash = pw
	pass.Id = user.Id
	//err = pw.Insert(pass)

	//Hack to get it to compile without error checking
	c.JSON(http.StatusOK, user )
}

func user_validate_email(c *gin.Context){
	//DB connection
	mon, err := mgo.Dial(mongo_address)
	if err != nil {
		c.AbortWithError(500, err)
	}
	defer mon.Close()
	db := mon.DB("user").C("users")

	//Get token
	token := c.Query("t")
	//Get user Id
	hxid := c.Param("uid")
	uid := bson.ObjectIdHex(hxid)

	//Fetch user from db
	user := User{}
	dberr := db.Find(bson.M{"_id": uid}).One(&user)
	if dberr != nil {
		c.String(200, "User not found " + string(uid) )
	}

	//If the user email is already valid return
	if !user.Istemp {
		c.String(200, "Already Auth")
		return
	}
	if user.EmailToken != token{
		c.JSON(401, gin.H{"status": "unauthorized"})
		return
	}

	//Unset temp status
	m := make(map[string]string)
	m["Istemp"] =  ""
	m["EmailToken"] =""
	dberr = db.UpdateId(uid, bson.M{"$unset": m})
	if dberr != nil {
		c.Error(dberr)
	}
	c.String(202, "Email valid")
}
/*
Handler for Get /user/:uid
*/
func user_get(c *gin.Context) {
	hxid := c.Param("uid")

	if auth(c, 3) == false {
		c.String(401, "Must login")
		return
	}

	//DB connection
	mon, err := mgo.Dial(mongo_address)
	if err != nil {
		panic("db error")
	}
	defer mon.Close()
	db := mon.DB("test").C("foo")

	//Fetch user from db
	user := User{}
	uid := bson.ObjectIdHex(hxid)
	dberr := db.Find(bson.M{"_id": uid}).One(&user)
	if dberr != nil {
		c.String(200, "User not found " + string(uid) )
	}

	c.JSON(200, user)
}

/*
Handler for Update /user/:uid

*/
func user_put(c *gin.Context) {
	hxid := c.Param("uid")


	if auth(c, 4) == false {
		c.String(403, "Must login")
		return
	}

	//DB connection
	mon, err := mgo.Dial(mongo_address)
	if err != nil {
		panic("db error")
	}
	defer mon.Close()
	db := mon.DB("test").C("foo")


	
	user := User{}
	c.Bind(&user)

	structs.DefaultTagName = "json"
	m := structs.Map(user)
	
	uid := bson.ObjectIdHex(hxid)
	dberr := db.UpdateId(uid, bson.M{"$set": m})
	if dberr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(200, m)
}

//------------------------------------
//     Utils
//------------------------------------

func parse_user(u User){
	//pares_email()
	//parse_username()
}
