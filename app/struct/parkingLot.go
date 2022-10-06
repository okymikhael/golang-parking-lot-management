package dataStruct

type MaxParkingLot struct {
	Max int
}

type ParkingLot struct {
	ID                  string `json:"id"`
	Color               string `json:"Color"`
	RegistrationNumbers string `json:"RegistrationNumbers"`
}