package id

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	// ID 提供 ID 生成的结构化入口。
	ID = facade{}

	snow   *snowflake.Node
	snowMu sync.RWMutex
)

type facade struct{}

func init() {
	ID.InitSnowflake(0)
}

// InitSnowflake 使用 node 初始化雪花 ID 生成器。
func (facade) InitSnowflake(node int) {
	newNode, err := snowflake.NewNode(int64(node))
	if err != nil {
		panic("init snowflake error : " + err.Error())
	}
	snowMu.Lock()
	snow = newNode
	snowMu.Unlock()
}

// SnowflakeID 返回一个新的雪花 ID。
func (facade) SnowflakeID() int64 {
	snowMu.RLock()
	node := snow
	snowMu.RUnlock()
	return node.Generate().Int64()
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
