package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: depify <repo URL> <path>")
		os.Exit(1)
	}
	url := os.Args[1]
	path := os.Args[2]
	err := depify(url, path)
	if err != nil {
		os.Exit(1)
	}
}

func depify(url, path string) error {
	tmpDir, err := ioutil.TempDir("", "depify")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	err = exec.Command("git", "clone", url, tmpDir).Run()
	if err != nil {
		fmt.Println("Error cloning!")
		return err
	}
	out, err := exec.Command("git", "-C", tmpDir, "log", "--all", "--format=%H").CombinedOutput()
	if err != nil {
		fmt.Println("Error getting SHAs!")
		return err
	}
	shas := strings.Split(strings.TrimSpace(string(out)), "\n")
	nearest, nearestSize := "", math.MaxInt64
	for _, sha := range shas {
		err = exec.Command("git", "-C", tmpDir, "checkout", sha).Run()
		if err != nil {
			fmt.Println("Error checking out!")
			return err
		}
		diff, _ := exec.Command("diff", "-r", "--exclude=.git", path, tmpDir).CombinedOutput()
		if len(diff) < nearestSize {
			nearest = sha
			nearestSize = len(diff)
		}
	}
	if nearestSize > 0 {
		fmt.Println("No match found! Closest match: ", nearest)
		return fmt.Errorf("no match")
	}
	fmt.Println(nearest)
	return nil
}
