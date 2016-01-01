package main

import(
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
//TODO: fix this so it can just store the sesssion in the interface
type db_obj struct {
	session *mgo.Session
}
func db_middle(con *mgo.Session) gin.HandlerFunc{
	return func(c *gin.Context) {
		session := con.Copy()
		c.DB = session
		defer session.Close()
		c.Next()
	}
}
func dial_db(c *gin.Context) *mgo.Session{
//Connet to the DB
	/*
	//con, err := mgo.Dial(mongo_address)
	con, err := c.Get("db")
	if err != false {
		c.AbortWithStatus(500)
		return nil
	}
	set := con.(db_obj).session
	*/
	return c.DB
}

//Gets the currently logged in user or returns nil if there is none
func (u *User) Current(c *gin.Context, db *mgo.Collection) *User {
	session := c.Sessions("user")
	hxid := session.GetString("uid", "")
	//If there is no loggedin user return nil
	if hxid == "" {
		return nil
	}
	uid := bson.ObjectIdHex(hxid)
	dberr := db.Find(bson.M{"_id": uid}).One(&u)
	if dberr != nil {
		return nil
	}
	return u
}

func (u User) getByEmail(email string, db *mgo.Collection) *User {
	//Query to send to the db
	q := bson.M{"email": email}
	dberr := db.Find(q).One(&u)
	if dberr != nil {
		return nil
	}
	return &u
}

func (u *User) getById(id bson.ObjectId, db *mgo.Collection) *User {
	//Query to send to the db
	q := bson.M{"_id": id}
	dberr := db.Find(q).One(&u)
	if dberr != nil {
		return nil
	}
	return u
}

func (p Pass)getById(id bson.ObjectId, db *mgo.Collection) *Pass {
	//Query to send to the db
	q := bson.M{"_id": id}
	dberr := db.Find(q).One(&p)
	if dberr != nil {
		return nil
	}
	return &p
}