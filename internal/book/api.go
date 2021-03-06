package book

import (
	"net/http"

	"github.com/dbijay/simple-gorest-crud/internal/errors"
	"github.com/dbijay/simple-gorest-crud/pkg/log"
	"github.com/dbijay/simple-gorest-crud/pkg/pagination"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type resource struct {
	service Service
	logger  log.Logger
}

// RequestHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler,
	logger log.Logger) {
	res := resource{service, logger}

	r.Get("/books/<id>", res.get)
	r.Get("/books", res.query)

	r.Use(authHandler)

	// the following endpoints requires a valid JWT
	r.Post("/books", res.create)
	r.Put("/books/<id>", res.update)
	r.Delete("/books/<id>", res.delete)
}

func (r resource) get(c *routing.Context) error {
	book, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}
	return c.Write(book)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	books, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = books
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateBookRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	book, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(book, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateBookRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	book, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}
	return c.Write(book)
}

func (r resource) delete(c *routing.Context) error {
	book, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}
	return c.Write(book)
}
