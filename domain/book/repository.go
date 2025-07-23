package book

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id uint) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(id uint) error
}
