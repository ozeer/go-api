# 测试命令
1、测试全部
```
go test ./tests -v
```

2、注册测试
```
go test -v -run TestRegister -timeout 7200s
```

# 利用wrk简单压力测试
1、命令
```
# 利用 wrk 对http://127.0.0.1:8888/base/fakePhone接口发起压力测试，线程数为 12，模拟 400 个并发请求，持续 30 秒
wrk -t12 -c400 -d30s --latency http://127.0.0.1:8888/base/fakePhone
```
引用文档：https://www.cnblogs.com/quanxiaoha/p/10661650.html