package home

import "encoding/json"

type Home struct {
}

type LHome struct {
	Home string
}

func Init() *Home {
	return &Home{}
}

func (h *Home) Index() []byte {
	lhome := &LHome{}
	lhome.Home = "halo anda berada di home"

	data, _ := json.Marshal(lhome)

	return data
}
