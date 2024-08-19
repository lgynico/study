local skynet = require "skynet"
local s = require "service"


s.client = {}

s.resp.client = function (source, fd, cmd, msg)
    if s.client[cmd] then
        local ret_msg = s.client[cmd](fd, msg, source)
        skynet.send(source, "lua", "send_by_fd", fd, ret_msg)
        return
    end

    skynet.error("s.resp.client fail", cmd)
end


s.client.login = function (fd, msg, source)
    skynet.error("login recv " .. msg[1] .. " " .. msg[2])
    
    local playerid = tonumber(msg[2])
    local password = tonumber(msg[3])
    local gateway = source
    local node = skynet.getenv("node")

    if password ~= 123 then
        return {"login", 1, "密码错误"}
    end

    local ok, agent = skynet.call("agentmgr", "lua", "reqlogin", playerid, node, gateway)
    if not ok then
        return {"login", 1, "请求 mgr 失败"}
    end

    local ok = skynet.call(gateway, "lua", "sure_agent", fd, playerid)
    if not ok then
        return {"login", 1, "gate 注册失败"}
    end

    skynet.error("login succ " .. playerid)
    return {"login", 0, "登录成功"}
end

s.start(...)