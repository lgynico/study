import requests
import logging
import pymongo


logging.basicConfig(level=logging.INFO,
                    format='%(asctime)s - %(levelname)s: %(message)s')

index_url = "https://spa1.scrape.center/api/movie/?limit={limit}&offset={offset}"
detail_url = "https://spa1.scrape.center/api/movie/{id}"
limit = 10

mongo_dsn = "mongodb://localhost:27017"
mongo_db = "python_spider"
mongo_collection = "movies"
client = pymongo.MongoClient(mongo_dsn)
db = client[mongo_db]
collection = db[mongo_collection]


def scrape_api(url):
    logging.info('scraping %s...', url)
    try:
        response = requests.get(url)
        if response.status_code == 200:
            return response.json()
        logging.error('get invalid status code %s while scraping %s',
                      response.status_code, url)
    except requests.RequestException:
        logging.error('error occurred while scraping %s', url, exc_info=True)


def scrape_index(page):
    url = index_url.format(limit=limit, offset=limit * (page - 1))
    return scrape_api(url)


def scrape_detail(id):
    url = detail_url.format(id=id)
    return scrape_api(url)


def save_data(data):
    collection.update_one({"name": data.get("name")},
                          {"$set": data}, upsert=True)


def main():
    for page in range(1, 4):
        index_data = scrape_index(page)
        for item in index_data.get('results'):
            data = scrape_detail(item.get('id'))
            # logging.info('get detail data %s', data)
            save_data(data)
            logging.info('data %s saved', item.get('id'))


if __name__ == '__main__':
    main()
