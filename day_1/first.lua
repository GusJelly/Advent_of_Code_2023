local numbers = {
    ["one"] = 1,
    ["two"] = 2,
    ["three"] = 3,
    ["four"] = 4,
    ["five"] = 5,
    ["six"] = 6,
    ["seven"] = 7,
    ["eight"] = 8,
    ["nine"] = 9
}

local getNumber = function(str)
    local result = "";
    for i = 1, #str do
        local char = str:sub(i, i)
        if tonumber(char) ~= nil then
            result = result .. tonumber(char)
        end
    end

    result = result:sub(1, 1) .. result:sub(-1, -1)

    return tonumber(result)
end

local exercise = function()
    local file = io.open("example", "r")
    local result = 0

    if file == nil then
        return
    end

    for line in file:lines() do
        repNrStr(line)
    end

    return result
end

print(exercise())
