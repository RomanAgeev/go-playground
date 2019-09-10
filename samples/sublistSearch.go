package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

type sublistSearchImpl = func(list *structs.LLNode, sub *structs.LLNode) int

func sublistSearchImpl_1(list *structs.LLNode, sub *structs.LLNode) int {
	if list == nil || sub == nil {
		return -1
	}

	listP := list
	subP := sub

	var p *structs.LLNode = nil

	m := false

	index, mIndex, pIndex := 0, 0, 0

	for {
		if m {
			switch {
			case subP == nil:
				return mIndex
			case listP == nil:
				return -1
			case listP.Data == subP.Data:
				if listP.Data == sub.Data {
					p = listP
					pIndex = index
				}
				listP = listP.Next
				subP = subP.Next
				index++
			case listP.Data != subP.Data:
				if p != nil {
					listP = p
					p = nil
					index = pIndex
				}
				subP = sub
				m = false
			}
		} else {
			switch {
			case listP == nil:
				return -1
			case listP.Data == subP.Data:
				subP = subP.Next
				mIndex = index
				m = true
			}
			listP = listP.Next
			index++
		}
	}
}

func sublistSearchImpl_2(list *structs.LLNode, sub *structs.LLNode) int {
	if list == nil || sub == nil {
		return -1
	}

	listP := list
	subP := sub

	var p *structs.LLNode = nil

	m := false

	index, pIndex, ppIndex := 0, 0, 0

	for listP != nil {
		pp := listP
		ppIndex = index

		for subP != nil {
			if listP == nil || listP.Data != subP.Data {
				m = false
				break
			}

			if listP.Data == sub.Data {
				if m {
					p = listP
					pIndex = index
				} else {
					m = true
				}
			}

			listP = listP.Next
			subP = subP.Next
			index++
		}

		if subP == nil {
			return ppIndex
		}
		if listP == nil {
			return -1
		}

		if p != nil {
			listP = p
			index = pIndex
			p = nil
		} else {
			listP = pp.Next
			index = ppIndex + 1
		}
		subP = sub
	}
	return -1
}

func sublistSearch(impl sublistSearchImpl, params []string) ([]string, error) {
	if len(params) != 2 {
		return nil, sampleError.InvalidParamCount(2, len(params))
	}

	listArr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	subArr, err := utils.ToIntegerArray(params[1])
	if err != nil {
		return nil, err
	}

	var list *structs.LLNode = nil
	for i := len(listArr) - 1; i >= 0; i-- {
		list = structs.NewLLNode(list, listArr[i])
	}

	var sub *structs.LLNode = nil
	for i := len(subArr) - 1; i >= 0; i-- {
		sub = structs.NewLLNode(sub, subArr[i])
	}

	res := impl(list, sub)

	return utils.ToStringArray(res), nil
}

func SublistSearch_1(params []string) ([]string, error) {
	return sublistSearch(sublistSearchImpl_1, params)
}

func SublistSearch_2(params []string) ([]string, error) {
	return sublistSearch(sublistSearchImpl_2, params)
}
