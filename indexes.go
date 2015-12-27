package main

import (
	"gopkg.in/mgo.v2"
	"time"
)

var (
user_new_index mgo.Index = mgo.Index{
	    Key: []string{"istemp"},
	    Unique: false,
	    DropDups: false,
	    Background: true, // See notes.
	    Sparse: true,
	    ExpireAfter: 5 * 24 * time.Hour, //5 days
	    //parttial on "partialFilterExpression: { insnew: true } }
	}
user_username_index mgo.Index = mgo.Index{
	    Key: []string{"path"},
	    Unique: true,
	    DropDups: false, // Makes sure an error is returned if insersion is attemted.
	    Background: true, 
	    Sparse: true,
	}
user_email_index mgo.Index = mgo.Index{
	    Key: []string{"email"},
	    Unique: true,
	    DropDups: false, // Makes sure an error is returned if insersion is attemted.
	    Background: true,
	    Sparse: true,
	} 

)
