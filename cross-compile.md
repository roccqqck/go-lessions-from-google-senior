# Cross-Compile

show all OS/ARCH
```
go tool dist list
```
```
aix/ppc64
darwin/amd64
darwin/arm64
linux/amd64
linux/arm64
windows/amd64
windows/arm64
```


### go install
```
GOOS=windows GOARCH=amd64 go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

ls ~/go/bin/windows_amd64
```

```$GOPATH/bin```
```
grpcurl.exe
```



### go build
```
git clone https://github.com/fullstorydev/grpcurl
cd grpcurl
GOOS=windows GOARCH=amd64 go build github.com/fullstorydev/grpcurl/cmd/grpcurl

ls ~/go/bin/windows_amd64
```
```
grpcurl.exe
```




