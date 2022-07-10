package services

import (
	"github.com/iloncar89/calculation-api/src/cache"
	"github.com/iloncar89/calculation-api/src/dao/calculation"
	"github.com/iloncar89/calculation-api/src/dto"
	"github.com/iloncar89/calculation-api/src/utils/enum"
	"github.com/iloncar89/calculation-api/src/utils/helper"
	"github.com/iloncar89/calculation-api/src/utils/logger"
	"github.com/iloncar89/calculation-api/src/utils/mapper"
	"github.com/iloncar89/calculation-api/src/utils/restErrors"
)

var CalculationService calculationServiceInterface = &calculationService{}

type calculationServiceInterface interface {
	BasicMathOperation(calc dto.CalculationRequestDto) (*dto.CalculationDto, restErrors.RestErr)
}

type calculationService struct {
}

//PerformBasicMathOperation function performs validation checks on calculation variables, call GetFromCache to check cache for result for given variables, basic math operation call if item was not stored in cache, and call for cache new result.
//Receiving arguments calculationRequestDto and returns calculationDTO if result was successfully calculated or retrieved from cache, or restError if something went wrong
func (c *calculationService) BasicMathOperation(calculationRequest dto.CalculationRequestDto) (*dto.CalculationDto, restErrors.RestErr) {
	logger.InfoLogger.Print(calculationRequest)
	calc := mapper.CalculationMapper.MapCalculationAction(calculationRequest.Action)
	xString := helper.CalculationHelper.RemovePlusOperatorFromString(calculationRequest.X)
	yString := helper.CalculationHelper.RemovePlusOperatorFromString(calculationRequest.Y)

	logger.InfoLogger.Print("Calling IsStringStartingWithTwoZeroDigits for x: ", xString, " and y: ", yString, "action: ", calc.Action)
	if helper.CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(xString) || helper.CalculationHelper.IsVariableNotFormattedCorrectlyWithLeadingZero(yString) {
		return nil, restErrors.NewBadRequestError(restErrors.ErrorLeadingZeros)
	}

	logger.InfoLogger.Print("Calling ConvertStringToNumber for x: ", xString)
	x, err := helper.CalculationHelper.ConvertStringToNumber(xString)
	if err != nil {
		return nil, restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)
	}
	calc.X = x
	logger.InfoLogger.Print("Calling ConvertStringToNumber for y: ", yString)
	y, err := helper.CalculationHelper.ConvertStringToNumber(yString)
	if err != nil {
		return nil, restErrors.NewBadRequestError(restErrors.ErrorParsingStringToNumber)
	}
	calc.Y = y

	logger.InfoLogger.Print("Calling CreateKeyForCache for x: ", calc.X, " and y: ", calc.Y, "action:", calc.Action)
	cacheKey := cache.AppCache.CreateKeyForCache(calc.X, calc.Y, calc.Action)

	if calc.Action == enum.Divide.String() {
		logger.InfoLogger.Print("Calling IsZeroValue for y:", calc.Y)
		isSecondParameterZero := helper.CalculationHelper.IsZeroValue(y)
		if isSecondParameterZero {
			return nil, restErrors.NewBadRequestError(restErrors.CannotDivideByZero)
		}
	}

	var ok bool
	logger.InfoLogger.Print("Calling GetFromCache for key: ", cacheKey)
	calc, ok = calculationDao.CalculationDao.GetFromCache(cacheKey, *calc)
	if ok {
		logger.InfoLogger.Print("Value got from cache", calc)
		return calc, nil
	}

	calc.Answer = helper.CalculationHelper.PerformBasicMathOperation(calc.X, calc.Y, calc.Action)
	logger.InfoLogger.Print("Math operation ", calc.Action, " performed on Float variables for x: ", x, " and y: ", y, " Result: ", calc.Answer)

	logger.InfoLogger.Print("Calling cacheKey for cache key: ", cacheKey, "and value ", calc.Answer)
	calculationDao.CalculationDao.CacheResult(cacheKey, calc.Answer)
	logger.InfoLogger.Print("Cached calculation", calc.Answer)

	return calc, nil
}
