package graph

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const size = 300
const expected = 50

const dataName = "data.txt"

type Graph struct {
	Matrix         [size][size]bool
	Cells          [size]Cell
	AvailableCells []*Cell
}

type Cell struct {
	Id         int
	Neighbours []*Cell
	IsUseful   bool
}

func (g *Graph) CliqueGenerate(size int) {
	nodes := make([]int, 0)
	for i := 0; len(nodes) < size; i++ {
		randNode := rand.Intn(size)
		isOk := true
		for _, v := range nodes {
			if v == randNode {
				isOk = false
				break
			}
		}
		if isOk {
			nodes = append(nodes, randNode)
		}
	}
	for _, v1 := range nodes {
		for _, v2 := range nodes {
			g.Matrix[v1][v2] = true
			g.Matrix[v2][v1] = true
		}
	}
}

func (c *Cell) GetRandomNeighbour() *Cell {
	return c.Neighbours[rand.Intn(len(c.Neighbours))]
}

func (g *Graph) GetRandomCell() *Cell {
	return g.AvailableCells[rand.Intn(len(g.AvailableCells))]
}

func (g *Graph) CellGenerate() {
	for i := range g.Matrix {
		g.Cells[i].Id = i
		for j := range g.Matrix[i] {
			if g.Matrix[i][j] {
				g.Cells[i].Neighbours = append(g.Cells[i].Neighbours, &g.Cells[j])
			}
		}
	}
}

func (g *Graph) SetUseful(cliqueSize int) {
	for i, v := range g.Cells {
		v.IsUseful = !(len(v.Neighbours) < cliqueSize)
		if v.IsUseful {
			g.AvailableCells = append(g.AvailableCells, &g.Cells[i])
		}
	}
}

func (g *Graph) Show() {
	for i := range g.Matrix {
		for j := range g.Matrix[i] {
			if g.Matrix[i][j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

func (g *Graph) Save() {
	open, err := os.Create(dataName)
	if err != nil {
		log.Fatal("unable to save matrix:", err.Error())
	}
	result := ""
	for i := range g.Matrix {
		for j := range g.Matrix[i] {
			if g.Matrix[i][j] {
				result += "1;"
			} else {
				result += "0;"
			}
		}
		result += "\n"
	}
	_, err = open.Write([]byte(result))
	if err != nil {
		log.Fatal("unable to write to file", err.Error())
	}
	err = open.Close()
	if err != nil {
		log.Fatal("unable to close file", err.Error())
	}
}

func (g *Graph) Load() {
	open, err := os.Open(dataName)
	if err != nil {
		log.Fatal("unable to open file:", err.Error())
	}
	scanner := bufio.NewScanner(open)
	i := 0
	for scanner.Scan() {
		sl := strings.Split(scanner.Text(), ";")
		for j := 0; j < len(sl)-1; j++ {
			parsed, err := strconv.Atoi(sl[j])
			if err != nil {
				log.Fatal("unable to parse datafile:", err.Error())
			}
			g.Matrix[i][j] = parsed == 1
		}
		i++
	}
}

func Generate() Graph {
	rand.Seed(time.Now().UnixNano())
	g := Graph{[size][size]bool{}, [size]Cell{}, []*Cell{}}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			random := getRandom()
			g.Matrix[i][j] = random
			g.Matrix[j][i] = g.Matrix[i][j]
		}
	}
	return g
}

func getRandom() bool {
	return float64(expected)/float64(size) > rand.Float64()
}
