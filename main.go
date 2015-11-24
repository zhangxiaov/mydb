package main

import (
	"mydb/group"
)

func main() {
	group := group.Group{1, "group1", nil}

	group.AddKey("userId")
	group.AddKey("userName")
	group.AddKey("age")

	group.Info()
}
