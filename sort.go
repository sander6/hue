package main

import (
	"github.com/sander6/hue/light"
)

type BulbsSortedByID light.Bulbs

func (b BulbsSortedByID) Len() int {
	return len(b)
}

func (b BulbsSortedByID) Less(i, j int) bool {
	return b[i].ID() < b[j].ID()
}

func (b BulbsSortedByID) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type BulbsSortedByName light.Bulbs

func (b BulbsSortedByName) Len() int {
	return len(b)
}

func (b BulbsSortedByName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func (b BulbsSortedByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type GroupsSortedByID light.Groups

func (g GroupsSortedByID) Len() int {
	return len(g)
}

func (g GroupsSortedByID) Less(i, j int) bool {
	return g[i].ID() < g[j].ID()
}

func (g GroupsSortedByID) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
