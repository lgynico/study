from pyquery import PyQuery


html = """
<div class="wrap">
    <div>
        <ul>
            <li class="item-0">first item</li>
            <li class="item-1"><a href="link2.html">second item</a></li>
            <li class="item-0 active"><a href="link3.html"><span class="bold">third item</span></a></li>
            <li class="item-1 active"><a href="link4.html">fourth item</a></li>
            <li class="item-0"><a href="link5.html">fifth item</a></li>
        </ul>
    </div>
</div>
"""
doc = PyQuery(html)
li = doc(".item-0.active")
print(li)
print(str(li))


lis = doc("li").items()
print(type(lis))
for li in lis:
    print(li, type(li))


a = doc(".item-0.active a")
print(a, type(a))
print(a.attr("href"))
print(a.attr.href)


a = doc("a")
print(a, type(a))
print(a.attr("href"))
print(a.attr.href)

for item in a.items():
    print(item.attr("href"))


a = doc(".item-0.active a")
print(a)
print(a.text())


li = doc(".item-0.active")
print(li)
print(li.html())


li = doc("li")
print(li.html())
print(li.text())
print(type(li.text()))
