function newCounter () 
    local count = 0

    return function()
        count = count + 1
        return count
    end
end

c1 = newCounter()
-- 1
print(c1())
-- 2
print(c1())

c2 = newCounter()

-- 1
print(c2())
-- 3
print(c1())
-- 2
print(c2())
