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
for SEASON in {1..5}; do \                                                             took 4s
  curl https://www.imdb.com/watch/_ajax/option -s -X POST --data-urlencode "minibar=$(curl https://www.imdb.com/title/$SHOW/episodes/_ajax\?season\=$SEASON -s | grep -o "/title/tt[0-9]\+" | sed 's|/title/||' | sort -u | tr '\n' ',' | sed 's/,$/\n/')" | jq -r '.minibar[]' | grep -o "www.netflix.com/watch/\d\+"; \
done
www.netflix.com/watch/70196257
www.netflix.com/watch/70196258
www.netflix.com/watch/70196255
www.netflix.com/watch/70196256
www.netflix.com/watch/70196253
www.netflix.com/watch/70196254
www.netflix.com/watch/70196252
www.netflix.com/watch/70196261
www.netflix.com/watch/70196260
www.netflix.com/watch/70196271
www.netflix.com/watch/70196270
www.netflix.com/watch/70196269
www.netflix.com/watch/70196268
www.netflix.com/watch/70196259
www.netflix.com/watch/70196267
www.netflix.com/watch/70196266
www.netflix.com/watch/70196265
www.netflix.com/watch/70196264
www.netflix.com/watch/70196263
www.netflix.com/watch/70196262
www.netflix.com/watch/70196279
www.netflix.com/watch/70196277
www.netflix.com/watch/70196276
www.netflix.com/watch/70196274
www.netflix.com/watch/70196273
www.netflix.com/watch/70196275
www.netflix.com/watch/70196272
www.netflix.com/watch/70196284
www.netflix.com/watch/70196283
www.netflix.com/watch/70196282
www.netflix.com/watch/70196281
www.netflix.com/watch/70196278
www.netflix.com/watch/70196280
www.netflix.com/watch/70236036
www.netflix.com/watch/70236037
www.netflix.com/watch/70236046
www.netflix.com/watch/70236047
www.netflix.com/watch/70236044
www.netflix.com/watch/70236045
www.netflix.com/watch/70236035
www.netflix.com/watch/70236042
www.netflix.com/watch/70236043
www.netflix.com/watch/70236040
www.netflix.com/watch/70236041
www.netflix.com/watch/70236038
www.netflix.com/watch/70236039
www.netflix.com/watch/70236428
www.netflix.com/watch/70236413
www.netflix.com/watch/70236414
www.netflix.com/watch/70236415
www.netflix.com/watch/70236416
www.netflix.com/watch/70236417
www.netflix.com/watch/70236418
www.netflix.com/watch/70236419
www.netflix.com/watch/70236421
www.netflix.com/watch/70236422
www.netflix.com/watch/70236423
www.netflix.com/watch/70236424
www.netflix.com/watch/70236425
www.netflix.com/watch/70236426
www.netflix.com/watch/70236427
www.netflix.com/watch/70236412
```