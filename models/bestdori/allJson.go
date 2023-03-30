package bestdori

import (
	"Bestdori-Proxy/config"
	"strings"
)

type BandAllJsonItem struct {
	BandName []*string `json:"bandName"`
}

type BandAllJson map[int]BandAllJsonItem
type BandNameMap map[int]string

func (p BandAllJson) Convert() (bandList BandNameMap) {
	bandList = make(BandNameMap)
	for key, value := range p {
		for _, s := range value.BandName {
			if s != nil {
				bandList[key] = strings.ToValidUTF8(*s, "")
				break
			}
		}
	}
	return bandList
}

type LLSifAllJsonItem struct {
	LiveIconAsset string `json:"live_icon_asset"`
	SoundAsset    string `json:"sound_asset"`
}

type LLSifAllJson map[int]LLSifAllJsonItem

func (p LLSifAllJsonItem) CoverUrl() string {
	return config.LLSifAssetPrefix + p.LiveIconAsset
}

func (p LLSifAllJsonItem) SongUrl() string {
	return config.LLSifAssetPrefix + p.SoundAsset
}

type ChartAllJson map[int]interface{}
