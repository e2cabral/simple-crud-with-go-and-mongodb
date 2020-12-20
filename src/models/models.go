package models

// User is a schema from MongoDB
type User struct {
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}
