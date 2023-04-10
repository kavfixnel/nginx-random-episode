math.randomseed(ngx.now()*1000)

-- Count the number of / in the path
-- service/show   => 1
-- service/show/2 => 2
local _, c = ngx.var.service_series:gsub("/","")

local file = ""
if c == 1 then
  -- Pick a random episode from all episodes
  file = "/var/episodes/" .. ngx.var.service_series .. "/all.episodes"
elseif c == 2 then
  -- Pick a random episode from a single season
  file = "/var/episodes/" .. ngx.var.service_series .. ".episodes"
end

-- Read in the episode file and pick an episode at random
local e = {}
for line in io.lines(file) do
  table.insert(e, line)
end

-- Redirect to episode link
ngx.say(e[math.random(1, table.getn(e))])
