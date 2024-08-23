package cars

const priceForTenCars = 95_000
const priceForOneCar = 10_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var rateInFloat float64 = float64(productionRate)

	return rateInFloat * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var carsPerHour = CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(carsPerHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupsOfTenCars = carsCount / 10
	var individualCars = carsCount % 10
	var result = (priceForTenCars * groupsOfTenCars) + (priceForOneCar * individualCars)

	return uint(result)
}
