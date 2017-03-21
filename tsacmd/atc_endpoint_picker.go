package tsa

import (
	"math/rand"
	"time"

	"github.com/concourse/atc"
	"github.com/tedsuo/rata"
)

type ATCEndpointPicker struct {
	ATCEndpoints []*rata.RequestGenerator
}

func NewATCEndpointPicker(atcURLFlags []URLFlag) *ATCEndpointPicker {
	atcEndpoints := []*rata.RequestGenerator{}
	for _, f := range atcURLFlags {
		atcURLs = append(atcURLs, rata.NewRequestGenerator(f.String(), atc.Routes))
	}

	rand.Seed(time.Now().Unix())

	return &ATCEndpointPicker{
		ATCEndpoints: atcEndpoints,
	}
}

func (p *ATCEndpointPicker) Pick() *rata.RequestGenerator {
	return p.ATCEndpoints[rand.Intn(len(p.ATCEndpoints))]
}
