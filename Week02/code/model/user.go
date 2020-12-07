package model

type User struct {
	Name string
	Age  int
}

func GenDefaultUser() *User {
	return &User{
		Name: "未知",
		Age:  18,
	}
}
