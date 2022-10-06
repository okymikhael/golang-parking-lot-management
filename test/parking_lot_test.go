package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func resetServer() {
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/reset", bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "Reset parking lot" {
		panic("Reset parking lot")
	}
}

func TestStoreParkingLot(t *testing.T) {
	lot := "6"

	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/create_parking_lot/"+lot, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "Created a parking lot with "+lot+" slots" {
		panic("Result is not 'Created a parking lot with " + lot + " slots'")
	}
}

func TestSetSlotParkingFirst(t *testing.T) {
	RegisNum := "B-1234-RFS"
	color := "Black"

	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/park/"+RegisNum+"/"+color, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "Allocated slot number: 1" {
		panic("Result is not 'Allocated slot number: 1'")
	}
}

func TestSetSlotParkingSec(t *testing.T) {
	RegisNum := "B-1111-RFS"
	color := "Black"

	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/park/"+RegisNum+"/"+color, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "Allocated slot number: 2" {
		panic("Result is not 'Allocated slot number: 2'")
	}
}

func TestEditSlotParking(t *testing.T) {
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/leave/1", bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "Slot number 1 is free" {
		panic("Result is not 'Slot number 1 is free'")
	}
}

func TestGetStatus(t *testing.T) {
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("GET", "http://localhost:8080/status", bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	expected := `Slot No. Registration No Colour
2 B-1111-RFS Black`
	if string(body) != expected {
		panic("Result is not '" + string(body) + "'")
	}
}

func TestGetCarsRegisByColour(t *testing.T) {
	color := "Black"
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("GET", "http://localhost:8080/cars_registration_numbers/colour/"+color, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "B-1111-RFS" {
		panic("Result is not 'B-1111-RFS'")
	}
}

func TestGetSlotByColour(t *testing.T) {
	color := "Black"
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("GET", "http://localhost:8080/cars_slot/colour/"+color, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "2" {
		panic("Result is not '2'")
	}
}

func TestGetSlotByRegisNumAvailable(t *testing.T) {
	regisNum := "B-1111-RFS"
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("GET", "http://localhost:8080/slot_number/car_registration_number/"+regisNum, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "2" {
		panic("Result is not '2'")
	}
}

func TestGetSlotByRegisNumNotFound(t *testing.T) {
	regisNum := "B-1233-RFS"
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("GET", "http://localhost:8080/slot_number/car_registration_number/"+regisNum, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "Not found" {
		panic("Result is not 'Not found'")
	}
	resetServer()
}
