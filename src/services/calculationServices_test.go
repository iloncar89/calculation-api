package services

import (
	"github.com/iloncar89/calculation-api/src/dto"
	"github.com/iloncar89/calculation-api/src/utils/constants/path/calculation"
	"github.com/iloncar89/calculation-api/src/utils/restErrors"
	"testing"
)

func TestAdd_Success_IntInt(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "3",
		Y:      "1",
		Action: calculation.AddCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "add",
		X:      3,
		Y:      1,
		Answer: 4,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Add with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Add with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestAdd_Success_FloatFloat(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "1.1",
		Y:      "1.2",
		Action: calculation.AddCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "add",
		X:      1.1,
		Y:      1.2,
		Answer: 2.3,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Add with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Add with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestAdd_Success_IntFloat(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "1",
		Y:      "1.05",
		Action: calculation.AddCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "add",
		X:      1,
		Y:      1.05,
		Answer: 2.05,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Add with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Add with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestAdd_Fail_IntFloat_LeadingZeros(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "1",
		Y:      "01.05",
		Action: calculation.AddCalculationUrlPath,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if calc != nil {
			if restErr.Message() != restErrors.ErrorLeadingZeros {
				t.Errorf("Add with arguments x = %s, y = %s did not returned rest error message %s, instead we got calculation %v and rest error %v", calculationRequest.X, calculationRequest.Y, restErrors.ErrorLeadingZeros, calc, restErr)
			}
		}
	}
}

func TestAddParameter_NotNumber_ParameterX(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "10.0s0",
		Y:      "5",
		Action: calculation.AddCalculationUrlPath,
	}

	expectedError := restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)

	_, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil && restErr.Message() != expectedError.Message() {
		t.Errorf("Add with arguments x = %s, y = %s fails, expected  rest error  %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedError, restErr)
	}
}

func TestAddParameter_NotNumber_ParameterY(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "10.00",
		Y:      "string",
		Action: calculation.AddCalculationUrlPath,
	}

	expectedError := restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)

	_, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil && restErr.Message() != expectedError.Message() {
		t.Errorf("Add with arguments x = %s, y = %s fails, expected  rest error  %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedError, restErr)
	}
}

func TestSubtract_Success_IntInt(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "1",
		Y:      "1",
		Action: calculation.SubtractCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "subtract",
		X:      1,
		Y:      1,
		Answer: 0,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Subtract with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Subtract with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestSubtract_Success_FloatFloat(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "5.5",
		Y:      "1.1",
		Action: calculation.SubtractCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "subtract",
		X:      5.5,
		Y:      1.1,
		Answer: 4.4,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Subtract with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Subtract with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestSubtract_Success_IntFloat(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "5.55",
		Y:      "1",
		Action: calculation.SubtractCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "subtract",
		X:      5.55,
		Y:      1,
		Answer: 4.55,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Subtract with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Subtract with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestSubtract_Fail_IntFloat_LeadingZeros(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "01",
		Y:      "0.05",
		Action: calculation.SubtractCalculationUrlPath,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if calc != nil {
			if restErr.Message() != restErrors.ErrorLeadingZeros {
				t.Errorf("Subtract with arguments x = %s, y = %s did not returned rest error message %s, instead we got calculation %v and rest error %v", calculationRequest.X, calculationRequest.Y, restErrors.ErrorLeadingZeros, calc, restErr)
			}
		}
	}
}

func TestSubtract_ParameterNotNumber_ParameterX(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "string",
		Y:      "9",
		Action: calculation.SubtractCalculationUrlPath,
	}

	expectedError := restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)

	_, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if restErr.Message() != expectedError.Message() {
			t.Errorf("Subtract with arguments x = %s , y = %s fails, expected  rest error  %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedError, restErr)
		}
	}
}

func TestSubtract_ParameterNotNumber_ParameterY(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "5.2",
		Y:      "test5",
		Action: calculation.SubtractCalculationUrlPath,
	}

	expectedError := restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)

	_, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if restErr.Message() != expectedError.Message() {
			t.Errorf("Subtract with arguments x = %s , y = %s fails, expected  rest error  %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedError, restErr)
		}
	}
}

