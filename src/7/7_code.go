package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/mattdsteele/advent-of-code"
)

func main() {
	silver()
	gold()
}

func silver() {
	lines := util.ReadFile("src/7/input.txt")
	fmt.Println(silverCalculate(lines))
}

func silverCalculate(input []string) string {
	total := 0
	rootDir := parseInput(input)
	s := rootDir.dirsUnder(100000)
	for _, d := range s {
		total += d.size()
	}
	return fmt.Sprint(total)
}

func gold() {
	lines := util.ReadFile("src/7/input.txt")
	fmt.Println(goldCalculate(lines))
}

func goldCalculate(lines []string) string {
	rootDir := parseInput(lines)
	rootSize := rootDir.size()
	availableSpace := 70000000
	neededSpace := 30000000
	remainingSpace := availableSpace - rootSize
	smallestSize := availableSpace
	neededToDeleteSpace := neededSpace - remainingSpace
	for _, d := range rootDir.flatDirs() {
		if d.size() > neededToDeleteSpace {
			if d.size() < smallestSize {
				smallestSize = d.size()
			}
		}
	}
	return fmt.Sprint(smallestSize)
}

func parseInput(lines []string) *Directory {
	rootDir := new(Directory)
	var currentDir *Directory
	for _, l := range lines {
		if strings.Contains(l, "$ cd") {
			newDir := l[5:]
			if newDir == "/" {
				currentDir = rootDir
			} else if newDir == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = currentDir.cd(newDir)
			}
		} else {
			currentDir.add(l)
		}
	}
	return rootDir
}

type Directory struct {
	name        string
	parent      *Directory
	directories []*Directory
	files       []*File
}

func (dir *Directory) dirsUnder(size int) (dirs []*Directory) {
	flat := dir.flatDirs()
	for _, d := range flat {
		if d.size() < size {
			dirs = append(dirs, d)
		}
	}
	return dirs
}
func (dir *Directory) flatDirs() (dirs []*Directory) {
	dirs = append(dirs, dir)
	for _, d := range dir.directories {
		dirs = append(dirs, d.flatDirs()...)
	}
	return dirs
}

func (dir *Directory) size() (s int) {
	for _, f := range dir.files {
		s += f.size
	}
	for _, d := range dir.directories {
		s += d.size()
	}
	return s
}
func (dir *Directory) add(line string) {
	s := strings.Split(line, " ")
	if s[0] == "dir" {
		newDir := new(Directory)
		newDir.name = s[1]
		newDir.parent = dir
		dir.directories = append(dir.directories, newDir)
	} else {
		newFile := new(File)
		size, _ := strconv.Atoi(s[0])
		newFile.size = size
		newFile.name = s[1]
		dir.files = append(dir.files, newFile)
	}
}
func (dir *Directory) cd(subfolder string) *Directory {
	for _, d := range dir.directories {
		if d.name == subfolder {
			return d
		}
	}
	panic("no known subfolder in dir")
}

type File struct {
	name string
	size int
}
