# Scraping IMDb

## Scraping for Netflix links

There are a couple of steps that are needed to scrape IMDb for Netflix links

1. Get the id in IMDb for the _show_ you are targeting
(Just find the page on https://www.imdb.com/ and copy from url). Should just be of form `tt[0-9]\+`
2. Get the ids of the _episodes_ in a single season
3. Translate said episode ids 

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
  grep -o "www.netflix.com/watch/\d\+" | \
  sed 's|www.netflix.com/watch/||'
```

And when we put it all together:

```bash
for SEASON in {1..5}; do \
  echo "=== $SEASON ==="; \
  curl https://www.imdb.com/watch/_ajax/option -s -X POST --data-urlencode "minibar=$(curl https://www.imdb.com/title/$SHOW/episodes/_ajax\?season\=$SEASON -s | grep -o "/title/tt[0-9]\+" | sed 's|/title/||' | sort -u | tr '\n' ',' | sed 's/,$/\n/')" | jq -r '.minibar[]' | grep -o "www.netflix.com/watch/\d\+" | sed 's|www.netflix.com/watch/||' | tee >(wc -l); \
done
=== 1 ===
70196257
70196258
70196255
70196256
70196253
70196254
70196252
       7
=== 2 ===
70196266
70196265
70196267
70196268
70196259
70196270
70196269
70196260
70196271
70196262
70196261
70196264
70196263
      13
=== 3 ===
70196279
70196277
70196276
70196274
70196273
70196275
70196272
70196284
70196283
70196282
70196281
70196278
70196280
      13
=== 4 ===
70236039
70236038
70236041
70236040
70236047
70236046
70236037
70236036
70236035
70236045
70236044
70236043
70236042
      13
=== 5 ===
70236428
70236413
70236414
70236415
70236416
70236417
70236418
70236419
70236421
70236422
70236423
70236424
70236425
70236426
70236427
70236412
      16
```