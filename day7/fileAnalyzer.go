package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)



type File struct {
	name string
	size int
	parent *Folder
}

type Folder struct {
	name string
	parent *Folder
	files []File
	folders []*Folder
}

func main() {
	cwd, _ := os.Getwd()
	path := (cwd + "/data/terminalCmds.txt")
	lines := readLines(path)

	allFolders := buildStructure(lines)

	// Part 1
	sizeLimit := 100000
	totalSize := 0
	for _, folder := range allFolders {
		size := folder.calculateFolderSize()
		//fmt.Printf("Folder %s has size %d\n", folder.name, size)
		if size <sizeLimit {
			totalSize += size
		}
	}
	fmt.Printf("Total size of all folders under %d bytes: %d\n", sizeLimit, totalSize)

	// Part 2
	usedSpace := 70000000-allFolders[0].calculateFolderSize()
	requiredSpace := 30000000 - usedSpace
	fmt.Printf("Required space to free up: %d\n", requiredSpace)
	// Find the folder with mininmun size larger than the requiredSpace
	eligibleFolders := []*Folder{}
	for _, folder := range allFolders {
		if folder.calculateFolderSize() > requiredSpace {
			eligibleFolders = append(eligibleFolders, folder)
		}
	}
	fmt.Printf("Number of eligible folders: %d\n", len(eligibleFolders))
	var foldeWithMinSize *Folder = allFolders[0]
	for _, folder := range eligibleFolders {
			if folder.calculateFolderSize() < foldeWithMinSize.calculateFolderSize() {
			foldeWithMinSize = folder
		}
	}
	fmt.Printf("Folder which shall be deledted to free up %d bytes is %s with size %d\n", requiredSpace, foldeWithMinSize.name, foldeWithMinSize.calculateFolderSize())

	

}

func buildStructure(lines []string) []*Folder {
	root := Folder{name: "root"}
	root.name = "root"
	currentFolder := &root
	allFolders := []*Folder{&root}
	for _, line := range lines {
	 	if line == "$ cd .." {
			// Command to go up one directory
			if currentFolder.parent != nil {
				currentFolder = currentFolder.parent
			} else {
				fmt.Println("Already at root")
			}
		} else if strings.HasPrefix(line, "$ cd") {
			// Command to change directory 
			//fmt.Printf("Changing directory to %s\n", line[5:])
			for _, folder := range currentFolder.folders{
				if folder.name == line[5:] {
					currentFolder = folder
					break
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			// Command to create a folder
			newFolder := Folder{name: line[4:], parent: currentFolder}
			// add the new folder to the folders of the current folder
			currentFolder.folders = append(currentFolder.folders, &newFolder)
			allFolders = append(allFolders, &newFolder)
			//fmt.Printf("Adding folder %s to folder %s\n", newFolder.name, currentFolder.name)
		} else if _, err := strconv.Atoi(string(line[0])); err == nil {
			// Command to create a file
			name := strings.Split(line, " ")[1]
			size,err := strconv.Atoi(strings.Split(line, " ")[0])
			if err != nil {
				fmt.Println("Error converting size to int")
			}
			newFile := File{name: name, size: size, parent: currentFolder}
			//fmt.Printf("Adding file %s to folder %s of size %d\n", newFile.name, currentFolder.name, newFile.size)
			currentFolder.files = append(currentFolder.files, newFile)
			}	
		}
	return allFolders
}

func (f *Folder) calculateFolderSize() int {
	size := 0
	//fmt.Printf("Folder %s has %d files and %d folders\n",f.name,len(f.files),len(f.folders))
	for _, file := range f.files {
		//fmt.Printf("File %s has size %d\n", file.name, file.size)
		size += file.size
	}
	for _, folder := range f.folders {
		size += folder.calculateFolderSize()
	}
	return size
}


// read a file and return the lines as an array
func readLines(path string) []string {
	var lines []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initializing
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Scon the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
