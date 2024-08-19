local skynet = require "skynet"
local cluster = require "skynet.cluster"

local M = {
    name = "",
    id = 0,

    init = nil,
    exit = nil,

    resp = {},
}

local function traceback(err)
    skynet.error(tostring(err))
    skynet.error(debug.traceback())
end

local function dispatch(session, address, cmd, ...)
    local f = M.resp[cmd]
    if not f then
        skynet.ret()
        return
    end

    local ret = table.pack(xpcall(f, traceback, address, ...))
    local ok = ret[1]

    if not ok then
        skynet.ret()
        return
    end

    skynet.retpack(table.unpack(ret, 2))
end

local function init()
    skynet.dispatch("lua", dispatch)
    if M.init then
        M.init()
    end
end


function M.start(name, id, ...)
    M.name = name
    M.id = tonumber(id)

    skynet.start(init)
end


function M.call(node, srv, ...)
    local mynode = skynet.getenv("node")
    if node == mynode then
        return skynet.call(srv, "lua", ...)
    end

    return cluster.call(node, srv, ...)
end


function M.send(node, srv, ...)
    local mynode = skynet.getenv("node")
    if node == mynode then
        return skynet.send(srv, "lua", ...)
    end

    return cluster.send(node, srv, ...)
end

return M