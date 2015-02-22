package desc

import (
	"encoding/json"
	"io/ioutil"
)

type MultiDockerDescReader struct {
	File string
}

func (reader *MultiDockerDescReader) NewMultiDockerDesc() (*MultiDockerDesc, error) {
	var data []byte
	data, err := ioutil.ReadFile(reader.File)
	if err != nil {
		return nil, err
	}

	var descriptor MultiDockerDesc
	err = json.Unmarshal(data, &descriptor)
	if err != nil {
		return nil, err
	}
	return &descriptor, nil
}
