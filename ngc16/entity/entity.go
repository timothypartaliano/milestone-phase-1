package entity

type User struct {
	Id			int
	Username 	string
	Password 	string
	Age			int
}

type Game struct {
	GameID  	int
	Title 		string
	Price 		int
}

type Order struct {
	OrderID 	int
	UserID		int
	GameID		int
	Amount 		int
}