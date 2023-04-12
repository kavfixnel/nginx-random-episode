package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func generateMetadata(rootPath, service, imgRef, overviewUrl string, numSeasons int) error {
	response, err := http.Get(showImg)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(fmt.Sprintf("%s/static/images/%s.jpeg", rootPath, showShort))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	// Write metadata.json file
	metadataFile, err := os.Create(fmt.Sprintf("%s/episodes/%s/%s/metadata.json", rootPath, service, showShort))
	if err != nil {
		return err
	}
	defer metadataFile.Close()

	metadataContent := showMetadata{
		Slug:        showSlug,
		Imgref:      imgRef,
		ShowPath:    showPath,
		OverviewUrl: overviewUrl,
		NumSeasons:  numSeasons,
	}
	metadata, err := json.MarshalIndent(metadataContent, "", "    ")
	if err != nil {
		return err
	}
	metadataFile.Write(metadata)

	return nil
}
