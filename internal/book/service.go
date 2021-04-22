package book

import (
	"context"
	"time"

	"github.com/dbijay/simple-gorest-crud/internal/entity"
	"github.com/dbijay/simple-gorest-crud/pkg/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Book represents the data about a Book.
type Book struct {
	entity.Book
}

// Service encapsulates usecase logic for books.
type Service interface {
	Get(ctx context.Context, id string) (Book, error)
	Query(ctx context.Context, offset, limit int) ([]Book, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateBookRequest) (Book, error)
	Update(ctx context.Context, id string, input UpdateBookRequest) (Book, error)
	Delete(ctx context.Context, id string) (Book, error)
}

// CreateBookRequest represents a book creation request.
type CreateBookRequest struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Validate validates the CreateBookRequest fields.
func (b CreateBookRequest) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required, validation.Length(0, 128)),
		validation.Field(&b.Author, validation.Required, validation.Length(0, 128)),
		validation.Field(&b.Publication, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateBookRequest represents a book update request
type UpdateBookRequest struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Validate validate the UpdateBookRequest fields.
func (b UpdateBookRequest) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required, validation.Length(0, 128)),
		validation.Field(&b.Author, validation.Required, validation.Length(0, 128)),
		validation.Field(&b.Publication, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new book service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the book with specified ID.
func (s service) Get(ctx context.Context, id string) (Book, error) {
	book, err := s.repo.Get(ctx, id)
	if err != nil {
		return Book{}, err
	}
	return Book{book}, nil
}

// Create creates a new Book.
func (s service) Create(ctx context.Context, req CreateBookRequest) (Book, error) {
	if err := req.Validate(); err != nil {
		return Book{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Book{
		ID:          id,
		Name:        req.Name,
		Author:      req.Author,
		Publication: req.Publication,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		return Book{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the book with specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateBookRequest) (Book, error) {
	if err := req.Validate(); err != nil {
		return Book{}, err
	}

	book, err := s.Get(ctx, id)
	if err != nil {
		return book, err
	}
	book.Name = req.Name
	book.Author = req.Author
	book.Publication = req.Publication
	book.UpdatedAt = time.Now()
	if err := s.repo.Update(ctx, book.Book); err != nil {
		return book, err
	}
	return book, nil
}

// Delete deletes the book with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Book, error) {
	book, err := s.Get(ctx, id)
	if err != nil {
		return Book{}, err
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return Book{}, err
	}
	return book, nil
}

// Count returns the number of books.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

func (s service) Query(ctx context.Context, offset, limit int) ([]Book, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Book{}
	for _, item := range items {
		result = append(result, Book{item})
	}
	return result, nil
}
