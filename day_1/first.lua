local numbers = {
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine"
}

local repNrStr = function(str)
    for i = 1, #str do
    end
end

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
        result = result + getNumber(line)
    end

    return result
end

print(exercise())
