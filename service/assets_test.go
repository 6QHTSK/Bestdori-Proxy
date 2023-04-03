package service

import (
	"github.com/6QHTSK/Bestdori-Proxy/errors"
	"github.com/6QHTSK/Bestdori-Proxy/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchAssetsUrl(t *testing.T) {
	// Bandori
	assets, err := FetchAssetsUrl(128, models.ServerBandori)
	assert.Nil(t, err)
	assert.Equal(t, "https://bestdori.com/assets/jp/musicjacket/musicjacket130_rip/assets-star-forassetbundle-startapp-musicjacket-musicjacket130-128_ichiyamonogatari-jacket.png", assets.Cover)
	assert.Equal(t, "https://bestdori.com/assets/jp/sound/bgm128_rip/bgm128.mp3", assets.Audio)
	// Caching
	assets, err = FetchAssetsUrl(128, models.ServerBandori)
	assert.Nil(t, err)
	assert.Equal(t, "https://bestdori.com/assets/jp/musicjacket/musicjacket130_rip/assets-star-forassetbundle-startapp-musicjacket-musicjacket130-128_ichiyamonogatari-jacket.png", assets.Cover)
	assert.Equal(t, "https://bestdori.com/assets/jp/sound/bgm128_rip/bgm128.mp3", assets.Audio)
	// LLSIF
	assets, err = FetchAssetsUrl(1, models.ServerLLSif)
	assert.Nil(t, err)
	assert.Equal(t, "https://card.niconi.co.ni/asset/assets/image/live/live_icon/l_jacket_001.png", assets.Cover)
	assert.Equal(t, "https://card.niconi.co.ni/asset/assets/sound/music/m_001.mp3", assets.Audio)
	// Bestdori
	assets, err = FetchAssetsUrl(50000, models.ServerBestdori)
	assert.Nil(t, err)
	assert.Equal(t, "https://bestdori.com/api/upload/file/dce2a181e0d0f8d3fd16ce7438fd5ebd18254706", assets.Cover)
	assert.Equal(t, "https://bestdori.com/api/upload/file/c3b0eedfa8c28cf131ccb4afef8961647f91e45b", assets.Audio)
	// Non Exist Bestdori
	assets, err = FetchAssetsUrl(10000, models.ServerBestdori)
	assert.Equal(t, errors.PostNotFound, err)
	// Non Exist Bandori
	assets, err = FetchAssetsUrl(-1, models.ServerBandori)
	assert.Equal(t, errors.PostNotFound, err)
	// Non Exist LLSIF
	assets, err = FetchAssetsUrl(-1, models.ServerLLSif)
	assert.Equal(t, errors.PostNotFound, err)
	// Non Exist Server
	assets, err = FetchAssetsUrl(1, -1)
	assert.Equal(t, errors.UnknownServerErr, err)
}
