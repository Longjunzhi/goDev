package 迭代器模式

func ExampleIterator() {
	var aggregate Aggregate
	aggregate = NewNumbers(1, 10)
	IteratorPrint(aggregate.Iterator())

}
