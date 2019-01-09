## iris + xorm
- iris 在 Golang Web Framework 号称 宇宙最快，堪比甄子丹。以简单，快速上手闻名。
- xorm 是一个简单而强大的Go语言 ORM 库. 通过它可以使数据库操作非常简便。

## 相关、学习文档
- iris(<https://www.studyiris.com/doc/>)
- xorm(<http://www.xorm.io/docs/>)
- go语言圣经(<https://books.studygolang.com/gopl-zh/>)
- go入门指南(<https://www.kancloud.cn/kancloud/the-way-to-go/72432>)

---

### 依赖安装

#### 1.iris(<https://github.com/kataras/iris>)
~~~
go get -u github.com/kataras/iris
~~~

#### 2.xorm(<https://github.com/go-xorm/xorm>)
~~~
go get -u github.com/go-xorm/xorm
~~~

#### 3.xorm工具(<https://github.com/go-xorm/cmd>)
~~~
go get -u github.com/go-xorm/cmd/xorm 
~~~

#### 4.mysql driver(<https://github.com/go-sql-driver/mysql>)
~~~
go get -u github.com/go-sql-driver/mysql
~~~

#### 5.其他
~~~
go get -u github.com/didip/tollbooth
go get -u github.com/iris-contrib/middleware/cors
go get -u github.com/dgrijalva/jwt-go
~~~

**注:**
- 默认情况下安装 cmd/xorm 会出现无法编译的情况(<https://blog.csdn.net/qq_22858601/article/details/81975380>)
- Bcrypt (<https://github.com/golang/crypto>)
   + GOPATH/src 下创建文件夹 golang.org/x/
   + 然后打开golang.org/x/ 
   + git clone https://github.com/golang/crypto.git
- 克隆本项目如果修改文件夹名称记得修改对应import的名称
- 按照写 PHP 的习惯来写的(刚学go不久)

---

###  翻转数据结构生成代码(go-xorm/cmd)
- 项目根目录 templates 是从 GOPATH/src/cmd/xorm/ 目录下拷贝过来的

- 修改 templates/goxorm/config
~~~
  lang=go
  genJson=1
  prefix=你的数据库的表前缀
~~~


- 执行命令 
~~~
  // 格式 xorm reverse mysql 用户名:密码@tcp(HOST:端口)/数据库名?charset=字符串编码 templates/goxorm
  xorm reverse mysql username:password@tcp\(127.0.0.1:3306\)/test_db?charset=utf8 template/goxorm
~~~
**注：有些数据表生成失败，修改表注释的特殊字符`(`如 () ，半角修改成全角即可`)`。**

- 将生成 models 文件夹移到 application 目录下(个人习惯问题)