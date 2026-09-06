// Package time 提供时间解析、格式化、时区转换、日期边界、时间区间、时间戳和工作日处理。
//
// 默认入口 Time 使用 Asia/Shanghai 时区和系统时钟。WithTimezone 和 WithNow
// 返回独立的 Helper 副本，可用于其他时区、稳定测试和任务回放。
//
// 涉及“今天”“昨天”“自然日”和“工作日”的方法均按 Helper 配置时区计算。
// 数据库时间条件优先使用以 HalfOpenRange 结尾的半开区间方法，避免依赖存储精度。
// 完整的分类函数索引和迁移表请参阅本目录 README.md。
package time
