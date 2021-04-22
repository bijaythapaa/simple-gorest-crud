package controllers

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestGetBook(t *testing.T) {
// 	req, err := http.NewRequest("Get", "/book/", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetBook)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	// check response body what we expected.
// 	// expected := `[{"ID":2,"CreatedAt":"2021-04-16T07:51:32.076+05:45","UpdatedAt":"2021-04-16T07:55:45.773+05:45","DeletedAt":null,"name":"Metamorphosis","author":"Franz Kafka","publication":"Goodreads"},{"ID":3,"CreatedAt":"2021-04-16T07:54:29.904+05:45","UpdatedAt":"2021-04-16T07:54:29.904+05:45","DeletedAt":null,"name":"The Prophet","author":"Kahlil Gibran","publication":"Penguine"},{"ID":5,"CreatedAt":"2021-04-16T08:05:15.385+05:45","UpdatedAt":"2021-04-16T08:05:15.385+05:45","DeletedAt":null,"name":"Abstract Chintan Pyaaj","author":"Shanker Lamichhane","publication":"BookHill"},{"ID":7,"CreatedAt":"2021-04-16T12:40:28.564+05:45","UpdatedAt":"2021-04-16T12:40:28.564+05:45","DeletedAt":null,"name":"To Kill a Mockingbird","author":"Harper Lee","publication":"J. B. Lippincott \u0026 Co."},{"ID":8,"CreatedAt":"2021-04-19T12:08:27.442+05:45","UpdatedAt":"2021-04-19T12:08:27.442+05:45","DeletedAt":null,"name":"Karnali Blues","author":"Buddhisagar","publication":"Book Hill"},{"ID":9,"CreatedAt":"2021-04-19T12:11:12.142+05:45","UpdatedAt":"2021-04-19T12:11:12.142+05:45","DeletedAt":null,"name":"FirFire","author":"Buddhisagar","publication":"Book Hill"},{"ID":10,"CreatedAt":"2021-04-19T12:11:36.893+05:45","UpdatedAt":"2021-04-19T12:11:36.893+05:45","DeletedAt":null,"name":"FirFire","author":"Buddhisagar","publication":"Book Hill"},{"ID":11,"CreatedAt":"2021-04-19T12:14:00.617+05:45","UpdatedAt":"2021-04-19T12:14:00.617+05:45","DeletedAt":null,"name":"FirFire","author":"Buddhisagar","publication":"Book Hill"},{"ID":12,"CreatedAt":"2021-04-19T12:24:46.894+05:45","UpdatedAt":"2021-04-19T12:24:46.894+05:45","DeletedAt":null,"name":"FirFire","author":"Buddhisagar","publication":"Book Hill"},{"ID":13,"CreatedAt":"2021-04-19T12:25:40.17+05:45","UpdatedAt":"2021-04-19T12:25:40.17+05:45","DeletedAt":null,"name":"FirFire","author":"Buddhisagar","publication":"Book Hill"},{"ID":14,"CreatedAt":"2021-04-19T12:27:29.145+05:45","UpdatedAt":"2021-04-19T12:27:29.145+05:45","DeletedAt":null,"name":"FirFire","author":"Buddhisagar","publication":"Book Hill"},{"ID":15,"CreatedAt":"2021-04-19T12:32:56.399+05:45","UpdatedAt":"2021-04-19T12:32:56.399+05:45","DeletedAt":null,"name":"FirFire","author":"Buddhisagar","publication":"Book Hill"},{"ID":16,"CreatedAt":"2021-04-19T12:34:37.982+05:45","UpdatedAt":"2021-04-19T12:34:37.982+05:45","DeletedAt":null,"name":"FirFire","author":"Buddhisagar","publication":"Book Hill"}]`
// 	// if rr.Body.String() != expected {
// 	// 	t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
// 	// }

// }

// func TestGetBookById(t *testing.T) {
// 	bookId := "3"
// 	req, err := http.NewRequest("Get", "/book/"+bookId, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// q := req.URL.Query()
// 	// q.Add("bookId", "2")
// 	// req.URL.RawQuery = q.Encode()

// 	rr := httptest.NewRecorder()
// 	// ts := httptest.NewServer(http.HandlerFunc(GetBookById))
// 	handler := http.HandlerFunc(GetBookById)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler return wrong status code: got: %v \n want: %v", status, http.StatusOK)
// 	}

// 	// expected := `{"ID":3,"CreatedAt":"2021-04-16T07:54:29.904+05:45","UpdatedAt":"2021-04-16T07:54:29.904+05:45","DeletedAt":null,"name":"The Prophet","author":"Kahlil Gibran","publication":"Penguine"}`
// 	// if rr.Body.String() != expected {
// 	// 	t.Errorf("handler returned unexpected body: got: %v \n want: %v", rr.Body.String(), expected)
// 	// }
// }

// func TestCreateBook(t *testing.T) {
// 	var jsonStr = []byte(`{"Name": "FirFire", "Author": "Buddhisagar", "Publication": "Book Hill"}`)
// 	req, err := http.NewRequest("POST", "/book/", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(CreateBook)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}
// 	// expected := `{"ID":9,"CreatedAt":"2021-04-19T12:08:27.442+05:45","UpdatedAt":"2021-04-19T12:08:27.442+05:45","DeletedAt":null,"name":"FirFire","author":"Buddhisagar","publication":"Book Hill"}`
// 	// if rr.Body.String() != expected {
// 	// 	t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
// 	// }
// }

// func TestUpdateBook(t *testing.T) {
// 	bookId := "3"
// 	var jsonStr = []byte(`{"name": "Norwegian Wood","Publication": "Goodreads"}`)
// 	req, err := http.NewRequest("PUT", "/book/"+bookId, bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(UpdateBook)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	// expected := `{"ID": 21,"CreatedAt": "2021-04-19T13:08:24.402+05:45","UpdatedAt": "2021-04-19T13:08:24.401+05:45","DeletedAt": null,"name": "Norwegian Wood","author": "","publication": "Goodreads"}`
// 	// if rr.Body.String() != expected {
// 	// 	t.Errorf("handler returned unexpected body: got %v want %v",
// 	// 		rr.Body.String(), expected)
// 	// }
// }

// func TestDeleteBook(t *testing.T) {

// 	bookId := "3"
// 	// var jsonStr = []byte(`{"name": "Norwegian Wood","Publication": "Goodreads"}`)
// 	req, err := http.NewRequest("PUT", "/book/"+bookId, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// q := req.URL.Query()
// 	// q.Add("id", "4")
// 	// req.URL.RawQuery = q.Encode()

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(DeleteBook)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := `"Deleted Sucessfully"`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }
