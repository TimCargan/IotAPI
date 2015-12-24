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
type User struct {
	V int 				`bson:"v" json:"v"`
	Id bson.ObjectId 	`bson:"_id" json:"id"`
	Username string 	`bson:"username" json:"username"`
	Email string 		`bson:"email" json:"email"`
	Dob	time.Time		`bson:"dob" json:"dob"`
	Name Name 			`bson:"name" json:"name"`
}

type Name struct{
	Fullname string 	`bson:"fullname" json:"fullname"`
	Nickname string 	`bson"nickname" json:"nickname"`
}

type Pass struct{
	V int 				`bson:"v" json: "v"`
	Id bson.ObjectId 	`bson:"_id" json: "id"`	
	Hash []byte 		`bson"has" json: "hash"`
}
