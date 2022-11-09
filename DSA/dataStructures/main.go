package dataStructures

type CustomStruct interface {
	GetValue(int) (int, error)
	GetValues() []int
	AddValue(int) (int, error)
	SearchValue(int) (int, error)
	RemoveValue(int) (int, error)
	AddValues(...int) (int, error)
}

type Structs struct {
	MinHeap    CustomStruct
	MaxHeap    CustomStruct
	Graph      CustomStruct
	Hashtable  CustomStruct
	LinkedList CustomStruct
	Queue      CustomStruct
	Stack      CustomStruct
	Tree       CustomStruct
}
