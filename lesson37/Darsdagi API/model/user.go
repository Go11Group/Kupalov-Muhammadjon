package model

type User struct {
	Id         string
	FirstName  string
	LastName   string
	Age        int
	Gender     string
	Nation     string
	Feild      string
	ParentName string
	City       string
}

type UserFilter struct {
	FirstName  *string
	LastName   *string
	Age        *int
	Gender     *string
	Nation     *string
	Feild      *string
	ParentName *string
	City       *string
	Limit      *int
	Offset     *int
}
