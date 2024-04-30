1. 安装 buf 
```cmd
go install github.com/bufbuild/buf/cmd/buf@v1.15.1
```
2. 如果没有buf.yaml 执行
```cmd 
buf mod init
```
3. 我们可以通过buf build命令检查当前module声明的合法性
```cmd
buf build
```
4. 创建buf.gen.yaml
```cmd
touch buf.gen.yaml
```
5. 生成 buf generate 目录
```cmd
buf generate raw
```
安装 grpc
```yaml
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```