package mapper

import (
	"github.com/iloncar89/calculation-api/src/dto"
	"github.com/iloncar89/calculation-api/src/utils/constants/path/calculation"
	"github.com/iloncar89/calculation-api/src/utils/enum"
)

var CalculationMapper calculationMapperInterface = &calculationMapper{}

type calculationMapperInterface interface {
	MapCalculationAction(path string) *dto.CalculationDto
}

type calculationMapper struct {
}

func (mapper *calculationMapper) MapCalculationAction(path string) *dto.CalculationDto {
	calc := dto.CalculationDto{}
	switch path {
	case calculation.AddCalculationUrlPath:
		calc.Action = enum.Add.String()
	case calculation.SubtractCalculationUrlPath:
		calc.Action = enum.Subtract.String()
	case calculation.MultiplyCalculationUrlPath:
		calc.Action = enum.Multiply.String()
	case calculation.DivideCalculationUrlPath:
		calc.Action = enum.Divide.String()
	}
	return &calc
}
