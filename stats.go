package main

import (
	"fmt"
	"sort"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)
type cell struct {
	time string
	value int
}

// stats calculates and prints the stats.
func stats(email string) {
	commits := processRepositories(email)
	printCommitsStats(commits)
}
// fillCommits given a repository found in `path`, gets the commits and
// puts them in the `commits` map, returning it when completed
func fillCommits(email string, path string, commits map[string]int) map[string]int {
	// instantiate a git repo object from path
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}
	// get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}
	// get the commits history starting from HEAD
	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}
	// iterate the commits
	//offset := calcOffset()
	err = iterator.ForEach(func(c *object.Commit) error {
		//daysAgo := countDaysSinceDate(c.Author.When) + offset


		if c.Author.Email != email {
			return nil
		}

		daysAgo := time.Now().Sub(c.Author.When).Hours() / 24
		if daysAgo <= 180 {
			commits[c.Author.When.Format("2006/01/02")]++
		}
		//if daysAgo != outOfRange {
		//	commits[daysAgo]++
		//}

		return nil
	})
	if err != nil {
		panic(err)
	}

	//fmt.Printf("commits: %v", commits)
	return commits
}

//get half of year days map
func getHalfYearDaysMap() map[string]int {
	m := make(map[string]int)
	m[time.Now().Format("2006/01/02")] = 0

	for{
		preDay := time.Now().Add(-(time.Duration(len(m)) * 24 * time.Hour))
		if time.Now().Sub(preDay).Hours() / 24 > 180{
			break
		}
		m[preDay.Format("2006/01/02")] = 0
	}
	return m
}
// processRepositories given an user email, returns the
// commits made in the last 6 months
func processRepositories(email string) map[string]int {
	filePath := getDotFilePath()
	repos := parseFileLinesToSlice(filePath)
	//daysInMap := daysInLastSixMonths
	//
	//
	//commits := make(map[int]int, daysInMap)
	//for i := daysInMap; i > 0; i-- {
	//	commits[i] = 0
	//}
	commits := getHalfYearDaysMap()

	for _, path := range repos {
		commits = fillCommits(email, path, commits)
	}

	return commits
}

// printCommitsStats prints the commits stats
func printCommitsStats(commits map[string]int) {
	printGraph(commits)
}

func printGraph(commits map[string]int) {
	keys := make([]string, 0)

	i := 0
	for k, v := range commits {
		if v == 0 {
			continue
		}
		keys = append(keys, k)
		i++
	}
	sort.Strings(keys)
	cells := make([]cell, len(keys))
	for i, k := range keys {
		cells[i] = cell{time: k, value: commits[k]}
	}
	for i := 0; i < len(cells); i++ {
		line := false
		if (i + 1) % 5 == 0 {
			line = true
		}
		printCell(cells[i], line)
	}
	fmt.Println()

}

func printCell(c cell, line bool) {
	val := c.value
	escape := "\033[0;37;30m"
	switch {
	case val > 0 && val < 5:
		escape = "\033[1;30;47m"
	case val >= 5 && val < 10:
		escape = "\033[1;30;43m"
	case val >= 10:
		escape = "\033[1;30;42m"
	}

	if line {
		fmt.Printf(escape+" %s \033[0m\n", fmt.Sprintf("%s(%d)", c.time, val))
	} else {
		fmt.Printf(escape+" %s \033[0m", fmt.Sprintf("%s(%d)", c.time, val))
	}
}