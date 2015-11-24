package group

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	GroupId   int
	GroupName string

	Keys []string
}

type Groups struct {
	Group []Group
}

func (group *Groups) LoadGroups(file string) bool {

	return true
}

func (group *Group) SaveGroup(file string) bool {

	return true
}

func (group *Group) AddKey(key string) bool {

	group.Keys = append(group.Keys, key)
	return true
}

func (group *Group) Info() {
	b, _ := json.Marshal(group)
	fmt.Println(string(b))
}
