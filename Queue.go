package gotool

import (
	"sync"

	LogTool "github.com/adimax2953/log-tool"
)

//***********************************************
//實現一個隊例, 需要有三個接口實現先進先出的Queue
//添加元素（從隊例尾部添加一個元素）
//彈出元素 (從隊首彈出一個元素)
//遍歷函數（傳入回調函數, 返回每一個元素）
//刪除元素（傳入元素, 從隊例中刪除這個元素）
//***********************************************

//先定義隊例節點
type QueueNode struct {
	value interface{}
	prev  *QueueNode
	next  *QueueNode
}

func (qn *QueueNode) Value() interface{} {
	return qn.value
}

func (qn *QueueNode) Next() *QueueNode {
	return qn.next
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
	size uint32 //隊例長度

	capacity uint32 //隊例容量, 如果不設默認為零, 則無限制

	lockDel sync.Mutex
	lockAdd sync.Mutex
}

func (q *Queue) Init() {
	q.Clear()
	q.capacity = 0
}

func (q *Queue) Size() uint32 {
	return q.size
}

//一開始就要設置, 如果已經有了內容, 則只能設置為和當前容量一樣大小
func (q *Queue) SetCapactiy(c uint32) uint32 {
	if q.size > c {
		q.capacity = q.size
	} else {
		q.capacity = c
	}

	return q.capacity

}

func (q *Queue) PushBack(elem interface{}) {
	q.lockAdd.Lock()
	defer q.lockAdd.Unlock()
	newNode := &QueueNode{value: elem}

	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		newNode.prev = q.tail
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++

	if q.size > 20 {
		LogTool.LogDebug("queue size is ", q.size)
		LogTool.LogDebug("queue capacity is ", q.capacity)
	}

	//如果設置了容量, 並且大小已經超過容量, 則前面的要彈出去
	if q.capacity > 0 && q.size > q.capacity {
		LogTool.LogDebug("queue size is", q.size, " need queue execute pop...")
		q.PopFront()
		LogTool.LogDebug("size of after queue pop", q.size)
	}

	newNode = nil
}

//取得首節點, 外部遍歷需要
func (q *Queue) Begin() *QueueNode {
	return q.head
}

//取得尾節點, 外部遍歷需要
func (q *Queue) End() *QueueNode {
	return nil
}

//取得尾元素
func (q *Queue) Back() interface{} {
	return q.tail.value
}

//取得首元素
func (q *Queue) Front() interface{} {
	return q.head.value
}

//彈出首元素
func (q *Queue) PopFront() interface{} {
	q.lockDel.Lock()
	defer q.lockDel.Unlock()

	var ret interface{}
	if q.size > 0 {
		ret = q.head.value
		q.head = q.head.next
		q.size--
	}

	return ret
}

//傳入一個元素, 如果鍊錶中有, 即移除
func (q *Queue) Remove(ele interface{}) bool {
	q.lockDel.Lock()
	defer q.lockDel.Unlock()

	ret := false

	checkNode := q.head
	for checkNode != nil {
		if checkNode.value == ele {
			//如果剛好是表頭, 則把表頭的下一節點作為新表頭
			//如果下一節點為空, 則表示已經是空表。如果下一節點有值, 那麼作為表頭prev節點清空
			if q.head == checkNode {
				q.head = checkNode.next
				if q.head != nil {
					q.head.prev = nil
				}
			} else if q.tail == checkNode { //如果剛好是表尾
				q.tail = checkNode.prev
				q.tail.next = nil
			} else { //前後的鏈節點直接相連
				checkNode.prev.next = checkNode.next
				checkNode.next.prev = checkNode.prev
			}

			ret = true
			q.size--
			break
		}

		checkNode = checkNode.next
	}

	return ret
}

//遍歷鍊錶
func (q *Queue) Range(callback func(v interface{}) bool) {
	for node := q.Begin(); node != q.End(); node = node.next {
		if callback(node.value) {
			//返回真即退出, 不再遍歷
			return
		}

	}
}

func (q *Queue) Clear() {
	q.head = nil
	q.tail = nil
	q.size = 0

	//清除並不重置capacity
	//q.capacity = 0
}

func (q *Queue) ShowSelf() {
	LogTool.LogDebug("Print Queue:")
	q.Range(func(v interface{}) bool {
		LogTool.LogDebug("", v)
		return false
	})
}

func CreateQueue() *Queue {
	q := &Queue{}
	q.Init()

	return q
}
