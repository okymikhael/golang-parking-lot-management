package controller

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"gitlab.okymikhael.io/playground/parking-lot-golang/app/helper"
	dataStruct "gitlab.okymikhael.io/playground/parking-lot-golang/app/struct"
)

type parkingLotHandlers struct {
	sync.Mutex
	store map[string]dataStruct.ParkingLot
	lot   dataStruct.MaxParkingLot
}

func (h *parkingLotHandlers) ResetParkingLot(w http.ResponseWriter, r *http.Request) {
	methodAllowed := "POST"
	if r.Method != methodAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(([]byte("method not allowed")))
		return
	}

	h.Lock()
	if h.lot.Max != 0 {
		for _, k := range h.store {
			delete(h.store, k.ID)
		}
		h.lot.Max = 0
	}
	h.Unlock()

	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reset parking lot"))
}

func (h *parkingLotHandlers) StoreParkingLot(w http.ResponseWriter, r *http.Request) {
	methodAllowed := "POST"
	if r.Method != methodAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(([]byte("method not allowed")))
		return
	}

	// Check if path not less or more than 3
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check if parking lot maximum was initialized
	if h.lot.Max > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte("parking lot was initialized, please restart server")))
		return
	}

	delete(h.store, "dump") // Delete initialization map

	h.Lock()
	slot := path[2]
	slotMax, err := strconv.Atoi(slot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	h.lot.Max = slotMax
	h.Unlock()

	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Created a parking lot with " + slot + " slots"))
}

func (h *parkingLotHandlers) SetSlotParking(w http.ResponseWriter, r *http.Request) {
	methodAllowed := "POST"
	if r.Method != methodAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(([]byte("method not allowed")))
		return
	}

	// Check if path not less or more than 4
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check if parking lot maximum have value
	if h.lot.Max == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte("parking lot not initialized")))
		return
	}

	totalParkingCar := len(h.store) + 1
	parkingID := strconv.Itoa(totalParkingCar)
	registrationNumbers := path[2]
	color := helper.CapitalFirstLetterEachWord(path[3])
	keys := helper.SortingMapParkingLot(h.store)

	// Parse all parked car list
	parkingLots := "Slot No. Registration No Colour\n"
	for _, parkingLot := range keys {
		parkinglot := h.store[parkingLot]
		parkingLots += parkinglot.ID + " " + parkinglot.RegistrationNumbers + " " + parkinglot.Color + "\n"
	}

	i := 0
	prevID := 0
	for _, parkingLot := range keys {
		parkinglot := h.store[parkingLot]

		// Check if car was parked by registration number
		if parkinglot.RegistrationNumbers == registrationNumbers {
			w.Header().Add("content-type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write(([]byte("Allocated slot number: " + parkinglot.ID)))
			return
		}

		// Check if parking available in center slot
		if strconv.Itoa(prevID+1) != parkinglot.ID {
			parkingID = strconv.Itoa(prevID)
		}

		i++
		prevID++
	}

	// Check if parking lot full
	if h.lot.Max < totalParkingCar {
		w.Header().Add("content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(([]byte("Sorry, parking lot is full")))
		return
	}

	h.Lock()
	ParkingLot := dataStruct.ParkingLot{
		ID:                  parkingID,
		Color:               color,
		RegistrationNumbers: registrationNumbers,
	}

	h.store[parkingID] = ParkingLot
	defer h.Unlock()

	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Allocated slot number: " + parkingID))
}

func (h *parkingLotHandlers) EditSlotParking(w http.ResponseWriter, r *http.Request) {
	methodAllowed := "POST"
	if r.Method != methodAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(([]byte("method not allowed")))
		return
	}

	// Check if path not less or more than 3
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	slot := path[2]
	delete(h.store, slot)
	h.Unlock()

	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Slot number " + slot + " is free"))
}

func (h *parkingLotHandlers) GetStatus(w http.ResponseWriter, r *http.Request) {
	methodAllowed := "GET"
	if r.Method != methodAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(([]byte("method not allowed")))
		return
	}

	h.Lock()
	keys := helper.SortingMapParkingLot(h.store)

	parkingLots := []string{"Slot No. Registration No Colour"}
	for _, parkingLot := range keys {
		parkinglot := h.store[parkingLot]
		parkingLots = append(parkingLots, parkinglot.ID+" "+parkinglot.RegistrationNumbers+" "+parkinglot.Color[:len(parkinglot.Color)-1])
	}
	parkingLotsStr := strings.Join(parkingLots[:], "\n")
	h.Unlock()

	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(parkingLotsStr))
}

