## ERD Title: Sportix Database Schema

### 1. Entities and Their Attributes

**Entity: users**
- Attributes:
    - id INT AUTO_INCREMENT PRIMARY KEY
    - username VARCHAR(50) NOT NULL
    - email VARCHAR(50) NOT NULL
    - password VARCHAR(100) NOT NULL
    - role ENUM('admin', 'user') default 'user'
    - created_at DATETIME
    - updated_at DATETIME

**Entity: deposits**
- Attributes:
    - id INT AUTO_INCREMENT PRIMARY KEY
    - saldo DECIMAL(10, 2) NOT NULL

**Entity: user_deposits**
- Attributes:
    - user_id INT
    - deposit_id INT
    - created_at DATETIME
    - FOREIGN KEY (user_id) REFERENCES users(id)
    - FOREIGN KEY (deposit_id) REFERENCES deposits(id)

**Entity: fields**
- Attributes:
    - id INT AUTO_INCREMENT PRIMARY KEY
    - name VARCHAR(50) NOT NULL
    - location VARCHAR(50) NOT NULL
    - price DECIMAL(10,2) NOT NULL
    - facilities TEXT
    - created_at DATETIME
    - updated_at DATETIME

**Entity: reservations**
- Attributes:
    - id INT AUTO_INCREMENT PRIMARY KEY
    - user_id INT
    - field_id INT
    - start_time DATETIME
    - end_time DATETIME
    - status ENUM('confirmed', 'cancelled')
    - created_at DATETIME
    - updated_at DATETIME
    - FOREIGN KEY (user_id) REFERENCES users(id)
    - FOREIGN KEY (field_id) REFERENCES fields(id)

**Entity: payments**
- Attributes:
    - id INT AUTO_INCREMENT PRIMARY KEY
    - reservation_id INT
    - amount DECIMAL(10,2) NOT NULL
    - payment_date DATETIME
    - status ENUM('completed', 'pending')
    - created_at DATETIME
    - updated_at DATETIME
    - FOREIGN KEY (reservation_id) REFERENCES reservations(id)

### 2. Relationships

**users to user_deposits:**
- **Type:** One to Many
- **Description:** Each user can have multiple deposits, but each deposit is assigned to only one user at a time.

**deposits to user_deposits:**
- **Type:** One to Many
- **Description:** Each deposit can be assigned to multiple users, but each user-deposit relationship is unique.

**fields to reservations:**
- **Type:** One to Many
- **Description:** Each field can have multiple reservations, but each reservation is for only one field.

**users to reservations:**
- **Type:** One to Many
- **Description:** Each user can make multiple reservations, but each reservation is made by one user.

**reservations to payments:**
- **Type:** One to One
- **Description:** Each reservation has one payment, and each payment is linked to one reservation.

### 3. Normalization

**1NF**:
- All columns contain atomic values.
- The `user_deposits` table separates the deposits assigned to users.

**2NF**:
- The `reservations` table uses foreign keys to reference `users` and `fields`, eliminating partial dependency between reservations and fields or users.

**3NF**:
- Each table stores data relevant to its entity, with non-primary attributes only depending on the primary key.
- Transitive dependencies are eliminated, with separate tables for user deposits and payments.
