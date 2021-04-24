package main

import (
	"fmt"
	"github.com/rickkvm/fops-task/linecount"
)

func main() {
	result := linecount.Count([]byte("1\n2\n3"))
}
