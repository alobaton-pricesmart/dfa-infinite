package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"dfa-infinite/pkg/dfa"

	log "github.com/sirupsen/logrus"
)

// main function of the program
func main() {
	args := os.Args[1:]
	fileName := args[0]

	if len(fileName) == 0 {
		log.Error("fileName is required")
		return
	}

	log.WithField("fileName", fileName).Info("reading the file...")

	// read the file
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.WithField("error", err).Error("error reading the file")
		return
	}

	// Just for testing purpose...
	// log.WithField("body", string(body)).Info()

	// Mapeamos el archivo en una estructura.
	d := dfa.DFA{}
	err = json.Unmarshal([]byte(body), &d)
	if err != nil {
		log.WithField("error", err).Error("error unmarshaling the file")
		return
	}

	// Just for testing purpose...
	// log.WithField("dfa", d).Info()

	// Validamos si el AFD es finito o infinito.
	finite := d.IsFinite(d.InitialState, "")
	log.WithField("finite", finite).Info("result")
}
