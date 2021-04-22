package book

import (
	"context"

	"github.com/dbijay/simple-gorest-crud/internal/entity"
	"github.com/dbijay/simple-gorest-crud/pkg/dbcontext"
	"github.com/dbijay/simple-gorest-crud/pkg/log"
)

type Repository interface {

	// Get returns the book of the specified book ID from database.
	Get(ctx context.Context, id string) (entity.Book, error)

	// Count returns total number of book available in database.
	Count(ctx context.Context) (int, error)

	// Query return list of the albums from database with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Book, error)

	// Create saves new book in the database.
	Create(ctx context.Context, book entity.Book) error

	// Update updates the book with the given ID in the database.
	Update(ctx context.Context, book entity.Book) error

	// Delete removes the book with given ID from database.
	Delete(ctx context.Context, id string) error
}

// repository persists books in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new book repository.
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the book with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Book, error) {
	var book entity.Book
	err := r.db.With(ctx).Select().Model(id, &book)
	return book, err
}

// Create saves a new book record in the database.
// It returns the ID of newly inserted book record.
func (r repository) Create(ctx context.Context, book entity.Book) error {
	return r.db.With(ctx).Model(&book).Insert()
}

// Update saves the changes to an book to the database.
func (r repository) Update(ctx context.Context, book entity.Book) error {
	return r.db.With(ctx).Model(&book).Update()
}

// Delete deletes a book with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	book, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&book).Delete()
}

// Count returns the number of book records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("book").Row(&count)
	return count, err
}

// Query retrives the album records with the specified offset
// and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.With(ctx).Select().OrderBy("id").
		Offset(int64(offset)).Limit(int64(limit)).All(&books)
	return books, err
}
