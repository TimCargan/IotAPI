#Users
This part of the API is just for dealing with users. Standard JSON user object that will be returned looks as follows:
```
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
```
The current auth flow is cookie bassed, however Once the auth middleware is develped in additon to cookies, tokens can and will be used. Once that is in place a better auth flow will be set up. However some things will like remain the same

* Session state will persist accross the token e.g
	Any API limits will apply to the token

* Rate limmits will exist be applied to the user and are token agnostic. This is in additon to standard IP rate limmiting etc.

<hr>

##/users/

###Post 
Crates a new temp user (They will need to validate their email or the account will stay as temp and be removed from the db). The user will have steps to complete in order to validate the account. See "User creation flow" for more info

Expected Body:
```
{
  "email": "son@gmail.com",
  "username": "T3",
  "dob": "0001-01-01T00:00:00Z",
  "name": {
    "fullname": "Tim Cargan",
    "nickname": "Tim"
  }
}
```
###Get
The logged in user will be returned (token owner). Else a 404 will be thrown

###Put
405

###Delete
405

<hr>

##/users/:id

###Post -- Unimplemented
In order to protect user privacy will return a 404

###Get -- Needs to be updated
Will get the given user provided there is a logged in user
If the user doesn’t exist a 404 will be returned.

###Put
Will update the given users according properties
Takes a JSON user object in the body of the request
```
{
  "username": "T3",
  "dob": "0001-01-01T00:00:00Z",
  "name": {
    "fullname": "Tim Cargan",
    "nickname": "Tim"
  }
}
```

####Known Bugs
* You can freely update the users email without having to re-validate.
* The username, if updated isnt passed and updated in the internal path property, means they have 2 usernames

###Delete -- Unimplemented
If the logged in user matches will remove the user