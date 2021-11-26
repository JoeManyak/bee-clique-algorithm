package main

import (
	"bee/bee"
	"bee/graph"
	"fmt"
	"math/rand"
	"time"
)

const cliqueSize = 8

func main() {
	rand.Seed(time.Now().UnixNano())
	g := graph.Generate()
	g.CliqueGenerate(cliqueSize)
	g.Save()
	g.CellGenerate()
	g.SetUseful(cliqueSize)
	h := bee.Hive{
		Graph:      g,
		CliqueSize: cliqueSize,
	}
	fmt.Println("finished:", h.Solve())
	//g.Show()
}
