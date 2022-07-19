package main

type BookType int

const (
	eBook BookType = iota
	Audiobook
	Hardback
	Paperback
	Encyclopedia
	Magazine
	Comic
)

type Book interface {
	Booktype() BookType
	Name() string
	Author() string
	Borrow(string) bool // Borrow accepts a username and attempts to borrow the book in that user's name.
	Return(string)      // Returns a boolean indicating the success of the borrow
}

type DigitalBook struct {
	bookType  BookType
	name      string
	author    string
	limit     int
	borrowers []string
}

func (d *DigitalBook) Booktype() BookType {
	return d.bookType
}

func (d *DigitalBook) Name() string {
	return d.name
}

func (d *DigitalBook) Author() string {
	return d.author
}

func (d *DigitalBook) Borrow(borrower string) bool {
	// If borrow slot is available, append borrower to list
	if len(d.borrowers) < d.limit {
		d.borrowers = append(d.borrowers, borrower)
		return true
	} else {
		// Else do not allow borrow
		return false
	}
}

//Removes the book
func RemoveBook(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func (d *DigitalBook) Return(borrower string) {
	index := 0
	for i := range d.borrowers {
		if d.borrowers[i] == borrower {
			index = i
			break
		}
	}

	RemoveBook(d.borrowers, index)
}

//constructor for DigitalBook
func NewDigitalBook(b_type BookType, name string, author string, limit int) *DigitalBook {
	return &DigitalBook{b_type, name, author, limit, make([]string, 0)}
}

type PhysicalBook struct {
	bookType BookType
	name     string
	author   string
	borrower string
}

func (b *PhysicalBook) Booktype() BookType {
	return b.bookType
}

func (b *PhysicalBook) Name() string {
	return b.name
}

func (b *PhysicalBook) Author() string {
	return b.author
}

func (p *PhysicalBook) Borrow(borrower string) bool {
	if p.borrower == "" { // Check if there is no current borrower
		p.borrower = borrower
		return true
	} else { // Else do not allow borrow
		return false
	}
}

func (p *PhysicalBook) Return(borrower string) {
	if p.borrower != "" {
		p.borrower = ""
	}
}

//constructor for PhysicalBook
func NewPhysicalBook(b_type BookType, name string, author string) *PhysicalBook {
	return &PhysicalBook{b_type, name, author, ""}
}

type Library struct {
	Books map[string]Book     //add books to the inventory
	Users map[string]struct{} //register new members to the userbase
}

func NewLibrary() *Library {
	return &Library{
		make(map[string]Book),
		make(map[string]struct{}),
	}
}

func (lib *Library) CheckUser(user string) bool {
	_, ok := lib.Users[user]
	return ok
}

func (lib *Library) AddUser(user string) {
	lib.Users[user] = struct{}{}
}

func (lib *Library) GetBook(bookName string) (Book, bool) {
	book, ok := lib.Books[bookName]
	return book, ok
}

func (lib *Library) AddBook(book Book) {
	lib.Books[book.Name()] = book
}

func (lib *Library) CheckBook(bookName string) Book {
	book := lib.Books[bookName]
	return book
}
