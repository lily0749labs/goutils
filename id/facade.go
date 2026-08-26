package id

// ID provides object-style access to the ID generation functions in this package.
var ID = facade{}

type facade struct{}

func (facade) InitSnowflake(node int) { InitSnowflake(node) }
func (facade) GetSnowflakeID() int64  { return GetSnowflakeID() }
func (facade) GetSnowflakeId() int64  { return GetSnowflakeId() }
func (facade) GetToken() string       { return GetToken() }
