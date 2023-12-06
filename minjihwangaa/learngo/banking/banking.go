package banking

// 대문자로 적으면 export 가능함
// Account struct
type Account struct {
	// 소문자이면 priviate
	// owner   string
	Owner   string
	Balance int
}