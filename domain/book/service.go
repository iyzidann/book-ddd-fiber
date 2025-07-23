package book

type Service interface {
	GetBooks() ([]Book, error)
	GetBook(id uint) (Book, error)
	CreateBook(book Book) (Book, error)
	UpdateBook(book Book) (Book, error)
	DeleteBook(id uint) error
}

type bookService struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &bookService{r}
}

func (s *bookService) GetBooks() ([]Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBook(id uint) (Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(book Book) (Book, error) {
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book Book) (Book, error) {
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
