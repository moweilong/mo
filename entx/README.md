# entx

## Go版本要求

entx 要求 Go 版本为 1.25 或更高版本。

## 子目录

- **update**：提供一组用于构建和执行 SQL 更新操作的工具函数，特别适用于处理 NULL 值和 JSON 字段的更新场景。
  [update](update/README.md)

## 测试

执行测试

```bash
go test -timeout 300s -run ^TestMySQL$ github.com/moweilong/mo/entx -gcflags=all=-N -gcflags=all=-l -count=1 -v
```