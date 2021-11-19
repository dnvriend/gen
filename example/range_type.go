// Generated code; DO NOT EDIT.
package main

// to is exclusive
func Range(from, to int) RangeIterator {
	return RangeIterator{from, to}
}

type RangeIterator struct {
	from int
	to int
}

func (rcv RangeIterator) ForEach(fn func()) {
	for i:=rcv.from; i < rcv.to; i++ {
		fn()
	}
}

func (rcv RangeIterator) ForEachWithIndex(fn func(int)) {
	for i:=rcv.from; i < rcv.to; i++ {
		fn(i)
	}
}

func (rcv RangeIterator) ForEachWithLastFlag(fn func(bool)) {
	for i:=rcv.from; i < rcv.to; i++ {
		fn(i+1 == rcv.to)
	}
}

func (rcv RangeIterator) ForEachWithReturnFunc(fn func(), ret func(int) bool) {
	for i:=rcv.from; i < rcv.to; i++ {
		if ret(i) {
			return
		} else {
			fn()
		}
	}
}

func (rcv RangeIterator) ForEachWithIndexWithReturnFunc(fn func(int), ret func(int) bool) {
	for i:=rcv.from; i < rcv.to; i++ {
		if ret(i) {
			return
		} else {
			fn(i)
		}
	}
}
