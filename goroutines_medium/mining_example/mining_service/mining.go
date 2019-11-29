package mining_service

import (
	"fmt"
	"strconv"
)

func Finder(mines []string) []string {
	ores := []string{}
	for _, rock := range mines {
		if rock == "ore" {
			ores = append(ores, rock)
		}
	}
	return ores
}

func Miner(ores []string) []string {
	minedOres := []string{}
	for range ores {
		minedOres = append(minedOres, "minedOre")
	}
	return minedOres
}

func Smelter(minedOres []string) []string {
	smelteredOres := []string{}
	for range minedOres {
		smelteredOres = append(smelteredOres, "smeltedOre")
	}
	return smelteredOres
}

func Finder1(mines []string) {
	for _, rock := range mines {
		if rock == "ore" {
			fmt.Println("Finder1 found the ore")
		}
	}
}

func Finder2(mines []string) {
	for _, rock := range mines {
		if rock == "ore" {
			fmt.Println("Finder2 found the ore")
		}
	}
}

func Execute() {
	rocks := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChanel := make(chan string)
	minedOreChanel := make(chan string)
	go func(rocks [5]string) {
		for index, rock := range rocks {
			if rock == "ore" {
				oreChanel <- rock + strconv.Itoa(index)
			}
		}
	}(rocks)

	go func() {
		index := 0
		for ore := range oreChanel {
			fmt.Println("From Finder:", ore)
			minedOreChanel <- "minedOre" + strconv.Itoa(index)
			index += 1
		}
	}()

	go func() {
		for minedOre := range minedOreChanel {
			fmt.Println("Mined Ore ", minedOre)
			fmt.Println("From Smelter:", minedOre, " is smelted")

		}
	}()
}
