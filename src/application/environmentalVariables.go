package application

import (
	"github.com/iloncar89/calculation-api/src/utils/logger"
	"os"
	"strconv"
)

const (
	cacheMemoryManagementEnabled = "CACHE_MEMORY_MANAGEMENT_ENABLED"
)

//getEnvironmentVariableForCache function is used to collect environment variable for cache.
//It returns true by default, and false if CACHE_MEMORY_MANAGEMENT_ENABLED environment variable contains false value.
func getEnvironmentVariableForCache() bool {
	cacheDeletionEnabled, err := strconv.ParseBool(os.Getenv(cacheMemoryManagementEnabled))
	logger.InfoLogger.Print("Environment variable for enabling cache periodical deletion (cacheMemoryManagementEnabled): ", cacheDeletionEnabled)
	if err != nil {
		cacheDeletionEnabled = true
		logger.WarningLogger.Print("Environment variable for enabling cache periodical deletion (cacheMemoryManagementEnabled) is not set, as default it is turned on")
	}
	return cacheDeletionEnabled
}
