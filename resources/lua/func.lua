local function add(a, b)
    return a + b
end

local function assert(v)
    if not v then fail() end
end

local r = add(1, 2)
assert(r == 3)
