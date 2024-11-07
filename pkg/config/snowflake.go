package config

import (
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func NewSnowFlake(env *env.Env) *snowflake.Node {
	node, err := snowflake.NewNode(env.SnowFlakeNode)
	if err != nil {
		zap.S().Fatalf("Unable to create snowflake node: %v", err)
	}

	return node
}
