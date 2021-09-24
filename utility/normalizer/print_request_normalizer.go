package normalizer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"threedee/entity"
)

type PrintRequestNormalizer struct {
}

func NewPrintRequestNormalizer() *PrintRequestNormalizer {
	return &PrintRequestNormalizer{}
}

func (*PrintRequestNormalizer) ReadAndNormalize(w http.ResponseWriter, r *http.Request) (*entity.PrintRequest, error) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, errors.New("failed to read request body")
	}

	// Unmarshal
	var output *entity.PrintRequest
	err = json.Unmarshal(b, &output)
	if err != nil {
		return nil, errors.New("failed to unmarshal request body")
	}

	return output, nil
}
