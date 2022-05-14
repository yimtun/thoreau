package auth
//

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetAccount ()  (map[string]string) {
	var account = make(map[string]string)
	test:=getConfig()
	for _,Resource :=range test.Resources{
		account[Resource.UserName]=Resource.PassWord
	}
	fmt.Println(account)
	return  account
}


type Resource struct {
	UserName string `json:"user_name"`
	PassWord       string `json:"pass_word"`
    UserId string `json:"user_id"`
}


type Resources struct {
	Resources []Resource `json:"resources"`
}

func getConfig()  Resources{
	var tmpConfig Resources
	filePtr, err := os.Open("a.json")
	if err != nil {
		panic(err)
	}

	var resources Resources
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&resources)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())

	} else {
		tmpConfig  = resources
	}
	return tmpConfig
}
