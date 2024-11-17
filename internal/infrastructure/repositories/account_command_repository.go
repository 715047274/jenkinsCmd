package repositories

type AccountCommandRepository interface {
	CreateAccount(userName string, balance int64, currency string) error
	UpdateAccount(userName string, balance int64, currency string) error
	RemoveAccount(userName string) error
}
