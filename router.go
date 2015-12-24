package main

import(
	 "github.com/plimble/ace"
	 "github.com/plimble/sessions/store/mongo"
	 "gopkg.in/mgo.v2"
	 "gopkg.in/mgo.v2/bson"
	 "golang.org/x/crypto/bcrypt"
	 )

const mongo_address = "localhost"
const MAX_LOGIN_ATTEMPTS = 20

func main() {
	a := ace.New()
	a.Use(ace.Logger())

	//Set DB connection
	session, err := mgo.Dial(mongo_address)
	if err != nil {
		panic("db error")
	}

	//Set session properties
	//TODO: Needs to be improved 
	so := ace.SessionOptions{"","localhost",300,true,false}
	store := mongo.NewMongoStore(session, "test", "session")
	a.Use(ace.Session(store, &so))


	a.GET("/", index)
	a.POST("/login", login_post)
	a.GET("/user/:uid/", user_get)

	//Set static folders
	a.Static("/static/folder", "/Users/tim/websites/cargancode.github.io")

	//Start the servers
	go func(){ a.Run(":8080") }()

	a.RunTLS(":5050","/Users/tim/websites/tim_cert.pem", "/Users/tim/websites/tim_key.pem")
	
}
/*
Handler for /
*/
func index(c *ace.C) {
	session := c.Sessions("user")
	uid := session.GetString("uid", "none")
	c.String(200, uid)
}
func index2(c *ace.C) {
    //get session name
    session1 := c.Sessions("user")
    session1.Set("user", "123 - i set this")

    
    sess, err := mgo.Dial(mongo_address)
	if err != nil {
		 c.String(500, "")
	}
	defer sess.Close()
	
	//sess.SetSafe(&mgo.Safe{})
	//collection := sess.DB("test").C("foo")
	//pw := sess.DB("pass").C("pass")
	user := User{ Id: bson.NewObjectId(), Username: "TimCargan",Email: "timcargan@gmail.com",Name: Name{
							Fullname:"Timothy R. Cargan",
							Nickname: "Tim"},

					}
	pass := Pass{Id: user.Id}
	dk, err := bcrypt.GenerateFromPassword([]byte("some password"), 12)
	pass.Hash = dk

	//err = collection.Insert(user)
	//err = pw.Insert(pass)

	//Hack to get it to compile without error checking
	if err != nil {}
	c.Writer.Header().Set("test", "test")
	c.String(200, string(dk[:32])  )
}

/*
Post handler for /login
*/
func login_post(c *ace.C) {
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

    //TODO: Read perams from post, might be worth it passing a token to ensure its the same user.
    //Help with the max login attempts.

	//name := c.Param("name")

    //Connet to the DB
    con, err := mgo.Dial(mongo_address)
    if err != nil {
    	c.Panic(err)
    }
    defer con.Close()

    //Load the user db and the pass db
	db := con.DB("test").C("foo")
	pass := con.DB("pass").C("pass")

	//Pass input
	login := Login{}
	c.ParseJSON(&login)


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
		session.Set("uid", user.Id.Hex())
		session.Set("login", true)
		c.String(200, "nope " + user.Id.Hex())
	}else{
    	c.String(200, "nope " + user.Id.String())
	}
}

/*
Handler for /test
*/
func test_cookieset(c *ace.C) {
    //session := c.Sessions("user")

	//name := c.Param("name")

    con, err := mgo.Dial(mongo_address)
    defer con.Close()
    if err != nil {
    	
    }

	db := con.DB("test").C("foo")
	pass := con.DB("pass").C("pass")

	email := "timcargan@gmail.com"
	val := User{Email: email}
	dberr := db.Find(bson.M{"email": email }).One(&val)
	if dberr != nil {
		c.String(200, "User not found")
	}

	hash := Pass{Id: val.Id}
	dberr_1 := pass.Find(bson.M{"_id": val.Id}).One(&hash)
	if dberr_1 != nil {
		c.String(200, "pass not found")

	}

	pass_valid := bcrypt.CompareHashAndPassword(hash.Hash, []byte("some password"))

	if pass_valid == nil {
		c.String(200, "Its a match ")
	}else{
    	c.String(200, "nope")
	}
}

/*
Handler for /thing/:name
testing handler for url perams
*/
func user_get(c *ace.C) {
	hxid := c.Param("uid")

	//Validate auth

	//Set DB connection
	mon, err := mgo.Dial(mongo_address)
	if err != nil {
		panic("db error")
	}
	db := mon.DB("test").C("foo")


	user := User{}
	uid := bson.ObjectIdHex(hxid)
	dberr := db.Find(bson.M{"_id": uid}).One(&user)
	if dberr != nil {
		c.String(200, "User not found " + string(uid))
	}

	c.JSON(200, user)
}