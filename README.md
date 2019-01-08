## iris + xorm
- iris 在 Golang Web Framework 号称 宇宙最快，堪比甄子丹。以简单，快速上手闻名。
- xorm 是一个简单而强大的Go语言 ORM 库. 通过它可以使数据库操作非常简便。

#### 依赖安装

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

**踩过的坑:**
- 默认情况下安装cmd/xorm 会出现无法编译的情况(<https://blog.csdn.net/qq_22858601/article/details/81975380>)
