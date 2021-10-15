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

```shell
file2proto -i ./ -o messages.proto
2021/10/14 18:57:10 output file written to /helloworld-service/srv/types/messages.proto
```

```shell
file2proto -i "mysql://myppdb:myppdb@tcp(127.0.0.1:3306)/db_psycinves"
2021/10/14 18:57:10 output file written to /helloworld-service/srv/types/messages.proto
```

```shell
file2proto -i ./test.json
2021/10/15 23:21:05 output file written to messages.proto
```

```shell
file2proto -i "https://silkroad.csdn.net/api/v2/assemble/list/channel/search_hot_word?channel_name=pc_hot_word&size=10&user_name=u013571243&platform=pc&imei=10_19003314510-1632359899251-804293&toolbarSearchExt=%7B%22landingWord%22%3A%5B%5D%2C%22queryWord%22%3A%22%22%2C%22tag%22%3A%5B%5D%2C%22title%22%3A%22Git%E5%86%B2%E7%AA%81%EF%BC%9Acommit+your+changes+or+stash+them+before+you+can+merge.%22%7D"
2021/10/15 23:16:32 output file written to messages.proto
```