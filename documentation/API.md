#Users
This part of the API is just for dealing with users. Standard JSON user object that will be returned looks as follows:

	{
	  "v": 0,
	  "id": "5681f9abd6f45f8f98c69dfe",
	  "istemp": false,
	  "email": "son@gmail.com",
	  "path": "t3",
	  "username": "T3",
	  "dob": "0001-01-01T00:00:00Z",
	  "name": {
	    "fullname": "Tim Cargan",
	    "nickname": "Tim"
	  }
	}


/users/

##Post 
	Crates a new temp user (They will need to validate their email or the account will stay as temp and be removed from the db). The user will have to 

##Get
	The logged in user will be returned (token owner). Else a 404 will be thrown

##Put
	404
##Delete
	404 

/users/:id
##Post -- Unimplemented
	In order to protect user privacy will return a 404

##Get
	Will get the given user provided there is a logged in user
	If the user doesn’t exist a 404 will be returned
##Put
	Will update the given users according properties
	Takes a JSON user object in the body of the request
		{

		Username string 	`bson:"username" json:"username,omitempty" structs:"username,omitempty"`
		Dob	time.Time		`bson:"dob" json:"dob,omitempty" structs:"dob,omitempty"`
			Name: {FullName: "Tim NewName Cargan", nickname: "New"}
		}

##Delete -- Unimplemented
	If the logged in user matches will remove the user