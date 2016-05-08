package chain

import (
	"io/ioutil"
	//	"github.com/buildertools/swarm-hooks/plugins"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadFile(f string) map[string]Chain {
	log.Info(`Loading chain configuration.`)
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("Unable to read a configuration file at: %s", f)
		os.Exit(1)
	}

	t := make(map[string]Chain)
	if err := yaml.Unmarshal([]byte(data), &t); err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
	log.Printf("Parsed configuration:\n%s\n\n", string(d))

	return t
}
