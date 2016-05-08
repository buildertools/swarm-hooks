package chain

import (
	log "github.com/Sirupsen/logrus"
)

type Chains map[string]Chain
type Processor map[string]Attributes
type Attributes map[string]string

type Chain struct {
	Include        []string    `yaml:"include,flow"`
	Exclude        []string    `yaml:"exclude,flow"`
	Preprocessors  []Processor `yaml:"pre,flow"`
	Postprocessors []Processor `yaml:"post,flow"`
}

func (c Chain) Preprocess(e PreEvent, ctx Context) {
	log.Printf("Preprocessing event: %s", e.Name())
}

func (c Chain) Postprocess(e PostEvent, ctx Context) {
	log.Printf("Postprocessing event: %s", e.Name())
}

func (c Chain) Rollback(e PreEvent, ctx Context) {
	log.Printf("Rollingback event: %s", e.Name())
}

func (c Chain) ListPreprocessors() {
	log.Printf("Preprocessors: %s", c.Preprocessors)
}