func TestMultiply_Success_IntInt(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "1",
		Y:      "1",
		Action: calculation.MultiplyCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "multiply",
		X:      1,
		Y:      1,
		Answer: 1,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Multiply with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Mubtract with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestMultiply_Success_IntFloat(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "1",
		Y:      "1.05",
		Action: calculation.MultiplyCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "multiply",
		X:      1,
		Y:      1.05,
		Answer: 1.05,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Multiply with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Mubtract with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestMultiply_Success_FloatFloat(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "8.5",
		Y:      "4.2",
		Action: calculation.MultiplyCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "multiply",
		X:      8.5,
		Y:      4.2,
		Answer: 35.7,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Multiply with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Mubtract with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestMultiply_Fail_IntFloat_LeadingZeros(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "001",
		Y:      "0.05",
		Action: calculation.MultiplyCalculationUrlPath,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if calc != nil {
			if restErr.Message() != restErrors.ErrorLeadingZeros {
				t.Errorf("Multiply with arguments x = %s, y = %s did not returned rest error message %s, instead we got calculation %v and rest error %v", calculationRequest.X, calculationRequest.Y, restErrors.ErrorLeadingZeros, calc, restErr)
			}
		}
	}
}

func TestMultiply_ParameterNotNumber_ParameterX(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "kk03",
		Y:      "-9",
		Action: calculation.MultiplyCalculationUrlPath,
	}

	expectedError := restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)

	_, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if restErr.Message() != expectedError.Message() {
			t.Errorf("Multiply with arguments x = %s, y = %s fails, expected  rest error  %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedError, restErr)
		}
	}
}

func TestMultiply_ParameterNotNumber_ParameterY(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "3",
		Y:      "-9test",
		Action: calculation.MultiplyCalculationUrlPath,
	}

	expectedError := restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)

	_, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if restErr.Message() != expectedError.Message() {
			t.Errorf("Multiply with arguments x = %s, y = %s fails, expected  rest error  %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedError, restErr)
		}
	}
}

func TestDivide_Success_IntInt(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "10",
		Y:      "1",
		Action: calculation.DivideCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "divide",
		X:      10,
		Y:      1,
		Answer: 10,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Divide with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Divide with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestDivide_Success_IntFloat(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "10",
		Y:      "2.5",
		Action: calculation.DivideCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "divide",
		X:      10,
		Y:      2.5,
		Answer: 4,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Divide with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Divide with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestDivide_Success_FloatFloat(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "6.82",
		Y:      "3.1",
		Action: calculation.DivideCalculationUrlPath,
	}

	expectedCalculation := dto.CalculationDto{
		Action: "divide",
		X:      6.82,
		Y:      3.1,
		Answer: 2.2,
		Cached: false,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		t.Errorf("Divide with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
	if calc.Cached != expectedCalculation.Cached || calc.X != expectedCalculation.X || calc.Y != expectedCalculation.Y || calc.Action != expectedCalculation.Action || calc.Answer != expectedCalculation.Answer {
		t.Errorf("Divide with arguments x = %s, y = %s fails, expected %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedCalculation, calc)
	}
}

func TestDivide_Fail_IntFloat_LeadingZeros(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "1",
		Y:      "-00.05",
		Action: calculation.DivideCalculationUrlPath,
	}

	calc, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if calc != nil {
			if restErr.Message() != restErrors.ErrorLeadingZeros {
				t.Errorf("Divide with arguments x = %s, y = %s did not returned rest error message %s, instead we got calculation %v and rest error %v", calculationRequest.X, calculationRequest.Y, restErrors.ErrorLeadingZeros, calc, restErr)
			}
		}
	}
}

func TestDivide_WithZero(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "10",
		Y:      "0",
		Action: calculation.DivideCalculationUrlPath,
	}

	_, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr == nil {
		t.Errorf("Divide with arguments x = %s, y = %s fails, expected empty rest error, got %+v", calculationRequest.X, calculationRequest.Y, restErr)
	}
}

func TestDivide_ParameterNotNumber_ParameterX(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "ssa05ttt",
		Y:      "11",
		Action: calculation.DivideCalculationUrlPath,
	}

	expectedError := restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)

	_, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if restErr.Message() != expectedError.Message() {
			t.Errorf("Divide with arguments x = %s, y = %s - fails, expected  rest error  %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedError, restErr)
		}
	}
}

func TestDivide_ParameterNotNumber_ParameterY(t *testing.T) {
	calculationRequest := &dto.CalculationRequestDto{
		X:      "77",
		Y:      "---5",
		Action: calculation.DivideCalculationUrlPath,
	}

	expectedError := restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)

	_, restErr := CalculationService.BasicMathOperation(*calculationRequest)

	if restErr != nil {
		if restErr.Message() != expectedError.Message() {
			t.Errorf("Divide with arguments x = %s, y = %s - fails, expected  rest error  %+v, got %+v", calculationRequest.X, calculationRequest.Y, expectedError, restErr)
		}
	}
}
