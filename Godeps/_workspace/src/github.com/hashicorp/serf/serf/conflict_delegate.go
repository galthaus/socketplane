package serf

import (
	"github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/memberlist"
)

type conflictDelegate struct {
	serf *Serf
}

func (c *conflictDelegate) NotifyConflict(existing, other *memberlist.Node) {
	c.serf.handleNodeConflict(existing, other)
}
