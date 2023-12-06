package account

// private
// Account struct
type Account struct {
	owner string
	balance int
}

// 2.0 NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// 복사본을 넘기지 않고 메모리 주소값을 return
	return &account
}


// Function - func functionName (argument) returnType {}
// Method - func (reciever) functionName (argument) returnType {}
// reciever => Account에 연동되는 함수, 보통 축약어는 struct 이름 앞 소문자

// 2.1 Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) WithDraw(amount int) {
	a.balance -= amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}


