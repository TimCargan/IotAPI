package home

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
func new(c *gin.Context) Home{
	h := Home{}
	h.V = HOME_VERSION
	h.Id =	bson.NewObjectId()
	//h.add_user(user.Current(c)) once the users package is nice and like this one
	//Once there is a user package
	//user := Users{}.Current()
	//h.add_user(user)
	return h
}

//will validate weather a given house object is valid/ fully formed
func (h Home) valid() bool{
	if h.V == HOME_VERSION && h.Id != nil && h.Name != "" {
		return true
	}
	return false
}
//Invers of valid, helps to make things a bit more readable
func (h Home) invalid() bool {
	return !h.valid()
}

//Is the user a valid editor i.e is the user in the slice of users
func (h Home) isuser(u User) bool {
	   for _, el := range h.Users {
        if  el.Id == u.Id {
            return true
        }
    }
    return false
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