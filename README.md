# thrift-web-go


thrift version = 0.11.0

### Use thrift
thrift --gen go *.thrift
thrift --gen js *.thrift
thrift --gen js:node *.thrift

### RUN go server

```bash
dep ensure
go run main.go
```

## RUN vue client
```bash
yarn 
yarn serve
```