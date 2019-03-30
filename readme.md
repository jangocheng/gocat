# GoCat -- 基于Go的高并发Web容器

------

欢迎使用**GoCat**,这是一款基于**Go**语言构建的Web容器(类Tomcat)，**GoCat**拥有以下特点：
<br>
> * 配置简单
> * 速度快 & 高并发
<br><br>
### 配置简单

目前版本仅支持静态HTML文件,配置文件-**configuration.xml**内也仅需要配置两个属性:
```XML
    <webapps>/webapps</webapps>
    <!--项目所在目录-->
    <port>8080</port>
    <!--端口-->
```
### 速度快 & 高并发
得益于Go原生的高并发特性,你使用GoCat时几乎不用考虑并发的问题

<code>注意:现在**GoCat**还处于测试版本,请谨慎选择使用GoCat部署项目</code>
  
------
  
## 如何使用?

> 1.在本repo下载GoCat的Go源码,切换至\$GOCAT_PATH\$/bin目录下编译gocat.go为可执行文件,将你的web项目放入webapps文件夹下,运行编译后的gocat.exe
>2.在[官网](http://www.teststudy.cn)上下载编译好的GoCat二进制文件,将web项目放入webapps文件夹下,运行bin目录下的gocat.exe
<pre>Tips:如果你想改变端口和项目目录,你可以自行配置configuration.xml</pre>
  
------
  
## 如何参与这个项目?

* 本软件采用MIT开源协议,欢迎Fork后自行修改,由此所带来的法律效应与本人无关.
* 你也可以fork后Pull request,我会酌情考虑合并
* 发现Bug后可以在Issues进行反馈或直接私聊联系我
  
  
作者 [赵不贪](http://www.teststudy.cn)    
2019 年 03月 30日
