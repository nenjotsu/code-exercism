package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) (total float64) {
	total = float64(productionRate) * successRate / 100
	return
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) (total int) {
	perHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	total = int(perHour) / 60
	return
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) (total uint) {
	base := carsCount / 10
	remainder := carsCount % 10
	total = uint(base*95000 + remainder*10000)
	return
}
