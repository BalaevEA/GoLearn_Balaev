package main

import (
	"fmt"
	"strings"
)

type Order struct {
	ID       int
	Customer *OrderCustomer
	Good     []*OrderGood
}

type OrderCustomer struct {
	Name     string
	SurName  string
	PhoneNum string
}
type OrderGood struct {
	Name  string
	Price float64
	Count int
}

func main() {

	Customer1 := OrderCustomer{
		Name:     "Эдуард",
		SurName:  "Балаев",
		PhoneNum: "+79143250311",
	}
	order := Order{
		ID:       1,
		Customer: &Customer1,
		Good:     nil,
	}
	goodslice := []*OrderGood{
		{"Protein", 3999.99, 2},
		{"Casein", 4999.99, 2},
		{"KFC", 200, 4},
	}
	order.Good = goodslice
	fmt.Println("Номер заказа: ", order.ID, "\nПокупатель: ", order.Customer, "\nТовары:\n", order.Good)

	Customer2 := OrderCustomer{
		Name:     "Карифанов",
		SurName:  "Поридж",
		PhoneNum: "+9999999999",
	}
	order2 := Order{
		ID:       2,
		Customer: &Customer2,
		Good:     nil,
	}
	goodslice2 := []*OrderGood{
		{"ГрандДеЛюкс", 299.99, 2},
		{"Вейп", 4999.99, 1},
		{"Самокат", 20000, 1},
	}
	order2.Good = goodslice2
	fmt.Println("Номер заказа: ", order2.ID, "\nПокупатель: ", order2.Customer, "\nТовары:\n", order2.Good)

}

func (og OrderGood) String() string {
	return strings.Join([]string{fmt.Sprintf("%d Шт.", og.Count), og.Name, "Цена", fmt.Sprintf("%.2f руб\n", og.Price)}, " ")
}

func (cs OrderCustomer) String() string {
	return fmt.Sprint(cs.Name, " ", cs.SurName, " ", cs.PhoneNum)
}

// func (og OrderGood) String() string {
// 	return fmt.Sprintln(og.Count, "Шт.", og.Name, "Цена", og.Price, "руб")
// }
