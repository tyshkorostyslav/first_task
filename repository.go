package main

type Operations interface {
	CreateUser()
	CreateLearningMaterial()
	CreateBook()
	CreatePage()
	ReadUser()
	ReadLearningMaterial()
	ReadBook()
	ReadPage()
	UpdateUser()
	UpdateLearningMaterial()
	UpdateBook()
	UpdatePage()
	DeleteUser()
	DeleteLearningMaterial()
	DeleteBook()
	DeletePage()
}
