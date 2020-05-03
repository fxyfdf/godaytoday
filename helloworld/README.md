
go build 
go build /dir/project_dir
go build main.go
go build -o name.exe

## linux 
SET CGO_ENABLED=0 //禁用CGO
SET CGOS=linux //目标平台是linux
SET GOARCH=amd64 //目标处理器是adm64

## win 
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build


