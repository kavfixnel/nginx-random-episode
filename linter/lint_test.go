package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/integralist/go-findroot/find"
	"github.com/stretchr/testify/assert"
)

type showMetadata struct {
	Slug       string `json:"slug"`
	Imgref     string `json:"imgRef"`
	NumSeasons int    `json:"numSeasons"`
}

func lintShow(t *testing.T, path string) {
	var metadata showMetadata

	b, err := os.ReadFile(path + "/metadata.json")
	assert.NoError(t, err, "metadata.json should exist.")

	err = json.Unmarshal(b, &metadata)
	assert.NoError(t, err, "metadata.json should be valid json.")

	assert.NotEqual(t, "", metadata.Slug, "metadata.Slug should not be an empty string.")
	assert.NotEqual(t, 0, metadata.NumSeasons, "metadata.NumSeasons should not be 0.")

	file, err := os.ReadFile(path + "/all.episodes")
	assert.NoError(t, err, "all.episodes should exist")

	all := make([]string, 0)
	n := 0
	for _, episode := range strings.Split(string(file), "\n") {
		if episode == "" {
			continue
		}

		all = append(all, episode)
		n++
	}
	assert.NotEqual(t, 0, n, "all.episodes should not be empty")

	seasons := make([]string, 0, len(all))
	for i := 1; i <= metadata.NumSeasons; i++ {
		file, err := os.ReadFile(fmt.Sprintf("%s/s%d.episodes", path, i))
		assert.NoError(t, err, fmt.Sprintf("%s/s%d.episodes should exist.", path, i))

		n := 0
		for _, episode := range strings.Split(string(file), "\n") {
			if episode == "" {
				continue
			}

			seasons = append(seasons, episode)
			n++
		}

		assert.NotEqual(t, 0, n, fmt.Sprintf("%s cannot be empty", path))
	}

	assert.ElementsMatch(t, all, seasons, "Episodes in all.episodes must equal the sum of all season episodes.")

	_, err = os.Stat(fmt.Sprintf("%s/s%d.episodes", path, metadata.NumSeasons+1))
	assert.Error(t, err, fmt.Sprintf("%s/s%d.episodes should not exist.", path, metadata.NumSeasons+1))

	_, err = os.Stat(fmt.Sprintf(path, metadata.Imgref))
	assert.Error(t, err, fmt.Sprintf("%s for %s should exists.", metadata.Imgref, fmt.Sprintf(path, metadata.Imgref)))
}

func TestLint(t *testing.T) {
	root, err := find.Repo()
	assert.NoError(t, err, "Should be in valid git repo.")

	baseFolder := root.Path + "/episodes"

	serviceDirs, err := os.ReadDir(baseFolder)
	assert.NoError(t, err, "Base folder should exist.")

	for _, serviceDir := range serviceDirs {
		showDirs, err := os.ReadDir(baseFolder + "/" + serviceDir.Name())
		assert.NoError(t, err, "Service folder should exist.")

		for _, showDir := range showDirs {
			t.Run(fmt.Sprintf("Linting %s/%s", serviceDir.Name(), showDir.Name()), func(t *testing.T) {
				lintShow(t, baseFolder+"/"+serviceDir.Name()+"/"+showDir.Name())
			})
		}
	}
}
