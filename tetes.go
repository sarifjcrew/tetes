package tetes

type RequestConfig struct {
	Voice        *Voice
	Encoding     string
	SpeakingRate float64
	Pitch        float64
	VolumeGain   float64
}

type Voice struct {
	Language string
	Gender   string
	Name     string
}

type SynthesizeKind string

const (
	SynthesizeText SynthesizeKind = "Text"
	SinthesizeSML  SynthesizeKind = "SML"
)

type Provider interface {
	Voices() ([]Voice, error)
	Synthesize(txt string, kind string, config *RequestConfig) ([]byte, error)
	Close()
}
