#api生成命令
goctl api go -api applet.api -dir ../  --style=goZero --home=../../../template
#rpc生成命令
goctl rpc protoc *.proto --go_out=./ --go-grpc_out=./  --zrpc_out=./ --style=goZero --home=../../../template

#xcode生成pd go代码命令
protoc --go_out=. ./pkg/xcode/types/status.proto