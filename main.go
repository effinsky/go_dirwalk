package main

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	p := "./"
	if err := printDirsAndFilesRec(p); err != nil {
		panic(err)
	}
}

// Write a program that, when given a string representing a path on a filesystem,
// prints out all the files and folders contained at that path. Your solution
// should recursively traverse directories if encountered.
func printDirsAndFilesRec(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("getting file info: %w", err)
	}

	fmt.Printf("%v\n", fi.Name())

	if fi.IsDir() {
		dirEntries, err := os.ReadDir(path)
		if err != nil {
			return fmt.Errorf("reading directory: %w", err)
		}

		for _, de := range dirEntries {
			fi, err := de.Info()
			if err != nil {
				return fmt.Errorf("getting file info: %w", err)
			}

			printDirsAndFilesRec(filepath.Join(path, fi.Name()))
		}
	}

	return nil
}

// iter using a queue from a linked list
func printDirEntriesBFSIter(path string) error {
	// Get file info for path
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("getting file info: %w", err)
	}

	type pathInfo struct {
		path     string
		fileInfo os.FileInfo
	}

	q := list.New()
	// Push it on the queue.
	q.PushBack(pathInfo{path, fi})

	// Process file info items until none are left in queue.
	for q.Len() > 0 {
		pi := q.Remove(q.Front()).(pathInfo)
		fmt.Printf("%v\n", pi.fileInfo.Name())

		if pi.fileInfo.IsDir() {
			dirEntries, err := os.ReadDir(pi.path)
			if err != nil {
				return fmt.Errorf("reading directory: %w", err)
			}
			for _, de := range dirEntries {
				dePath := filepath.Join(pi.path, de.Name())
				fi, err := de.Info()
				if err != nil {
					return fmt.Errorf("getting file info: %w", err)
				}
				// Push every dirEntry's file info on the queue.
				q.PushBack(pathInfo{dePath, fi})
			}
		}
	}
	return nil
}

// DFS using slice (stack), popping off the end, iterating over dir entries
// back to front
func printDirsAndFilesIterDFS(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("getting file info: %w", err)
	}

	type pathInfo struct {
		path     string
		fileInfo os.FileInfo
	}

	stack := []pathInfo{{path, fi}}

	for len(stack) > 0 {
		// Pop from the stack
		last := len(stack) - 1
		pi := stack[last]
		stack = stack[:last]

		fmt.Printf("%v\n", pi.fileInfo.Name())

		if pi.fileInfo.IsDir() {
			dirEntries, err := os.ReadDir(pi.path)
			if err != nil {
				return fmt.Errorf("reading directory: %w", err)
			}

			// Push entries onto the stack in reverse order
			// This ensures we process them in the original order when popping
			for i := len(dirEntries) - 1; i >= 0; i-- {
				de := dirEntries[i]
				dePath := filepath.Join(pi.path, de.Name())
				fi, err := de.Info()
				if err != nil {
					return fmt.Errorf("getting file info: %w", err)
				}
				stack = append(stack, pathInfo{dePath, fi})
			}
		}
	}

	return nil
}
