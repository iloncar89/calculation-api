package calculationDao

import (
	"github.com/iloncar89/calculation-api/src/cache"
	"github.com/iloncar89/calculation-api/src/dto"
	"github.com/iloncar89/calculation-api/src/utils/logger"
	"sync"
)

var CalculationDao calculationDaoInterface = &calculationDao{}
var M = &sync.RWMutex{}

type calculationDaoInterface interface {
	CacheResult(key string, value float64)
	GetFromCache(string, dto.CalculationDto) (*dto.CalculationDto, bool)
}

type calculationDao struct {
}

//CacheResult function used for call Put to cache item.
//Receive key and value as arguments.
func (calcDao *calculationDao) CacheResult(key string, value float64) {
	logger.InfoLogger.Print("Calling Put method on cache.")
	cache.AppCache.Put(key, value, M)
}

//GetFromCache function used for call GetFromCache to get item from cache.
//Receive key and calculationDto as arguments. Returns calculationDto with result and boolean value if item is found in cache.
func (calcDao *calculationDao) GetFromCache(key string, calc dto.CalculationDto) (*dto.CalculationDto, bool) {
	logger.InfoLogger.Print("Calling Get method on cache.")
	result, ok := cache.AppCache.Get(key, M)
	if ok {
		calc.Cached = true
		calc.Answer = result
		logger.InfoLogger.Print("Received value from cache, for key: ", key, " is received value: ", result)
		return &calc, true
	}
	logger.InfoLogger.Print("Item for key: ", key, " was not found in cache.")
	calc.Cached = false
	return &calc, false
}
