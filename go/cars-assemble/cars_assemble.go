package cars
import "fmt"
// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    var r: float64 = float64(float64(productionRate) * float64(successRate))
    fmt.Print(r)
	return float64(float64(productionRate) * float64(successRate))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(successRate) * productionRate
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	panic("CalculateCost not implemented")
}
