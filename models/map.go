package models

import "fmt"

type BestdoriV2Note struct {
	BestdoriV2BasicNote
	Type        string                `json:"type"`
	BPM         float64               `json:"bpm,omitempty"`
	Connections []BestdoriV2BasicNote `json:"connections,omitempty"`
	Direction   string                `json:"direction,omitempty"`
	Width       int                   `json:"width,omitempty"`
}

type BestdoriV2BasicNote struct {
	Beat_  interface{} `json:"beat,omitempty"`
	Lane_  interface{} `json:"lane,omitempty"`
	Flick  bool        `json:"flick,omitempty"`
	Hidden bool        `json:"hidden,omitempty"`
}

type BestdoriCustomMap struct {
	Result bool `json:"result"`
	Post   struct {
		CategoryName string          `json:"categoryName"`
		CategoryId   string          `json:"categoryId"`
		Chart        BestdoriV2Chart `json:"chart"`
	} `json:"post"`
}

func (note BestdoriV2Note) Beat() (value float64) {
	var ok bool
	if len(note.Connections) == 0 {
		value, ok = note.Beat_.(float64)
	} else {
		value, ok = note.Connections[0].Beat_.(float64)
	}
	if !ok {
		return 0
	}
	return value
}

func (note BestdoriV2Note) Lane() (value float64) {
	var ok bool
	if len(note.Connections) == 0 {
		value, ok = note.Lane_.(float64)
	} else {
		value, ok = note.Connections[0].Lane_.(float64)
	}
	if !ok {
		return 0
	}
	return value
}

type BestdoriV2Chart []BestdoriV2Note

func (formatChart BestdoriV2Chart) Len() int {
	return len(formatChart)
}

func (formatChart BestdoriV2Chart) Less(i, j int) bool {
	if formatChart[i].Beat() == formatChart[j].Beat() {
		return formatChart[i].Lane() < formatChart[j].Lane()
	}
	return formatChart[i].Beat() < formatChart[j].Beat()
}

func (formatChart BestdoriV2Chart) Swap(i, j int) {
	formatChart[i], formatChart[j] = formatChart[j], formatChart[i]
}
func (formatChart BestdoriV2Chart) MapCheck() (result bool, err error) {
	BPMStartCorrect := false // BPM需出现在第0拍上
	for _, formatNote := range formatChart {
		switch formatNote.Type {
		case "Directional":
			if formatNote.Direction != "Left" && formatNote.Direction != "Right" {
				return false, fmt.Errorf("无法识别侧划音符的标识符")
			}
			if formatNote.Width < 0 || formatNote.Width > 3 {
				return false, fmt.Errorf("侧划音符超限")
			}
			fallthrough
		case "Single":
			currentBeat, ok := formatNote.Beat_.(float64)
			if !ok || currentBeat < 0.0 {
				return false, fmt.Errorf("无法解析某个音符的节拍/节拍数小于0")
			}
			_, ok = formatNote.Lane_.(float64) // 音符轨道可以被修正，在后续Decode部分修正
			if !ok {
				return false, fmt.Errorf("无法解析某个音符的轨道")
			}
			if len(formatNote.Connections) != 0 {
				return false, fmt.Errorf("单键错误的拥有Connections字段")
			}
		case "Long":
			fallthrough
		case "Slide":
			// 绿条长度可以被修正，在后续Decode部分修正
			for _, formatTick := range formatNote.Connections {
				currentBeat, ok := formatTick.Beat_.(float64)
				if !ok || currentBeat < 0.0 {
					return false, fmt.Errorf("无法解析某个绿条节点的节拍/节拍数小于0")
				}
				_, ok = formatTick.Lane_.(float64)
				if !ok {
					return false, fmt.Errorf("无法解析某个绿条节点的轨道")
				}
			}
		case "BPM":
			currentBeat, ok := formatNote.Beat_.(float64)
			if !ok || currentBeat < 0.0 {
				return false, fmt.Errorf("无法解析BPM节点的节拍/节拍数小于0")
			}
			if formatNote.Beat_.(float64) == 0.0 {
				BPMStartCorrect = true
			}
			if len(formatNote.Connections) != 0 {
				return false, fmt.Errorf("BPM错误的拥有Connections字段")
			}
			// BPM的正负会在Decode部分修正
		default:
			// 不知道的音符会在Decode部分扔掉
		}
	}
	if !BPMStartCorrect {
		return false, fmt.Errorf("BPM不在Beat 0上")
	}
	return true, nil
}
