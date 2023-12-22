package purchase

import (
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	stringSlice := [2]string{option1, option2}
	sort.Strings(stringSlice[:])
	return stringSlice[0] + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3.0:
		return originalPrice * 0.80
	case age >= 10:
		return originalPrice * 0.50
	default:
		return originalPrice * 0.70
	}
}
