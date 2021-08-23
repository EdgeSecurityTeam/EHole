package finger

import (
	"encoding/json"
	"io/ioutil"
)

type Packjson struct {
	Fingerprint []Fingerprint
}

type Fingerprint struct {
	Cms      string
	Method   string
	Location string
	Keyword  []string
}

var (
	Webfingerprint *Packjson
)

func LoadWebfingerprint(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var config Packjson
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	Webfingerprint = &config
	return nil
}

func GetWebfingerprint() *Packjson {
	return Webfingerprint
}
