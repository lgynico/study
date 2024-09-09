from pyquery import PyQuery


html = """
<div id="container">
    <ul class="list">
        <li class="item-0">first item</li>
        <li class="item-1"><a href="link2.html">second item</a></li>
        <li class="item-0 active"><a href="link3.html"><span class="bold">third item</span></a></li>
        <li class="item-1 active"><a href="link4.html">fourth item</a></li>
        <li class="item-0"><a href="link5.html">fifth item</a></li>
    </ul>
</div>
"""
doc = PyQuery(html)
print(doc("#container .list li"))
print(type(doc("#container .list li")))


for item in doc("#container .list li").items():
    print(item.text())


items = doc(".list")
print(type(items))
print(items)

lis = items.find("li")
print(type(lis))
print(lis)

lis = items.children()
print(type(lis))
print(lis)

lis = items.children(".active")
print(lis)


html = """
<div class="wrap">
    <div id="container">
        <ul class="list">
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
items = doc(".list")
container = items.parent()
print(type(container))
print(container)


parents = items.parents()
print(type(parents))
print(parents)

parent = items.parents(".wrap")
print(parent)


li = doc(".list .item-0.active")
print(li.siblings())
print(li.siblings(".active"))


li = doc("li:first-child")
print(li)
li = doc("li:last-child")
print(li)
li = doc("li:nth-child(2)")
print(li)
li = doc("li:gt(2)")
print(li)
li = doc("li:nth-child(2n)")
print(li)
li = doc("li:contains(second)")
print(li)
