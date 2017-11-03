package main

import ("fmt")

type User struct {
	UserID int
  UserName string
  preferredSize string := ["M","L"]
}
type test struct {
  id string
  name string
}


func main()  {
  user1 :=  User {UserID:1234,UserName:"xyz"}
  fmt.Println(user1)

}
