# entx

## Go版本要求

entx 要求 Go 版本为 1.25 或更高版本。

## 子目录

- **update**：提供一组用于构建和执行 SQL 更新操作的工具函数，特别适用于处理 NULL 值和 JSON 字段的更新场景。
  [update](update/README.md)
- **mixin**：提供一组用于扩展 ent 实体的 mixin 实现，包括操作人字段、创建时间字段、更新时间字段、软删除字段等。
  [mixin](mixin/README.md)
- **query**：提供一组用于构建和执行 SQL 查询操作的工具函数，特别适用于处理 JSON 字段的查询场景。
  [query](query/README.md)



## 测试

执行测试

```bash
go test -timeout 300s -run ^TestMySQL$ github.com/moweilong/mo/entx -gcflags=all=-N -gcflags=all=-l -count=1 -v
```