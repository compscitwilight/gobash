local math = require("math")
local poll_time = 1000
local polling = true
print("Searching for Equestrian magic sources..")
while polling do
    poll_time = poll_time - math.random(0, 1)
    if poll_time == 0 then
        polling = false
    end
end