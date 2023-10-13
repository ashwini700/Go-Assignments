// In this assignment, you will work with file reading and error handling in Go. You'll create a simple program to read a text file, count the number of words, and handle potential errors that may occur during the process.

// Instructions:
// Create a text file named "sample.txt" with some content. The content can be any text or a short paragraph.
// Write a Go program that reads the content of the "sample.txt" file and counts the number of words in it.
// Use the ioutil.ReadFile function to read the file's contents into a byte slice. Handle the error that may occur during file reading and print an error message if it occurs.
// Implement a function called CountWords that takes the file content as a string and returns the number of words in the content.
// Handle potential errors within the CountWords function. Specifically, account for cases where the input string is empty or contains only whitespace characters.
// In the main program, call the CountWords function with the file content and print the number of words. Handle any errors returned by the function and print an error message in case of an error.
// Test your program with different content in the "sample.txt" file, including empty files or files containing only spaces, to ensure error handling is working correctly.
// At the end try to delete the “sample.txt” file & check if your program is able to handle the errors.

package main

import (
	"fmt"
	"os"
	"strings"
)

func CountWords(content string) (int, error) {
	if strings.TrimSpace(content) == "" {
		return 0, fmt.Errorf("string is empty")
	}

	words := strings.Fields(content)
	return len(words), nil
}

func main() {
	content, err := os.ReadFile("text.txt") //filename is text
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileContents := string(content) //convert byte slice to string

	numWords, err := CountWords(fileContents) //count number of words by calling cw func
	if err != nil {
		fmt.Println("Error counting words:", err)
		return
	}

	fmt.Printf("Number of words in text file: %d\n", numWords)

	//to delete file

	// err = DeleteFile("text.txt")
	// if err != nil {
	// 		fmt.Println("Error deleting file:", err)
	// 	}
	// }

	// func DeleteFile(filename string) error { //deletes file else return error
	// 	err := deleteFile(filename)
	// 	if err != nil {
	// 		return fmt.Errorf("fails to delete : %v", err)
	// 	}
	// 	fmt.Printf("File \"%s\" deleted successfully.\n", filename)
	// 	return nil
	// }

	// func deleteFile(filename string) error {
	// 	return os.WriteFile(filename, []byte{}, 0)
	// }
}
