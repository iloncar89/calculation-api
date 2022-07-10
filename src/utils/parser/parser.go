package parser

import (
	"github.com/iloncar89/calculation-api/src/dto"
	"github.com/iloncar89/calculation-api/src/utils/logger"
	"github.com/iloncar89/calculation-api/src/utils/restErrors"
	"net/http"
)

var CalculationParser calculationParserInterface = &calculationParser{}

type calculationParserInterface interface {
	ParseCalculationPathVariables(r http.Request) (*dto.CalculationRequestDto, restErrors.RestErr)
	TrimLeftChar(s string) string
	GetFirstChar(s string) string
}

type calculationParser struct{}

//ParseCalculationPathVariables function is used for parsing constants variables from GET or POST call.
//Receives http.Request as argument and returns calculationRequest if everything succeed or restError if not.
func (parser *calculationParser) ParseCalculationPathVariables(r http.Request) (*dto.CalculationRequestDto, restErrors.RestErr) {
	pathVariables := r.URL.Query()
	xParams, xOk := pathVariables["x"]
	yParams, yOk := pathVariables["y"]
	if len(xParams) == 0 && len(yParams) == 0 {
		restErr := restErrors.NewBadRequestError(restErrors.MissingBothCalculationParamsErrorMessage)
		logger.ErrorLogger.Print("Error while performing ParseCalculationPathVariables ", restErr)
		return nil, restErr
	}
	if len(xParams) > 1 && len(yParams) > 1 {
		restErr := restErrors.NewBadRequestError(restErrors.FoundMoreThanOneBothParamsErrorMessage)
		logger.ErrorLogger.Print("Error while performing ParseCalculationPathVariables ", restErr)
		return nil, restErr
	}
	var x string
	var y string
	if xOk {
		if len(xParams) == 1 {
			x = xParams[0]
		} else {
			restErr := restErrors.NewBadRequestError(restErrors.FoundMoreThanOneXParamsErrorMessage)
			logger.ErrorLogger.Print("Error while performing ParseCalculationPathVariables ", restErr)
			return nil, restErr
		}
	} else {
		restErr := restErrors.NewBadRequestError(restErrors.MissingParameterXErrorMessage)
		logger.ErrorLogger.Print("Error while performing ParseCalculationPathVariables ", restErr)
		return nil, restErr
	}
	if yOk {
		if len(yParams) == 1 {
			y = yParams[0]
		} else {
			restErr := restErrors.NewBadRequestError(restErrors.FoundMoreThanOneYParamsErrorMessage)
			logger.ErrorLogger.Print("Error while performing ParseCalculationPathVariables ", restErr)
			return nil, restErr
		}
	} else {
		restErr := restErrors.NewBadRequestError(restErrors.MissingParameterYErrorMessage)
		logger.ErrorLogger.Print("Error while performing ParseCalculationPathVariables ", restErr)
		return nil, restErr
	}

	var calculationParsedVars dto.CalculationRequestDto
	calculationParsedVars.X = x
	calculationParsedVars.Y = y
	return &calculationParsedVars, nil

}

func (parser *calculationParser) TrimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func (parser *calculationParser) GetFirstChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[:1]
		}
	}
	return s
}
