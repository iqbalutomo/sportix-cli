package repository

import (
	"database/sql"                // Paket untuk interaksi dengan database SQL
	"fmt"                         // Paket untuk format string dan kesalahan
	"sportix-cli/internal/entity" // Paket yang berisi definisi entitas, seperti User

	"golang.org/x/crypto/bcrypt" // Paket untuk hashing dan verifikasi password
)

// UserRepo adalah interface yang mendefinisikan operasi yang bisa dilakukan pada entitas User
type UserRepo interface {
	CreateUser(user *entity.User) error                        // Method untuk membuat pengguna baru
	FindUserByEmail(email string) (*entity.User, error)        // Method untuk menemukan pengguna berdasarkan email
	ValidateUser(email, password string) (*entity.User, error) // Method untuk memvalidasi pengguna berdasarkan email dan password
	FindBalanceByEmail(email string) (float64, error)          // Method untuk menemukan saldo pengguna berdasarkan email
	UpdateBalance(userID uint, deposit float64) error          // Method untuk memperbarui saldo pengguna
}

// userRepo adalah struct yang menyimpan referensi ke database SQL
type userRepo struct {
	db *sql.DB // Referensi ke database SQL
}

// NewUserRepo adalah konstruktor untuk membuat instance userRepo baru
func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db} // Mengembalikan instance userRepo dengan referensi ke database
}

// CreateUser membuat pengguna baru dan juga membuat dompet untuk pengguna tersebut
func (u *userRepo) CreateUser(user *entity.User) error {
	// Memulai transaksi database
	tx, err := u.db.Begin()
	if err != nil {
		return err // Mengembalikan kesalahan jika tidak bisa memulai transaksi
	}

	// Menyisipkan pengguna baru ke tabel users
	query := `INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?);`
	result, err := tx.Exec(query, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		tx.Rollback() // Membatalkan transaksi jika ada kesalahan
		return fmt.Errorf("error creating user table: %v", err)
	}

	// Mengambil ID pengguna yang baru dimasukkan
	userID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback() // Membatalkan transaksi jika ada kesalahan
		return fmt.Errorf("error getting last inserted user ID: %v", err)
	}

	user.UserID = uint(userID) // Mengatur ID pengguna yang baru dimasukkan

	// Membuat dompet baru untuk pengguna
	query = `INSERT INTO wallets (user_id) VALUES (?);`
	_, err = tx.Exec(query, user.UserID)
	if err != nil {
		tx.Rollback() // Membatalkan transaksi jika ada kesalahan
		return fmt.Errorf("error creating wallet table: %v", err)
	}

	// Menyelesaikan transaksi
	err = tx.Commit()
	if err != nil {
		tx.Rollback() // Membatalkan transaksi jika ada kesalahan
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil // Mengembalikan nil jika transaksi berhasil
}

// FindUserByEmail menemukan pengguna berdasarkan email
func (u *userRepo) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User // Membuat variabel untuk menyimpan data pengguna

	// Query untuk mengambil data pengguna berdasarkan email
	query := `SELECT user_id, username, password, email, role FROM users WHERE email = ?;`
	rows := u.db.QueryRow(query, email)

	// Memindahkan data dari query ke variabel user
	if err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Mengembalikan nil jika tidak ditemukan baris
		}

		return nil, err // Mengembalikan kesalahan jika ada masalah
	}

	return &user, nil // Mengembalikan pengguna yang ditemukan
}

// ValidateUser memvalidasi pengguna berdasarkan email dan password
func (u *userRepo) ValidateUser(email, password string) (*entity.User, error) {
	user, err := u.FindUserByEmail(email) // Mencari pengguna berdasarkan email
	if err != nil {
		return nil, err // Mengembalikan kesalahan jika terjadi masalah dalam pencarian pengguna
	}
	if user == nil {
		return nil, fmt.Errorf("user not found") // Mengembalikan kesalahan jika pengguna tidak ditemukan
	}

	// Memverifikasi password yang dimasukkan dengan password yang tersimpan
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password") // Mengembalikan kesalahan jika password tidak valid
	}

	return user, nil // Mengembalikan pengguna yang sudah terverifikasi
}

// FindBalanceByEmail menemukan saldo pengguna berdasarkan email
func (u *userRepo) FindBalanceByEmail(email string) (float64, error) {
	var wallet float64 // Variabel untuk menyimpan saldo

	// Query untuk mengambil saldo pengguna berdasarkan email
	query := `SELECT w.balance
			FROM users u
			JOIN wallets w ON u.user_id = w.user_id
			WHERE email = ?;`
	rows := u.db.QueryRow(query, email)

	// Memindahkan data dari query ke variabel wallet
	if err := rows.Scan(&wallet); err != nil {
		if err == sql.ErrNoRows {
			return 0, err // Mengembalikan 0 jika tidak ditemukan baris
		}

		return 0, err // Mengembalikan kesalahan jika ada masalah
	}

	return wallet, nil // Mengembalikan saldo yang ditemukan
}

// UpdateBalance memperbarui saldo pengguna berdasarkan ID pengguna
func (u *userRepo) UpdateBalance(userID uint, totalBalance float64) error {
	query := `UPDATE wallets SET balance = ? WHERE user_id = ?;`

	// Menjalankan query untuk memperbarui saldo
	_, err := u.db.Exec(query, totalBalance, userID)
	if err != nil {
		return err // Mengembalikan kesalahan jika ada masalah
	}

	return nil // Mengembalikan nil jika saldo berhasil diperbarui
}
