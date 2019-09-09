package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func sublistSearch(list *structs.LLNode, sub *structs.LLNode) int {
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

func SublistSearch(params []string) ([]string, error) {
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

	res := sublistSearch(list, sub)

	return utils.ToStringArray(res), nil
}
