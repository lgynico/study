from lxml import etree


html = etree.parse("./test.html", etree.HTMLParser())
result = html.xpath("//*")
print(result)


result = html.xpath("//li")
print(result)
print(result[0])


result = html.xpath("//li/a")
print(result)


result = html.xpath("//ul//a")
print(result)


result = html.xpath("//ul/a")
print(result)


result = html.xpath("//a[@href='link4.html']/../@class")
print(result)


result = html.xpath("//a[@href='link4.html']/parent::*/@class")
print(result)


result = html.xpath("//li[1]/a/text()")
print(result)
result = html.xpath("//li[last()]/a/text()")
print(result)
result = html.xpath("//li[position()<3]/a/text()")
print(result)
result = html.xpath("//li[last()-2]/a/text()")
print(result)


result = html.xpath("//li[1]/ancestor::*")
print(result)
result = html.xpath("//li[1]/ancestor::div")
print(result)
result = html.xpath("//li[1]/attribute::*")
print(result)
result = html.xpath("//li[1]/child::a[@href='link1.html']")
print(result)
result = html.xpath("//li[1]/descendant::span")
print(result)
result = html.xpath("//li[1]/following::*[2]")
print(result)
result = html.xpath("//li[1]/following-sibling::*")
print(result)
