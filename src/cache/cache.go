package cache

import (
	"fmt"
	"github.com/iloncar89/calculation-api/src/utils/enum"
	"github.com/iloncar89/calculation-api/src/utils/logger"
	"sync"
	"time"
)

var AppCache cacheInterface = &cache{cache: make(map[string]Entry)}

type cacheInterface interface {
	Get(key string, m *sync.RWMutex) (float64, bool)
	Put(key string, value float64, m *sync.RWMutex)
	Delete(key string, m *sync.RWMutex)
	CreateKeyForCache(x, y float64, action string) string
	DeleteUnusedCacheItems(m *sync.RWMutex)
	SetCacheMemoryManagement(cacheEnabled bool, m *sync.RWMutex)
}

type cache struct {
	cache map[string]Entry
}

type Entry struct {
	value float64
	time  time.Time
}

//Get function is used for search in cache for given key.
//Arguments of function are key, and mutex for locking cache while reading.
//Return values are value from cache as string, and bool for check if value exists for given key
func (c *cache) Get(key string, m *sync.RWMutex) (float64, bool) {
	m.RLock()
	result, ok := c.cache[key]
	m.RUnlock()
	if !ok {
		logger.InfoLogger.Print("Item not found in cache. Key: ", key)
		return 0, false
	}

	c.cache[key] = Entry{
		value: result.value,
		time:  time.Now(),
	}

	logger.InfoLogger.Print("Item found in cache. Key: ", key, "Value: ", result.value)

	return result.value, true
}

//Put function is used for set value in cache for given key.
//Arguments of function are key and value as string, and mutex for locking cache while setting new entry in cache.
func (c *cache) Put(key string, value float64, m *sync.RWMutex) {
	entry := Entry{
		value: value,
		time:  time.Now(),
	}
	m.Lock()
	c.cache[key] = entry
	m.Unlock()
	logger.InfoLogger.Print("New item added to cache. Key: ", key, " Value: ", value)
}

//Delete function is used for deleting in cache for given key.
//Arguments of function are key, and mutex for locking cache while deleting.
func (c *cache) Delete(key string, m *sync.RWMutex) {
	m.Lock()
	delete(c.cache, key)
	m.Unlock()
	logger.InfoLogger.Print("Item deleted from cache. Key: ", key)
}

func (c *cache) CreateKeyForCache(x, y float64, action string) string {
	var key string
	switch action {
	case enum.Add.String():
		key = c.createKeyForCacheAddAndMultiplyActions(x, y, action)
	case enum.Multiply.String():
		key = c.createKeyForCacheAddAndMultiplyActions(x, y, action)
	case enum.Subtract.String():
		key = c.createKeyForCacheSubtractAndDivideActions(x, y, action)
	case enum.Divide.String():
		key = c.createKeyForCacheSubtractAndDivideActions(x, y, action)
	}
	return key
}

//createKeyForCacheAddAndMultiplyActions function is used for creating key for cache for add and multiply actions.
//Arguments of function are x and y variables as float64, and calculation action as string.
//Return values is generated key which contains greater float variable, second float variable and action
func (c *cache) createKeyForCacheAddAndMultiplyActions(x, y float64, action string) string {
	var cacheKey string
	if x >= y {
		cacheKey = fmt.Sprintf("%f%f%s", x, y, action)
	} else {
		cacheKey = fmt.Sprintf("%f%f%s", y, x, action)
	}
	return cacheKey
}

//createKeyForCacheSubtractAndDivideActions function is used for creating key for cache for subtract and divide actions.
//Arguments of function are x and y variables as float64, and calculation action as string.
//Return values is generated key which contains x float variable, y second float variable and action
func (c *cache) createKeyForCacheSubtractAndDivideActions(x, y float64, action string) string {
	var cacheKey string
	cacheKey = fmt.Sprintf("%f%f%s", x, y, action)
	return cacheKey
}

//DeleteUnusedCacheItems function is used for compare if item stored in cache was not used for 1 minute. If so it calls delete function to delete those items.
func (c *cache) DeleteUnusedCacheItems(m *sync.RWMutex) {
	if len(c.cache) != 0 {
		for key, element := range c.cache {
			if 60*time.Second <= (time.Since(element.time)) {
				c.Delete(key, m)
				logger.InfoLogger.Print("Time spent in cache for item with key: ", time.Since(element.time))
				logger.InfoLogger.Print("Item deleted from cache, item key:", key, " , item value:", element)
			}
		}
	}
}

//SetCacheMemoryManagement function is used for setting go routine and infinite loop which will every 30 seconds call DeleteUnusedCacheItems to clean up unused cached items.
func (c *cache) SetCacheMemoryManagement(cacheEnabled bool, m *sync.RWMutex) {
	if cacheEnabled {
		go func(m *sync.RWMutex) {
			logger.InfoLogger.Print("Go routine for memory management is set.")
			for true {
				logger.InfoLogger.Print("Go routine calls DeleteUnusedCacheItems for checking unused cached items")
				c.DeleteUnusedCacheItems(m)
				time.Sleep(30 * time.Second)
			}
		}(m)
	}
}
