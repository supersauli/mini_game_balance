package tool_uuid

import (
	"github.com/bwmarrin/snowflake"
)

var snowflakeNode *snowflake.Node

func init() {
	snowflakeNode, _ = snowflake.NewNode(1)
}

func Generate() int64 {
	return int64(snowflakeNode.Generate())
}
