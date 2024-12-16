package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error opening the file: %v\n", err)
	}

	fragments := strings.Split(string(file), "")

	var files []string
	var sizes []string

	for i := 0; i < len(fragments); i++ {
		if i%2 == 0 {
			files = append(files, fragments[i])
		} else {
			sizes = append(sizes, fragments[i])
		}
	}

	var disk []string

	for i, file := range files {
		sizesCount := 0
		fileCount, _ := strconv.Atoi(file)
		if i >= len(sizes) {
			sizesCount = 0
		} else {
			sizesCount, _ = strconv.Atoi(sizes[i])
		}
		fileIdStr := strconv.Itoa(i)
		for i := 0; i < fileCount; i++ {
			disk = append(disk, fileIdStr)
		}

		for i := 0; i < sizesCount; i++ {
			disk = append(disk, ".")
		}
	}

	compactedDisk := compactDisk(disk)

	checksum := calculateChecksum(compactedDisk)

	fmt.Println("Checksum: ", checksum)
}

func calculateChecksum(disk []string) int {
	checksum := 0
	for pos, block := range disk {
		if block != "." {
			fileID, _ := strconv.Atoi(block)
			checksum += pos * fileID
		}
	}

	return checksum
}

func compactDisk(disk []string) []string {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != "." {
			for j := 0; j < i; j++ {
				if disk[j] == "." {
					disk[j], disk[i] = disk[i], disk[j]
					break
				}
			}
		}
	}

	return disk
}
