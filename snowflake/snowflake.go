// Copyright 2022 Sun XY <xy.sunviv@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(epoch string, machineID int64) (err error) {
	var startTime time.Time
	startTime, err = time.Parse("2006-01-02", epoch)
	if err != nil {
		return
	}
	snowflake.Epoch = startTime.UnixMilli()
	node, err = snowflake.NewNode(machineID)
	return
}

// GenID 生成 int64 snowflake ID
func GenID() int64 {
	return node.Generate().Int64()
}
