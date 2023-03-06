## Proto生成脚本

# 当前目录
srcDir=$(pwd )
# go文件生成根目录
goOutDir=$srcDir/gen

# 生成go、validate文件
# 将go文件生成在`gen`目录中
protoc --validate_out="lang=go:$goOutDir" --go_out=$goOutDir ./proto/*.proto


