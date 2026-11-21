# Changelog

本项目的重要变更都会记录在此文件中。版本号遵循 [Semantic Versioning](https://semver.org/)。

## [Unreleased]

### Added

- 时间门面导出为 `time.Helper`，并新增时区时间、无溢出月份计算、显式错误和半开时间区间 API。
- 新增 `pbtime` 子包：`pbtime.Format(ts, loc)` 把 protobuf 时间戳按指定时区格式化为 `YYYY-MM-DD HH:mm:ss`，未设置（nil）时返回空字符串；`pbtime.FromCarbon(t)` 把 carbon 时间转成 protobuf 时间戳，nil 和零时间返回 nil。单独成包，只有用到它的项目才会链接 `google.golang.org/protobuf`。
- 新增 `dbtypes` 子包：泛型 `dbtypes.Array[T]` 对应 PostgreSQL 一维数组列（`text[]`、`smallint[]`、`integer[]`、`bigint[]`、`boolean[]`、`real[]`、`double precision[]`，`uint64` 用 `numeric(20,0)[]` 保存完整范围）。实现 `sql.Scanner` 和 `driver.Valuer`，按 PostgreSQL 数组文本格式读写，nil 切片写成 NULL、空切片写成 `{}`；只支持一维数组，元素不能为 NULL；只依赖标准库。

### Fixed

- 修复在月末计算上月范围时可能返回当月的问题。
- `NowTime` 保留 Go 单调时钟，相对时间解析遵循 `WithNow` 注入时钟。

## [0.3.0] - 2026-08-26

### Added

- 新增 `Money.FormatCents` 金额格式化门面。
- 新增泛型 `slice`、`maps` 和 `queue` 子包。
- 新增 Unicode 安全的 `String` 字符串门面。
- 新增安全随机、严格解析和现代加密入口。

### Changed

- 强化雪花 ID、时间处理和现有门面 API。

[Unreleased]: https://github.com/lily0749labs/goutils/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/lily0749labs/goutils/releases/tag/v0.3.0
