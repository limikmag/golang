package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/sha3"
)

func main() {

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "get is a tool for list content files and printing their keccak hash",
		Run: func(cmd *cobra.Command, args []string) {
			size, _ := cmd.Flags().GetInt("size")
			_, err := HandleGetInput(args, size)
			if err != nil {
				fmt.Println(err)
				return
			}
		},
	}

	var lsCmd = &cobra.Command{
		Use:   "ls",
		Short: "ls is a tool for listing files in directory and printing their keccak hash",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := HandleLsInput(args)
			if err != nil {
				fmt.Println(err)
				return
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "byhash"}
	getCmd.PersistentFlags().Int("size", 100, "size piece of file to print in bytes")
	rootCmd.AddCommand(lsCmd, getCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic("fatal error")
	}

}

// HandleLsInput function hande input for ls command
func HandleLsInput(args []string) (int, error) {
	if len(args) < 1 {
		return 1, fmt.Errorf("you do not write any directory")
	}

	for _, arg := range args {

		arg = lsInputToCorrect(arg)
		fmt.Printf("Directory: %v\n\n", arg)
		files, err := ioutil.ReadDir(arg)
		if err != nil {
			return 2, fmt.Errorf("problem while reading directory %v: %v", arg, err)
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}

			data, err := ioutil.ReadFile(arg + file.Name())
			if err != nil {
				return 3, fmt.Errorf("problem while reading file %v: %v", arg+file.Name(), err)
			}
			fmt.Printf("%-40v -> %64x\n", file.Name(), sha3.Sum256([]byte(data)))

		}
	}
	return 0, nil
}

//HandleGetInput function handle input for get command
func HandleGetInput(args []string, size int) (int, error) {
	if len(args) < 1 {
		return 1, fmt.Errorf("you do not write any file")
	}

	for _, arg := range args {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			return 2, fmt.Errorf("problem while reading file %v: %v", arg, err)
		}
		if len(data) < size {
			size = len(data)
		}
		fmt.Printf("%-40v -> %64x\n", string(data[:size]), sha3.Sum256([]byte(data)))

	}
	return 0, nil

}

func lsInputToCorrect(arg string) string {
	matchAll, _ := regexp.MatchString("^(./).*[/]$", arg)
	if !matchAll {
		matchOnlyBegin, _ := regexp.MatchString("^(./).*", arg)
		matchOnlyEnd, _ := regexp.MatchString("^[^./].*/$", arg)
		if matchOnlyBegin {
			arg += "/"
		}
		if matchOnlyEnd {
			arg = "./" + arg
		}
		if !matchOnlyBegin && !matchOnlyEnd {
			arg = "./" + arg + "/"
		}

	}
	return arg
}
