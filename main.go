package main

import (
	"encoding/json"
	"log"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	ID uint64
}

func main() {

	overFlowError()
	noOverflow()

}

//CreateUserRequest parameters
type UserType struct {
	Email string `json:"email"`
	ID    int64  `json:"id"`
}

func overFlowError() {

	var id int64 = 512281913614499841

	user := UserType{
		Email: "example",
		ID:    id,
	}

	message := map[string]interface{}{
		"data": user,
	}

	//mashal to byte to simulate service payload
	servicePayload, err := json.Marshal(message)

	if err != nil {
		log.Println(err)
	}

	var receivedMessage map[string]interface{}

	json.Unmarshal(servicePayload, &receivedMessage)

	var myUser UserType

	mapstructure.Decode(receivedMessage["data"], &myUser)

	log.Println("---receivedMessage:", myUser.ID) //prints 512281913614499840 - incorrect
}

func noOverflow() {

	var id int64 = 512281913614499841

	message := map[string]UserType{}

	message["data"] = UserType{
		Email: "yusuf",
		ID:    id,
	}

	byteMessage, err := json.Marshal(message)

	if err != nil {
		log.Println(err)
	}

	var msgMap map[string]UserType

	json.Unmarshal(byteMessage, &msgMap)

	log.Println("---myUser:", msgMap["data"].ID) // prints 512281913614499841 - correct
}
