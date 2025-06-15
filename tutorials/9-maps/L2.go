package main

import (
	"errors"
	"fmt"
)

func L2() {
	users := map[string]user2{
		"alex": {name: "alex", number: 103633, scheduledForDeletion: false},
		"john": {name: "john", number: 1243424, scheduledForDeletion: true},
		"peter": {name: "peter", number: 3424234, scheduledForDeletion: true},
		"martin": {name: "martin", number: 23424343 ,scheduledForDeletion: false},
	}

	for _, nm := range [5]string{"alex","john","martin", "peter", "arthur"} {
		test, err := deleteIfNecessary(users, nm)
		if err != nil {
			fmt.Printf("Error deleting %s message : %s \n", nm, err)
		}
		
		if test {
			fmt.Printf("message are deleted for %s \n", nm)
		} else {
			fmt.Printf("messages not deleted for %s \n", nm)
		}
	}
}
func deleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {

	elm, ok := users[name]
	if !ok {
		return false, errors.New("this user does not exist")
	}

	if elm.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}

	return false, nil

}

type user2 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
