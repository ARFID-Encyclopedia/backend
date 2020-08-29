# ARFID Encyclopedia backend

[![Go Report Card](https://goreportcard.com/badge/github.com/ARFID-Encyclopedia/backend)](https://goreportcard.com/report/github.com/ARFID-Encyclopedia/backend)

## `main.go`

This file bootstrap the server, housing the request router and the main function

At the moment `main()` contains some dummy data for testing purposes

## `api.go`

`api.go` defines how the different API routes should behave as per this table:

| Endpoint | method | function    | description |
|----------|--------|-------------|-------------|
| `/`      | GET    | `homePage`  | This is temporary and just lets a client know the server is up and operational without returning any data|
| `/foods` | GET    | `returnAll` | Returns data on all food in the database, as of right now the only data in the program is defined in main and is, again, temporary |
| `/food`  | GET   | `createNewFood` | Adds a new food to the database |
| `/food/{name}` | POST | `returnByName` | Gets a given foods data |
| `/login` | POST | `login` | Brokers the user an API token to access certain api endpoints. The way in which all the user stuff is done needs to be fleshed out |


## `auth.go`

Manages authorisation of users

This API uses JWT or javascript web tokens to ensure only users of the correct access level can access certain functions of the API. The details of this can be viewed [here](#access-levels)

## `database.go`

The database will use scribble mainly as it is very light

There will be 2 collections:
* Users
* Foods

### Users

A user will look like this:
```go
type user struct {
	ID           uint64 `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passHash"`
	AccessLevel  int    `json:"accessLevel"`
}
```

with IDs being unique user identifiers, usernames, passwords and access levels
An access level dictates what a user can and cannot do

#### Access Levels

| Access level | Alias | Description |
|--------------|-------|-------------|
| 0            | ANON  | The user hasn't logged in and has no rights to add edit or remove content |
| 1            | USER  | The user has logged in and has access to adding editing and removing but the actions must be approved by a mod or admin |
| 2            | MOD   | Mods have the right to approve edits and add new foods |
| 3            | ADMIN | Admins can do everything a mod can do plus the ability do delete foods |

### Foods

Foods are described under this `struct`
```go
type food struct {
	Name      string   `json:"Name"`
	Category  string   `json:"Category"`
	Visual    string   `json:"Visual"`
	Texture   string   `json:"Texture"`
	Smell     string   `json:"Smell"`
	Taste     string   `json:"Taste"`
	Nutrients []string `json:"Nutrients"`
}
```

They follow the format dictated in the google doc