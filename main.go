package main

import (
	"bufio"
	"fmt"
    "io"
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
		cmd := exec.Command("lux", link)

        // Create pipes to capture the standard output and standard error
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error creating stdout pipe:", err)
			return
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Println("Error creating stderr pipe:", err)
			return
		}
        
        // Start the command
		err = cmd.Start()
		if err != nil {
			fmt.Printf("Error executing 'lux' with link %s: %v\n", link, err)
		}

		// Create a combined reader to read both standard output and standard error
		reader := io.MultiReader(stdout, stderr)

		// Display the output from the command
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

        // Wait for the command to complete before proceeding to the next link
		err = cmd.Wait()
        if err != nil {
			fmt.Printf("Error executing 'lux' with link %s: %v\n", link, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

