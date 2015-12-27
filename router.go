package main

import(
	 "github.com/gin-gonic/gin"
	 "net/http"
	 )

func make_routs(r *gin.Engine) {
	r.GET("/", index)
	r.GET("/login", login_get)
	r.POST("/login", login_post)
	r.GET("/user/:uid/auth", user_validate_email)
	r.GET("/user/:uid/", user_get)
	r.PUT("/user/:uid/", user_put)
	r.POST("/user", user_post)
}

/*
Handler for /
*/
func index(c *gin.Context) {
	session := c.Sessions("user")
	uid := session.GetString("uid", "none")
	c.HTML(http.StatusOK, "index.html", gin.H{
            "title": uid,
     	})
}
