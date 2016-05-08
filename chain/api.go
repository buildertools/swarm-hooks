package chain

import ()

type Context interface {
	Cancelled() <-chan struct{}
	Value(k string) (string, error)
}

type PreEvent interface {
	Name() string
	// container options like name, privs, volumes, network, etc
}

type PostEvent interface {
	Name() string
	CID() string
	// container options like name, privs, volumes, network, etc
}

type Preprocessor interface {
	Preprocess(e PreEvent, c Context)
	Rollback(e PreEvent, c Context)
}

type Postprocessor interface {
	Postprocess(e PostEvent, c Context)
}
