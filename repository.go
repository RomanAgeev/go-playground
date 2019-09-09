package main

import "github.com/RomanAgeev/playground/samples"

type sample = func(params []string) ([]string, error)

var repository = map[string]sample{
	"binarySearch":        samples.BinarySearch,
	"jumpSearch":          samples.JumpSearch,
	"interpolationSearch": samples.InterpolationSearch,
	"exponentialSearch":   samples.ExponentialSearch,
	"fibonacciSearch":     samples.FibonacciSearch,
	"sublistSearch1":      samples.SublistSearch_1,
	"sublistSearch2":      samples.SublistSearch_2,
}
