# ARFID Encyclopedia backend

[![Go Report Card](https://goreportcard.com/badge/github.com/ARFID-Encyclopedia/backend)](https://goreportcard.com/report/github.com/ARFID-Encyclopedia/backend)

## `main.go`

## `api.go`

## `auth.go`

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