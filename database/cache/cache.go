package cache

import (
	"GoBot-Recode/core/logger"
	"github.com/patrickmn/go-cache"
	"time"
)

var dataCache *cache.Cache

/*
	Initialize cache
 */
func Init() {
	logger.LogModuleNoNewline(logger.TypeInfo, "GoBot/Cache", "Initializing cache...")
	dataCache = cache.New(5*time.Minute, 10*time.Minute)
	logger.AppendDone()
}

/*
	Get a key stored in cache
       - returns: interface{}
*/
func GetKey(key string) interface{} {
	str, _ := dataCache.Get(key)
	return str
}

/*
	Deletes a key stored in cache
*/
func DeleteKey(key string) {
	dataCache.Delete(key)
}

/*
	Sets a key in cache
*/
func SetKey(key string, value interface{}) {
	dataCache.Set(key, value, cache.DefaultExpiration)
}