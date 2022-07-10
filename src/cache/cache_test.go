package cache

import (
	"github.com/iloncar89/calculation-api/src/utils/enum"
	"sync"
	"testing"
)

var Mutex = &sync.RWMutex{}

func TestCache_CreateKeyForCacheAddAndMultiplyActionsFloatVariablesInArgumentSuccess(t *testing.T) {
	var num1 float64
	var num2 float64
	num1 = 1.11
	num2 = 5.51
	expectedRes := "5.5100001.110000add"
	action := enum.Add.String()
	key := AppCache.CreateKeyForCache(num1, num2, action)
	if key != expectedRes {
		t.Errorf("Key was created %s, but expected is %s", key, expectedRes)
	}
}

func TestCache_Delete(t *testing.T) {
	key := "test"
	var value float64
	value = 6.4
	AppCache.Put(key, value, Mutex)
	resp, cached := AppCache.Get(key, Mutex)
	if value != resp || !cached {
		t.Errorf("Cached key %s and value %f, but after that not found in cache.", key, value)
	}
	AppCache.Delete(key, Mutex)
	valueAfterDelete, cachedAfterDelete := AppCache.Get(key, Mutex)
	if value == valueAfterDelete || cachedAfterDelete {
		t.Errorf("Cached key %s and value %f, received after caching and deleted but received from cache, value %f, and bool cached %t", key, value, valueAfterDelete, cachedAfterDelete)
	}
}

func TestCache_GetSuccessAfterPuttingItemInCache(t *testing.T) {
	key := "test3"
	var value float64
	value = 5.3
	AppCache.Put(key, value, Mutex)
	resp, cached := AppCache.Get(key, Mutex)
	if value != resp || !cached {
		t.Errorf("Cached key %s and value %f, but after that not found in cache.", key, value)
	}

}

func TestCache_GetSomethingNotCached(t *testing.T) {
	key := "test4"
	resp, cached := AppCache.Get(key, Mutex)
	if cached || resp != 0 {
		t.Errorf("Received somethnig from cache %f but it is expected that cache is empty", resp)
	}
}
