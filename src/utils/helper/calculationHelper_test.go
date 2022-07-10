package helper

import "testing"

func TestCalculationHelper_ConvertStringToNumber_Float(t *testing.T) {
	var number string
	number = "-64.5555"
	var expected float64
	expected = -64.5555
	res, err := CalculationHelper.ConvertStringToNumber(number)
	if res != expected && err != nil {
		t.Errorf("Conversion string type %s to number type float64 interface was not successful, got response %f,  error %v, but expected is no err and response %f", number, res, err, expected)
	}
}

func TestCalculationHelper_ConvertStringToNumber_FloatWithPlusOperator(t *testing.T) {
	var number string
	number = "+64.5555"
	var expected float64
	expected = 64.5555
	res, err := CalculationHelper.ConvertStringToNumber(number)
	if res != expected && err != nil {
		t.Errorf("Conversion string type %s to number type float64 interface was not successful, got response %f,  error %v, but expected is no err and response %f", number, res, err, expected)
	}
}

func TestCalculationHelper_ConvertStringToNumber_ErrorNotDigits(t *testing.T) {
	var number string
	number = "test"
	res, err := CalculationHelper.ConvertStringToNumber(number)
	if res != 0 && err == nil {
		t.Errorf("Conversion of %s to number was succesfull and run without errors", number)
	}
}

func TestCalculationHelper_IsVariableNotFormattedCorrectlyWithLeadingZero_TwoZerosWithoutDotTrue(t *testing.T) {
	var number string
	number = "005"
	res := CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(number)
	if !res {
		t.Errorf("IsVariableNotFormattedCorrectlyWithLeadingZero returned false in response, but expected true for argument %s", number)
	}
}

func TestCalculationHelper_IsVariableNotFormattedCorrectlyWithLeadingZero_OneZeroWithoutDotTrue(t *testing.T) {
	var number string
	number = "05"
	res := CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(number)
	if !res {
		t.Errorf("IsVariableNotFormattedCorrectlyWithLeadingZero returned false in response, but expected true for argument %s", number)
	}
}

func TestCalculationHelper_IsVariableNotFormattedCorrectlyWithLeadingZero_OneZeroWithDotTrue(t *testing.T) {
	var number string
	number = "05.05"
	res := CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(number)
	if !res {
		t.Errorf("IsVariableNotFormattedCorrectlyWithLeadingZero returned false in response, but expected true for argument %s", number)
	}
}

func TestCalculationHelper_IsVariableNotFormattedCorrectlyWithLeadingZero_TwoZeroWithDotTrue(t *testing.T) {
	var number string
	number = "005.05"
	res := CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(number)
	if !res {
		t.Errorf("IsVariableNotFormattedCorrectlyWithLeadingZero returned false in response, but expected true for argument %s", number)
	}
}

func TestCalculationHelper_IsVariableNotFormattedCorrectlyWithLeadingZero_MinusOperatorTwoZeroWithDotTrue(t *testing.T) {
	var number string
	number = "-005.05"
	res := CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(number)
	if !res {
		t.Errorf("IsVariableNotFormattedCorrectlyWithLeadingZero returned false in response, but expected true for argument %s", number)
	}
}

func TestCalculationHelper_IsVariableNotFormattedCorrectlyWithLeadingZero_MinusOperatorOneZeroWithDotTrue(t *testing.T) {
	var number string
	number = "-05.05"
	res := CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(number)
	if !res {
		t.Errorf("IsVariableNotFormattedCorrectlyWithLeadingZero returned false in response, but expected true for argument %s", number)
	}
}

func TestCalculationHelper_IsVariableNotFormattedCorrectlyWithLeadingZero_MinusOperatorOneZeroWithoutDotTrue(t *testing.T) {
	var number string
	number = "-0505"
	res := CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(number)
	if !res {
		t.Errorf("IsVariableNotFormattedCorrectlyWithLeadingZero returned false in response, but expected true for argument %s", number)
	}
}

func TestCalculationHelper_IsVariableNotFormattedCorrectlyWithLeadingZero_MinusOperatorTwoZeroWithoutDotTrue(t *testing.T) {
	var number string
	number = "-0505"
	res := CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(number)
	if !res {
		t.Errorf("IsVariableNotFormattedCorrectlyWithLeadingZero returned false in response, but expected true for argument %s", number)
	}
}

func TestCalculationHelper_IsVariableNotFormattedCorrectlyWithLeadingZero_IntegerFalse(t *testing.T) {
	var number string
	number = "5"
	res := CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(number)
	if res {
		t.Errorf("IsVariableNotFormattedCorrectlyWithLeadingZero returned true in response, but expected false for argument %s", number)
	}
}

func TestCalculationHelper_IsZeroValue_ZeroTrue(t *testing.T) {
	var number float64
	number = 0
	res := CalculationHelper.IsZeroValue(number)

	if !res {
		t.Errorf("For number, %f, got %t when calling IsZeroValue", number, res)
	}
}

func TestCalculationHelper_RemovePlusOperatorFromString_WhichContainsPlusOperator(t *testing.T) {
	testString := "+5852"
	expectedString := "5852"
	res := CalculationHelper.RemovePlusOperatorFromString(testString)
	if res != expectedString {
		t.Errorf("Trying to test TestCalculationHelper_RemovePlusOperatorFromString_WhichContainsPlusOperator, for inputed %s, got %s insted of %s", testString, res, expectedString)
	}
}

func TestCalculationHelper_RemovePlusOperatorFromString_WhichContainsMinusOperator(t *testing.T) {
	testString := "-5852"
	res := CalculationHelper.RemovePlusOperatorFromString(testString)
	if res != testString {
		t.Errorf("Tested string with TestCalculationHelper_RemovePlusOperatorFromString_WhichContainsMinusOperator %s is not same in resoult %s", testString, res)
	}
}

func TestCalculationHelper_RemovePlusOperatorFromString_WhichDoesNotContainsPlusOperator(t *testing.T) {
	testString := "5852"
	res := CalculationHelper.RemovePlusOperatorFromString(testString)
	if res != testString {
		t.Errorf("Tested string with TestCalculationHelper_RemovePlusOperatorFromString_WhichDoesNotContainsPlusOperator %s is not same in resoult %s", testString, res)
	}
}

func TestCalculationHelper_RemovePlusOperatorFrom_String_Empty(t *testing.T) {
	testString := ""
	res := CalculationHelper.RemovePlusOperatorFromString(testString)
	if res != testString {
		t.Errorf("Tested string with TestCalculationHelper_RemovePlusOperatorFrom_String_Empty %s is not same in resoult %s", testString, res)
	}
}
