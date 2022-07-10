package parser

import (
	"github.com/iloncar89/calculation-api/src/dto"
	"github.com/iloncar89/calculation-api/src/utils/restErrors"
	"net/http"
	"testing"
)

func TestCalculationParser_ParseCalculation_ZeroFloat(t *testing.T) {
	expectedRes := dto.CalculationRequestDto{
		X: "0",
		Y: "7.1",
	}
	var url string
	url = "http://localhost:8080/multiply?&x=0&y=7.1"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_ZeroFloat test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if res != nil {

			if res.X != expectedRes.X || res.Y != expectedRes.Y || restErr != nil {
				t.Errorf("For url %s, got in response %v and rest error %v, but expected result is %v", url, res, restErr, expectedRes)
			}
		}
	}
}

func TestCalculationParser_ParseCalculation_FloatFloat(t *testing.T) {
	expectedRes := dto.CalculationRequestDto{
		X: "0.88888",
		Y: "-7.1",
	}
	var url string
	url = "http://localhost:8080/multiply?&x=0.88888&y=-7.1"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_FloatFloat test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if res != nil {
			if res.X != expectedRes.X || res.Y != expectedRes.Y || restErr != nil {
				t.Errorf("For url %s, got in response %v and rest error %v, but expected result is %v", url, res, restErr, expectedRes)
			}
		}
	}
}

func TestCalculationParser_ParseCalculation_IntInt(t *testing.T) {
	expectedRes := dto.CalculationRequestDto{
		X: "1",
		Y: "-1",
	}
	var url string
	url = "http://localhost:8080/multiply?&x=1&y=-1"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_IntInt test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if res != nil {
			if res.X != expectedRes.X || res.Y != expectedRes.Y || restErr != nil {
				t.Errorf("For url %s, got in response %v and rest error %v, but expected result is %v", url, res, restErr, expectedRes)
			}
		}
	}
}

func TestCalculationParser_ParseCalculation_IntFloat(t *testing.T) {
	expectedRes := dto.CalculationRequestDto{
		X: "1.558",
		Y: "-1",
	}
	var url string
	url = "http://localhost:8080/multiply?&x=1.558&y=-1"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_IntFloat test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if res != nil {
			if res.X != expectedRes.X || res.Y != expectedRes.Y || restErr != nil {
				t.Errorf("For url %s, got in response %v and rest error %v, but expected result is %v", url, res, restErr, expectedRes)
			}
		}
	}
}

func TestCalculationParser_ParseCalculation_StringFloat(t *testing.T) {
	expectedRes := dto.CalculationRequestDto{
		X: "string",
		Y: "-1.5",
	}
	var url string
	url = "http://localhost:8080/multiply?&x=string&y=-1.5"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_StringFloat test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if res != nil {
			if res.X != expectedRes.X || res.Y != expectedRes.Y || restErr != nil {
				t.Errorf("For url %s, got in response %v and rest error %v, but expected result is %v", url, res, restErr, expectedRes)
			}
		}
	}
}

func TestCalculationParser_ParseCalculation_NoQueryParams(t *testing.T) {
	var url string
	url = "http://localhost:8080/multiply"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_NoQueryParams test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if restErr != nil && restErr.Message() != restErrors.MissingBothCalculationParamsErrorMessage {
			t.Errorf("For url %s, should get error message %s but got rest error %v and response %v", url, restErrors.MissingBothCalculationParamsErrorMessage, restErr, res)
		}
	}
}

func TestCalculationParser_ParseCalculation_NoQueryParamX(t *testing.T) {
	var url string
	url = "http://localhost:8080/multiply?&y=-1.5"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_NoQueryParamX test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if restErr != nil && restErr.Message() != restErrors.MissingParameterXErrorMessage {
			t.Errorf("For url %s, should get error message %s but got rest error %v and response %v", url, restErrors.MissingParameterXErrorMessage, restErr, res)
		}
	}
}

