package main

import(
	 "github.com/gin-gonic/gin"
	 "gopkg.in/mgo.v2"
	 "gopkg.in/mgo.v2/bson"
	 
	 )
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


	if auth(c, 4) == true {
		c.String(403, "Must login")
		return
	}

	//DB connection
	mon, err := mgo.Dial(mongo_address)
	if err != nil {
		panic("db error")
	}
	db := mon.DB("test").C("foo")

	//Fetch user from db
	user := User{}
	uid := bson.ObjectIdHex(hxid)
	dberr := db.Find(bson.M{"_id": uid}).One(&user)
	if dberr != nil {
		c.String(200, "User not found " + string(uid))
	}
	c.JSON(200, user)
}