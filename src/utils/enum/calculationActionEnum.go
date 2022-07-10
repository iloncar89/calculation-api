package enum

// Weekday - Custom type to hold value for CalculationAction ranging from 1-4
type CalculationAction int

// Declare related constants for each CalculationAction starting with index 1
const (
	Add      CalculationAction = iota + 1 // EnumIndex = 1
	Subtract                              // EnumIndex = 2
	Multiply                              // EnumIndex = 3
	Divide                                // EnumIndex = 4

)

// String - Creating common behavior - give the type a String function
func (a CalculationAction) String() string {
	return [...]string{"add", "subtract", "multiply", "divide"}[a-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (a CalculationAction) EnumIndex() int {
	return int(a)
}
