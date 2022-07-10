package application

import (
	"github.com/iloncar89/calculation-api/src/cache"
	calculationDao "github.com/iloncar89/calculation-api/src/dao/calculation"
	"github.com/iloncar89/calculation-api/src/utils/logger"
	"net/http"
)

//StartApplication function contains calls for get environment variable for cache memory management, setting memory management, mapping urls and starting server.
func StartApplication() {

	cacheMemoryManagementEnabled := getEnvironmentVariableForCache()

	cache.AppCache.SetCacheMemoryManagement(cacheMemoryManagementEnabled, calculationDao.M)

	mapUrls()

	logger.InfoLogger.Printf("About to start the application...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
