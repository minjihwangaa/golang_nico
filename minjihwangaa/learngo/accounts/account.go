package account

// private
// Account struct
type Account struct {
	owner string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// 복사본을 넘기지 않고 메모리 주소값을 return
	return &account
}