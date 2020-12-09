package main

type Handlers interface {
	AvailableLearmingMaterials()
	AvailableBooks()
	AvailablePages()
	Commitment()
}
