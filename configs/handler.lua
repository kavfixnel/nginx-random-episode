math.randomseed(ngx.now()*1000)

local file = "/var/episodes/" .. ngx.var.service_series

-- Read in the episode file and pick an episode at random
local e = {}
for line in io.lines(file) do
  table.insert(e, line)
end

-- Redirect to episode link
ngx.redirect(e[math.random(1, table.getn(e))])
