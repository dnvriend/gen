// Generated code; DO NOT EDIT.
package main

// to is exclusive
func Range(from, to int) RangeIterator {
	return RangeIterator{
		from: from,
		to:   to,
		step: 1,
	}
}

type RangeIterator struct {
	from int
	to int
	step int
}

func (rcv RangeIterator) WithStep(step int) RangeIterator {
	rcv.step = step
	return rcv
}

func (rcv RangeIterator) ForEach(fn func()) {
	for i:=rcv.from; i < rcv.to; i = i + rcv.step {
		fn()
	}
}

func (rcv RangeIterator) ForEachWithIndex(fn func(int)) {
	for i:=rcv.from; i < rcv.to; i = i + rcv.step {
		fn(i)
	}
}

func (rcv RangeIterator) ForEachWithLastFlag(fn func(bool)) {
	for i:=rcv.from; i < rcv.to; i = i + rcv.step {
		fn(i+1 == rcv.to)
	}
}

func (rcv RangeIterator) ForEachWithReturnFunc(fn func(), ret func(int) bool) {
	for i:=rcv.from; i < rcv.to; i = i + rcv.step {
		if ret(i) {
			return
		} else {
			fn()
		}
	}
}

func (rcv RangeIterator) ForEachWithIndexWithReturnFunc(fn func(int), ret func(int) bool) {
	for i:=rcv.from; i < rcv.to; i = i + rcv.step {
		if ret(i) {
			return
		} else {
			fn(i)
		}
	}
}

func (rcv RangeIterator) ToIntList() IntList {
	xs := make([]int, 0)
	rcv.ForEachWithIndex(func(i int) {
		xs = append(xs, i)
	})
	return xs
}
