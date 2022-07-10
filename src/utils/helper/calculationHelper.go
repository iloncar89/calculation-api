package helper

import (
	"github.com/iloncar89/calculation-api/src/utils/enum"
	"github.com/iloncar89/calculation-api/src/utils/logger"
	"github.com/iloncar89/calculation-api/src/utils/parser"
	"strconv"
	"strings"
)

var CalculationHelper calculationHelperInterface = &calculationHelper{}

type calculationHelperInterface interface {
	ConvertStringToNumber(variable string) (float64, error)
	IsZeroValue(number float64) bool
	IsVariableNotFormattedCorrectlyWithLeadingZero(variable string) bool
	RemovePlusOperatorFromString(variable string) string
	PerformBasicMathOperation(x, y float64, action string) float64
}
type calculationHelper struct{}

//IsVariableNotFormattedCorrectlyWithLeadingZeros function is used for validation to check if string variable is starting with -00... or 00... which is not correct, returns correct boolean
func (ch *calculationHelper) IsVariableNotFormattedCorrectlyWithLeadingZero(variable string) bool {
	if len(variable) > 1 {
		firstChar := parser.CalculationParser.GetFirstChar(variable)
		if firstChar == "-" && len(variable) > 2 {
			variableWithoutPrefix := parser.CalculationParser.TrimLeftChar(variable)
			secondChar := parser.CalculationParser.GetFirstChar(variableWithoutPrefix)
			if secondChar == "0" {
				variableWithoutPrefixOperatorAndZero := parser.CalculationParser.TrimLeftChar(variableWithoutPrefix)
				thirdChar := parser.CalculationParser.GetFirstChar(variableWithoutPrefixOperatorAndZero)
				if thirdChar != "." {
					logger.InfoLogger.Print("Variable, ", variable, " contains -0 at beginning of string")
					return true
				}
			}
		}
		if firstChar == "0" && len(variable) > 1 {
			variableWithoutLeadingZero := parser.CalculationParser.TrimLeftChar(variable)
			secondChar := parser.CalculationParser.GetFirstChar(variableWithoutLeadingZero)
			if secondChar != "." && strings.Contains(variable, ".") {
				logger.InfoLogger.Print("Variable, ", variable, " contains 0 at beginning of string")
				return true
			} else if secondChar == "0" && !strings.Contains(variable, ".") && len(variable) > 2 {
				logger.InfoLogger.Print("Variable, ", variable, " contains 0 at beginning of string")
				return true
			} else if secondChar != "0" && !strings.Contains(variable, ".") && len(variable) == 2 {
				logger.InfoLogger.Print("Variable, ", variable, " contains 0 at beginning of string")
				return true
			}
		}
	}
	return false
}

//ConvertStringToNumber is used to convert string to number, returns number as interface if succeed and error if not.
func (ch *calculationHelper) ConvertStringToNumber(variable string) (float64, error) {
	number, err := strconv.ParseFloat(variable, 64)
	if err != nil {
		logger.ErrorLogger.Print("Error occur while executing ParseInt and ParseFloat. Variable ", variable, " is not an integer or float number")
		return 0, err
	}
	return number, nil
}

//IsZeroValue returns true if float64 is zero
func (ch *calculationHelper) IsZeroValue(number float64) bool {
	return number == 0
}
func (ch *calculationHelper) RemovePlusOperatorFromString(variable string) string {
	if len(variable) > 1 {
		firstChar := parser.CalculationParser.GetFirstChar(variable)
		if firstChar == "+" || firstChar == " " {
			variable = parser.CalculationParser.TrimLeftChar(variable)
		}
		return variable
	}
	return variable
}

func (ch *calculationHelper) PerformBasicMathOperation(x, y float64, action string) float64 {
	var result float64
	switch action {
	case enum.Add.String():
		result = x + y
	case enum.Subtract.String():
		result = x - y
	case enum.Multiply.String():
		result = x * y
	case enum.Divide.String():
		result = x / y
	}
	return result
}
