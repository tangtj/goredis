package list

type List struct {
	//listNode *head;
	//listNode *tail;
	//void *(*dup)(void *ptr);
	//void (*free)(void *ptr);
	//int (*match)(void *ptr, void *key);
	//unsigned int len;
	head *Node
	tail *Node
	len  int
}

//list *listCreate(void);
//void listRelease(list *list);
//list *listAddNodeHead(list *list, void *value);
//list *listAddNodeTail(list *list, void *value);
//list *listInsertNode(list *list, listNode *old_node, void *value, int after);
//void listDelNode(list *list, listNode *node);
//listIter *listGetIterator(list *list, int direction);
//listNode *listNext(listIter *iter);
//void listReleaseIterator(listIter *iter);
//list *listDup(list *orig);
//listNode *listSearchKey(list *list, void *key);
//listNode *listIndex(list *list, int index);
//void listRewind(list *list, listIter *li);
//void listRewindTail(list *list, listIter *li);

func (l *List) AddNodeHead(value interface{}) *List {
	node := &Node{
		prev:  nil,
		next:  nil,
		value: value,
	}
	if l.len <= 0 {
		l.head, l.tail = node, node
	} else {
		l.head.prev = node
		node.next = l.head
		l.head = node
	}
	l.len++
	return l
}

func (l *List) AddNodeTail(value interface{}) *List {
	node := &Node{
		prev:  nil,
		next:  nil,
		value: value,
	}
	if l.len <= 0 {
		l.head, l.tail = node, node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.len++
	return l
}

func (l *List) DelNode(node *Node) *List {
	if node == nil {
		return l
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
	l.len--
	return l
}

func (l *List) GetIdx(idx int) interface{} {
	if idx < 0 || idx >= l.len {
		return nil
	}
	node := l.head
	for i := 0; i < idx; i++ {
		node = node.next
	}
	return node.value
}

func (l *List) GetNode(idx int) *Node {
	if idx < 0 || idx >= l.len {
		return nil
	}
	node := l.head
	for i := 0; i < idx; i++ {
		node = node.next
	}
	return node
}

func (l List) GetLen() uint {
	return uint(l.len)
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
		len:  0,
	}
}
