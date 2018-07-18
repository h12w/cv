package cv

import (
	"encoding/json"
	"os"
	"path"
)

var DefaultDataPath = path.Join(os.Getenv("GOPATH"), "src/h12.io/cv-data")

func JSON(filename string, data interface{}) error {
	cvFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer cvFile.Close()
	return json.NewDecoder(cvFile).Decode(data)
}
