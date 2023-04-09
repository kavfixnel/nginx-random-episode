# Scraping Peacock for show links

There are a couple of steps that are needed to scrape Peacock for episode links

1. Get the path from Peacock for the show you are scraping. On peacock.com on the show
page if the url is everything after (but not including query paraters) `https://www.peacocktv.com/watch/asset`.
So if the full path was `https://www.peacocktv.com/watch/asset/tv/parks-and-recreation/5883799404534408112`,
the `SHOW_PATH` variable should be set to `/tv/parks-and-recreation/5883799404534408112`.
2. Query Peacock's API to find the episode links.
4. Place the resulting list in a file in `episodes/peacock/<show-name>`

```bash
SHOW_PATH="/tv/parks-and-recreation/5883799404534408112" # Parks and Rec's path

# Get all episodes from Parks and Rec's
curl "https://atom.peacocktv.com/adapter-calypso/v3/query/node?slug=$SHOW_PATH&represent=(items(items))&features=upcoming" \
    --header 'x-skyott-proposition: NBCUOTT' \
    --header 'x-skyott-territory: US' | \
    jq -r '.relationships.items.data[].relationships.items.data[].attributes | "https://www.peacocktv.com/watch/playback/vod/" + .formats.HD.contentId + "/" + .providerVariantId' | \
    sort -u
```

Simply place or pipe the shows into a new file `episodes/peacock` folder.