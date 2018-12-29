package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	api := &Api{
		url:    "http://localhost:4000",
		client: client,
	}

	userData, err := ParseCmd()
	if err != nil {
		panic(err)
	}

	if userData.Command == cmdList {
		handleList(api, userData.ListData)
	} else if userData.Command == cmdCreate {
		handleCreate(api, userData.CreateData)
	}
}

func handleList(api *Api, listData *ListData) {
	if listData.EventID == "" {
		events, err := api.GetAllEvents()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, event := range events {
			fmt.Println(event)
		}

	} else {
		event, err := api.GetEvent(listData.EventID)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(event)
	}
}

func handleCreate(api *Api, createData *CreateData) {

}
