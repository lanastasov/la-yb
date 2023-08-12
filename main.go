package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Open the text file containing links
	filePath := "links.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the links line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		link := scanner.Text()
        
		// Run the "lux" executable with the link as an argument
		cmd := exec.Command("glux", link)
		err := cmd.Run()

		if err != nil {
			fmt.Printf("Error executing 'lux' with link %s: %v\n", link, err)
		}

        // Wait for the command to complete before proceeding to the next link
		cmd.Wait()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

