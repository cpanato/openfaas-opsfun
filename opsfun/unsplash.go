package function

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	imageURL = "https://api.unsplash.com/photos/random?query=%s&client_id=a5eb0efe067ab368afab5c0cb1887bed4ece479187ff1da89ecb818e1f2a0e89"
)

func getImage(category string) (image, thumb string) {
	for i := 0; i < 5; i++ {
		image, thumb, err := readImage(category)
		if err != nil {
			continue
		}
		return image, thumb
	}

	return "", ""
}

type imageResult struct {
	ImageURLS imageImages `json:"urls"`
}

type imageImages struct {
	Small string `json:"small"`
	Thumb string `json:"thumb"`
}

func readImage(category string) (image string, thumb string, err error) {
	client := http.Client{}
	uri := fmt.Sprintf(imageURL, category)
	resp, err := client.Get(uri)
	if err != nil {
		return "", "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("no princess found")
	}
	var a imageResult
	if err = json.NewDecoder(resp.Body).Decode(&a); err != nil {
		return "", "", fmt.Errorf("failed to decode response: %v", err)
	}
	return a.ImageURLS.Small, a.ImageURLS.Thumb, nil
}
