package router

import (
	"net/http"

	"gitlab.okymikhael.io/playground/parking-lot-golang/app/controller"
)

func Router() {
	parkingLotHandlers := controller.NewParkingLotHandlers()
	http.HandleFunc("/create_parking_lot/", parkingLotHandlers.StoreParkingLot)
	http.HandleFunc("/park/", parkingLotHandlers.SetSlotParking)
	http.HandleFunc("/leave/", parkingLotHandlers.EditSlotParking)
	http.HandleFunc("/status", parkingLotHandlers.GetStatus)
	http.HandleFunc("/cars_registration_numbers/", parkingLotHandlers.GetCarsRegisByColour)
	http.HandleFunc("/cars_slot/", parkingLotHandlers.GetSlotByColour)
	http.HandleFunc("/slot_number/", parkingLotHandlers.GetSlotByRegisNum)
	http.HandleFunc("/bulk", parkingLotHandlers.StoreBulkData)
	http.HandleFunc("/reset", parkingLotHandlers.ResetParkingLot)
}
