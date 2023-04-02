package bestdori

import (
	"Bestdori-Proxy/config"
	"math"
	"strconv"
	"strings"
)

type BandoriPost struct {
	MusicTitle  []*string `json:"musicTitle"`
	BandID      int       `json:"bandId"`
	BgmFile     string    `json:"bgmFile"`
	JacketImage []string  `json:"jacketImage"`
	Description []*string `json:"Description"`
	PublishedAt []*string `json:"publishedAt"`
	Difficulty  map[int]struct {
		PlayLevel int `json:"playLevel"`
	} `json:"difficulty"`
}

func (p *BandoriPost) GetRegion() string {
	var region = []string{"jp", "en", "tw", "cn", "kr"}
	for i, s := range p.MusicTitle {
		if s != nil {
			return region[i]
		}
	}
	return region[0]
}

func (p *BandoriPost) Title() string {
	for _, s := range p.MusicTitle {
		if s != nil {
			return strings.ToValidUTF8(*s, "")
		}
	}
	return ""
}

func (p *BandoriPost) CoverUrl(musicID int) string {
	var jacket string
	for _, s := range p.JacketImage {
		if s != "" {
			jacket = strings.ToValidUTF8(s, "")
		}
	}
	var bundle = int(10 * math.Ceil(float64(musicID)/10.0))
	return config.BestdoriCoverUrl(p.GetRegion(), bundle, jacket)
}

func (p *BandoriPost) AudioUrl(musicID int) string {
	return config.BestdoriAudioUrl(p.GetRegion(), musicID)
}

func (p *BandoriPost) Content() string {
	for _, s := range p.MusicTitle {
		if s != nil {
			return strings.ToValidUTF8(*s, "")
		}
	}
	return ""
}

func (p *BandoriPost) Time() int64 {
	for _, s := range p.PublishedAt {
		if s != nil {
			value, err := strconv.ParseInt(*s, 10, 64)
			if err != nil {
				continue
			}
			return value / 1000
		}
	}
	return 0
}
