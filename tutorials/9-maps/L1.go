package main

import (
	"errors"
	"fmt"
)

func L1(){
	names := []string{"alex","john","peter","martin"}
	phones := []int{0312321,1243424,3424234,23424343}

	maps,err := getUserMap(names,phones);
if err != nil{
	fmt.Printf("Execution Error : %s",err);
	return
}

fmt.Println(maps)


}
func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers){
		return nil,errors.New("invalid sizes")
	}
	maps := make(map[string] user)

	for i := 0; i < len(names); i++ {
		maps[fmt.Sprintf("user%d", i)]= user{
			name: names[i],
			phoneNumber: phoneNumbers[i],
		}
	}

	return maps,nil;
}

type user struct {
	name        string
	phoneNumber int
}
