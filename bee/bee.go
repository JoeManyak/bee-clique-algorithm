package bee

import (
	"bee/bee/area"
	"bee/graph"
	"fmt"
	"sort"
)

const scout = 1000

//const worker = 30
const step = 5000
const bestAreaWorkers = 20
const goodAreaWorkers = 5
const bestAreas = 10
const goodAreas = 30

type Hive struct {
	Graph      graph.Graph
	BestZone   [bestAreas]area.Area
	GoodZone   [goodAreas]area.Area
	Areas      []area.Area
	Swarm      []Bee
	CliqueSize int
}

type Bee struct {
	Area    area.Area
	IsScout bool
	Home    *Hive
}

func (b *Bee) GetRandomClique() {
	b.Area = area.Area{
		Clique: []area.CellHelper{},
	}
	randCell := b.Home.Graph.GetRandomCell()
	b.Area.Clique = append(b.Area.Clique, area.CellHelper{
		Cell:      randCell,
		CellPrice: 0,
	})
	for len(b.Area.Clique) < b.Home.CliqueSize {
		randCell = randCell.GetRandomNeighbour()
		isUsed := false
		for _, v := range b.Area.Clique {
			if v.Cell == randCell {
				isUsed = true
				break
			}
		}
		if !isUsed {
			b.Area.Clique = append(b.Area.Clique, area.CellHelper{
				Cell:      randCell,
				CellPrice: 0,
			})
		}
	}
}

func (h *Hive) Solve() string {
	h.setupSwarm()
	solvedPrice := h.CliqueSize * (h.CliqueSize - 1)
	canChangeFrom := 0
	for i := 0; canChangeFrom != h.CliqueSize; i++ {
		if i%2500 == 0 {
			fmt.Printf("Step #[%5d]\n", i)
			fmt.Println(h.BestZone[0].Price)
			fmt.Println(h.BestZone[0].GetResult())
		}
		if h.BestZone[0].Price >= solvedPrice {
			return h.BestZone[0].GetResult()
		}
		h.Areas = make([]area.Area, 0)
		for _, v := range h.BestZone {
			for repeater := 0; repeater < bestAreaWorkers; repeater++ {
				for j := canChangeFrom; j < h.CliqueSize; j++ {
					h.Areas = append(h.Areas, v.ChangeToRandomFrom(canChangeFrom))
				}
			}
		}
		for _, v := range h.GoodZone {
			for repeater := 0; repeater < goodAreaWorkers; repeater++ {
				for j := canChangeFrom; j < h.CliqueSize; j++ {
					h.Areas = append(h.Areas, v.ChangeToRandomFrom(canChangeFrom))
				}
			}
		}
		for scoutI := 0; scoutI < scout; scoutI++ {
			tempScout := Bee{Home: h}
			tempScout.GetRandomClique()
			tempScout.Area.GetPrice()
			h.Areas = append(h.Areas, tempScout.Area)
		}
		h.SortAreas()
		if (i+1)%step == 0 {
			canChangeFrom++
		}
	}
	return "Не вдалось знайти кліку"
}

func (h *Hive) setupSwarm() {
	h.Swarm = make([]Bee, 0)
	h.Areas = make([]area.Area, 0)

	for i := 0; i < scout; i++ {
		tempScout := Bee{Home: h}
		tempScout.GetRandomClique()
		tempScout.Area.GetPrice()
		h.Areas = append(h.Areas, tempScout.Area)
	}
	h.SortAreas()
}

func (h *Hive) SortAreas() {
	sort.Slice(h.Areas, func(i, j int) bool {
		return h.Areas[i].Price > h.Areas[j].Price
	})
	for _, v := range h.Areas {
		sort.Slice(v.Clique, func(i, j int) bool {
			return v.Clique[i].CellPrice > v.Clique[j].CellPrice
		})
	}
	for i := 0; i < bestAreas; i++ {
		h.BestZone[i] = h.Areas[i]
	}
	for i := bestAreas; i < bestAreas+goodAreas; i++ {
		h.GoodZone[i-bestAreas] = h.Areas[i]
	}
	h.Areas = make([]area.Area, 0)
}
