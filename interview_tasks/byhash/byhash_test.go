package main

import (
	"os"
	"testing"
)

var lstests = []struct {
	in  []string
	out int
}{
	{
		[]string{}, 1,
	},
	{
		[]string{"./one/"}, 0,
	},
	{
		[]string{"./one/", "./two/"}, 0,
	},
	{
		[]string{"./one", "./two"}, 0,
	},

	{
		[]string{"/one", "./two"}, 2,
	},
	{
		[]string{"one", "two/"}, 0,
	},
	{
		[]string{"./one", "two"}, 0,
	},
	{
		[]string{"one", "two"}, 0,
	},
}

func TestHandleLsInput(t *testing.T) {

	for _, args := range lstests {
		for _, dir := range args.in {
			os.MkdirAll(dir, os.ModePerm)
		}
		code, _ := HandleLsInput(args.in)
		if code != args.out {
			t.Errorf("HandleLsInput(%v) != %v", args.in, args.out)
		}

		for _, dir := range args.in {
			os.RemoveAll(dir)
		}
	}

}

var getTests = []struct {
	in  []string
	out int
}{
	{
		[]string{}, 1,
	},
	{
		[]string{"test"}, 0,
	},
	{
		[]string{"test", "test2"}, 0,
	},
}

func TestHandleGetInput(t *testing.T) {
	for _, args := range getTests {
		for _, file := range args.in {
			f, _ := os.Create(file)
			defer f.Close()
		}
		code, _ := HandleGetInput(args.in, 100)
		if code != args.out {
			t.Errorf("HandleLsInput(%v) != %v", args.in, args.out)
		}

		for _, file := range args.in {
			os.Remove(file)
		}
	}
}

func TestHandleGetInputWithSizeFlag(t *testing.T) {
	code, _ := HandleGetInput([]string{"byhash_test.go"}, 20)
	if code != 0 {
		t.Errorf("HandleLsInput(%v) != %v", 0, code)
	}
}

