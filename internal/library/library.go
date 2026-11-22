// internal/library/library.go
package library
 
type Library struct {
    name    string    // неэкспортировано — используем через методы
    address string
    books   []Book    // неэкспортировано — доступ через AddBook/FindBook
    members []Member
}
 
func NewLibrary(name, address string) *Library {
	return 
}
func (lib *Library) AddBook(b Book) {
	return
}
func (lib *Library) FindBook(title string) (*Book, bool) {
	return
}
func (lib *Library) BorrowBook(memberEmail, bookTitle string) error {
	return
}
 
func (lib *Library) findMember(email string) (*Member, bool) {
	return
} // неэкспортировано