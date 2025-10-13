# entx

## Go版本要求

entx 要求 Go 版本为 1.25 或更高版本。

## 测试

执行测试

```bash
go test -timeout 300s -run ^TestMySQL$ github.com/moweilong/mo/entx -gcflags=all=-N -gcflags=all=-l -count=1 -v
```