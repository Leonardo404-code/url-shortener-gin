package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func BenchmarkStoreInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assert.True(b, testStoreService.redisClient != nil)
	}
}
