// import "log"

// type myQueue struct {
// 	in  stack
// 	out stack
// }

// func (q *myQueue) add(item int) {
// 	q.in.push(item)
// }

// func (q *myQueue) remove() int {
// 	if q.out.isEmpty() {
// 		q.reshuffle()
// 	}
// 	return q.out.pop()
// }

// func (q *myQueue) peek() int {
// 	if q.isEmpty() {
// 		log.Fatal("queue is empty!")
// 	} else if q.out.isEmpty() {
// 		q.reshuffle()
// 	}
// 	return q.out.peek()
// }

// func (q *myQueue) reshuffle() {
// 	for !q.in.isEmpty() {
// 		q.out.push(q.in.pop())
// 	}
// }

// func (q *myQueue) isEmpty() bool {
// 	return len(q.in.data) == 0 && len(q.out.data) == 0
// }
