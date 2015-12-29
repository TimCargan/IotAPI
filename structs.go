package main

import( 
	"time"
	"gopkg.in/mgo.v2/bson"
	)

type Login struct {
	Email string 		`json:"email"`
	Username string 	`json:"username"`
	Pass string 		`json:"pass"`
}
const USER_V = 0
type User struct {
	V int 				`bson:"v" json:"v" structs:"-"`
	Id bson.ObjectId 	`bson:"_id" json:"id,omitempty" structs:"-"`

	//Auth flow from temp to email verification
	Istemp bool			`bson:"istemp,omitempty" json:"istemp,omitempty" structs:"-"`
	EmailToken string	`bson:"emailtoken,omitempty" json:"emailtoken,omitempty" structs:"-"`
	Email string 		`bson:"email" json:"email,omitempty" structs:"email,omitempty"`

	//username, path is lowercase
	Path string			`bson:"path" json:"path,omitempty" structs:"-"`
	Username string 	`bson:"username" json:"username,omitempty" structs:"username,omitempty"`
	
	Dob	time.Time		`bson:"dob" json:"dob,omitempty" structs:"dob,omitempty"`
	Name Name 			`bson:"name" json:"name,omitempty" structs:"name,omitempty"`
}

type Name struct{
	Fullname string 	`bson:"fullname,omitempty" json:"fullname,omitempty" structs:"fullname,omitempty"`
	Nickname string 	`bson:"nickname,omitempty" json:"nickname,omitempty" structs:"nickname,omitempty"`
}
type settings  struct {}

const PASS_V = 0
type Pass struct{
	V int 				`bson:"v" json: "v"`
	Id bson.ObjectId 	`bson:"_id" json: "id"`	
	Hash []byte 		`bson"hash" json: "hash"`

	Istemp bool			`bson:"istemp,omitempty" json:"istemp"`

	ResetPasswordToken string `bson"resetToken" json: "resetPasswordToken"`
  	ResetPasswordExpires time.Time  `bson"resetExpires" json: "resetPasswordExpires"`
}
