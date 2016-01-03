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
	r.GET("/user/", user_get)
	r.POST("/user", user_post)
}

/*
Handler for /
*/
func index(c *gin.Context) {
	session := get_session(c).Get("user")
	uid := session.GetString("uid", "none")
	c.HTML(http.StatusOK, "index.html", gin.H{
            "title": uid,
     	})
}
