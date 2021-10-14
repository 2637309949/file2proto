### Install
```shell
go get -u github.com/2637309949/file2proto
```

### Syntax
```shell
file2proto
  -i string
        Fully qualified path of packages to analyse
  -o string
        Protobuf output file. (default ".")
```

### Example
```shell
file2proto -i ./ -o messages.proto
2021/10/14 18:57:10 output file written to /helloworld-service/srv/types/messages.proto
```
