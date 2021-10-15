file2proto, a utility that can convert struct, mysql, JSON and HTTP into proto

### Install
```shell
go get -u github.com/2637309949/file2proto
```

### Syntax
```shell
file2proto
  -i string
        Fully qualified uri path
  -o string
        Protobuf output file. (default "messages.proto")
```

### Example

- from struct
```shell
file2proto -i ./ -o messages.proto
2021/10/14 18:57:10 output file written to /helloworld-service/srv/types/messages.proto
```

- from mysql
```shell
file2proto -i "mysql://myppdb:myppdb@tcp(127.0.0.1:3306)/db_psycinves"
2021/10/14 18:57:10 output file written to /helloworld-service/srv/types/messages.proto
```

- from json
```shell
file2proto -i ./test.json
2021/10/15 23:21:05 output file written to messages.proto
```

- from http
```shell
file2proto -i "https://help.aliyun.com/api/v2/recommend/doc?idList=209974&limit=6&isMobile=false"
2021/10/15 23:16:32 output file written to messages.proto
```