func (h *parkingLotHandlers) GetCarsRegisByColour(w http.ResponseWriter, r *http.Request) {
	methodAllowed := "GET"
	if r.Method != methodAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(([]byte("method not allowed")))
		return
	}

	// Check if path not less or more than 4
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	color := helper.CapitalFirstLetterEachWord(path[3])
	keys := helper.SortingMapParkingLot(h.store)

	// Parse all parked car list
	parkingLots := []string{}
	for _, parkingLot := range keys {
		parkinglot := h.store[parkingLot]
		if parkinglot.Color == color {
			parkingLots = append(parkingLots, parkinglot.RegistrationNumbers)
		}
	}
	parkingLotsStr := strings.Join(parkingLots[:], ", ")
	h.Unlock()

	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(parkingLotsStr))
}

func (h *parkingLotHandlers) GetSlotByColour(w http.ResponseWriter, r *http.Request) {
	methodAllowed := "GET"
	if r.Method != methodAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(([]byte("method not allowed")))
		return
	}

	// Check if path not less or more than 4
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	color := helper.CapitalFirstLetterEachWord(path[3])
	keys := helper.SortingMapParkingLot(h.store)

	// Parse all parked car list
	parkingLots := []string{}
	for _, parkingLot := range keys {
		parkinglot := h.store[parkingLot]
		if parkinglot.Color == color {
			parkingLots = append(parkingLots, parkinglot.ID)
		}
	}
	parkingLotsStr := strings.Join(parkingLots[:], ", ")
	h.Unlock()

	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(parkingLotsStr))
}

func (h *parkingLotHandlers) GetSlotByRegisNum(w http.ResponseWriter, r *http.Request) {
	methodAllowed := "GET"
	if r.Method != methodAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(([]byte("method not allowed")))
		return
	}

	// Check if path not less or more than 4
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	h.Lock()
	registrationNumbers := path[3]
	parkingLots := ""
	for _, parkingLot := range h.store {
		if parkingLot.RegistrationNumbers == registrationNumbers {
			parkingLots = parkingLot.ID
			break
		}
	}
	h.Unlock()

	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if parkingLots != "" {
		w.Write([]byte(parkingLots))
		return
	}
	w.Write([]byte("Not found"))
}

func (h *parkingLotHandlers) StoreBulkData(w http.ResponseWriter, r *http.Request) {
	methodAllowed := "POST"
	if r.Method != methodAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(([]byte("method not allowed")))
		return
	}

	// Read body form
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	host := "http://localhost:8080/"
	tempResponse := []string{}
	translator := map[string]string{ // Translate param to known path
		"registration_numbers_for_cars_with_colour": "cars_registration_numbers/colour/",
		"slot_numbers_for_cars_with_colour":         "cars_slot/colour/",
		"slot_number_for_registration_number":       "slot_number/car_registration_number/",
	}

	// Send request one by one
	i := 0
	addMissingMethod := []string{"POST", "POST", "POST", "POST", "POST", "POST", "POST", "POST", "GET", "POST", "POST", "GET", "GET", "GET", "GET"}
	bodyToList := strings.Split(string(bodyByte), "\n")
	for _, command := range bodyToList {
		if command != "" {
			getCommand := strings.Split(string(command), " ")
			value, ok := translator[getCommand[0]]
			if ok {
				command = value + getCommand[1]
			}
			request := host + strings.ReplaceAll(command, " ", "/")

			// Send Request same as needed method
			if addMissingMethod[i] == "POST" {
				var jsonStr = []byte(`{"dummy":"dummy data"}`)
				req, err := http.NewRequest("POST", request, bytes.NewBuffer(jsonStr))

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				tempResponse = append(tempResponse, string(body))
			} else {
				resp, err := http.Get(request)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				tempResponse = append(tempResponse, string(body))
			}
			i++
		}
	}

	tempResponseStr := strings.Join(tempResponse[:], "\n") + "\n"
	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tempResponseStr))
}

func NewParkingLotHandlers() *parkingLotHandlers {
	return &parkingLotHandlers{
		store: map[string]dataStruct.ParkingLot{
			"dump": {
				ID:                  "dump",
				Color:               "dump",
				RegistrationNumbers: "dump",
			},
		},
	}
}
