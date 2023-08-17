package entity

type User struct {
	Id			int
	Username 	string
	Email 		string
	Password 	string
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