package models

type UseDefaultStruct struct {
	UseDefault bool `json:"useDefault"`
}

type Engine struct {
	Engine        map[string]interface{} `json:"engine"`
	UseSkin       UseDefaultStruct       `json:"useSkin"`
	UseBackground UseDefaultStruct       `json:"useBackground"`
	UseEffect     UseDefaultStruct       `json:"useEffect"`
	UseParticle   UseDefaultStruct       `json:"useParticle"`
}

func (e *Engine) SetTrue() {
	e.UseSkin.UseDefault = true
	e.UseBackground.UseDefault = true
	e.UseEffect.UseDefault = true
	e.UseParticle.UseDefault = true
}
