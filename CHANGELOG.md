# Changelog

本项目的重要变更都会记录在此文件中。版本号遵循 [Semantic Versioning](https://semver.org/)。

## [Unreleased]

### Added

- 时间门面导出为 `time.Helper`，并新增时区时间、无溢出月份计算、显式错误和半开时间区间 API。

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
