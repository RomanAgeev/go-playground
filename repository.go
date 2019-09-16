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
	"selectionSort":       samples.SelectionSort,
	"insertionSort":       samples.InsertionSort,
	"bubbleSort":          samples.BubbleSort,
	"mergeSort":           samples.MergeSort,
	"mergeSortList":       samples.MergeSortList,
	"mergeSortIterative":  samples.MergeSortIterative,
	"quickSort":           samples.QuickSort,
	"quickSortIterative":  samples.QuickSortIterative,
	"tripleSort1":         samples.TripleSort1,
	"tripleSort2":         samples.TripleSort2,
	"heapSort":            samples.HeapSort,
	"countingSort":        samples.CountingSort,
	"countingSortInplace": samples.CountingSortInplace,
	"countingSortAZ":      samples.CountingSortAZ,
	"lsdSort":             samples.LSDSort,
	"lsdSortAZ":           samples.LSDSortAZ,
	"msdSortAZ":           samples.MSDSortAZ,
}
