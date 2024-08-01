## ERD Title: Sportix

## A. Entities

### 1. Entity: users
##### Attributes:
- user_id INT (PK, AI)
- username VARCHAR(50) NOT NULL
- email VARCHAR(50) NOT NULL
- password VARCHAR(100) NOT NULL
- role ENUM('owner', 'user') DEFAULT 'user'
- created_at DATETIME
- updated_at DATETIME

### 2. Entity: wallets
##### Attributes:
- wallet_id INT (PK, AI)
- user_id INT (FK)
- saldo DECIMAL(10, 2) NOT NULL
- created_at DATETIME
- updated_at DATETIME

### 3. Entity: deposits
##### Attributes:
- deposit_id INT (PK, AI)
- amount DECIMAL(10, 2) NOT NULL
- wallet_id INT (FK)
- created_at DATETIME
- updated_at DATETIME

### 4. Entity: fields
##### Attributes:
- field_id INT (PK, AI)
- name VARCHAR(50) NOT NULL
- price DECIMAL(10, 2) NOT NULL
- category_id INT (FK)
- location_id INT (FK)
- address TEXT NOT NULL
- facility_id INT (FK)
- available_hour_id INT (FK)
- created_by INT (FK)
- created_at DATETIME
- updated_at DATETIME

### 5. Entity: locations
##### Attributes:
- location_id INT (PK, AI)
- name VARCHAR(50) NOT NULL

### 6. Entity: facilities
##### Attributes:
- facility_id INT (PK, AI)
- bathroom INT NOT NULL
- cafeteria BOOL NOT NULL
- vehicle_park INT NOT NULL
- prayer_room BOOL NOT NULL
- changing_room INT NOT NULL
- cctv BOOL NOT NULL

### 7. Entity: available_hours
##### Attributes:
- available_hour_id INT (PK, AI)
- start_time DATETIME NOT NULL
- end_time DATETIME NOT NULL
- status ENUM('available', 'booked') DEFAULT 'available' NOT NULL

### 8. Entity: categories
##### Attributes:
- category_id INT (PK, AI)
- name VARCHAR(50) NOT NULL

### 9. Entity: reservations
##### Attributes:
- reservation_id INT (PK, AI)
- user_id INT (FK)
- field_id INT (FK)
- status ENUM('confirmed', 'canceled', 'completed') NOT NULL
- reservation_date DATE NOT NULL
- created_at DATETIME DEFAULT NOW()
- updated_at DATETIME DEFAULT NOW() ON UPDATE NOW()

### 10. Entity: payments
##### Attributes:
- payment_id INT (PK, AI)
- reservation_id INT (FK)
- amount DECIMAL(10, 2) NOT NULL
- payment_date DATETIME NOT NULL DEFAULT NOW()
- status ENUM('completed', 'pending') NOT NULL
- created_at DATETIME
- updated_at DATETIME

## B. Relationships:

### 1. users to wallets
- **Type:** One to One
- **Description: Each user has one wallet, and each wallet belongs to one user.**

### 2. wallets to deposits
- **Type:** One to Many
- **Description: One wallet can have multiple deposits, but each deposit is linked to only one wallet.**

### 3. users to fields
- **Type:** One to Many
- **Description: One user (owner) can create multiple fields, but each field is created by only one user.**

### 4. fields to categories
- **Type:** Many to One
- **Description: Multiple fields can belong to the same category, but each field belongs to only one category.**

### 5. fields to locations
- **Type:** Many to One
- **Description: Multiple fields can be in the same location, but each field is in only one location.**

### 6. fields to facilities
- **Type:** One to One
- **Description: Each field has one set of facilities, and each set of facilities belongs to one field.**

### 7. fields to available_hours
- **Type:** One to Many
- **Description: One field can have multiple available hours, but each available hour is linked to only one field.**

### 8. users to reservations
- **Type:** One to Many
- **Description: One user can make multiple reservations, but each reservation is made by only one user.**

### 9. fields to reservations
- **Type:** One to Many
- **Description: One field can have multiple reservations, but each reservation is for only one field.**

### 10. reservations to payments
- **Type:** One to One
- **Description: Each reservation has one payment, and each payment is for one reservation.**

## C. Integrity Constraints:
- The email in users should be unique.
- The balance in wallets and amount in deposits and payments should be positive numbers.
- The status in reservations and payments should only allow specified enum values.
- Foreign keys must reference existing entries in their respective tables.

## D. Additional Notes:
- The system supports user roles (owner and user).
- The system includes a digital wallet feature for users.
- Fields have various attributes including category, location, and facilities.
- The reservation system includes status tracking and payment processing.