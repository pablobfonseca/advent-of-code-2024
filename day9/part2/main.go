package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type DiskElement interface{}

type File struct {
	id   string
	size int
}

type Space struct {
	size int
}

type Disk struct {
	elements []DiskElement
}

func (d *Disk) compact() {
	for i := len(d.elements) - 1; i >= 0; i-- {
		currentFile, ok := d.elements[i].(File)
		if !ok {
			continue
		}

		for j := 0; j < i; j++ {
			space, ok := d.elements[j].(Space)
			if !ok || space.size < currentFile.size {
				continue
			}

			remainingSpace := space.size - currentFile.size

			// Replace the file's original position with a space
			d.elements[i] = Space{size: currentFile.size}

			// Insert the file before the space
			d.elements = append(d.elements[:j], append([]DiskElement{currentFile}, d.elements[j:]...)...)
			if remainingSpace > 0 {
				d.elements[j+1] = Space{size: remainingSpace}
			} else {
				d.elements = append(d.elements[:j+1], d.elements[j+2:]...)
			}

			break
		}
	}
}

func (d *Disk) checksum() int {
	checksum := 0
	position := 0

	for _, element := range d.elements {
		switch e := element.(type) {
		case File:
			fileID, err := strconv.Atoi(e.id)
			if err != nil {
				log.Fatalf("Invalid file ID: %v", err)
			}
			for i := 0; i < e.size; i++ {
				checksum += position * fileID
				position++
			}
		case Space:
			position += e.size
		}
	}

	return checksum
}

func NewDisk(input string) Disk {
	fragments := strings.Split(input, "")
	var elements []DiskElement
	fileID := 0

	for i := 0; i < len(fragments); i += 2 {
		fileSize, _ := strconv.Atoi(fragments[i])
		spaceSize, _ := strconv.Atoi(fragments[i+1])
		elements = append(elements, File{id: fmt.Sprint(fileID), size: fileSize})
		elements = append(elements, Space{size: spaceSize})
		fileID++
	}

	return Disk{elements: elements}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error opening the file: %v\n", err)
	}

	disk := NewDisk(string(input))

	disk.compact()

	fmt.Printf("Checksum: %d\n", disk.checksum())
}
