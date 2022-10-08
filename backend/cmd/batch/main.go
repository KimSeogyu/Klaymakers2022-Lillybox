package main

import (
	"lillybox-backend/internal/batch"
)

func main() {
	c := batch.Client{}

	c.ConfigConnection()
	c.SetBatchAccessLog()
	c.SetBatchErrorLog()
	c.BatchLoop()
}
