package models

import (
	"sync/atomic"
	"unsafe"
)

// Node 节点结构体
type Node struct {
	value Message
	next  *Node
}

// LockFreeQueue 无锁队列结构体
type LockFreeQueue struct {
	head *Node
	tail *Node
}

// NewLockFreeQueue 创建一个新的无锁队列
func NewLockFreeQueue() *LockFreeQueue {
	dummy := &Node{} // 创建一个哑节点
	return &LockFreeQueue{
		head: dummy,
		tail: dummy,
	}
}

// Enqueue 入队操作
func (q *LockFreeQueue) Enqueue(value Message) {
	newNode := &Node{value: value}
	for {
		tail := q.tail
		next := tail.next
		if tail == q.tail { // 检查 tail 是否未改变
			if next == nil { // tail 是最后一个节点
				if atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&tail.next)),
					unsafe.Pointer(next),
					unsafe.Pointer(newNode),
				) {
					// 将 tail 移动到新的节点
					atomic.CompareAndSwapPointer(
						(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
						unsafe.Pointer(tail),
						unsafe.Pointer(newNode),
					)
					return
				}
			} else {
				// 将 tail 移动到 next
				atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
					unsafe.Pointer(tail),
					unsafe.Pointer(next),
				)
			}
		}
	}
}

// Dequeue 出队操作
func (q *LockFreeQueue) Dequeue() (Message, bool) {
	for {
		head := q.head
		tail := q.tail
		next := head.next
		if head == q.head { // 检查 head 是否未改变
			if head == tail {
				if next == nil {
					// 队列为空
					return Message{}, false
				}
				// 将 tail 移动到 next
				atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&q.tail)),
					unsafe.Pointer(tail),
					unsafe.Pointer(next),
				)
			} else {
				value := next.value
				if atomic.CompareAndSwapPointer(
					(*unsafe.Pointer)(unsafe.Pointer(&q.head)),
					unsafe.Pointer(head),
					unsafe.Pointer(next),
				) {
					return value, true
				}
			}
		}
	}
}
