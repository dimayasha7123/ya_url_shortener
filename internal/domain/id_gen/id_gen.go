package id_gen

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

const (
	stdEpoch = 1682813374925
)

type IDGenerator struct {
	node *snowflake.Node
}

func New(nodeID int64) (IDGenerator, error) {
	snowflake.Epoch = stdEpoch
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return IDGenerator{}, fmt.Errorf("can't create new snowflake node: %v", err)
	}
	return IDGenerator{node: node}, nil
}

func (g IDGenerator) Generate() uint64 {
	return uint64(g.node.Generate())
}
