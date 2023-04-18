package repositories

import (
	"database/sql"
	"errors"
	"mnc_test/models"
	"time"
)

type paymentRepo struct {
	db *sql.DB
}

type PaymentRepo interface {
	Create(data *models.Payment) (*models.Payment, error)
}

func NewPaymentRepo(db *sql.DB) PaymentRepo {
	return &paymentRepo{
		db: db,
	}
}

func (r *paymentRepo) Create(data *models.Payment) (*models.Payment, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return &models.Payment{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		} else {
			err = tx.Commit()
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}()

	var balance float32
	err = tx.QueryRow("SELECT balance FROM customers WHERE id = $1", data.Customer_Id).Scan(&balance)
	if err != nil {
		return &models.Payment{}, err
	}

	if balance < 1 {
		return &models.Payment{}, errors.New("user has no balance to complete payment")
	}

	var id int
	if balance >= data.Bill {
		err := tx.QueryRow("UPDATE customers SET balance = $1 WHERE id = $2 RETURNING id", (balance - data.Bill), data.Customer_Id).Scan(&id)
		if err != nil {
			return &models.Payment{}, err
		}

		var date time.Time
		query := "INSERT INTO payments (customer_id, bill, description) VALUES ($1, $2, $3) RETURNING id, date"
		err = tx.QueryRow(query, data.Customer_Id, data.Bill, data.Description).Scan(&id, &date)
		if err != nil {
			return &models.Payment{}, err
		}

		data.ID = id
		data.Date = date

		return data, nil
	}
	return &models.Payment{}, errors.New("not enought balance to complete payment")
}
