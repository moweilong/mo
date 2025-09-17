# 设计

## AddCallerSkip

AddCallerSkip 用于添加调用者的跳过层数，用于在日志中显示正确的调用位置。

- 调用链分析 ：
    - 用户代码调用全局日志函数，如 mlog.Info(...)
    - 全局函数内部调用 std.Info(...) （std是全局logger实例）
    - std.Info(...) 最终调用底层的 zap.Logger.Info(...)

- 为什么需要跳过2层 ：
    - 第一层跳过：全局日志函数（如Info、Debug等）
    - 第二层跳过：zapLogger结构体的方法实现

- 设置AddCallerSkip(2)的目的 ：
    - 默认情况下，zap会将日志中的调用者信息显示为直接调用zap方法的代码位置
    - 在这个多层封装的日志系统中，如果不调整，日志会显示mlog包内部的函数位置，而不是用户代码中的实际调用位置
    - 通过设置 AddCallerSkip(2) ，我们告诉zap跳过两层调用栈，从而显示用户代码中真正调用日志函数的位置