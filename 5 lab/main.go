package main

import (
	"fmt"
	"os"
)

type Product struct {
	name     string
	price    float64
	currency Currency
	quantity int64
	producer string
	weight   float64
}

func CreateProduct(name string, price float64, currency Currency, quantity int64, producer string, weight float64) Product {
	return Product{name, price, currency, quantity, producer, weight}
}

func (p *Product) SetName(name string) {
	p.name = name
}

func (p Product) GetName() string {
	return p.name
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p *Product) SetCurrency(currency Currency) {
	p.currency = currency
}

func (p Product) GetCurrency() Currency {
	return p.currency
}

func (p *Product) SetQuantity(quantity int64) {
	p.quantity = quantity
}

func (p Product) GetQuantity() int64 {
	return p.quantity
}

func (p *Product) SetProducer(producer string) {
	p.producer = producer
}

func (p Product) GetProducer() string {
	return p.producer
}

func (p *Product) SetWeight(weight float64) {
	p.weight = weight
}

func (p Product) GetWeight() float64 {
	return p.weight
}

func (p Product) GetPriceInUAH() float64 {
	return p.currency.ExchangeRate * p.price
}

func (p Product) GetTotalPriceInUAH() float64 {
	return p.GetPriceInUAH() * float64(p.GetQuantity())
}

func (p Product) GetTotalWeight() float64 {
	return p.GetWeight() * float64(p.GetQuantity())
}

type Currency struct {
	Name         string
	ExchangeRate float64
}

func CreateCurrency(name string, exchangeRate float64) Currency {
	return Currency{name, exchangeRate}
}

func (c Currency) String() string {
	return fmt.Sprintf("%s: %.2f грн", c.Name, c.ExchangeRate)
}

func readProducts() []Product {
	var products []Product
	var n int
	for {
		fmt.Println("Введіть кількість продуктів: ")
		_, err := fmt.Scanf("%d", &n)
		if err != nil || n <= 0 {
			fmt.Println("Будь ласка, введіть коректне позитивне число!")
			_, _ = fmt.Fscanf(os.Stdin, "%*s") 
		} else {
			break
		}
	}

	var name, currencyName, producer string
	var price, currencyRate, weight float64
	var quantity int64

	for i := 0; i < n; i++ {
		fmt.Println("Введіть через пробіл назву, ціну, валюту, курс валюти, кількість, виробника та вагу: ")
		_, err := fmt.Scan(&name, &price, &currencyName, &currencyRate, &quantity, &producer, &weight)

		if err != nil {
			fmt.Println("Некоректні дані!")
			os.Exit(1)
		}

		currency := CreateCurrency(currencyName, currencyRate)
		product := CreateProduct(name, price, currency, quantity, producer, weight)
		products = append(products, product)
	}

	return products
}

func printProduct(p Product) {
	fmt.Println("Назва:", p.GetName())
	fmt.Printf("Ціна: %.2f\n", p.GetPrice())
	fmt.Println("Валюта:", p.GetCurrency().Name)
	fmt.Printf("Курс валюти: %.2f\n", p.GetCurrency().ExchangeRate)
	fmt.Printf("Кількість: %d\n", p.GetQuantity())
	fmt.Println("Виробник:", p.GetProducer())
	fmt.Printf("Вага: %.2f\n", p.GetWeight())
}

func printAllProducts(products []Product) {
	for i, p := range products {
		fmt.Printf("\nПродукт №%d:\n", i+1)
		printProduct(p)
	}
}
func getMinMaxPriceProducts(products []Product) (Product, Product) {
	min, max := products[0], products[0]
	for _, p := range products {
		if p.GetPriceInUAH() < min.GetPriceInUAH() {
			min = p
		}
		if p.GetPriceInUAH() > max.GetPriceInUAH() {
			max = p
		}
	}
	return min, max
}

func main() {
	products := readProducts()
	printAllProducts(products)

	minProduct, maxProduct := getMinMaxPriceProducts(products)

	fmt.Println("\nПродукт з мінімальною ціною:")
	printProduct(minProduct)
	fmt.Println("\nПродукт з максимальною ціною:")
	printProduct(maxProduct)
}
