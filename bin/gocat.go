package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//Configuration配置文件容器
type Configuration struct {
	Webapps string `xml:"webapps"`
	Port    string `xml:"port"`
}

func main() {
	now := time.Now()
	fmt.Println("正在初始化")
	configuration := &Configuration{} //新建Configuration容器
	fmt.Println("正在读取配置文件")
	bytes, err := ioutil.ReadFile("../conf/configuration.xml") //读取配置文件
	CheckErr(err)
	fmt.Println("正在解析配置文件")
	err = xml.Unmarshal(bytes, configuration) //解析配置文件
	CheckErr(err)
	fmt.Println("正在加载项目")
	server := http.FileServer(http.Dir("../" + configuration.Webapps)) //项目所在目录
	http.Handle("/", server)                                           //路由
	fmt.Printf("所有配置已加载,您的配置如下:\n端口:%s,\n项目目录:%s\n用时:%s\n" +
		"你可以:\n\t--访问locahost:%v/项目名来浏览效果\n\t--访问localhost:%v来查看所有项目目录\n",
		configuration.Port,configuration.Webapps,time.Since(now),configuration.Port,configuration.Port)
	err = http.ListenAndServe(":"+configuration.Port, nil) //端口
	CheckErr(err)


}


//日志
func CheckErr(err error) {
	if err != nil {
		fd, _ := os.OpenFile("../log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		fd_time := time.Now().Format("2006-01-02 15:04:05");
		fd_content := strings.Join([]string{"[", fd_time, "]", err.Error(), "\r\n"}, "")
		buf := []byte(fd_content)
		fd.Write(buf)
		fd.Close()
	}
}
