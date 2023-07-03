package bestdori

import (
	"github.com/6QHTSK/Bestdori-Proxy/errors"
	"math"
	"sort"
)

type V2Note struct {
	Beat        *float64       `json:"beat,omitempty"`
	Lane        *float64       `json:"lane,omitempty"`
	Flick       bool           `json:"flick,omitempty"`
	Type        string         `json:"type"`
	BPM         float64        `json:"bpm,omitempty"`
	Connections []V2Connection `json:"connections,omitempty"`
	Direction   string         `json:"direction,omitempty"`
	Width       int            `json:"width,omitempty"`
}

type V2Connection struct {
	Beat   *float64 `json:"beat,omitempty"`
	Lane   *float64 `json:"lane,omitempty"`
	Flick  bool     `json:"flick,omitempty"`
	Hidden bool     `json:"hidden,omitempty"`
}

func (note V2Note) GetBeat() (value float64) {
	if len(note.Connections) == 0 {
		return *note.Beat
	} else {
		return *note.Connections[0].Beat
	}
}

func (note V2Note) GetLane() (value float64) {
	if note.Type == "BPM" {
		return 0
	}
	if len(note.Connections) == 0 {
		return *note.Lane
	} else {
		return *note.Connections[0].Lane
	}
}

type V2Chart []V2Note

func (chart V2Chart) Len() int {
	return len(chart)
}

func (chart V2Chart) Less(i, j int) bool {
	if chart[i].GetBeat() == chart[j].GetBeat() {
		return chart[i].GetLane() < chart[j].GetLane()
	}
	return chart[i].GetBeat() < chart[j].GetBeat()
}

func (chart V2Chart) Swap(i, j int) {
	chart[i], chart[j] = chart[j], chart[i]
}
func fixLane(lane float64, isHidden bool) float64 {
	if isHidden {
		return math.Max(-0.5, math.Min(6.5, lane))
	}
	return math.Max(0.0, math.Min(6.0, math.Round(lane)))
}

// ChartCheck 尝试检查并修复谱面，并将官谱转化为自制谱格式
func (chart V2Chart) ChartCheck() (err error) {
	BPMStartCorrect := false // BPM需出现在第0拍上
	var i int
	for _, note := range chart {
		toRemove := false
		switch note.Type {
		case "Directional":
			if note.Direction != "Left" && note.Direction != "Right" {
				return errors.DirectionNoteTypeErr
			}
			note.Width = int(math.Max(1, math.Min(3, float64(note.Width))))
			fallthrough
		case "Single":
			if *note.Beat < 0.0 {
				return errors.BeatLessThanZero
			}
			*note.Lane = fixLane(*note.Lane, false)
			note.Connections = []V2Connection{}
		case "Long":
			note.Type = "Slide"
			fallthrough
		case "Slide":
			// 绿条长度可以被修正，在后续Decode部分修正
			sort.Slice(note.Connections, func(i, j int) bool {
				if *note.Connections[i].Beat == *note.Connections[j].Beat {
					return *note.Connections[i].Lane < *note.Connections[j].Lane
				}
				return *note.Connections[i].Beat < *note.Connections[j].Beat
			})
			for j, formatTick := range note.Connections {
				if *formatTick.Beat < 0.0 {
					return errors.BeatLessThanZero
				}
				*formatTick.Lane = fixLane(*formatTick.Lane, formatTick.Hidden)
				if j != len(note.Connections)-1 { // 不是最后一个节点
					formatTick.Flick = false
				} else {
					formatTick.Hidden = false
				}
			}
		case "BPM":
			if *note.Beat < 0.0 {
				return errors.BeatLessThanZero
			}
			if *note.Beat == 0.0 {
				BPMStartCorrect = true
			}
			note.Connections = []V2Connection{}
			// note.BPM = math.Abs(note.BPM) // Sonolus服务器优化，不再修复负BPM问题
		default:
			toRemove = true
		}
		if !toRemove {
			chart[i] = note
			i++
		}
	}
	if !BPMStartCorrect {
		return errors.BPMNotAtBeatZero
	}
	sort.Slice(chart, func(i, j int) bool {
		if chart[i].GetBeat() == chart[j].GetBeat() {
			return chart[i].GetLane() < chart[j].GetLane()
		}
		return chart[i].GetBeat() < chart[j].GetBeat()
	})
	return nil
}
