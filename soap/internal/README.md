# internal

## encoding/xml

解决encode/decode时element prefix的影响!

- 基于go 1.19.9源码
- 添加local()去掉prefix
- 修改read.go里面涉及"xxx.Local"的比较处. 