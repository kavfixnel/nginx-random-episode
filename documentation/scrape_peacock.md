# Scraping Peacock for show links

There are a couple of steps that are needed to scrape Peacock for episode links

1. Get the path from Peacock for the show you are scraping. On peacock.com on the show
page if the url is everything after (but not including query paraters) `https://www.peacocktv.com/watch/asset`.
So if the full path was `https://www.peacocktv.com/watch/asset/tv/parks-and-recreation/5883799404534408112`,
the `SHOW_PATH` variable should be set to `/tv/parks-and-recreation/5883799404534408112`.
2. Query Peacock's API to find the episode links.
4. Place the resulting list in a file in `episodes/peacock/<show-name>`

## Get all episodes

In the `scraper/` subdirectory run the following command

```bash
go run main.go peacock \
    --showPath /tv/that-70s-show/8674399279598141112 \
    --showShort the70sshow \
    --showImg https://m.media-amazon.com/images/M/MV5BN2RkZGE0MjAtZGVkOS00MzVhLTg0OWItZTc4OGRjOTQ1ZTM4XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_.jpg \
    --showSlug "The '70s Show"
```
