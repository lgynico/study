local skynet = require "skynet"
local s = require "service"

local STATUS = {
    LOGIN = 2, GAMING = 3, LOGOUT = 4
}

-- [playerid] -> mgrplayer
local players = {}

local function mgrplayer()
    local m = {
        playerid = nil,
        node = nil,
        agent = nil,
        status = nil,
        gate = nil,
    }

    return m
end


s.resp.reqlogin = function (source, playerid, node, gate)
    local player = players[playerid]

    if player and player.status ~= STATUS.GAMING then
        skynet.error("reqlogin fail: " .. playerid .. " " .. player.status)
        return false
    end

    if player then
        local oldagent = player.agent
        local oldnode = player.node
        s.call(oldnode, oldagent, "kick")
        s.send(oldnode, oldagent, "exit")
        s.send(oldnode, player.gate, "send", playerid, {"kick", "顶替下线"})
        s.call(oldnode, player.gate, "kick", playerid)
    end

    player = mgrplayer()
    player.playerid = playerid
    player.status = STATUS.LOGIN

    players[playerid] = player

    player.gate = gate
    player.node = node

    local agent = s.call(node, "nodemgr", "newservice", "agent", "agent", playerid)
    player.agent = agent
    player.status = STATUS.GAMING

    return true, agent
end


s.resp.reqkick = function (source, playerid, reason)
    
    local player = players[playerid]
    if not player then
        return false
    end

    player.status = STATUS.LOGOUT

    local node = player.node
    local agent = player.agent
    local gate = player.gate
    
    s.call(node, agent, "kick")
    s.send(node, agent, "exit")
    s.send(node, gate, "kick", playerid)

    player[playerid] = nil
    return true
end

s.start(...)