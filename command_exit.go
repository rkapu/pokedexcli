package main

import "os"

func callbackExit(cfg *Config) error {
	os.Exit(0)
	return nil
}
