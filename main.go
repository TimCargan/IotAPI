package main

import (
	"github.com/gin-gonic/gin"
	"github.com/plimble/sessions/store/mongo"
	"gopkg.in/mgo.v2"
)

const mongo_address = "localhost"
const MAX_LOGIN_ATTEMPTS = 20

var (
	Mongo *mgo.Session
	)
func main() {
	a := gin.Default()

	//Set DB connection
	Mongo, err := mgo.Dial(mongo_address)
	if err != nil {
		panic("db error")
	}

	//Set session properties
	//TODO: Needs to be improved 
	so := gin.SessionOptions{"","localhost",300,true,false}
	store := mongo.NewMongoStore(Mongo.Copy(), "test", "session")
	a.Use(gin.Session(store, &so))

	a.LoadHTMLGlob("/Users/tim/websites/iot/*")
	make_routs(a)
	//Set static folders
	a.Static("/static/folder", "/Users/tim/websites/cargancode.github.io")

	//Start the servers
	go func(){ a.Run(":8080") }()
	a.RunTLS(":5050","/Users/tim/websites/tim_cert.pem", "/Users/tim/websites/tim_key.pem")
	
}