# Core

- 提供核心功能，如配置管理、Copier 功能封装、Gin 请求统一处理等。

## 配置管理
### 功能
- 提供配置管理功能，支持从文件、环境变量和命令行参数加载配置。
- 支持配置的热重载，当配置文件发生变化时，自动加载新的配置。

### 代码示例
```go
// 初始化配置函数，在每个命令运行时调用
cobra.OnInitialize(core.OnInitialize(&configFile, "MO", searchDirs(), defaultConfigName))

const defaultHomeDir = ".mo"
func searchDirs() []string {
	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	// 如果获取用户主目录失败，则打印错误信息并退出程序
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultHomeDir), "."}
}
```

## Copier 功能封装
### 功能
- 提供 Copier 功能封装，用于对象的深拷贝和浅拷贝。
- 支持自定义 Copier 函数，用于处理特殊的拷贝场景。

### 代码示例
```go
// gorm 模型定义
type UserM struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    string    `gorm:"column:userID;not null;uniqueIndex:idx_user_userID;comment:用户唯一 ID" json:"userID"`       // 用户唯一 ID
	Username  string    `gorm:"column:username;not null;uniqueIndex:idx_user_username;comment:用户名（唯一）" json:"username"` // 用户名（唯一）
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:current_timestamp;comment:用户创建时间" json:"createdAt"`    // 用户创建时间
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;default:current_timestamp;comment:用户最后修改时间" json:"updatedAt"`  // 用户最后修改时间
}

// Protobuf 定义生成代码
type User struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// userID 表示用户 ID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// username 表示用户名称
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// updatedAt 表示用户最后更新时间
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

userModel := &UserM{
    UserID:   "123",
    Username: "test",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
}
protoUser := &User{}

core.CopyWithConverters(&userModel, protoUser)
```

## Gin 统一请求处理

### 功能
- 提供 Gin 统一请求处理函数 `HandleJSONRequest`，用于处理 JSON 请求并返回 JSON 响应。
- 支持自定义验证函数和绑定函数，用于对请求数据进行验证和绑定。
- 支持自定义错误处理函数，用于处理请求过程中出现的错误。

### 代码示例

```go
func (h *Handler) Login(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.UserV1().Login, h.val.ValidateLoginRequest)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.UserV1().RefreshToken)
}
```