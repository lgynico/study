local skynet = require "skynet"
local s = require "service"

s.client = {}
s.gate = nil

s.init = function ()
    -- TODO: load data from database
    skynet.sleep(200)

    s.data = {
        coin = 100,
        hp = 200,
    }
end


s.resp.client = function (source, cmd, msg)
    s.gate = source

    if not s.client[cmd] then
        skynet.error("s.resp.client fail", cmd)
        return
    end

    local ret_msg = s.client[cmd](msg, source)
    if ret_msg then
        skynet.send(source, "lua", "send", s.id, ret_msg)
    end
end


s.resp.kick = function (source)
    -- TODO: save data to db
    skynet.sleep(200)
end

s.resp.exit = function (source)
    skynet.exit()
end


s.client.work =function (msg)
    s.data.coin = s.data.coin + msg[2]
    return {"work", s.data.coin}
end

s.start(...)