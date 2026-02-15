package api

// coin balance params
type CoinBalanceParams struct {
	UserName string
}

// coin balance response
type CoinBalanceResponse struct {
	//success code usually 200
	Code int

	//account balance
	balance int64
}

type Error struct {
	//error code
	Code int

	//error message
	Message string
}
