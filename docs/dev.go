package docs

import "gopkg.in/mgo.v2/bson"



const HOME_VERSION =  1

type Home struct{
	V int 				`bson:"v" json:"v" structs:"-"`
	Id bson.ObjectId 	`bson:"_id" json:"id,omitempty" structs:"-"`
	Name string 		`bson:"name" json: "name" structs:"name"`
	Users []string 		`bson:"users" json: "users" structs:"-"`
	Devs  []Dev			`bson:"devs" json: "devs" structs:"-"`
	Rooms []Room		`bson:"rooms" json: "rooms" structs:"-"`
}

type Room struct{
	V int 				`bson:"v" json:"v" structs:"-"`
	Id bson.ObjectId 	`bson:"_id" json:"id,omitempty" structs:"-"`
	Name 	string
	users	[]string
	dev 	[]Dev
}


type Dev struct {
	id 		bson.ObjectId
	v		string
	mac		string
	props	[]Propertie
}

type Propertie struct {
	Id 		bson.ObjectId
	name 	string
	value	int
}

func (h Home) new(){
	h.V = HOME_VERSION
	h.Id =	bson.NewObjectId()
}

/*

//TODO: need to see how append works
//Adds a new user to the end of the user list in home 
See for more info https://blog.golang.org/slices
For helpful tips see https://github.com/golang/go/wiki/SliceTricks
func (h *Home) add_user(new_user User){
	h.Users = append(h.Users, new_user)
}
*/