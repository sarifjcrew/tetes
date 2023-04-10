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

func (eng *Engine) Synthesize(txt string, kind SynthesizeKind, config *RequestConfig) error {
	var err error

	if eng.provider == nil {
		return errors.New("tts provider is nil")
	}

	if eng.writeFn == nil {
		return errors.New("write function is nil")
	}

	bs := []byte{}
	if bs, err = eng.provider.Synthesize(txt, string(kind), config); err != nil {
		return fmt.Errorf("synthesize error: %s", err.Error())
	}

	if err = eng.writeFn(bs); err != nil {
		return fmt.Errorf("write content error: %s", err.Error())
	}

	return nil
}

func (eng *Engine) Close() {
	if eng.provider != nil {
		eng.provider.Close()
	}
}
