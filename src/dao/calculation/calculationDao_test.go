package calculationDao

import (
	"github.com/iloncar89/calculation-api/src/dto"
	"testing"
)

func TestCalculationDao_CacheResultReceivedAfterCaching(t *testing.T) {
	key := "test"
	var value float64
	value = 5.8
	CalculationDao.CacheResult(key, value)
	response, foundInCache := CalculationDao.GetFromCache(key, dto.CalculationDto{})
	if value != response.Answer || foundInCache != true {
		t.Errorf("Cached key %s and value %f, but after that not found in cache, got calculationDto %v and bool receivedFromCache %t", key, value, response, foundInCache)
	}

}

func TestCalculationDao_GetFromCacheSomethingThatIsNotCached(t *testing.T) {
	key := "123"
	response, receivedFromCache := CalculationDao.GetFromCache(key, dto.CalculationDto{})
	if response.Answer != 0 || receivedFromCache {
		t.Errorf("Cache empty, but received something from cache, got calculationDto %v and bool receivedFromCache %t", response, receivedFromCache)
	}
}
