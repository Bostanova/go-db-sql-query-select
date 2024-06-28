package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT product, volume, date FROM sales WHERE id = :id", sql.Named("id", client))

	sales = append(sales, Sale{})
	err = row.Scan(&sales[0].Product, &sales[0].Volume, &sales[0].Date)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
