package bestdori

import (
	"Bestdori-Proxy/config"
	"math"
	"strings"
)

type OfficialPost struct {
	MusicTitle  []*string `json:"musicTitle"`
	BandID      int       `json:"bandId"`
	BgmFile     string    `json:"bgmFile"`
	JacketImage []string  `json:"jacketImage"`
	Description []*string `json:"Description"`
	PublishAt   []*int64  `json:"publishAt"`
	Difficulty  map[int]struct {
		PlayLevel int `json:"playLevel"`
	} `json:"difficulty"`
}

func (p *OfficialPost) GetRegion() string {
	var region = []string{"jp", "en", "tw", "cn", "kr"}
	for i, s := range p.MusicTitle {
		if s != nil {
			return region[i]
		}
	}
	return region[0]
}

func (p *OfficialPost) Title() string {
	for _, s := range p.MusicTitle {
		if s != nil {
			return strings.ToValidUTF8(*s, "")
		}
	}
	return ""
}

func (p *OfficialPost) CoverUrl(musicID int) string {
	var jacket string
	for _, s := range p.JacketImage {
		if s != "" {
			jacket = strings.ToValidUTF8(s, "")
		}
	}
	var bundle = int(10 * math.Ceil(float64(musicID)/10.0))
	return config.BestdoriCoverUrl(p.GetRegion(), bundle, jacket)
}

func (p *OfficialPost) SongUrl(musicID int) string {
	return config.BestdoriBGMUrl(p.GetRegion(), musicID)
}

func (p *OfficialPost) Content() string {
	for _, s := range p.MusicTitle {
		if s != nil {
			return strings.ToValidUTF8(*s, "")
		}
	}
	return ""
}

func (p *OfficialPost) Time() int64 {
	for _, s := range p.PublishAt {
		if s != nil {
			return *s
		}
	}
	return 0
}
