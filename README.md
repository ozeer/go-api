## 编译执行
```
go build -o bin/server main.go
./bin/server
```

## 性能分析
### 一、CPU分析
1、获取最近10秒程序运行的CPU profile，-seconds参数不填默认为30。
```
go tool pprof -seconds 10 http://127.0.0.1:8888/debug/pprof/profile
```
2、运行编写好的请求服务的脚本
```
cd tests
go test -v -run TestRegister -timeout 7200s
```
3、可视化展示
```
go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz
```
### 二、内存分析
1、获取程序运行的内存profile
```
go tool pprof http://127.0.0.1:8888/debug/pprof/heap
```
2、可视化展示
```
go tool pprof -http=:8081 ~/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
```

## 分布式唯一id
- https://github.com/sony/sonyflake
- https://sqids.org/?hashids