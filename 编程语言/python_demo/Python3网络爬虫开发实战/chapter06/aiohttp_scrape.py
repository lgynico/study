import asyncio
import aiohttp
import logging
import json
from motor.motor_asyncio import AsyncIOMotorClient

logging.basicConfig(level=logging.INFO,
                    format='%(asctime)s - %(levelname)s - %(message)s')

index_url = "https://spa5.scrape.center/api/book/?limit=18&offset={offset}"
detail_url = "https://spa5.scrape.center/api/book/{id}"

page_size = 18
page_number = 100
concurrency = 5

semaphore = asyncio.Semaphore(concurrency)
session = None

mongo_dsn = "mongodb://localhost:27017"
mongo_db = "python_spider"
mongo_collection = "book"

client = AsyncIOMotorClient(mongo_dsn)
db = client[mongo_db]
collection = db[mongo_collection]


async def save_data(data):
    logging.info('saving data %s', data)
    if data:
        return await collection.update_one({'id': data.get('id')}, {'$set': data}, upsert=True)


async def scrape_detail(id):
    url = detail_url.format(id=id)
    data = await scrape_api(url)
    await save_data(data)


async def scrape_api(url):
    async with semaphore:
        try:
            logging.info('scraping %s', url)
            async with session.get(url) as response:
                return await response.json()
        except aiohttp.ClientError:
            logging.error('error occurred while scraping %s',
                          url, exc_info=True)


async def scrape_index(page):
    url = index_url.format(offset=(page-1) * page_size)
    return await scrape_api(url)


async def main():
    global session
    session = aiohttp.ClientSession()
    tasks = [asyncio.ensure_future(scrape_index(page))
             for page in range(1, page_number + 1)]
    results = await asyncio.gather(*tasks)
    logging.info('results', json.dumps(results, ensure_ascii=False, indent=2))
    ids = []
    for index_data in results:
        for item in index_data.get('results'):
            ids.append(item.get('id'))

    scrape_detail_tasks = [asyncio.ensure_future(
        scrape_detail(id)) for id in ids]
    await asyncio.wait(scrape_detail_tasks)
    await session.close()


if __name__ == '__main__':
    asyncio.run(main())
