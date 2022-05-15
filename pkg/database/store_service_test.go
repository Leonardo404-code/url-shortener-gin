package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://github.com/Leonardo404-code"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	SaveUrl(shortURL, initialLink, userUUId)

	originalURL := RetrieveOriginalUrl(shortURL)

	assert.Equal(t, initialLink, originalURL)
}

func BenchmarkInsertionAndRetrieval(b *testing.B) {
	initialLink := "https://github.com/Leonardo404-code"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	for i := 0; i < b.N; i++ {
		SaveUrl(shortURL, initialLink, userUUId)

		originalURL := RetrieveOriginalUrl(shortURL)

		assert.Equal(b, initialLink, originalURL)
	}
}
