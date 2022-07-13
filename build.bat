SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -gcflags "all=-N -l" -o ./order ./main.go