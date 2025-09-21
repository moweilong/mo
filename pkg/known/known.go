package known

const (
	// XTraceID 用来定义上下文中的键，代表请求 ID.
	XTraceID = "x-trace-id"

	// XUserID 用来定义上下文的键，代表请求用户 ID. UserID 整个用户生命周期唯一.
	XUserID = "x-user-id"

	// XUsername 用来定义上下文的键，代表请求用户名.
	XUsername = "x-username"
)
