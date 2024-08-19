local skynet = require "skynet"
local socket = require "skynet.socket"
local s = require "service"
local runconfig = require "runconfig"


-- [fd] -> conn
local conns = {}
-- [playerid] -> gateplayer
local players = {}

local function conn()
    local m = {
        fd = nil,
        playerid = nil,
    }
    return m
end

local function gateplayer()
    local m = {
        playerid = nil,
        conn = nil,
        agent = nil,
    }
    return m
end


local function str_pack(cmd, msg)
    return table.concat(msg, ",") .. "\r\n"
end

local function str_unpack(msgstr)
    local msg = {}
    while true do
        local arg, rest = string.match(msgstr, "(.-),(.*)")
        if arg then
            msgstr = rest
            table.insert(msg, arg)
        else
            table.insert(msg, msgstr)
            break
        end
    end
    return msg[1], msg
end


local function process_msg(fd, msgstr)
    local cmd, msg = str_unpack(msgstr)
    skynet.error("(" .. fd .. ") recv msg [" .. cmd .. "] -> {" .. table.concat(msg, ",") .. "}")

    
    local c = conns[fd]
    local playerid = c.playerid
    if not playerid then
        local node = skynet.getenv("node")
        local nodecfg = runconfig[node]
        local loginid = math.random(1, #nodecfg.login)
        local login = "login" .. loginid
        skynet.send(login, "lua", "client", fd, cmd, msg)
    else
        local player = players[playerid]
        local agent = player.agent
        skynet.send(agent, "lua", "client", cmd, msg)
    end
end


local function process_buff(fd, buff)
    while true do
        local msgstr, rest = string.match(buff, "(.-)\r\n(.*)")
        if msgstr then
            buff = rest
            process_msg(fd, msgstr)
        else
            return buff
        end
    end
end


local function disconnect(fd)
    local c = conns[fd]
    if not c then
        return
    end

    local playerid = c.playerid
    if not playerid then
        return
    end

    players[playerid] = nil
    skynet.call("agentmgr", "lua", "reqkick", playerid, "断线")
end


local function recv_loop(fd)
    socket.start(fd)
    skynet.error("socket connected " .. fd)

    local readbuff = ""
    while true do
        local recvstr = socket.read(fd)
        if recvstr then
            readbuff = readbuff .. recvstr
            readbuff = process_buff(fd, readbuff)
        else
            skynet.error("socket close " .. fd)
            disconnect(fd)
            socket.close(fd)
            return
        end
    end
end


local function connect(fd, addr)
    skynet.error("connect from " .. addr .. " " .. fd)
    local c = conn()
    c.fd = fd
    conns[fd] = c
    skynet.fork(recv_loop, fd)
end


s.init = function ()
    skynet.error("[service start] " .. s.name .. " " .. s.id)

    local node = skynet.getenv("node")
    local nodecfg = runconfig[node]
    local host = "0.0.0.0"
    local port = nodecfg.gateway[s.id].port

    local listenfd = socket.listen(host, port)
    skynet.error("Listen socket: ", host, port)
    socket.start(listenfd, connect)
end


s.resp.send_by_fd = function (source, fd, msg)
    if not conns[fd] then
        return
    end

    local msgstr = str_pack(msg[1], msg)
    skynet.error("(" .. fd .. ") send msg [" .. msg[1] .. "] {" .. table.concat(msg, ",") .. "}")
    socket.write(fd, msgstr)
end

s.resp.send = function (source, playerid, msg)
    local player = players[playerid]
    if not player then
        return
    end

    local c = player.conn
    if not c then
        return
    end

    s.resp.send_by_fd(nil, c.fd, msg)
end

s.resp.sure_agent = function (source, fd, playerid, agent)
    local c = conns[fd]
    if not c then
        skynet.call("agentmgr", "lua", "reqkick", playerid, "未完成登录即下线")
        return false
    end

    local player = gateplayer()
    player.playerid = playerid
    player.conn = c
    player.agent = agent
    
    players[playerid] = player
    
    return true
end

s.resp.kick = function (source, playerid)
    local player = players[playerid]
    if not player then
        return
    end

    players[playerid] = nil

    local c = player.conn
    if not c then
        return
    end

    conns[c.fd] = nil
    disconnect(c.fd)
    socket.close(c.fd)
end

s.start(...)