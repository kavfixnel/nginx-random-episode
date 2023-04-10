package cmd

type peacockResp struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	ID         string `json:"id"`
	Type       string `json:"type"`
	ChildTypes struct {
		Images struct {
			NodeTypes []string `json:"nodeTypes"`
			Count     int      `json:"count"`
		} `json:"images"`
		FirstEp struct {
			NodeTypes []string `json:"nodeTypes"`
			Count     int      `json:"count"`
		} `json:"first_ep"`
		Shortforms struct {
			NodeTypes []string `json:"nodeTypes"`
			Count     int      `json:"count"`
		} `json:"shortforms"`
		FreeEpisodes struct {
			NodeTypes []interface{} `json:"nodeTypes"`
			Count     int           `json:"count"`
		} `json:"free_episodes"`
		Collections struct {
			NodeTypes []string `json:"nodeTypes"`
			Count     int      `json:"count"`
		} `json:"collections"`
		Clips struct {
			NodeTypes []string `json:"nodeTypes"`
			Count     int      `json:"count"`
		} `json:"clips"`
		Items struct {
			NodeTypes []string `json:"nodeTypes"`
			Count     int      `json:"count"`
		} `json:"items"`
		Latest struct {
			NodeTypes []string `json:"nodeTypes"`
			Count     int      `json:"count"`
		} `json:"latest"`
		Trailers struct {
			NodeTypes []string `json:"nodeTypes"`
			Count     int      `json:"count"`
		} `json:"trailers"`
	} `json:"childTypes"`
	Attributes struct {
		AvailableEpisodeCount int `json:"availableEpisodeCount"`
		AvailableSeasonCount  int `json:"availableSeasonCount"`
		Badging               struct {
			VideoFormats []struct {
				VideoFormat  string   `json:"videoFormat"`
				ColourSpaces []string `json:"colourSpaces"`
			} `json:"videoFormats"`
			AudioTracks []string `json:"audioTracks"`
		} `json:"badging"`
		Cast    []string `json:"cast"`
		Channel struct {
			AccessChannel string `json:"accessChannel"`
			Name          string `json:"name"`
		} `json:"channel"`
		ChildNodeTypes       []string `json:"childNodeTypes"`
		Classification       []string `json:"classification"`
		ContentSegments      []string `json:"contentSegments"`
		CreatedDate          int64    `json:"createdDate"`
		DescLongSeo          string   `json:"descLongSeo"`
		DeviceAvailabilities []struct {
			Format             string `json:"format"`
			MediaType          string `json:"mediaType"`
			OfferStage         string `json:"offerStage"`
			OfferStartTs       int64  `json:"offerStartTs"`
			OfferEndTs         int64  `json:"offerEndTs"`
			ContentSegment     string `json:"contentSegment"`
			VideoFormat        string `json:"videoFormat"`
			VideoFormatVariant string `json:"videoFormatVariant"`
			ColourSpace        string `json:"colourSpace"`
		} `json:"deviceAvailabilities"`
		DeviceAvailability struct {
			Available bool `json:"available"`
		} `json:"deviceAvailability"`
		FanCriticRating []struct {
			Source      string   `json:"source"`
			FanScore    int      `json:"fanScore"`
			CriticScore int      `json:"criticScore"`
			Tags        []string `json:"tags"`
		} `json:"fanCriticRating"`
		Formats struct {
			HD struct {
				AudioTracks struct {
					Eng []string `json:"eng"`
				} `json:"audioTracks"`
				EventStage   string `json:"eventStage"`
				Availability struct {
					Available        bool   `json:"available"`
					MediaType        string `json:"mediaType"`
					OfferStage       string `json:"offerStage"`
					OfferStartTs     int64  `json:"offerStartTs"`
					OfferEndTs       int64  `json:"offerEndTs"`
					AvailableDevices []struct {
						Type     string `json:"type"`
						Platform string `json:"platform"`
					} `json:"availableDevices"`
					ExtendedOfferStartTs int64  `json:"extendedOfferStartTs"`
					ExtendedOfferEndTs   int64  `json:"extendedOfferEndTs"`
					ContentSegment       string `json:"contentSegment"`
					VideoFormat          string `json:"videoFormat"`
					VideoFormatVariant   string `json:"videoFormatVariant"`
					ColourSpace          string `json:"colourSpace"`
				} `json:"availability"`
				StartOfCredits int `json:"startOfCredits"`
			} `json:"HD"`
		} `json:"formats"`
		GenreList []struct {
			Subgenre []string `json:"subgenre"`
			Genre    []string `json:"genre"`
		} `json:"genreList"`
		Genres            []string `json:"genres"`
		GracenoteSeriesID string   `json:"gracenoteSeriesId"`
		Images            []struct {
			URL  string `json:"url"`
			Type string `json:"type"`
		} `json:"images"`
		MediaTypes []struct {
			MediaType string `json:"mediaType"`
			Count     int    `json:"count"`
			Latest    int64  `json:"latest"`
		} `json:"mediaTypes"`
		MerlinSeriesID    string `json:"merlinSeriesId"`
		OttCertificate    string `json:"ottCertificate"`
		ProviderSeriesID  string `json:"providerSeriesId"`
		ReverseOrder      bool   `json:"reverseOrder"`
		SectionNavigation string `json:"sectionNavigation"`
		SeriesUUID        string `json:"seriesUuid"`
		Slug              string `json:"slug"`
		SmartCallToAction string `json:"smartCallToAction"`
		SortTitle         string `json:"sortTitle"`
		SynopsisBrief     string `json:"synopsisBrief"`
		SynopsisLong      string `json:"synopsisLong"`
		SynopsisShort     string `json:"synopsisShort"`
		Title             string `json:"title"`
		TitleMedium       string `json:"titleMedium"`
		TitleMediumSeo    string `json:"titleMediumSeo"`
		TitleSeo          string `json:"titleSeo"`
	} `json:"attributes"`
	Relationships struct {
		Items struct {
			Data []struct {
				Links struct {
					Self string `json:"self"`
				} `json:"links"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				ChildTypes struct {
					Images struct {
						NodeTypes []string `json:"nodeTypes"`
						Count     int      `json:"count"`
					} `json:"images"`
					Items struct {
						NodeTypes []string `json:"nodeTypes"`
						Count     int      `json:"count"`
					} `json:"items"`
				} `json:"childTypes"`
				Attributes struct {
					Badging struct {
						VideoFormats []struct {
							VideoFormat  string   `json:"videoFormat"`
							ColourSpaces []string `json:"colourSpaces"`
						} `json:"videoFormats"`
						AudioTracks []string `json:"audioTracks"`
					} `json:"badging"`
					Channel struct {
						AccessChannel string `json:"accessChannel"`
						Name          string `json:"name"`
					} `json:"channel"`
					ChildNodeTypes       []string `json:"childNodeTypes"`
					Classification       []string `json:"classification"`
					ContentSegments      []string `json:"contentSegments"`
					CreatedDate          int64    `json:"createdDate"`
					DescLongSeo          string   `json:"descLongSeo"`
					DeviceAvailabilities []struct {
						Format             string `json:"format"`
						MediaType          string `json:"mediaType"`
						OfferStage         string `json:"offerStage"`
						OfferStartTs       int64  `json:"offerStartTs"`
						OfferEndTs         int64  `json:"offerEndTs"`
						ContentSegment     string `json:"contentSegment"`
						VideoFormat        string `json:"videoFormat"`
						VideoFormatVariant string `json:"videoFormatVariant"`
						ColourSpace        string `json:"colourSpace"`
					} `json:"deviceAvailabilities"`
					DeviceAvailability struct {
						Available bool `json:"available"`
					} `json:"deviceAvailability"`
					Formats struct {
						HD struct {
							AudioTracks struct {
								Eng []string `json:"eng"`
							} `json:"audioTracks"`
							EventStage   string `json:"eventStage"`
							Availability struct {
								Available        bool   `json:"available"`
								MediaType        string `json:"mediaType"`
								OfferStage       string `json:"offerStage"`
								OfferStartTs     int64  `json:"offerStartTs"`
								OfferEndTs       int64  `json:"offerEndTs"`
								AvailableDevices []struct {
									Type     string `json:"type"`
									Platform string `json:"platform"`
								} `json:"availableDevices"`
								ExtendedOfferStartTs int64  `json:"extendedOfferStartTs"`
								ExtendedOfferEndTs   int64  `json:"extendedOfferEndTs"`
								ContentSegment       string `json:"contentSegment"`
								VideoFormat          string `json:"videoFormat"`
								VideoFormatVariant   string `json:"videoFormatVariant"`
								ColourSpace          string `json:"colourSpace"`
							} `json:"availability"`
							StartOfCredits int `json:"startOfCredits"`
						} `json:"HD"`
					} `json:"formats"`
					GenreList []struct {
						Subgenre []string `json:"subgenre"`
						Genre    []string `json:"genre"`
					} `json:"genreList"`
					Genres []string `json:"genres"`
					Images []struct {
						URL  string `json:"url"`
						Type string `json:"type"`
					} `json:"images"`
					OttCertificate    string `json:"ottCertificate"`
					ProviderSeasonID  string `json:"providerSeasonId"`
					ProviderSeriesID  string `json:"providerSeriesId"`
					SeasonNumber      int    `json:"seasonNumber"`
					SeasonUUID        string `json:"seasonUuid"`
					SectionNavigation string `json:"sectionNavigation"`
					SeriesID          string `json:"seriesId"`
					SeriesName        string `json:"seriesName"`
					Slug              string `json:"slug"`
					SortTitle         string `json:"sortTitle"`
					SynopsisLong      string `json:"synopsisLong"`
					Title             string `json:"title"`
					TitleMedium       string `json:"titleMedium"`
					TitleMediumSeo    string `json:"titleMediumSeo"`
					TitleSeo          string `json:"titleSeo"`
				} `json:"attributes"`
				Relationships struct {
					Items struct {
						Data []struct {
							Links struct {
								Self string `json:"self"`
							} `json:"links"`
							ID         string `json:"id"`
							Type       string `json:"type"`
							ChildTypes struct {
								Images struct {
									NodeTypes []string `json:"nodeTypes"`
									Count     int      `json:"count"`
								} `json:"images"`
							} `json:"childTypes"`
							Attributes struct {
								AudioDescribed bool `json:"audioDescribed"`
								Badging        struct {
									VideoFormats []struct {
										VideoFormat  string   `json:"videoFormat"`
										ColourSpaces []string `json:"colourSpaces"`
									} `json:"videoFormats"`
									AudioTracks []string `json:"audioTracks"`
								} `json:"badging"`
								Cast    []string `json:"cast"`
								Channel struct {
									AccessChannel string `json:"accessChannel"`
									Name          string `json:"name"`
								} `json:"channel"`
								ChildNodeTypes       []string `json:"childNodeTypes"`
								Classification       []string `json:"classification"`
								ClosedCaptioned      bool     `json:"closedCaptioned"`
								ContentSegments      []string `json:"contentSegments"`
								CreatedDate          int64    `json:"createdDate"`
								DeviceAvailabilities []struct {
									Format             string `json:"format"`
									MediaType          string `json:"mediaType"`
									OfferStage         string `json:"offerStage"`
									OfferStartTs       int64  `json:"offerStartTs"`
									OfferEndTs         int64  `json:"offerEndTs"`
									Streamable         bool   `json:"streamable"`
									Downloadable       bool   `json:"downloadable"`
									ContentSegment     string `json:"contentSegment"`
									VideoFormat        string `json:"videoFormat"`
									VideoFormatVariant string `json:"videoFormatVariant"`
									ColourSpace        string `json:"colourSpace"`
								} `json:"deviceAvailabilities"`
								DeviceAvailability struct {
									Available bool `json:"available"`
								} `json:"deviceAvailability"`
								Director             []string `json:"director"`
								DurationMilliseconds int      `json:"durationMilliseconds"`
								DurationMinutes      int      `json:"durationMinutes"`
								DurationSeconds      int      `json:"durationSeconds"`
								EditorialWarningText string   `json:"editorialWarningText"`
								EpisodeName          string   `json:"episodeName"`
								EpisodeNameLong      string   `json:"episodeNameLong"`
								EpisodeNumber        int      `json:"episodeNumber"`
								Formats              struct {
									HD struct {
										AudioTracks struct {
											Eng []string `json:"eng"`
										} `json:"audioTracks"`
										EventStage   string `json:"eventStage"`
										ContentID    string `json:"contentId"`
										Availability struct {
											Available        bool   `json:"available"`
											MediaType        string `json:"mediaType"`
											OfferStage       string `json:"offerStage"`
											OfferStartTs     int64  `json:"offerStartTs"`
											OfferEndTs       int64  `json:"offerEndTs"`
											Streamable       bool   `json:"streamable"`
											Downloadable     bool   `json:"downloadable"`
											AvailableDevices []struct {
												Type     string `json:"type"`
												Platform string `json:"platform"`
											} `json:"availableDevices"`
											ExtendedOfferStartTs int64  `json:"extendedOfferStartTs"`
											ExtendedOfferEndTs   int64  `json:"extendedOfferEndTs"`
											ContentSegment       string `json:"contentSegment"`
											VideoFormat          string `json:"videoFormat"`
											VideoFormatVariant   string `json:"videoFormatVariant"`
											ColourSpace          string `json:"colourSpace"`
										} `json:"availability"`
										StartOfCredits int `json:"startOfCredits"`
										Markers        struct {
											SOCR int `json:"SOCR"`
											SOI  int `json:"SOI"`
											HSI  int `json:"HSI"`
											SPI  int `json:"SPI"`
										} `json:"markers"`
									} `json:"HD"`
								} `json:"formats"`
								GenreList []struct {
									Subgenre []string `json:"subgenre"`
									Genre    []string `json:"genre"`
								} `json:"genreList"`
								Genres            []string `json:"genres"`
								GracenoteID       string   `json:"gracenoteId"`
								GracenoteSeriesID string   `json:"gracenoteSeriesId"`
								Images            []struct {
									URL  string `json:"url"`
									Type string `json:"type"`
								} `json:"images"`
								MerlinAlternateID string   `json:"merlinAlternateId"`
								MerlinID          string   `json:"merlinId"`
								MerlinSeriesID    string   `json:"merlinSeriesId"`
								OttCertificate    string   `json:"ottCertificate"`
								Producer          []string `json:"producer"`
								ProgrammeUUID     string   `json:"programmeUuid"`
								ProviderID        string   `json:"providerId"`
								ProviderSeasonID  string   `json:"providerSeasonId"`
								ProviderSeriesID  string   `json:"providerSeriesId"`
								ProviderVariantID string   `json:"providerVariantId"`
								Runtime           string   `json:"runtime"`
								SeasonID          string   `json:"seasonId"`
								SeasonNumber      int      `json:"seasonNumber"`
								SectionNavigation string   `json:"sectionNavigation"`
								SeriesID          string   `json:"seriesId"`
								SeriesName        string   `json:"seriesName"`
								Slug              string   `json:"slug"`
								SortTitle         string   `json:"sortTitle"`
								Subtitled         bool     `json:"subtitled"`
								Synopsis          string   `json:"synopsis"`
								SynopsisBrief     string   `json:"synopsisBrief"`
								SynopsisLong      string   `json:"synopsisLong"`
								SynopsisShort     string   `json:"synopsisShort"`
								Title             string   `json:"title"`
								TitleLong         string   `json:"titleLong"`
								TitleMedium       string   `json:"titleMedium"`
								Uriid             string   `json:"uriid"`
								Year              int      `json:"year"`
							} `json:"attributes"`
						} `json:"data"`
					} `json:"items"`
				} `json:"relationships"`
			} `json:"data"`
		} `json:"items"`
	} `json:"relationships"`
}

type showMetadata struct {
	Slug   string `json:"slug"`
	Imgref string `json:"imgref"`
}
