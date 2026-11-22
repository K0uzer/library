// internal/library/book.go
package library
 
type Book struct {
    Title     string // экспортировано — нужно из main
    Author    string
    Year      int
    ISBN      string
    Available bool
}
 
func NewBook(title, author string, year int, isbn string) (*Book, error) {
	return 
}