# Unity 学习笔记

## 视图
### Project 视图
![Project View](<Unity 学习笔记/ProjectView.png>)
|编号|描述|备注|
|-:|:-|:-|
|1|文件列表||
|2|新建文件||
|3|当前文件夹内容||
|4|文件搜索框||
|5|按类型搜索||
|6|按标签搜索||
|7|按日志搜索|可以跟 `AssetPostprocessor` 配合搜索错误和警告的资源|
|8|保存当前搜索||
|9|Packages 的数量|点击可以隐藏 Packages|
|10| ？？？||
```CSharp
public class MyAsset : AssetPostprocessor
{
    void OnPostprocessTexture(Texture2D texture)
    {
        // context.LogImportError("错误");
        // context.LogImportWarning("警告");
    }
}
```

### Hierarchy 视图
![Hierarchy View](<Unity 学习笔记/HierarchyView.png>)
|编号|描述|备注|
|-:|:-|:-|
|1|场景游戏对象列表||
|2|新建 GameObject||
|3|搜索 GameObject||
|4|隐藏 GameObject||
|5|???||

### Inspector 视图
![Inspector View](<Unity 学习笔记/InspectorView.png>)
|编号|描述|备注|
|-:|:-|:-|
|1|显示/隐藏资源||
|2|资源名称||
|3|???||
|4|选择资源标签||
|5|选择资源层级||
|6|资源的组件列表||
|7|给资源增加组件||

### Scene 视图
![Scene View](<Unity 学习笔记/SceneView.png>)

### Game 视图
![Game View](<Unity 学习笔记/GameView.png>)

## 文件夹

## 脚本


## UI