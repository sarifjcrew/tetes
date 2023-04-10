package tetes_test

import (
	"io/ioutil"
	"testing"

	"github.com/sebarcode/codekit"
	"github.com/sebarcode/tetes"
	"github.com/sebarcode/tetes/gtts"
	"github.com/smartystreets/goconvey/convey"
)

func TestGetVoice(t *testing.T) {
	convey.Convey("prepare provider", t, func() {
		pro := gtts.New()
		voices, err := pro.Voices()
		convey.So(err, convey.ShouldBeNil)
		convey.So(len(voices), convey.ShouldBeGreaterThan, 0)
		convey.Println("voices: " + codekit.JsonString(voices[:3]))
	})
}

func TestSynthesize(t *testing.T) {
	convey.Convey("get provider", t, func() {
		pro := gtts.New()
		eng := tetes.NewEngine(pro, func(bs []byte) error {
			e := ioutil.WriteFile("d:\\sevo1.mp3", bs, 0644)
			return e
		})
		defer eng.Close()

		err := eng.Synthesize("hai selamat pagi, nama saya sevo", tetes.SynthesizeText, &tetes.RequestConfig{
			Voice: &tetes.Voice{
				Language: "id-ID",
				Name:     "id-ID-Standard-C",
			},
			Encoding: "MP3",
		})

		convey.So(err, convey.ShouldBeNil)
	})
}
