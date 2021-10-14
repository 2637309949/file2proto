### Install
```shell
go get -u github.com/2637309949/go2proto
```

### Syntax
```shell
go2proto
  -i string
        Fully qualified path of packages to analyse
  -o string
        Protobuf output file. (default ".")
```

### Example
```shell
go2proto -i ./ -o messages.proto
2021/10/14 18:57:10 output file written to /helloworld-service/srv/types/messages.proto
```
