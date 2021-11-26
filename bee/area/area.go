package area

import (
	"bee/graph"
	"fmt"
	"math/rand"
)

type Area struct {
	Clique []CellHelper
	Price  int
}

func (a Area) ChangeToRandomFrom(id int) Area {
	for {
		change := rand.Intn(len(a.Clique)-id) + id
		randCell := rand.Intn(len(a.Clique[change].Cell.Neighbours))
		randNeighbour := a.Clique[change].Cell.Neighbours[randCell]
		isOk := true
		for _, v := range a.Clique {
			if randNeighbour == v.Cell {
				isOk = false
			}
		}
		if isOk {
			a.Clique[change].Cell = a.Clique[change].Cell.GetRandomNeighbour()
			a.GetCellPrice(&a.Clique[change])
			return a
		}
	}
}

func (a *Area) GetResult() string {
	result := ""
	for _, v := range a.Clique {
		result += fmt.Sprintf("%d;", v.Cell.Id)
	}
	return result
}

type CellHelper struct {
	Cell      *graph.Cell
	CellPrice int
}

func (a *Area) GetPrice() int {
	a.Price = 0
	for _, v1 := range a.Clique {
		a.GetCellPrice(&v1)
		a.Price += v1.CellPrice
	}
	return a.Price
}

func (a *Area) GetCellPrice(cellHelper *CellHelper) {
	cellHelper.CellPrice = 0
	for _, v2 := range cellHelper.Cell.Neighbours {
		for _, v3 := range a.Clique {
			if v3.Cell == v2 {
				cellHelper.CellPrice++
				break
			}
		}
	}
}
