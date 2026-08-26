package id

import (
	"errors"
	"fmt"
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	// ID 提供 ID 生成的结构化入口。
	ID = id{}

	snow      *snowflake.Node
	snowNodes = make(map[int]*snowflake.Node)
	snowMu    sync.RWMutex
)

var (
	// ErrInvalidNode 表示雪花节点编号不在实现支持的范围内。
	ErrInvalidNode = errors.New("invalid snowflake node")
	// ErrInvalidCount 表示批量生成数量不是正数。
	ErrInvalidCount = errors.New("id count must be positive")
)

type id struct{}

func init() {
	ID.InitSnowflake(0)
}

// InitSnowflakeE 使用 node 初始化雪花 ID 生成器，失败时返回错误。
func (id) InitSnowflakeE(node int) error {
	newNode, err := snowflake.NewNode(int64(node))
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidNode, err)
	}
	snowMu.Lock()
	if existingNode, exists := snowNodes[node]; exists {
		snow = existingNode
	} else {
		snowNodes[node] = newNode
		snow = newNode
	}
	snowMu.Unlock()
	return nil
}

// InitSnowflake 使用 node 初始化雪花 ID 生成器。
// 初始化失败时保持旧行为并触发 panic；新代码推荐使用 InitSnowflakeE。
func (f id) InitSnowflake(node int) {
	if err := f.InitSnowflakeE(node); err != nil {
		panic(err)
	}
}

// SnowflakeID 返回一个新的雪花 ID。
func (id) SnowflakeID() int64 {
	snowMu.RLock()
	node := snow
	snowMu.RUnlock()
	return node.Generate().Int64()
}

// SnowflakeIDs 批量生成 count 个雪花 ID。
func (f id) SnowflakeIDs(count int) ([]int64, error) {
	if count <= 0 {
		return nil, ErrInvalidCount
	}
	values := make([]int64, count)
	for index := range values {
		values[index] = f.SnowflakeID()
	}
	return values, nil
}

// InitSnowflake 使用 node 初始化雪花 ID 生成器。
// Deprecated: 使用 ID.InitSnowflake。
func InitSnowflake(node int) { ID.InitSnowflake(node) }

// GetSnowflakeID 返回一个新的雪花 ID。
// Deprecated: 使用 ID.SnowflakeID。
func GetSnowflakeID() int64 { return ID.SnowflakeID() }

// GetSnowflakeId 返回一个新的雪花 ID。
// Deprecated: 使用 ID.SnowflakeID。
func GetSnowflakeId() int64 { return ID.SnowflakeID() }
