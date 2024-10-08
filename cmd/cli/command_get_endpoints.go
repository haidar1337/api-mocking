package main

import (
	"fmt"
)

func commandGet(cfg *config, args ...string) error {
	endpoints, err := getEndpoints(cfg)
	if err != nil {
		return err
	}
	fmt.Println(structureEndpoints(endpoints))

	return nil
}
