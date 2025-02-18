package graphs

import "robtrizzo/data-and-algo-practice/arrays"

type vertex struct {
	val      int
	distance int
}

type vertexMinHeap struct {
	Length int
	Data   []vertex
	Index  map[int]int
}

func NewVertexMinHeap() vertexMinHeap {
	return vertexMinHeap{
		Length: 0,
		Data:   []vertex{},
		Index:  make(map[int]int),
	}
}

func (vmh *vertexMinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	v := vmh.Data[idx].distance

	p := arrays.Parent(idx)
	pv := vmh.Data[p].distance

	if pv > v {
		vmh.swap(p, idx)
		vmh.heapifyUp(p)
	}
}

func (vmh *vertexMinHeap) heapifyDown(idx int) {
	lIdx := arrays.LeftChild(idx)
	rIdx := arrays.RightChild(idx)

	if idx >= vmh.Length || lIdx >= vmh.Length {
		return
	}

	lv := vmh.Data[lIdx].distance
	rv := vmh.Data[rIdx].distance
	v := vmh.Data[idx].distance

	if lv > rv && v > rv {
		vmh.swap(idx, rIdx)
		vmh.heapifyDown(rIdx)
	} else if rv > lv && v > lv {
		vmh.swap(idx, lIdx)
		vmh.heapifyDown(lIdx)
	}
}

func (vmh *vertexMinHeap) Insert(val vertex) {
	vmh.Data = append(vmh.Data, val)
	vmh.Index[val.val] = vmh.Length
	vmh.heapifyUp(vmh.Length)
	vmh.Length++
}

func (vmh *vertexMinHeap) Delete() (vertex, error) {
	if vmh.Length == 0 {
		return vertex{}, arrays.ErrDeleteFromEmpty
	}

	out := vmh.Data[0]
	vmh.Length--

	if vmh.Length == 0 {
		vmh.Data = []vertex{}
		return out, nil
	}

	delete(vmh.Index, out.val)
	vmh.Data[0] = vmh.Data[vmh.Length]
	vmh.heapifyDown(0)
	return out, nil
}

func (vmh *vertexMinHeap) swap(i, j int) {
	arrays.Swap(&vmh.Data, i, j)
	vmh.Index[vmh.Data[i].val] = i
	vmh.Index[vmh.Data[j].val] = j
}
