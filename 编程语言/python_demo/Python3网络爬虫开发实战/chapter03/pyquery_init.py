from pyquery import PyQuery


html = """
<div>
    <ul>
        <li class="item-0">first item</li>
        <li class="item-1"><a href="link2.html">second item</a></li>
        <li class="item-0 active"><a href="link3.html"><span class="bold">third item</span></a></li>
        <li class="item-1 active"><a href="link4.html">fourth item</a></li>
        <li class="item-0"><a href="link5.html">fifth item</a></li>
    </ul>
</div>
"""
doc = PyQuery(html)
print(doc("li"))


doc = PyQuery(url="https//cuiqingcai.com")
# doc = PyQuery(requests.get("https//cuiqingcai.com").text)
print(doc("title"))


doc = PyQuery(filename="test.html")
print(doc("li"))
