package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Client struct {
	ID       int
	FIO      string
	Login    string
	Birthday string
	Email    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Client.
// Теперь, если передать объект Client в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (c Client) String() string {
	return fmt.Sprintf("ID: %d FIO: %s Login: %s Birthday: %s Email: %s",
		c.ID, c.FIO, c.Login, c.Birthday, c.Email)
}

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	newClient := Client{
		FIO:      "Денисов Денис Денисович",
		Login:    "den777",
		Birthday: "19860107",
		Email:    "denisovden1986@mail.ru",
	}

	id, err := insertClient(db, newClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	newLogin := "denisovdd007"
	err = updateClientLogin(db, newLogin, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err = selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	// удаление клиента
	err = deleteClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	_, err = selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func insertClient(db *sql.DB, client Client) (int64, error) {
	res, err := db.Exec("INSERT INTO clients (FIO, Login, Birthday, Email) VALUES (:FIO, :Login, :Birthday, :Email)",
		sql.Named("FIO", client.FIO),
		sql.Named("Login", client.Login),
		sql.Named("Birthday", client.Birthday),
		sql.Named("Email", client.Email))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func updateClientLogin(db *sql.DB, login string, id int64) error {
	_, err := db.Exec("UPDATE clients SET login = :login WHERE id = :id",
		sql.Named("login", login),
		sql.Named("id", id))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func deleteClient(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM clients WHERE id = :id", sql.Named("id", id))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func selectClient(db *sql.DB, id int64) (Client, error) {
	client := Client{}

	row := db.QueryRow("SELECT id, fio, login, birthday, email FROM clients WHERE id = :id", sql.Named("id", id))
	err := row.Scan(&client.ID, &client.FIO, &client.Login, &client.Birthday, &client.Email)

	return client, err
}
