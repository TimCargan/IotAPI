package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/gzip"
	//"github.com/goibibo/gin-gzip"
	"github.com/plimble/sessions/store/mongo"
	"gopkg.in/mgo.v2"
)

const mongo_address = "localhost"
const MAX_LOGIN_ATTEMPTS = 20


func main() {
	a := gin.Default()

	//Set DB connection
	mongo_s, err := mgo.Dial(mongo_address)
	if err != nil {
		panic("db error")
	}
	
	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode.
	mongo_s.SetMode(mgo.Monotonic, true)

	//Set session properties
	//TODO: Needs to be improved 
	so := gin.SessionOptions{"","localhost",300,false,false}
	store := mongo.NewMongoStore(mongo_s.Copy(), "test", "session")

	a.Use(gzip.Gzip(gzip.DefaultCompression))
	a.Use(db_middle(mongo_s.Copy()))
	a.Use(gin.Session(store, &so))

	a.LoadHTMLGlob("/Users/tim/websites/iot/*")
	make_routs(a)
	//Set static folders
	a.Static("/static/folder", "/Users/tim/websites/cargancode.github.io")

	//Start the servers
	go func(){ a.Run(":8081") }()
	a.RunTLS(":5050","/Users/tim/websites/tim_cert.pem", "/Users/tim/websites/tim_key.pem")
}