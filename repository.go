package main

import "github.com/RomanAgeev/playground/samples"

type sample = func(params []string) ([]string, error)

var repository = map[string]sample{
	"binarySearch":        samples.BinarySearch,
	"jumpSearch":          samples.JumpSearch,
	"interpolationSearch": samples.InterpolationSearch,
}
