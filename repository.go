package main

type Operations interface {
	GetAllAviableLearmingMaterials()
	GetAllAviableBooks()
	GetAllAviablePages()
	Commitment()

	createUser()
	createLearningMaterial()
	createBook()
	createPage()
	readUser()
	readLearningMaterial()
	readBook()
	readPage()
	updateUser()
	updateLearningMaterial()
	updateBook()
	updatePage()
	deleteUser()
	deleteLearningMaterial()
	deleteBook()
	deletePage()
}
