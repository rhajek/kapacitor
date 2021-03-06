// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: circularqueue.gen.go.tmpl

//lint:file-ignore U1000 this is generated code
package kapacitor

// srcPointCircularQueue defines a circular queue, always use the contructor to create one.
type srcPointCircularQueue struct {
	data []srcPoint
	head int
	tail int
	l    int
}

// newSrcPointconstructs a Circular Queue
// with a given buffer buf.  It is ok for buf to be nil.
func newSrcPointCircularQueue(buf ...srcPoint) *srcPointCircularQueue {
	// if we have a useless buffer, make one that is at least useful
	if cap(buf) < 4 {
		buf = append(make([]srcPoint, 0, 4), buf...)
	}
	return &srcPointCircularQueue{
		data: buf[:cap(buf)],
		tail: len(buf), // tail is here we insert
		l:    len(buf),
	}
}

// Enqueue adds an item to the queue.
func (q *srcPointCircularQueue) Enqueue(v srcPoint) {
	// if full we must grow and insert together. This is an expensive op
	if cap(q.data) > q.l { // no need to grow
		if q.tail == len(q.data) {
			q.tail = 0
		}
		q.data[q.tail] = v
	} else { // we need to grow
		buf := make([]srcPoint, cap(q.data)*2)
		if q.head < q.tail {
			copy(buf, q.data[q.head:q.tail])
		} else {
			partialWriteLen := copy(buf, q.data[q.head:])
			copy(buf[partialWriteLen:], q.data[:q.tail])
		}
		q.head = 0
		q.tail = cap(q.data)
		buf[q.tail] = v
		q.data = buf
	}
	q.l++
	q.tail++
	return
}

// Dequeue removes n items from the queue. If n is longer than the number of the items in the queue it will clear them all out.
func (q *srcPointCircularQueue) Dequeue(n int) {
	if n <= 0 {
		return
	}
	if q.l <= n {
		n = q.l
	}
	ni := n
	var fill srcPoint
	if q.head > q.tail {
		for i := q.head; i < len(q.data) && ni > 0; i++ {
			q.data[i] = fill
			ni--
		}
		for i := 0; i < q.tail && ni > 0; i++ {
			q.data[i] = fill
			ni--
		}
	} else {
		for i := q.head; i < q.tail && ni > 0; i++ {
			q.data[i] = fill
			ni--
		}
	}
	q.head += n
	if q.head > len(q.data) {
		q.head -= len(q.data)
	}
	q.l -= n
	if q.l == 0 {
		q.head = 0
		q.tail = 0
	}
	return
}

// Peek peeks i ahead of the current head of queue.  It should be used in conjunction with .Len() to prevent a panic.
func (q *srcPointCircularQueue) Peek(i int) srcPoint {
	if i < 0 || i >= q.l {
		panic("peek index is out of bounds")
	}
	p := q.head + i

	if p >= len(q.data) {
		p -= len(q.data)
	}
	return q.data[p]
}

// Len returns the current number of items in the queue.
func (q *srcPointCircularQueue) Len() int {
	return q.l
}

// joinsetPtrCircularQueue defines a circular queue, always use the contructor to create one.
type joinsetPtrCircularQueue struct {
	data []joinsetPtr
	head int
	tail int
	l    int
}

// newJoinsetPtrconstructs a Circular Queue
// with a given buffer buf.  It is ok for buf to be nil.
func newJoinsetPtrCircularQueue(buf ...joinsetPtr) *joinsetPtrCircularQueue {
	// if we have a useless buffer, make one that is at least useful
	if cap(buf) < 4 {
		buf = append(make([]joinsetPtr, 0, 4), buf...)
	}
	return &joinsetPtrCircularQueue{
		data: buf[:cap(buf)],
		tail: len(buf), // tail is here we insert
		l:    len(buf),
	}
}

// Enqueue adds an item to the queue.
func (q *joinsetPtrCircularQueue) Enqueue(v joinsetPtr) {
	// if full we must grow and insert together. This is an expensive op
	if cap(q.data) > q.l { // no need to grow
		if q.tail == len(q.data) {
			q.tail = 0
		}
		q.data[q.tail] = v
	} else { // we need to grow
		buf := make([]joinsetPtr, cap(q.data)*2)
		if q.head < q.tail {
			copy(buf, q.data[q.head:q.tail])
		} else {
			partialWriteLen := copy(buf, q.data[q.head:])
			copy(buf[partialWriteLen:], q.data[:q.tail])
		}
		q.head = 0
		q.tail = cap(q.data)
		buf[q.tail] = v
		q.data = buf
	}
	q.l++
	q.tail++
	return
}

// Dequeue removes n items from the queue. If n is longer than the number of the items in the queue it will clear them all out.
func (q *joinsetPtrCircularQueue) Dequeue(n int) {
	if n <= 0 {
		return
	}
	if q.l <= n {
		n = q.l
	}
	ni := n
	var fill joinsetPtr
	if q.head > q.tail {
		for i := q.head; i < len(q.data) && ni > 0; i++ {
			q.data[i] = fill
			ni--
		}
		for i := 0; i < q.tail && ni > 0; i++ {
			q.data[i] = fill
			ni--
		}
	} else {
		for i := q.head; i < q.tail && ni > 0; i++ {
			q.data[i] = fill
			ni--
		}
	}
	q.head += n
	if q.head > len(q.data) {
		q.head -= len(q.data)
	}
	q.l -= n
	if q.l == 0 {
		q.head = 0
		q.tail = 0
	}
	return
}

// Peek peeks i ahead of the current head of queue.  It should be used in conjunction with .Len() to prevent a panic.
func (q *joinsetPtrCircularQueue) Peek(i int) joinsetPtr {
	if i < 0 || i >= q.l {
		panic("peek index is out of bounds")
	}
	p := q.head + i

	if p >= len(q.data) {
		p -= len(q.data)
	}
	return q.data[p]
}

// Len returns the current number of items in the queue.
func (q *joinsetPtrCircularQueue) Len() int {
	return q.l
}
