package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"aoc2025/internal/day01"
	"aoc2025/internal/day02"
	"aoc2025/internal/day03"
	"aoc2025/internal/day04"
)

type solver func()

var solvers = map[int]solver{
	1: day01.SolveDay01,
	2: day02.SolveDay02,
	3: day03.SolveDay03,
	4: day04.SolveDay04,
}

func main() {
	day := flag.Int("day", 0, "day to run (e.g. 2)")
	list := flag.Bool("list", false, "list available days")
	all := flag.Bool("all", false, "run all available days in ascending order")
	flag.Parse()

	if *list {
		fmt.Println("Available days:")
		for _, d := range sortedDays() {
			fmt.Printf("  %d\n", d)
		}
		return
	}

	if *all {
		for _, d := range sortedDays() {
			fmt.Printf("== Day %02d ==\n", d)
			solvers[d]()
		}
		return
	}

	if *day == 0 {
		fmt.Fprintln(os.Stderr, "error: missing -day (or use -list / -all)")
		flag.Usage()
		os.Exit(2)
	}

	run, ok := solvers[*day]
	if !ok {
		fmt.Fprintf(os.Stderr, "error: unknown day %d (use -list)\n", *day)
		os.Exit(2)
	}

	run()
}

func sortedDays() []int {
	days := make([]int, 0, len(solvers))
	for d := range solvers {
		days = append(days, d)
	}
	sort.Ints(days)
	return days
}

