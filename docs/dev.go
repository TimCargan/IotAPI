package docs

import "gopkg.in/mgo.v2/bson"



const HOME_VERSION =  1

type Home struct{
	V int 				`bson:"v" json:"v" structs:"-"`
	Id bson.ObjectId 	`bson:"_id" json:"id,omitempty" structs:"-"`
	Name string 		`bson:"name" json: "name" structs:"name,omitempty"`
	Users []string 		`bson:"users" json: "users" structs:"-"`
	Devs  []Dev			`bson:"devs" json: "devs" structs:"-"`
	Rooms []Room		`bson:"rooms" json: "rooms" structs:"-"`
}

type Room struct{
	V int 				`bson:"v" json:"v" structs:"-"`
	Id bson.ObjectId 	`bson:"_id" json:"id,omitempty" structs:"-"`
	Name string			`bson:"name" json:"name" structs:"name,omitempty"`
	users []string		`bson:"users" json:"users" structs:"-"`
	Devs []Dev 			`bson:"devs" json:"devs" structs:"-"`
}


type Dev struct {
	Id 	bson.ObjectId 	`bson:"_id" json:"id" structs:"-"`
	V	string			`bson:"v" json:"v" structs:"-"`
	Mac	string			`bson:"mac" json:"mac" structs:"-"`
	Props []Propertie	`bson:"props" json:"props" structs:"-"`
}

type Propertie struct {
	Id 	bson.ObjectId 	`bson:"_id" json:"id" structs:"-"`
	Name string 		`bson:"name" json:"name" structs:"name,omitempty"`
	Value int 			`bson:"value" json:"vlaue" structs:"value,omitempty"`
	kind int 			`bson:"kind" json:"kind" structs:"-"`
}


///////////////////////////////////////////////////////////////////////////
//-------------------------Functions------------------------------------//
/////////////////////////////////////////////////////////////////////////
func (h *Home) new(){
	h.V = HOME_VERSION
	h.Id =	bson.NewObjectId()
}

/*
TODO: Coppy for add dev and add room
TOOD: Create remove user, room and dev
//TODO: need to see how append works
//Adds a new user to the end of the user list in home 
See for more info https://blog.golang.org/slices
For helpful tips see https://github.com/golang/go/wiki/SliceTricks
func (h *Home) add_user(new_user User){
	h.Users = append(h.Users, new_user)
}
*/