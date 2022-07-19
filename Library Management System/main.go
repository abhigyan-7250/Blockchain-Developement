package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input int
	option := true
	lib := NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Welcome To The Library\n\n")

	for option {
		fmt.Printf("Mention Choice: \n\n")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Borrow")
		fmt.Println("4. Return")
		fmt.Println("5. Exit")
		fmt.Printf("\n")
		fmt.Printf("Enter your choice: ")
		fmt.Scanln(&input)

		switch input {
		//case for adding books
		case 1:
			var (
				choice int
				btype  BookType
			)

			fmt.Print("Enter Name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Println()
			fmt.Println("Enter BookType: \n1. E-book\n2. AudioBook\n3. HardBack\n4. PaperBack\n5. Encyclopedia\n6. Magazine\n7. Comic: ")
			fmt.Scanln(&btype)
			fmt.Println()
			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()
			fmt.Println()
			fmt.Println("Enter choice of the book: \n1. Physical\n2. Digital")
			fmt.Scanln(&choice)
			fmt.Println()

			if choice == 1 {
				lib.AddBook(NewPhysicalBook(btype, name, author))
			} else {
				var limit int //digitalbooks have limit
				fmt.Print("Enter the borrowing limit: ")
				fmt.Scanln(&limit)

				lib.AddBook(NewDigitalBook(btype, name, author, limit))
			}
			fmt.Println(lib)
			fmt.Println()

		//case for adding members
		case 2:
			var name string
			fmt.Print("Enter Name:")
			fmt.Scanln(&name)
			lib.AddUser(name)
			fmt.Println(lib)
			fmt.Println()

		//case for borrowing a book
		case 3:
			var (
				username string
			)

			fmt.Print("Enter your name: ")
			fmt.Scanln(&username)

			ok := lib.CheckUser(username)
			if !ok {
				fmt.Println("User Does Not Exist")
				return
			} else {
				fmt.Println("Registered User")
				var bookname string
				fmt.Println("Enter Name Of Borrowing Book: ")
				fmt.Scanln(&bookname)

				book, ok := lib.GetBook(bookname)
				if !ok {
					fmt.Println("Book Does Not Exist")
					return
				} else {
					if success := book.Borrow(username); success {
						fmt.Println("Book Borrowed Successfully!")
					} else {
						fmt.Println("Borrow Failed!")
					}
				}
			}

		//case for returning the borrowed book
		case 4:
			var (
				username string
			)

			fmt.Print("Enter your name: ")
			fmt.Scanln(&username)

			ok := lib.CheckUser(username)
			if !ok {
				fmt.Println("User Does Not Exist")
				return
			} else {
				fmt.Println("Registered User")
				var bookname string
				fmt.Print("Enter Name of Returning Book ")
				fmt.Scanln(&bookname)

				book := lib.CheckBook(bookname)
				book.Return(username)

				fmt.Println("Book Returned Successfully")
			}

		//any case apart from above exists
		default:
			option = false
		}
	}
}
