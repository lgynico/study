local skynet = require "skynet"
local runconfig = require "runconfig"

skynet.start(function ()
    
    skynet.error("[start main] " .. runconfig.agentmgr.node)

    skynet.newservice("debug_console", "8000")

    skynet.newservice("gateway", "gateway", 1)

    skynet.newservice("login", "login", 1)
    skynet.newservice("login", "login", 2)

    skynet.newservice("agentmgr", "agentmgr", 0)

    skynet.newservice("nodemgr", "nodemgr", 0)

    skynet.exit()

end)