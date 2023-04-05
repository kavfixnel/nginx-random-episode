# Scraping IMDb

## Scraping for Netflix links

There are a couple of steps that are needed to scrape IMDb for Netflix links

1. Get the id in IMDb for the _show_ you are targeting
(Just find the page on https://www.imdb.com/ and copy from url). Should just be of form `tt[0-9]\+`
2. Get the ids of the _episodes_ in a single season
3. Translate said episode ids into links
4. Place the resulting list in a file in `episodes/netflix/<show-name>`

```bash
SHOW=tt0903747 # Breaking Bad
SEASON=1

# Get all episodes from Breaking Bad
curl https://www.imdb.com/title/$SHOW/episodes/_ajax\?season\=$SEASON -s | \
  grep -o "/title/tt[0-9]\+" | \
  sed 's|/title/||' | \
  sort -u | \
  tr '\n' ',' | sed 's/,$/\n/'

# Get the Netflix ids of those episodes
curl https://www.imdb.com/watch/_ajax/option -s -X POST \
  --data-urlencode "minibar=tt0959621,tt1054724,tt1054725,tt1054726,tt1054727,tt1054728,tt1054729" | \
  jq -r '.minibar[]' | \
  grep -o "www.netflix.com/watch/\d\+"
```

And when we put it all together:

```bash
for SEASON in {1..5}; do \
  curl https://www.imdb.com/watch/_ajax/option -s -X POST --data-urlencode "minibar=$(curl https://www.imdb.com/title/$SHOW/episodes/_ajax\?season\=$SEASON -s | grep -o "/title/tt[0-9]\+" | sed 's|/title/||' | sort -u | tr '\n' ',' | sed 's/,$/\n/')" | jq -r '.minibar[]' | grep -o "https\://www.netflix.com/watch/\d\+"; \
done
https://www.netflix.com/watch/70196257
https://www.netflix.com/watch/70196258
https://www.netflix.com/watch/70196255
https://www.netflix.com/watch/70196256
https://www.netflix.com/watch/70196253
https://www.netflix.com/watch/70196254
https://www.netflix.com/watch/70196252
https://www.netflix.com/watch/70196261
https://www.netflix.com/watch/70196260
https://www.netflix.com/watch/70196271
https://www.netflix.com/watch/70196270
https://www.netflix.com/watch/70196269
https://www.netflix.com/watch/70196268
https://www.netflix.com/watch/70196259
https://www.netflix.com/watch/70196267
https://www.netflix.com/watch/70196266
https://www.netflix.com/watch/70196265
https://www.netflix.com/watch/70196264
https://www.netflix.com/watch/70196263
https://www.netflix.com/watch/70196262
https://www.netflix.com/watch/70196279
https://www.netflix.com/watch/70196277
https://www.netflix.com/watch/70196276
https://www.netflix.com/watch/70196274
https://www.netflix.com/watch/70196273
https://www.netflix.com/watch/70196275
https://www.netflix.com/watch/70196272
https://www.netflix.com/watch/70196284
https://www.netflix.com/watch/70196283
https://www.netflix.com/watch/70196282
https://www.netflix.com/watch/70196281
https://www.netflix.com/watch/70196278
https://www.netflix.com/watch/70196280
https://www.netflix.com/watch/70236036
https://www.netflix.com/watch/70236037
https://www.netflix.com/watch/70236046
https://www.netflix.com/watch/70236047
https://www.netflix.com/watch/70236044
https://www.netflix.com/watch/70236045
https://www.netflix.com/watch/70236035
https://www.netflix.com/watch/70236042
https://www.netflix.com/watch/70236043
https://www.netflix.com/watch/70236040
https://www.netflix.com/watch/70236041
https://www.netflix.com/watch/70236038
https://www.netflix.com/watch/70236039
https://www.netflix.com/watch/70236428
https://www.netflix.com/watch/70236413
https://www.netflix.com/watch/70236414
https://www.netflix.com/watch/70236415
https://www.netflix.com/watch/70236416
https://www.netflix.com/watch/70236417
https://www.netflix.com/watch/70236418
https://www.netflix.com/watch/70236419
https://www.netflix.com/watch/70236421
https://www.netflix.com/watch/70236422
https://www.netflix.com/watch/70236423
https://www.netflix.com/watch/70236424
https://www.netflix.com/watch/70236425
https://www.netflix.com/watch/70236426
https://www.netflix.com/watch/70236427
https://www.netflix.com/watch/70236412
```