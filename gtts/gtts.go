package gtts

import (
	"context"
	"errors"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/sarifjcrew/tetes"
)

type GttsEngine struct {
}

func New() *GttsEngine {
	eng := new(GttsEngine)
	return eng
}

func (g *GttsEngine) Voices() ([]tetes.Voice, error) {
	c, e := newClient(nil)
	if e != nil {
		return []tetes.Voice{}, e
	}
	defer c.Close()

	res, e := c.ListVoices(context.Background(), &texttospeechpb.ListVoicesRequest{})
	if e != nil {
		return []tetes.Voice{}, e
	}

	voices := make([]tetes.Voice, len(res.Voices))
	for idx, voice := range res.Voices {
		voices[idx] = tetes.Voice{
			Language: voice.LanguageCodes[0],
			Name:     voice.Name,
			Gender:   voice.SsmlGender.String(),
		}
	}
	return voices, nil
}

func (g *GttsEngine) Synthesize(txt string, kind string, config *tetes.RequestConfig) ([]byte, error) {
	ctx := context.Background()
	c, e := newClient(ctx)
	if e != nil {
		return nil, e
	}
	defer c.Close()

	ac := &texttospeechpb.AudioConfig{
		AudioEncoding: texttospeechpb.AudioEncoding_MP3,
	}

	if config.Pitch != 0 {
		ac.Pitch = config.Pitch
	}
	if config.SpeakingRate != 0 {
		ac.SpeakingRate = config.SpeakingRate
	}
	if config.VolumeGain != 0 {
		ac.VolumeGainDb = config.VolumeGain
	}

	req := texttospeechpb.SynthesizeSpeechRequest{
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: config.Voice.Language,
			Name:         config.Voice.Name,
		},
		AudioConfig: ac,
	}

	switch kind {
	case string(tetes.SinthesizeSML):
		req.Input = &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Ssml{Ssml: txt},
		}

	case string(tetes.SynthesizeText):
		req.Input = &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: txt},
		}

	default:
		return nil, errors.New("invalid sythesize kind")
	}

	resp, e := c.SynthesizeSpeech(ctx, &req)
	if e != nil {
		return nil, errors.New("fail to synthesize: " + e.Error())
	}
	return resp.AudioContent, nil
}

func (g *GttsEngine) Close() {
}

func newClient(ctx context.Context) (*texttospeech.Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	c, err := texttospeech.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}
