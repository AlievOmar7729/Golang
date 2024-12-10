package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Bank struct {
	Name      string
	bankMoney float64
	Deposit   float64
	Credit    float64
	Clients   map[string]*Client
	mu        sync.Mutex
}

func CreateBank(name string, bankMoney, deposit, credit float64) *Bank {
	return &Bank{
		Name:      name,
		bankMoney: bankMoney,
		Deposit:   deposit,
		Credit:    credit,
		Clients:   make(map[string]*Client),
	}
}

func (b *Bank) SetBankMoney(money float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.bankMoney = money
}

func (b *Bank) GetBankMoney() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.bankMoney
}

func (b *Bank) AddClient(c *Client) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Clients[c.Surname] = c
}

func (b *Bank) GetClientBySurname(surname string) *Client {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Clients[surname]
}

type Client struct {
	Name          string
	Surname       string
	AccountNumber string
	cDeposit      float64
	cCredit       float64
	bank          *Bank
	mu            sync.Mutex
}

func CreateClient(name, surname, accountNumber string, bank *Bank) *Client {
	return &Client{
		Name:          name,
		Surname:       surname,
		AccountNumber: accountNumber,
		cDeposit:      0,
		cCredit:       0,
		bank:          bank,
	}
}

func (c *Client) DepositMoney(amount float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cDeposit += amount
	c.bank.SetBankMoney(c.bank.GetBankMoney() + amount)
}

func (c *Client) WithdrawDeposit(amount float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cDeposit >= amount {
		c.cDeposit -= amount
		c.bank.SetBankMoney(c.bank.GetBankMoney() - amount)
	} else {
		fmt.Println("Недостатньо коштів на депозиті.")
	}
}

func (c *Client) TakeCredit(amount float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.bank.GetBankMoney() >= amount {
		c.cCredit += amount
		c.bank.SetBankMoney(c.bank.GetBankMoney() - amount)
	} else {
		fmt.Println("Недостатньо коштів у банку для надання кредиту.")
	}
}

func (c *Client) RepayCredit(amount float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cCredit >= amount {
		c.cCredit -= amount
		c.bank.SetBankMoney(c.bank.GetBankMoney() + amount)
	} else {
		fmt.Println("Сума погашення перевищує наданий кредит.")
	}
}

func (c *Client) GetInfo() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("Клієнт: %s %s\n", c.Name, c.Surname)
	fmt.Printf("Номер рахунку: %s\n", c.AccountNumber)
	fmt.Printf("Депозит: %.2f\n", c.cDeposit)
	fmt.Printf("Кредит: %.2f\n", c.cCredit)
}

func (c *Client) RandomOperations() {
	for {
		time.Sleep(1 * time.Second)
		if rand.Intn(2) == 0 {
			amount := float64(rand.Intn(100) + 1)
			c.DepositMoney(amount)
			fmt.Printf("%s поповнив депозит на %.2f\n", c.Name, amount)
		} else {
			amount := float64(rand.Intn(100) + 1)
			c.TakeCredit(amount)
			fmt.Printf("%s взяв кредит на %.2f\n", c.Name, amount)
		}
		if c.cDeposit > 0 && c.cCredit > 0 {
			break
		}
	}
	fmt.Printf("Операції для клієнта %s завершено.\n", c.Name)
}

func printMenu() {
	fmt.Println("\nМеню:")
	fmt.Println("1. Створити банк")
	fmt.Println("2. Створити клієнта для роботи з кредитами")
	fmt.Println("3. Створити клієнта для роботи з депозитами")
	fmt.Println("4. Вивести інформацію про клієнта за прізвищем")
	fmt.Println("5. Вивести інформацію про клієнта за номером рахунку")
	fmt.Println("6. Завершити")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var bank *Bank
	var wg sync.WaitGroup

	for {
		printMenu()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var bankName string
			var bankMoney, deposit, credit float64
			fmt.Println("Введіть назву банку:")
			fmt.Scan(&bankName)
			fmt.Println("Введіть початкову суму коштів у банку:")
			fmt.Scan(&bankMoney)
			fmt.Println("Введіть суму депозитів:")
			fmt.Scan(&deposit)
			fmt.Println("Введіть суму кредитів:")
			fmt.Scan(&credit)

			bank = CreateBank(bankName, bankMoney, deposit, credit)
			fmt.Printf("Банк '%s' створений!\n", bank.Name)

		case 2:
			var name, surname, accountNumber string
			fmt.Println("Введіть ім'я, прізвище та номер рахунку клієнта:")
			fmt.Scan(&name, &surname, &accountNumber)
			client := CreateClient(name, surname, accountNumber, bank)
			bank.AddClient(client)
			wg.Add(1)
			go func() {
				defer wg.Done()
				client.RandomOperations()
			}()
			fmt.Printf("Клієнт %s %s створений для роботи з кредитами\n", name, surname)

		case 3:
			var name, surname, accountNumber string
			fmt.Println("Введіть ім'я, прізвище та номер рахунку клієнта:")
			fmt.Scan(&name, &surname, &accountNumber)
			client := CreateClient(name, surname, accountNumber, bank)
			bank.AddClient(client)
			wg.Add(1)
			go func() {
				defer wg.Done()
				client.RandomOperations()
			}()
			fmt.Printf("Клієнт %s %s створений для роботи з депозитами\n", name, surname)

		case 4:
			var surname string
			fmt.Println("Введіть прізвище клієнта:")
			fmt.Scan(&surname)
			client := bank.GetClientBySurname(surname)
			if client != nil {
				client.GetInfo()
			} else {
				fmt.Println("Клієнта з таким прізвищем не знайдено.")
			}

		case 5:
			var accountNumber string
			fmt.Println("Введіть номер рахунку клієнта:")
			fmt.Scan(&accountNumber)
			var found bool
			for _, client := range bank.Clients {
				if client.AccountNumber == accountNumber {
					client.GetInfo()
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Клієнта з таким номером рахунку не знайдено.")
			}

		case 6:
			wg.Wait()
			fmt.Println("Програма завершена.")
			return
		}
	}
}
