package arrays

import "errors"

type MinHeap struct {
	Length int
	data   []int
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return idx*2 + 1
}

func rightChild(idx int) int {
	return idx*2 + 2
}

func (mh *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	v := mh.data[idx]

	p := parent(idx)
	pv := mh.data[p]

	if pv > v {
		swap(&mh.data, p, idx)
		mh.heapifyUp(p)
	}
}

func (mh *MinHeap) heapifyDown(idx int) {
	lIdx := leftChild(idx)
	rIdx := rightChild(idx)

	if idx >= mh.Length || lIdx >= mh.Length {
		return
	}

	lv := mh.data[lIdx]
	rv := mh.data[rIdx]
	v := mh.data[idx]

	if lv > rv && v > rv {
		swap(&mh.data, idx, rIdx)
		mh.heapifyDown(rIdx)
	} else if rv > lv && v > lv {
		swap(&mh.data, idx, lIdx)
		mh.heapifyDown(lIdx)
	}
}

func (mh *MinHeap) Insert(val int) {
	mh.data = append(mh.data, val)
	mh.heapifyUp(mh.Length)
	mh.Length++
}

var ErrDeleteFromEmpty = errors.New("cannot delete from empty heap")

func (mh *MinHeap) Delete() (int, error) {
	if mh.Length == 0 {
		return -1, ErrDeleteFromEmpty
	}

	out := mh.data[0]
	mh.Length--

	if mh.Length == 0 {
		mh.data = []int{}
		return out, nil
	}

	mh.data[0] = mh.data[mh.Length]
	mh.heapifyDown(0)
	return out, nil
}
