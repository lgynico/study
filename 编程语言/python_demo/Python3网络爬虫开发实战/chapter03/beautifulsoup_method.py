from bs4 import BeautifulSoup
import re


html = """
<div class="panel">
    <div class="panel-heading">
        <h4>Hello</h4>
    </div>
    <div class="panel-body">
        <ul class="list" id="list-1">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
            <li class="element">Jay</li>
        </ul>
        <ul class="list list-small" id="list-2">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
        </ul>
    </div>
</div>
"""
soup = BeautifulSoup(html, "lxml")
print(soup.find_all(name="ul"))
print(type(soup.find_all(name="ul")[0]))

for ul in soup.find_all(name="ul"):
    print(ul.find_all(name="li"))
    for li in ul.find_all(name="li"):
        print(li.string)


html = """
<div class="panel">
    <div class="panel-heading">
        <h4>Hello</h4>
    </div>
    <div class="panel-body">
        <ul class="list" id="list-1" name="elements">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
            <li class="element">Jay</li>
        </ul>
        <ul class="list list-small" id="list-2">
            <li class="element">Foo</li>
            <li class="element">Bar</li>
        </ul>
    </div>
</div>
"""
soup = BeautifulSoup(html, "lxml")
print(soup.find_all(attrs={"id": "list-1"}))
print(soup.find_all(attrs={"name": "elements"}))


print(soup.find_all(id="list-1"))
print(soup.find_all(class_="element"))


print(soup.find(name="ul"))
print(type(soup.find(name="ul")))
print(soup.find(class_="list"))


html = """
<div class="panel">
    <div class="panel-body">
        <a>Hello, this is a link</a>
        <a>Hello, this is a link, too</a>
    </div>
</div>
"""
soup = BeautifulSoup(html, "lxml")
# print(soup.find_all(text=re.compile("link")))
print(soup.find_all(string=re.compile("link")))
