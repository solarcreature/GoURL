package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func InitializeTest(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func IOTest(t *testing.T) {
	initialLink := "https://github.com/solarcreature/GoURL"
	userID := "test12345"
	shortURL := "TESTabcdef"

	StoreMap(shortURL,initialLink,userID)

	result := RetrieveOriginalURL(shortURL)

	assert.Equal(t,initialLink,result)
}