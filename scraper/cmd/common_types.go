package cmd

type showMetadata struct {
	Slug        string `json:"slug"`
	Imgref      string `json:"imgRef"`
	ShowPath    string `json:"showPath"`
	OverviewUrl string `json:"overviewUrl"`
	NumSeasons  int    `json:"numSeasons"`
}
