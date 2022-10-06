package helper

import (
	"sort"

	dataStruct "gitlab.okymikhael.io/playground/parking-lot-golang/app/struct"
)

func SortingMapParkingLot(store map[string]dataStruct.ParkingLot) []string {
	keys := make([]string, 0, len(store))
	for k := range store {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
