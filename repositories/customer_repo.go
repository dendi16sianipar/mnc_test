package repositories

import (
	"database/sql"
	"mnc_test/models"
	"mnc_test/utils"

	"golang.org/x/crypto/bcrypt"
)

type customerRepo struct {
	db *sql.DB
}

type CustomerRepo interface {
	Create(newCustomer *models.Customer) (*models.Customer, error)
	Login(email string, password string) (string, error)
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepo{
		db: db,
	}
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *customerRepo) Create(data *models.Customer) (*models.Customer, error) {
	query := "INSERT INTO customers (name, email, password, balance) VALUES ($1, $2, $3, $4) RETURNING id"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return &models.Customer{}, err
	}

	var id int
	err = r.db.QueryRow(query, data.Name, data.Email, string(hashedPassword), data.Balance).Scan(&id)
	if err != nil {
		return &models.Customer{}, err
	}

	data.ID = id
	return data, nil
}

func (r *customerRepo) Login(email string, password string) (string, error) {
	u := models.Customer{}

	query := "SELECT id, email, password FROM customers WHERE email = $1"
	row := r.db.QueryRow(query, email)

	err := row.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.CreateTokenString(email)
	if err != nil {
		return "", err
	}

	return token, nil
}
