package formatter

import "io/ioutil"

//go:generate go run -tags=dev assets_gen.go

func getAssetContent(path string) ([]byte, error) {
	f, err := Assets.Open("/html.html")
	if err != nil {

		return nil, err
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {

		return nil, err
	}
	return content, nil
}
