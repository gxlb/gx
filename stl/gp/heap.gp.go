package gp

//#GOGP_FILE_BEGIN

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)

////////////////////////////////////////////////////////////////////////////////

type GOGPGlobalNamePrefixHeap struct {
	b      []GOGPValueType         //data buffer
	limitN int                     //if limitN>0, heap size must<=limitN
	cmp    CmpGOGPGlobalNamePrefix //if top is max value
}

//new object
func NewGOGPHeapNamePrefixHeap(capacity int, maxTop bool) (r *GOGPGlobalNamePrefixHeap) {
	r = &GOGPGlobalNamePrefixHeap{}
	r.Init(capacity, 0, maxTop)
	return
}

//initialize
func (this *GOGPGlobalNamePrefixHeap) Init(capacity, limitN int, maxTop bool) {
	if cap(this.b) < capacity {
		this.b = make([]GOGPValueType, 0, capacity)
	}
	this.b = this.b[:0]

	this.limitN = limitN
	this.cmp = this.cmp.CreateByBool(!maxTop)
}

//push heap value
func (this *GOGPGlobalNamePrefixHeap) PushHeap(b []GOGPValueType, v GOGPValueType) []GOGPValueType {
	b = append(b, v)
	this.adjustUp(b, len(b)-1, v)
	return b
}

//pop heap top
func (this *GOGPGlobalNamePrefixHeap) PopHeap(b []GOGPValueType) (h []GOGPValueType, top GOGPValueType, ok bool) {
	l := len(b)
	if ok = l > 0; ok {
		top = b[0]
		if l > 1 {
			b[0], b[l-1] = b[l-1], b[0]
			this.adjustDown(b[:l-1], 0, b[0])
		}
		h = b[:l-1]
	}
	return
}

//check if b is a valid heap
func (this *GOGPGlobalNamePrefixHeap) CheckHeap(b []GOGPValueType) bool {
	for i := len(b) - 1; i > 0; i-- {
		p := this.parent(i)
		if !this.cmpV(b[i], b[p]) {
			return false
		}
	}
	return true
}

//adjust heap to select a proper hole to set v
func (this *GOGPGlobalNamePrefixHeap) adjustDown(b []GOGPValueType, hole int, v GOGPValueType) {
	size := len(b)
	//#GOGP_IFDEF GOGP_ImproveHeap
	//try to improve STL's adjust down algorithm
	//adjust heap to select a proper hole to set v
	for l := this.lchild(hole); l < size; l = this.lchild(hole) {
		c := l                                              //index that need compare with hole
		if r := l + 1; r < size && !this.cmpV(b[r], b[l]) { //let the most proper child to compare with v
			c = r
		}
		if this.cmpV(b[c], v) { //v is the most proper root, finish adjust
			break
		} else { //c is the most proper root, swap with hole, and continue adjust
			b[hole], hole = b[c], c
		}
	}
	b[hole] = v //put v to last hole
	//#GOGP_ELSE
	//C++ stl's adjust down algorithm
	//it seems to cost more move, to get probably less cmpare
	for l := this.lchild(hole); l < size; l = this.lchild(hole) {
		c := l                                              //index that need to be new root
		if r := l + 1; r < size && !this.cmpV(b[r], b[l]) { //let the most proper child to compare with v
			c = r
		}
		b[hole], hole = b[c], c
	}
	this.adjustUp(b, hole, v) //adjust up from leaf hole
	//#GOGP_ENDIF
}

//adjust heap to select a proper hole to set v
func (this *GOGPGlobalNamePrefixHeap) adjustUp(b []GOGPValueType, hole int, v GOGPValueType) {
	for hole > 0 {
		if parent := this.parent(hole); !this.cmpV(v, b[parent]) {
			b[hole], hole = b[parent], parent
		} else {
			break
		}
	}
	b[hole] = v //put v to last hole
}

//make b as a heap
func (this *GOGPGlobalNamePrefixHeap) MakeHeap(b []GOGPValueType) {
	if l := len(b); l > 1 {
		for i := l / 2; i >= 0; i-- {
			this.adjustDown(b, i, b[i])
		}
	}
}

//reverse order of b
func (this *GOGPGlobalNamePrefixHeap) Reverse(b []GOGPValueType) {
	l := len(b) - 1
	for i := l / 2; i >= 0; i-- {
		b[i], b[l-i] = b[l-i], b[i]
	}
}

//sort slice use heap algorithm
func (this *GOGPGlobalNamePrefixHeap) SortHeap(b []GOGPValueType) []GOGPValueType {
	this.MakeHeap(b)
	for t := b; len(t) > 1; {
		t, _, _ = this.PopHeap(t)
	}
	return b
}

//heap slice
func (this *GOGPGlobalNamePrefixHeap) Buffer() []GOGPValueType {
	return this.b
}

//push
func (this *GOGPGlobalNamePrefixHeap) Push(v GOGPValueType) {
	if this.limitN > 0 && this.Size() >= this.limitN {
		if top, ok := this.Top(); ok && this.cmpV(top, v) {
			this.Pop()
		}
	}
	this.b = this.PushHeap(this.b, v)
}

//pop
func (this *GOGPGlobalNamePrefixHeap) Pop() (top GOGPValueType, ok bool) {
	this.b, top, ok = this.PopHeap(this.b)
	return
}

//heap top
func (this *GOGPGlobalNamePrefixHeap) Top() (top GOGPValueType, ok bool) {
	if ok = !this.Empty(); ok {
		top = this.b[0]
	}
	return
}

//cmpare value
func (this *GOGPGlobalNamePrefixHeap) cmpV(c, p GOGPValueType) (ok bool) {
	ok = !this.cmp.F(p, c)
	return
}

//get parent index
func (this *GOGPGlobalNamePrefixHeap) parent(idx int) int {
	return (idx - 1) / 2
}

//get left child index
func (this *GOGPGlobalNamePrefixHeap) lchild(idx int) int {
	return 2*idx + 1
}

//get right child index
//func (this *GOGPHeapNamePrefixHeap) rchild(idx int) int {
//	return 2*idx + 2
//}

//func (this *GOGPHeapNamePrefixHeap) children(idx int) (l, r int) {
//	l = 2*idx + 1
//	r = l + 1
//	return
//}

//size
func (this *GOGPGlobalNamePrefixHeap) Size() int {
	return len(this.b)
}

//empty
func (this *GOGPGlobalNamePrefixHeap) Empty() bool {
	return this.Size() == 0
}

//#GOGP_FILE_END
