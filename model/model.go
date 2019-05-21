package model

import (
	"github.com/ninorain22/gintest/manager"
	"fmt"
	)

func init()  {
	// 同步数据库结构
	fmt.Printf("sync table...")
	if err := manager.Engine.Sync2(new(User)); err != nil {
		fmt.Printf("failed to sync table: %s\n", err)
	}
	fmt.Printf("sync table success!")
}
