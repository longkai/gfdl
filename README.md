# gfdl
Google Fonts Downloader

### 网页使用了 Google Fonts 但是被墙了...
我的[个人网站][1]上使用了 Google fonts，这些字体效果非常棒。然而忽略了一个事实：大多数人都没法科学上网。好在今天一个朋友说网站一直在转半天打不开，于是打开浏览器开发者工具一看，发现 Google fonts 被墙，加载超时了...

Google 了一下解决办法，360做了一个镜像，但是看到知乎网友评论时不时抽风，最关键的是，全是 http 的！还是着手自己写一个下载器吧。

以上。


### 用法
```shell
# gfdl src [dest]
# `src`: google fonts css url; `desc`: optional, where to put the download css file and its fonts.

$ gfdl https://example.org/fonts.css path/to/fonts.css
```

### 安装
```shell
# go get github.com/longkai/gfdl
```

或者[戳这里下载][2]

### License
MIT

[1]: https://xiaolongtongxue.com
[2]: http://dl.xiaolongtongxue.com/gfdl/
