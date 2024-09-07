from lxml import etree

html = etree.parse("./test.html", etree.HTMLParser())


result = html.xpath("//li[@class='item-0']")
print(result)


result = html.xpath("//li/a/@href")
print(result)


text = """
<li class="li li-first" name="item"><a href="link.html">first item</a></li>
"""
html2 = etree.HTML(text)

result = html2.xpath("//li[contains(@class, 'li')]/a/text()")
print(result)


result = html2.xpath("//li[contains(@class, 'li') and @name='item']/a/text()")
print(result)
