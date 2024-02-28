package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//Race condition masalah yang dihadapi karena banyak goroutine berjalan secara concurrent.

func TestRaceConditionIssue(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x)
	/*! hasilnya akan berbeda2 karena manipulasi goroutines yang berjalan berbarengan*/
}

// Mutex
// digunakan untuk melakukan locking dan unlocking
// jika ada beberapa goroutine melakukan lock, maka hanya 1 goroutine yang diperbolehkan,
//
//	setelah unlock goroutine selanjutnya diperbolehkan lock lagi.
func TestRaceConditionIssue_Mutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x)
	//! hasilnya akan sesuai 100000
}

// RWMutex (Read Write Mutex)
// perubahan lock tidak hanya pada mengubah data, tapi membaca data juga
// jika hanya Mutex saja maka akan rebutan antara proses membaca dan mengubah data
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (receiver *BankAccount) AddBalance(amount int) {
	receiver.RWMutex.Lock()
	receiver.Balance = receiver.Balance + amount
	receiver.RWMutex.Unlock()
}
func (receiver *BankAccount) GetBalance() int {
	receiver.RWMutex.Lock()
	balance := receiver.Balance
	receiver.RWMutex.Unlock()
	return balance
}
func TestRaceConditionIssue_RWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance : ", account.GetBalance())
}

// DEADLOCK
// situasi dimana goroutine saling menunggu lock, sehingga tidak ada goroutine yang bisa jalan
// hati hati dalam penggunaan Mutex / RWMutex
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (receiver *UserBalance) Lock() {
	receiver.Mutex.Lock()
}
func (receiver *UserBalance) Unlock() {
	receiver.Mutex.Unlock()
}
func (receiver *UserBalance) Change(amount int) {
	receiver.Balance = receiver.Balance + amount
}
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(-amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Iqbal",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "Fauzan",
		Balance: 100000,
	}
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)
	time.Sleep(3 * time.Second)
	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}
