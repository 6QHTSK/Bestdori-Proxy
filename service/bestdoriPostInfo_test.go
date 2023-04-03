package service

import (
	"github.com/6QHTSK/Bestdori-Proxy/errors"
	"github.com/6QHTSK/Bestdori-Proxy/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestBandoriChart GET /bestdori/post/128/chart?diff=1
// 六兆年 Normal谱面有271个物件
func TestBandoriChart(t *testing.T) {
	postInfo, err := FetchBandoriPost(128, 1, models.MethodChart)
	assert.Nil(t, err)
	assert.Nil(t, postInfo.Info)
	assert.NotNil(t, postInfo.Chart)
	assert.Equal(t, len(*postInfo.Chart), 271)
}

// TestBandoriPostInfo GET /bestdori/post/10001?diff=4&official&method=info
// 这里应该返回官谱的彩虹节拍
func TestBandoriPostInfo(t *testing.T) {
	postInfo, err := FetchBandoriPost(10001, 4, models.MethodInfo)
	assert.Nil(t, err)
	assert.Equal(t, postInfo.Title, "彩虹节拍")
	assert.Equal(t, postInfo.Artists, "肥皂菌×易言×音阙诗听/赵方婧×西瓜Kune")
	assert.Equal(t, postInfo.Username, "craftegg")
	assert.Equal(t, postInfo.Nickname, "")
	assert.Equal(t, postInfo.Rating, 26)
	assert.Equal(t, postInfo.CoverUrl, "https://bestdori.com/assets/cn/musicjacket/musicjacket10010_rip/assets-star-forassetbundle-startapp-musicjacket-musicjacket10010-10001_caihongjiepai-jacket.png")
	assert.Equal(t, postInfo.AudioUrl, "https://bestdori.com/assets/cn/sound/bgm10001_rip/bgm10001.mp3")
	assert.Equal(t, postInfo.Content, "彩虹节拍")
	assert.Nil(t, postInfo.Chart)
	assert.Equal(t, postInfo.Time, int64(1673586000))
	// 测试缓存
	postInfo, err = FetchBandoriPost(10001, 3, models.MethodInfo)
	assert.Nil(t, err)
}

// TestBestdoriPost GET /bestdori/post/10001
// 这里返回彩绫与6QHTSK@psk2019的よいまちカンターレ
func TestBestdoriPost(t *testing.T) {
	postInfo, err := FetchBestdoriPost(10001, models.MethodFull)
	assert.Nil(t, err)
	assert.Equal(t, postInfo.Title, "YOIMACHI CANTARE(よいまちカンターレ)")
	assert.Equal(t, postInfo.Artists, "CORO MACHIKADO(コーロまちカド)")
	assert.Equal(t, postInfo.Username, "psk2019")
	assert.NotEqual(t, postInfo.Nickname, "")
	assert.Equal(t, postInfo.Rating, 27)
	assert.Equal(t, postInfo.AudioUrl, "https://bestdori.com/api/upload/file/f7ca6e9d4aceb6f3ccc77c4408896ee4f9637440")
	assert.Equal(t, postInfo.CoverUrl, "https://img.moegirl.org/common/a/a6/Yoimachi_Cantare.jpg")
	assert.Equal(t, postInfo.Time, int64(1583387691))
}

// TestBestdoriPostBandoriTypeAsset GET /bestdori/post/1005/info
// 测试官谱 musicURl 和 coverURL 接口
func TestBestdoriPostBandoriTypeAsset(t *testing.T) {
	postInfo, err := FetchBestdoriPost(1005, models.MethodInfo)
	assert.Nil(t, err)
	assert.Equal(t, postInfo.AudioUrl, "https://bestdori.com/assets/jp/sound/bgm166_rip/bgm166.mp3")
	assert.Equal(t, postInfo.CoverUrl, "https://bestdori.com/assets/jp/musicjacket/musicjacket170_rip/assets-star-forassetbundle-startapp-musicjacket-musicjacket170-166_maware_setsugekka-jacket.png")
	// 测试缓存
	postInfo, err = FetchBestdoriPost(1005, models.MethodInfo)
	assert.Nil(t, err)
}

// TestFanmadeLLSifTypeAsset GET /bestdori/post/2401/info
// 测试 LLSIF musicURL 和 coverURL 接口
func TestFanmadeLLSifTypeAsset(t *testing.T) {
	postInfo, err := FetchBestdoriPost(2401, models.MethodInfo)
	assert.Nil(t, err)
	assert.Equal(t, postInfo.AudioUrl, "https://card.niconi.co.ni/asset/assets/sound/music/m_001.mp3")
	assert.Equal(t, postInfo.CoverUrl, "https://card.niconi.co.ni/asset/assets/image/live/live_icon/l_jacket_001.png")
	assert.Equal(t, postInfo.Content, "...\n")
	postInfo, err = FetchBestdoriPost(2401, models.MethodInfo)
	assert.Nil(t, err)
}

// TestNonExistFanmadePost GET /bestdori/post/10000/info
// 测试不存在的自制谱
func TestNonExistFanmadePost(t *testing.T) {
	_, err := FetchBestdoriPost(10000, models.MethodInfo)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), errors.PostNotFound.Error())
}

// TestNonExistOfficialPost GET /bestdori/post/900/info
// 测试不存在的官谱
func TestNonExistOfficialPost(t *testing.T) {
	_, err := FetchBandoriPost(900, 3, models.MethodInfo)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), errors.PostNotFound.Error())
}
