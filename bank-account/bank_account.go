package account

import "sync"

type Account struct {
	balance int64
	isOpen  bool
	mu      sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		balance: amount,
		isOpen:  true,
	}
}

func (a *Account) Balance() (int64, bool) {
	return a.balance, a.isOpen
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	newBalance := int64(0)
	status := false
	if a.isOpen {
		newBalance = a.balance + amount
		if newBalance >= 0 {
			a.balance = newBalance
			status = true
		}
	}
	defer a.mu.Unlock()
	return newBalance, status
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	closingBalance := int64(0)
	status := false
	if a.isOpen {
		status = true
		closingBalance = a.balance
		a.balance = 0
		a.isOpen = false
	}
	defer a.mu.Unlock()
	return closingBalance, status
}
