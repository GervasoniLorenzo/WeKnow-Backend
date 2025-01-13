package adapter

import (
	"fmt"

	"weKnow/config"
)

type KnownAdapter struct {
	config *config.KnownConfig
}

func NewAdapter() *KnownAdapter {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
	}
	return &KnownAdapter{
		config: c,
	}
}
