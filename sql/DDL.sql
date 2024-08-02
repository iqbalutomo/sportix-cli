-- Clear the database if it exists
DROP DATABASE IF EXISTS sportix_development;
-- Create the database if it does not exist
CREATE DATABASE IF NOT EXISTS sportix_development;
-- Use the database
USE sportix_development;


-- Create the users table
CREATE TABLE users
(
    user_id    INT PRIMARY KEY AUTO_INCREMENT,
    username   VARCHAR(50)  NOT NULL,
    email      VARCHAR(50)  NOT NULL UNIQUE,
    password   VARCHAR(100) NOT NULL,
    role       ENUM ('owner', 'user') DEFAULT 'user',
    created_at TIMESTAMP              DEFAULT NOW(),
    updated_at TIMESTAMP              DEFAULT NOW() ON UPDATE NOW()
);

-- Create the wallets table
CREATE TABLE wallets
(
    wallet_id  INT PRIMARY KEY AUTO_INCREMENT,
    user_id    INT            NOT NULL,
    balance    DECIMAL(10, 2) NOT NULL DEFAULT 1000000,
    created_at DATETIME                DEFAULT NOW(),
    updated_at DATETIME                DEFAULT NOW() ON UPDATE NOW()
);

-- Create the deposits table
CREATE TABLE deposits
(
    deposit_id INT PRIMARY KEY AUTO_INCREMENT,
    amount     DECIMAL(10, 2) NOT NULL,
    wallet_id  INT            NOT NULL,
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW() ON UPDATE NOW()
);

-- Create the fields table
CREATE TABLE fields
(
    field_id    INT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(50)    NOT NULL,
    price       DECIMAL(10, 2) NOT NULL,
    category_id INT            NOT NULL,
    location_id INT            NOT NULL,
    address     TEXT           NOT NULL,
    facility_id INT            NOT NULL,
    created_by  INT,
    created_at  DATETIME DEFAULT NOW(),
    updated_at  DATETIME DEFAULT NOW() ON UPDATE NOW()
);

-- Create the location table
CREATE TABLE locations
(
    location_id INT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(50) NOT NULL
);

-- Create the facilities table
CREATE TABLE facilities
(
    facility_id   INT PRIMARY KEY AUTO_INCREMENT,
    bathroom      INT  NOT NULL,
    cafeteria     BOOL NOT NULL,
    vehicle_park  INT  NOT NULL,
    prayer_room   BOOL NOT NULL,
    changing_room INT  NOT NULL,
    cctv          BOOL NOT NULL
);

-- Create the available_hours table
CREATE TABLE available_hours
(
    available_hour_id INT PRIMARY KEY AUTO_INCREMENT,
    start_time        TIME NOT NULL,
    end_time          TIME NOT NULL
);

-- Create the field_available_hours table
CREATE TABLE field_available_hours
(
    field_available_hour_id INT PRIMARY KEY AUTO_INCREMENT,
    field_id                INT                                              NOT NULL,
    available_hour_id       INT                                              NOT NULL,
    status                  ENUM ('available', 'booked') DEFAULT 'available' NOT NULL,
    created_at              DATETIME                     DEFAULT NOW(),
    updated_at              DATETIME                     DEFAULT NOW() ON UPDATE NOW()
);

-- Create the categories table
CREATE TABLE categories
(
    category_id INT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(50) NOT NULL
);

-- Create the reservations table
CREATE TABLE reservations
(
    reservation_id   INT PRIMARY KEY AUTO_INCREMENT,
    user_id          INT                                         NOT NULL,
    field_id         INT                                         NOT NULL,
    status           ENUM ('confirmed', 'canceled', 'completed') NOT NULL,
    reservation_date DATE                                        NOT NULL,
    created_at       DATETIME DEFAULT NOW(),
    updated_at       DATETIME DEFAULT NOW() ON UPDATE NOW()
);

-- Create the payments table
CREATE TABLE payments
(
    payment_id     INT PRIMARY KEY AUTO_INCREMENT,
    reservation_id INT                           NOT NULL,
    amount         DECIMAL(10, 2)                NOT NULL,
    payment_date   DATETIME DEFAULT NOW()        NOT NULL,
    status         ENUM ('completed', 'pending') NOT NULL,
    created_at     DATETIME DEFAULT NOW(),
    updated_at     DATETIME DEFAULT NOW() ON UPDATE NOW()
);

-- Foreign Key Constraints

-- Foreign key constraint for wallets table
ALTER TABLE wallets
    ADD CONSTRAINT fk_wallets_user_id FOREIGN KEY (user_id) REFERENCES users (user_id);

-- Foreign key constraint for deposits table
ALTER TABLE deposits
    ADD CONSTRAINT fk_deposits_wallet_id FOREIGN KEY (wallet_id) REFERENCES wallets (wallet_id);

-- Foreign key constraint for fields table
ALTER TABLE fields
    ADD CONSTRAINT fk_fields_category_id FOREIGN KEY (category_id) REFERENCES categories (category_id),
    ADD CONSTRAINT fk_fields_location_id FOREIGN KEY (location_id) REFERENCES locations (location_id),
    ADD CONSTRAINT fk_fields_facility_id FOREIGN KEY (facility_id) REFERENCES facilities (facility_id),
    ADD CONSTRAINT fk_fields_created_by FOREIGN KEY (created_by) REFERENCES users (user_id);

-- Foreign key constraint for field_available_hours table
ALTER TABLE field_available_hours
    ADD CONSTRAINT fk_field_available_hours_field_id FOREIGN KEY (field_id) REFERENCES fields (field_id),
    ADD CONSTRAINT fk_field_available_hours_available_hour_id FOREIGN KEY (available_hour_id) REFERENCES available_hours (available_hour_id);

-- Foreign key constraint for reservations table
ALTER TABLE reservations
    ADD CONSTRAINT fk_reservations_user_id FOREIGN KEY (user_id) REFERENCES users (user_id),
    ADD CONSTRAINT fk_reservations_field_id FOREIGN KEY (field_id) REFERENCES fields (field_id);

-- Foreign key constraint for payments table
ALTER TABLE payments
    ADD CONSTRAINT fk_payments_reservation_id FOREIGN KEY (reservation_id) REFERENCES reservations (reservation_id);

-- Check Constraints

-- Check constraint for wallets table
ALTER TABLE wallets
    ADD CONSTRAINT chk_wallet_balance CHECK (balance >= 0);

-- Check constraint for deposits table
ALTER TABLE deposits
    ADD CONSTRAINT chk_deposit_amount CHECK (amount >= 0);

-- Check constraint for fields table
ALTER TABLE fields
    ADD CONSTRAINT chk_field_price CHECK (price >= 0);

-- Check constraint for payments table
ALTER TABLE payments
    ADD CONSTRAINT chk_payment_amount CHECK (amount >= 0);

