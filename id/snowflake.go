package id

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	snow   *snowflake.Node
	snowMu sync.RWMutex
)

func init() {
	InitSnowflake(0)
}

// InitSnowflake 使用 node 初始化雪花 ID 生成器。
func InitSnowflake(node int) {
	newNode, err := snowflake.NewNode(int64(node))
	if err != nil {
		panic("init snowflake error : " + err.Error())
	}
	snowMu.Lock()
	snow = newNode
	snowMu.Unlock()
}

// GetSnowflakeID 返回一个新的雪花 ID。
func GetSnowflakeID() int64 {
	snowMu.RLock()
	node := snow
	snowMu.RUnlock()
	return node.Generate().Int64()
}

// GetSnowflakeId 返回一个新的雪花 ID。
// Deprecated: 请使用 GetSnowflakeID。
func GetSnowflakeId() int64 {
	return GetSnowflakeID()
}
