from bs4 import BeautifulSoup


html = """
<html>
    <head>
        <title>The Dormouse's story</title>
    </head>
    <body>
        <p class="story">
            Once upon a time there were three little sisters; and their names were
            <a href="http://example.com/elsie" class="sister" id="link1">
                <span>Elsie</span>
            </a>
            <a href="http://example.com/lacie" class="sister" id="link2">Lacie</a>
            and
            <a href_"http://example.com/tillie" class="sister" id="link3">Tillie</a>
            and they lived at the bottom of a well.
        </p>
        <p class="story">...</p>
"""

soup = BeautifulSoup(html, "lxml")
print(soup.p.contents)


print(soup.p.children)
for i, child in enumerate(soup.p.children):
    print(i, child)


print(soup.p.descendants)
for i, child in enumerate(soup.p.descendants):
    print(i, child)


html = """
<html>
    <head>
        <title>The Dormouse's story</title>
    </head>
    <body>
        <p class="story">
            Once upon a time there were three little sisters; and their names were
            <a href="http://example.com/elsie" class="sister" id="link1">
                <span>Elsie</span>
            </a>
        </p>
        <p class="story">...</p>
"""
soup = BeautifulSoup(html, "lxml")
print(soup.p.parent)


html = """
<html>
    <body>
        <p class="story">
            <a href="http://example.com/elsie" class="sister" id="link1">
                <span>Elsie</span>
            </a>
        </p>
"""
soup = BeautifulSoup(html, "lxml")
print(type(soup.a.parents))
print(list(enumerate(soup.a.parents)))


html = """
<html>
    <body>
        <p class="story">
            Once upon a time there were three little sisters; and their names were
            <a href="http://example.com/elsie" class="sister" id="link1">
                <span>Elsie</span>
            </a>
            <a href="http://example.com/lacie" class="sister" id="link2">Lacie</a>
            and
            <a href_"http://example.com/tillie" class="sister" id="link3">Tillie</a>
            and they lived at the bottom of a well.
        </p>
"""
soup = BeautifulSoup(html, "lxml")
print("Next Sibling", soup.a.next_sibling)
print("Prev Sibling", soup.a.previous_sibling)
print("Next Siblings", list(enumerate(soup.a.next_siblings)))
print("Prev Siblings", list(enumerate(soup.a.previous_siblings)))


html = """
<html>
    <body>
        <p class="story">
            Once upon a time there were three little sisters; and their names were
            <a href="http://example.com/elsie" class="sister" id="link1">Bob</a><a href=
                    "http://example.com/lacie" class="sister" id="link2">Lacie</a>
        </p>
"""
soup = BeautifulSoup(html, "lxml")
print("Next Sibling:")
print(type(soup.a.next_sibling))
print(soup.a.next_sibling)
print(soup.a.next_sibling.string)

print("parent:")
print(type(soup.a.parents))
print(list(soup.a.parents)[0])
print(list(soup.a.parents)[0].attrs["class"])
