### Install
```shell
go get -u github.com/2637309949/file2proto
```

### Syntax
```shell
file2proto
  -i string
        Fully qualified path of uri
  -o string
        Protobuf output file. (default ".")
```

### Example

```shell
file2proto -i ./ -o messages.proto
2021/10/14 18:57:10 output file written to /helloworld-service/srv/types/messages.proto
```


```shell
file2proto -i "mysql://myppdb:myppdb@tcp(127.0.0.1:3306)/db_psycinves"
2021/10/14 18:57:10 output file written to /helloworld-service/srv/types/messages.proto
```
