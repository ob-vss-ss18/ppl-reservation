package main

import (
	"testing"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func TestTest(t *testing.T) {
	myPrint("test")
}

func TestTestTest(t *testing.T) {

	type Reservation struct {
		ID     int
		CID    int
		ITEMID int
	}

	var res Reservation
	var res2 Reservation

	var client http.Client
	resp, err := client.Get("https://ppl-reservation-beta.herokuapp.com/reservations?query={reservations(cId:2){id,cId,itemId}}")
	if err != nil {
		// err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	println(resp.StatusCode)
	fmt.Printf("\n")

	json.Unmarshal(bodyBytes, &res)

	println(res.ID)
	println(res.CID)
	println(res.ITEMID)
	fmt.Printf("\n")


	res2.CID = 1
	res2.ID = 3
	res2.ITEMID = 14

	data, err := json.Marshal(res2)

	data2 := string(data)
	println(data2)


}
