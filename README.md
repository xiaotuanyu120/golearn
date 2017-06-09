# golearn

---

### 1. gowiki
golang example 01: wiki that can store page in file named by wiki title and return wiki body.

---

### 2. goweb
golang example 02: simple web that return "hello".

---

### 3. gocmdb
golang example 03: simple cmdb implemented with restful api and mysql
https://github.com/motyar/restgomysql/blob/master/server.go
http://go-database-sql.org/retrieving.html
https://github.com/go-sql-driver/mysql/wiki/Examples
https://stackoverflow.com/questions/17779204/how-to-pass-type-into-an-http-handler
https://thenewstack.io/make-a-restful-json-api-go/
https://gist.github.com/andreagrandi/97263aaf7f9344d3ffe6 - 解决获取json post data问题
http://blog.csdn.net/wangjun5159/article/details/47781443 - 解释了为什么post data获取到是空值的问题，类型选错

POST example `http://192.168.33.10:8080/api/server`
```
{  
   "uuid":"395eed99-9e66-469f-af2b-b2da116e",
   "sn":"server-170530001",
   "ip":"192.168.100.111",
   "cpu":"xeon E3",
   "memory":"48GB",
   "disktype":"SSD",
   "disksize":"1024G",
   "nic":"eth0,eth1",
   "manufacturer":"HP",
   "model":"model-test",
   "expiredate":"2017-05-31T00:00:00Z",
   "idc":"HK",
   "comment":"comments sample"
}
```
> postman中选择raw格式，然后再post
