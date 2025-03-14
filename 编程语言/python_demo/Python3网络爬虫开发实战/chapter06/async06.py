import asyncio
import requests


async def request():
    url = "https://www.baidu.com"
    status = requests.get(url)
    return status

corutine = request()
task = asyncio.ensure_future(corutine)
print('Task:', task)

loop = asyncio.get_event_loop()
loop.run_until_complete(task)
print('Task:', task)
print('Task Result:', task.result())
