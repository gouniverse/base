package helpers

import (
	"errors"
	"io"
	"net/http"
	"os"
)

// DownloadURL will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadURL(url string, localFilepath string) error {

	// Get the data
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp == nil {
		return errors.New("no response")
	}

	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(localFilepath)

	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	return err
}
