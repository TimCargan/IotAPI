package main

import(
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func dial_db(c *gin.Context) *mgo.Session{
//Connet to the DB
	con, err := mgo.Dial(mongo_address)
	if err != nil {
		c.AbortWithError(500, err)
		return nil
	}
	return con
}


func (u User)getByEmail(email string, db *mgo.Collection) *User {
	//Query to send to the db
	q := bson.M{"email": email}
	dberr := db.Find(q).One(&u)
	if dberr != nil {
		return nil
	}
	return &u
}

func (u User)getById(id string, db *mgo.Collection) *User {
	//Query to send to the db
	q := bson.M{"email": id}
	dberr := db.Find(q).One(&u)
	if dberr != nil {
		return nil
	}
	return &u
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