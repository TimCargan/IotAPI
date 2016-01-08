package home
import(
	 "github.com/gin-gonic/gin"
	 "gopkg.in/mgo.v2"
	 "net/http"
	 //"strconv"
	 )

/*
Handler for POST /home/:id
Handler for creating a home
*/
func home_post(c *gin.Context) {
	home := home.New(c)
	c.Bind(&home)

	if home.invalid() {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
	}
	
	mon:= dial_db(c)
	db := mon.DB("home").C("homes")
	
	dberr := db.Insert(home)
	if dberr != nil {
		c.AbortWithError(500, dberr)
	}

	c.JSON(http.StatusOK, home)
}

/*
Handler for Get /home/:id
Handler for getting a home
*/
func home_post(c *gin.Context) {
	home := home.New(c)
	c.Bind(&home)

	if home.invalid() {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
	}
	
	mon:= dial_db(c)
	db := mon.DB("home").C("homes")
	
	dberr := db.Insert(home)
	if dberr != nil {
		c.AbortWithError(500, dberr)
	}

	c.JSON(http.StatusOK, home)
}

func user_get(c *gin.Context) {
	hxid := c.Param("id")
	//DB connection
	mon := dial_db(c)
	db := mon.DB("user").C("users")
	user := Home{}

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