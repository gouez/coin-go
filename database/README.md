# database
    使用 bitcask 存储模型，存储k/v键值对

### 使用

```golang

db, _ := Open("./data.db")
db.Put([]byte("key"), []byte("value"))
db.Get([]byte("key"))
db.Delete([]byte("key"))

```

### 内部结构

#### 日志格式

#### 索引格式