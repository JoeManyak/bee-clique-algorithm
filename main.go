package main

import (
	"bee/bee"
	"bee/graph"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	g := graph.Generate()
	g.Save()
	g.CellGenerate()
	g.SetUseful(3)
	h := bee.Hive{
		Graph:      g,
		CliqueSize: 4,
	}
	fmt.Println("finished:", h.Solve())
	//g.Show()
}
