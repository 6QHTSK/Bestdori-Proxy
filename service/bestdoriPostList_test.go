package service

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestFetchBestdoriPostList(t *testing.T) {
	offset := rand.Int() % 10000
	limit := rand.Int()%10 + 10
	count, list, err := FetchBestdoriPostList(uint64(offset), uint64(limit), "")
	assert.Nil(t, err)
	assert.NotEqual(t, count, 0)
	assert.Equal(t, limit, len(list))
	// With Username
	count, list, err = FetchBestdoriPostList(0, uint64(limit), "psk2019")
	assert.NotEqual(t, count, 0)
	assert.Nil(t, err)
	assert.Equal(t, limit, len(list))
}

func TestFetchBandoriPostList(t *testing.T) {
	limit := rand.Int()%10 + 10
	count, list, err := FetchBandoriPostList(0, uint64(limit))
	assert.NotEqual(t, count, 0)
	assert.Nil(t, err)
	assert.Equal(t, limit, len(list))
}
