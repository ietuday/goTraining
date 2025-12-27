package main

import (
	"example.com/struct-pratice/user"
)

func main() {
	appUser := user.NewUserFromInput()
	admin := user.NewAdmin("test@gmail.com", "test@123890")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
	
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}
