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
li.removeClass("active")
print(li)
li.addClass("active")
print(li)


html = """
<ul class="list">
    <li class="item-0 active"><a href="link3.html"><span class="bold">third item</span></a></li>
</ul>
"""
doc = PyQuery(html)
li = doc(".item-0.active")
print(li)
li.attr("name", "link")
print(li)
li.text("changed item")
print(li)
li.html("<span>changed item</span>")
print(li)


html = """
<div class="wrap">
    Hello, World
    <p>This is paragraph.</p>
</div>
"""
doc = PyQuery(html)
wrap = doc(".wrap")
print(wrap.text())

wrap.find("p").remove()
print(wrap.text())
