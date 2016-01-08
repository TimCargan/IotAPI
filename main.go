package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/secure"
	
	"github.com/plimble/sessions/store/mongo"
	"gopkg.in/mgo.v2"
)

const mongo_address = "localhost"
const MAX_LOGIN_ATTEMPTS = 20


func main() {

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

	a := gin.Default()
	//Enable GZip on requets
	a.Use(gzip.Gzip(gzip.DefaultCompression))
	//DB midle were so sessions can access mongo
	a.Use(db_middle(mongo_s.Copy()))
	//Enables sessions
	//TODO: Need to move out of Gin
	a.Use(gin.Session(store, &so))
	//Sets some usefull headers
	
	a.Use(secure.Secure(secure.Options{
		IsDevelopment:		   gin.IsDebugging(),
		AllowedHosts:          []string{"example.com", "ssl.example.com"},
		SSLRedirect:           false,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'; style-src http://cargancode.io;",
	}))
	
	a.LoadHTMLGlob("/Users/tim/websites/iot/*")
	make_routs(a)


	//Set static folders
	//a.Static("/static/folder", "/Users/tim/websites/cargancode.github.io")

	//Start the servers
	go func(){ a.Run(":8081") }()
	a.RunTLS(":5050","/Users/tim/websites/tim_cert.pem", "/Users/tim/websites/tim_key.pem")
}