package banking

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance int
}

// 직접 에러 핸들링을 할 때 어떤 에러에 대한 것인지 따로 모아서 관리하는 방법이 사용되는 듯 하다.
// lint 규칙으로 에러는 err로 시작하는 네이밍을 해야 한다.
var errBalanceError = errors.New("Error")

func NewAccount(owner string) *BankAccount {
	bankAccount := BankAccount{owner: owner, balance: 0}
	return &bankAccount
}

// BankAccount 타입의 receiver 구조체의 첫글자를 소문자로 하는 것이 규칙
// 리시버에 종속된 함수로 선언이 되는 것이므로 위 객체의 메소드가 된다.
// 리시버를 포인터로 선언하면 이미 생성한 구조체에 주소로 접근하고 객체의 멤버 변수의 값을 변경한다.
func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
}

// 리시버를 포인터로 선언하지 않으면 복사본에 대해 명령을 수행한다.
// 이 경우 값을 변경하면 복사본의 값이 변경된다.
func (b BankAccount) GetBalance() int {
	return b.balance
}

func (b *BankAccount) Withdraw(amount int) error {
	if b.balance < amount {
		return errBalanceError
	}
	b.balance -= amount
	return nil
}

func (b BankAccount) String() string {
	return fmt.Sprint(b.owner, b.balance)
}