func TestCalculationParser_ParseCalculation_NoQueryParamY(t *testing.T) {
	var url string
	url = "http://localhost:8080/multiply?&x=5"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_NoQueryParamY test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if restErr != nil && restErr.Message() != restErrors.MissingParameterYErrorMessage {
			t.Errorf("For url %s, should get error message %s but got rest error %v and response %v", url, restErrors.MissingParameterYErrorMessage, restErr, res)
		}
	}
}

func TestCalculationParser_ParseCalculation_MoreThanOneBothQueryParams(t *testing.T) {
	var url string
	url = "http://localhost:8080/multiply?&x=5&x=4&y=-5&y=7.8"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_MoreThanOneBothQueryParams test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if restErr != nil && restErr.Message() != restErrors.FoundMoreThanOneBothParamsErrorMessage {
			t.Errorf("For url %s, should get error message %s but got rest error %v and response %v", url, restErrors.FoundMoreThanOneBothParamsErrorMessage, restErr, res)
		}
	}
}

func TestCalculationParser_ParseCalculation_MoreThanOneXQueryParams(t *testing.T) {
	var url string
	url = "http://localhost:8080/multiply?&x=5&x=-5&y=7.8"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_MoreThanOneXQueryParams test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if restErr != nil && restErr.Message() != restErrors.FoundMoreThanOneXParamsErrorMessage {
			t.Errorf("For url %s, should get error message %s but got rest error %v and response %v", url, restErrors.FoundMoreThanOneXParamsErrorMessage, restErr, res)
		}
	}
}

func TestCalculationParser_ParseCalculation_MoreThanOneYQueryParams(t *testing.T) {
	var url string
	url = "http://localhost:8080/multiply?&x=5&y=-5&y=7.8"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Errorf("Internal server error while creating new request in TestCalculationParser_ParseCalculation_MoreThanOneYQueryParams test")
	}

	if request != nil {
		res, restErr := CalculationParser.ParseCalculationPathVariables(*request)

		if restErr != nil && restErr.Message() != restErrors.FoundMoreThanOneYParamsErrorMessage {
			t.Errorf("For url %s, should get error message %s but got rest error %v and response %v", url, restErrors.FoundMoreThanOneYParamsErrorMessage, restErr, res)
		}
	}
}

func TestCalculationParser_GetFirstChar_NormalWord(t *testing.T) {
	test := "string"
	expected := "s"

	result := CalculationParser.GetFirstChar(test)

	if result != expected {
		t.Errorf("Tested string %s with TestCalculationParser_GetFirstChar_NormalWord, expected %s, got %s", test, expected, result)
	}
}

func TestCalculationParser_GetFirstChar_EmptyString(t *testing.T) {
	test := ""
	expected := ""

	result := CalculationParser.GetFirstChar(test)

	if result != expected {
		t.Errorf("Tested string %s with TestCalculationParser_GetFirstChar_EmptyString, expected %s, got %s", test, expected, result)
	}
}

func TestCalculationParser_GetFirstChar_OnlySpace(t *testing.T) {
	test := " "
	expected := " "

	result := CalculationParser.GetFirstChar(test)

	if result != expected {
		t.Errorf("Tested string %s with TestCalculationParser_GetFirstChar_OnlySpace, expected %s, got %s", test, expected, result)
	}
}

func TestCalculationParser_TrimLeftChar_NormalWord(t *testing.T) {
	test := "string"
	expected := "tring"

	result := CalculationParser.TrimLeftChar(test)

	if result != expected {
		t.Errorf("Tested string %s with TestCalculationParser_TrimLeftChar_NormalWord, expected %s, got %s", test, expected, result)
	}
}

func TestCalculationParser_TrimLeftChar_EmptyString(t *testing.T) {
	test := ""
	expected := ""

	result := CalculationParser.TrimLeftChar(test)

	if result != expected {
		t.Errorf("Tested string %s with TestCalculationParser_TrimLeftChar_EmptyString, expected %s, got %s", test, expected, result)
	}
}

func TestCalculationParser_TrimLeftChar_OnlySpace(t *testing.T) {
	test := " "
	expected := ""

	result := CalculationParser.TrimLeftChar(test)

	if result != expected {
		t.Errorf("Tested string %s with TestCalculationParser_TrimLeftChar_OnlySpace, expected %s, got %s", test, expected, result)
	}
}
