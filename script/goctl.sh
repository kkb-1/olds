#api生成命令
goctl api go -api applet.api -dir ../  --style=goZero --home=../../../template
#rpc生成命令
goctl rpc protoc *.proto --go_out=./ --go-grpc_out=./  --zrpc_out=./ --style=goZero --home=../../../template

#xcode生成pd go代码命令
protoc --go_out=. ./pkg/xcode/types/status.proto

# 配置拉取
go mod init

# 启动服务命令
# 启动user同步mq服务
go run ../apps/user-manage/mq/main.go
# 启动user-manage rpc服务
go run ../apps/user-manage/rpc/usermanage.go
# 启动user rpc服务
go run ../apps/user/rpc/user.go
# 启动applet api服务
go run ../apps/applet/user.go