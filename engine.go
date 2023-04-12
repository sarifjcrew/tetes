package tetes

import (
	"errors"
	"fmt"
)

type Engine struct {
	provider Provider
	writeFn  func([]byte) error
}

func NewEngine(provider Provider, writeFn func([]byte) error) *Engine {
	eng := new(Engine)
	eng.provider = provider
	eng.writeFn = writeFn
	return eng
}

func (eng *Engine) Synthesize(txt string, kind SynthesizeKind, config *RequestConfig) ([]byte, error) {
	var err error
	var res []byte
	if eng.provider == nil {
		return res, errors.New("tts provider is nil")
	}

	if eng.writeFn == nil {
		return res, errors.New("write function is nil")
	}

	// bs := []byte{}
	if res, err = eng.provider.Synthesize(txt, string(kind), config); err != nil {
		return res, fmt.Errorf("synthesize error: %s", err.Error())
	}

	if err = eng.writeFn(res); err != nil {
		return res, fmt.Errorf("write content error: %s", err.Error())
	}

	return res, nil
}

func (eng *Engine) Close() {
	if eng.provider != nil {
		eng.provider.Close()
	}
}
