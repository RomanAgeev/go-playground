package samples

import (
	"github.com/RomanAgeev/playground/sampleError"
	"github.com/RomanAgeev/playground/structs"
	"github.com/RomanAgeev/playground/utils"
)

func mergeLists(left *structs.LLNode, right *structs.LLNode) *structs.LLNode {
	pLeft := left
	pRight := right

	var result, pResult, p *structs.LLNode = nil, nil, nil

	for pLeft != nil && pRight != nil {
		if pLeft.Data.(int) <= pRight.Data.(int) {
			p = pLeft
			pLeft = pLeft.Next
		} else {
			p = pRight
			pRight = pRight.Next
		}

		if result == nil {
			result = p
		} else {
			pResult.Next = p
		}
		pResult = p
	}

	for pLeft != nil {
		pResult.Next = pLeft
		pResult = pLeft
		pLeft = pLeft.Next
	}

	for pRight != nil {
		pResult.Next = pRight
		pResult = pRight
		pRight = pRight.Next
	}

	return result
}

func mergeSortList(list *structs.LLNode) *structs.LLNode {
	if list == nil || list.Next == nil {
		return list
	}

	p, p1, p2 := list, list, list

	for p2 != nil && p2.Next != nil {
		p = p1
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	p.Next = nil

	left := mergeSortList(list)
	right := mergeSortList(p1)

	return mergeLists(left, right)
}

func MergeSortList(params []string) ([]string, error) {
	if len(params) != 1 {
		return nil, sampleError.InvalidParamCount(1, len(params))
	}

	listArr, err := utils.ToIntegerArray(params[0])
	if err != nil {
		return nil, err
	}

	var list *structs.LLNode = nil
	for i := len(listArr) - 1; i >= 0; i-- {
		list = structs.NewLLNode(list, listArr[i])
	}

	res := mergeSortList(list)

	resArr := make([]int, len(listArr))
	for i := 0; res != nil; i++ {
		resArr[i] = res.Data.(int)
		res = res.Next
	}

	return utils.ToStringArray(resArr...), nil
}
