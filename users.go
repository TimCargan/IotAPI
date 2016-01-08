package main

import(
	 "github.com/gin-gonic/gin"
	 "gopkg.in/mgo.v2"
	 "gopkg.in/mgo.v2/bson"
	 "github.com/fatih/structs"
	 "net/http"
	 "strings"
	 "golang.org/x/crypto/bcrypt"
	 //"strconv"
	 )

/*
Handler for POST /user/
Handler for creating a user
Takes a json object (of type user and login) and will add it to the database
with the temp user flag and send an email to the user to validate the email.
If the user doesnt respond to the email the user should be removed from the db.
*/
func user_post(c *gin.Context) {
	//DB connection
	mon:= dial_db(c)
	db := mon.DB("user").C("users")
	pw := mon.DB("pass").C("pass")

	user := User{}
	c.Bind(&user)

	//Set user info
	user.V = USER_V
	user.Id = bson.NewObjectId()
	user.Path = strings.ToLower(user.Username)
	user.Istemp = true
	user.EmailToken = genToken()

	//Attept to add user to db
	dberr := db.Insert(user)
	if mgo.IsDup(dberr){
		if strings.Contains(dberr.Error(), " email_1 ") {
			c.String(http.StatusOK, "email")
		}else{
			c.String(http.StatusOK, "user")
		}
		//return
	}else if dberr != nil{
		c.AbortWithError(http.StatusInternalServerError, dberr)
		//return
	}

	//Create pass object
	pass := Pass{}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Pass), 12)
	if err != nil {
		//TODO: Should prob invalidate user account
		c.AbortWithError(500, err)
	}
	pass.Hash = hash
	pass.Id = user.Id
	pass.Istemp = true
	//Insert password into db
	dberr = pw.Insert(pass)
	if dberr != nil {
		c.AbortWithError(500, dberr)
	}

	c.JSON(http.StatusOK, user )
}

func user_validate_email(c *gin.Context){
	//DB connection
	mon:= dial_db(c)
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
	/*
	if auth(c, 3) == false {
		c.JSON(401, gin.H{"status": "Unauthorized"})
		return
	}
	*/
	hxid := c.Param("uid")
	//DB connection
	mon := dial_db(c)
	db := mon.DB("user").C("users")
	user := &User{}

	if hxid != "" {
		//Fetch user from db
		if !bson.IsObjectIdHex(hxid) {
			c.JSON(404, gin.H{"status": "User not found"})
			return
		}
		uid := bson.ObjectIdHex(hxid)
		user = user.getById(uid, db)
		if user == nil {
			c.JSON(404, gin.H{"status": "User not found"})
			return
		}
	}else{
		if user.Current(c, db) == nil {
			c.JSON(404, gin.H{"status": "User not found"})
			return
		}
	}
	c.IndentedJSON(200, user)
}


/*
Handler for Update /user/:uid
TODO: BUG FIX, users can just update there email whenever... Should fix
*/
func user_put(c *gin.Context) {
	hxid := c.Param("uid")


	if auth(c, 4) == false {
		c.String(403, "Must login")
		return
	}

	//DB connection
	mon := dial_db(c)
	db := mon.DB("user").C("users")


	user := User{}
	c.Bind(&user)

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
