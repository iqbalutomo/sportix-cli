INSERT INTO users (username, email, password, role, created_at, updated_at) VALUES
    ('johndoe', 'owner1@example.com', 'password123', 'user', NOW(), NOW()),
    ('ameliawatson', 'amelia@example.com', 'password123', 'owner', NOW(), NOW()),
    ('ceresfauna', 'fauna@example.com', 'password123', 'owner', NOW(), NOW()),
    ('bigfish', 'bigfish@example.com', 'password123', 'owner', NOW(), NOW()),
    ('regisaltare', 'altare@example.com', 'password123', 'owner', NOW(), NOW());

INSERT INTO categories (name) VALUES
    ('Soccer'),
    ('Badminton'),
    ('Tennis');
   
INSERT INTO locations (name) VALUES
	('Jakarta'),
	('Bogor'),
	('Depok'),
	('Tangerang'),
	('Bekasi');

INSERT INTO available_hours (start_time, end_time) VALUES
    ('09:00:00', '09:59:59'),
    ('10:00:00', '10:59:59'),
    ('11:00:00', '11:59:59'),
    ('12:00:00', '12:59:59'),
    ('13:00:00', '13:59:59'),
    ('14:00:00', '14:59:59'),
    ('15:00:00', '15:59:59'),
    ('16:00:00', '16:59:59'),
    ('17:00:00', '17:59:59'),
    ('18:00:00', '18:59:59'),
    ('19:00:00', '19:59:59'),
    ('20:00:00', '20:59:59'),
    ('21:00:00', '21:59:59');

INSERT INTO fasilites (bathroom, cafetaria, vehichle_park, prayer_room, changing_room, cctv) VALUES
	(2, true, 15, true, 2, true),
	(1, false, 10, false, 1, true),
	(2, true, 20, true, 2, false),
	(1, true, 18, true, 1, true),
	(3, false, 25, false, 3, true),
	(1, true, 12, true, 1, false),
	(2, true, 22, true, 2, true),
	(1, false, 14, false, 1, true),
	(2, true, 16, true, 2, true),
	(1, true, 20, true, 1, false),
	(2, false, 15, false, 2, true),
	(1, true, 18, true, 1, true),
	(2, true, 24, true, 2, false),
	(1, false, 12, true, 1, true),
	(3, false, 40, true, 2, true),
	(2, true, 35, true, 2, false),
	(3, true, 50, false, 3, true),
	(2, false, 30, true, 2, true),
	(4, true, 45, true, 3, false),
	(3, true, 38, false, 3, true),
	(2, true, 25, true, 2, false),
	(4, false, 50, true, 4, true),
	(3, true, 40, false, 2, true),
	(2, false, 30, true, 2, false),
	(4, true, 45, true, 3, true),
	(3, true, 35, true, 2, false),
	(4, true, 60, true, 3, true),
	(3, false, 50, true, 2, true),
	(4, true, 55, false, 3, false),
	(3, true, 65, true, 2, true),
	(4, true, 70, true, 3, true),
	(3, false, 45, false, 2, true),
	(4, true, 60, true, 4, false);

INSERT INTO fields (name, category_id, location_id, address, price, facility_id, created_at, updated_at, created_by) VALUES
    ('Grand Sport Kuningan Center', 2, 1, 'Jl. Karet Pedurenan Masjid No.45, RT.6/RW.7, Kuningan, Karet Kuningan, Kecamatan Setiabudi, Kota Jakarta Selatan', 80000.00, 1, NOW(), NOW(), 2),
    ('Supreme Arena Badminton', 2, 1, 'Jalan Karet Pedurenan Al Barokah #12A, RT.4/RW.6, Kuningan, Karet Kuningan, Kecamatan Setiabudi, Kota Jakarta Selatan', 120000.00, 2, NOW(), NOW(), 3),
    ('Badminton Tomang Field', 2, 1, 'Gedung Apotik Tomang Raya, Jl. Tomang Raya No.25 Lt. 3', 75000.00, 3, NOW(), NOW(), 4),
    ('Fullbelly Sports', 2, 2, 'Jl. Kaum Sari No.1, Cibuluh, Kec. Bogor Utara', 65000.00, 4, NOW(), NOW(), 2),
    ('Ginza Sport', 2, 2, 'Jl. Moro No.36, Ciangsana, Kec. Gn. Putri', 60000.00, 5, NOW(), NOW(), 3),
    ('ASBC Sport Center', 2, 2, 'JW7R+FPQ, Nagrak, Gunung Putri', 140000.00, 6, NOW(), NOW(), 4),
    ('GOR MG Sport', 2, 5, 'Kp Jl. Selang Tengah, RT.03/RW02, Wanasari, Kec. Cibitung', 30000.00, 7, NOW(), NOW(), 5),
    ('Kampiun Sports Center MM2100', 2, 5, 'Cikedokan, Cikarang Barat', 35000.00, 8, NOW(), NOW(), 2),
    ('Sudarsono Badminton Hall', 2, 4, 'Jl. Gotong Royong, Larangan Indah, Kec. Larangan', 50000.00, 9, NOW(), NOW(), 3),
    ('Goat Badminton', 2, 4, 'Jl. Pd. Betung Raya No.9, Pd. Betung, Kec. Pd. Aren', 50000.00, 10, NOW(), NOW(), 4),
    ('Boss Badminton Hall', 2, 4, 'Serua Asri - Block D, Banten 15414', 31000.00, 11, NOW(), NOW(), 5),
    ('GOR Badminton Kukusan', 2, 3, 'Jl. M.H. Sanim No.3, RT.2/RW.7, Kukusan, Kecamatan Beji', 50000.00, 12, NOW(), NOW(), 2),
    ('GOR Ibnu Mandiri', 2, 3, 'Jl. Gg. Langgar, RT.003/RW.002, Kemiri Muka, Kecamatan Beji', 45000.00, 13, NOW(), NOW(), 3),
    ('GOR Raden Tegar', 2, 3, 'HQRQ+WRC, Jl. Raya Keadilan, Rangkapan Jaya Baru, Kec. Pancoran Mas', 40000.00, 14, NOW(), NOW(), 4),
    ('Rawasi Sport Center', 1, 3, 'Jl. Ismaya, RT.04/RW.07, Cinere, Kec. Cinere', 50000.00, 15, NOW(), NOW(), 5),
    ('Cahaya Futsal', 1, 3, 'rt.02/06, Jl. Gandaria I, Ratu Jaya, Kec. Cipayung', 125000.00, 16, NOW(), NOW(), 2),
    ('Heis Futsal Depok', 1, 3, 'Rangkapan Jaya, Pancoran Mas', 80000.00, 17, NOW(), NOW(), 3),
    ('Futsal Corner Bekasi', 1, 5, 'Jl. Sultan Agung No.28, RT.002/RW.001, Medan Satria, Kecamatan Medan Satria', 65000.00, 18, NOW(), NOW(), 4),
    ('Estadio Futsal', 1, 5, 'Jl. Raya Perjuangan No.66, RT.003/RW.008, Marga Mulya, Kec. Bekasi Utara', 150000.00, 19, NOW(), NOW(), 5),
    ('Ralumbu Futsal Center', 1, 5, 'Jl. Dasa Dharma II No.23, RT.001/RW.017, Pengasinan, Kec. Rawalumbu', 80000.00, 20, NOW(), NOW(), 2),
    ('Futsal Center 27 Meruya', 1, 1, 'Jl. Meruya Selatan No.27, RT.5/RW.1, Joglo, Kec. Kembangan, Kota Jakarta Barat', 65000.00, 21, NOW(), NOW(), 3),
    ('Champion Futsal', 1, 1, 'Jl. Rw. Belong No.13 1, RT.1/RW.9, Kb. Jeruk, Kec. Kb. Jeruk, Kota Jakarta Barat', 175000.00, 22, NOW(), NOW(), 4),
    ('Parama Futsal Arena', 1, 1, 'Jalan Daan Mogot Km.17,4 No.10B, RT.1/RW.8, Kalideres, Kec. Kalideres, Kota Jakarta Barat', 140000.00, 23, NOW(), NOW(), 5),
    ('Welco Futsal', 1, 4, 'Sebelum Arah Strada Damos, Jl pabuaran Indah No.4, RT.003/RW.002, Pabuaran, Kec. Karawaci', 50000.00, 24, NOW(), NOW(), 2),
    ('End Sports Center', 1, 2, 'Jl. KH. R. Abdullah Bin Nuh No.23. b, RT.04/RW.07, Sindangbarang, Kec. Bogor Barat', 100000.00, 25, NOW(), NOW(), 3),
    ('Lafutsal', 1, 2, 'Jl. Raya Laladon, RT.01Rw01, Pagelaran, Kec. Ciomas', 100000.00, 26, NOW(), NOW(), 4),
	('GOAT Tennis Arena', 3, 4, 'Selatan, Banten, Jl. Pd. Betung Raya No.9, Bintaro, Kec. Pd. Aren', 100000.00, 27, NOW(), NOW(), 5),
	('The Racquet Club', 3, 1, 'Jl. Raya Kby. Lama No.64D, Grogol Sel., Kec. Kby. Lama, Kota Jakarta Selatan', 225000.00, 28, NOW(), NOW(), 2),
	('Bukit Pratama Tennis Court', 3, 1, 'Jl. Bukit Pratama Raya No.6, RT.7/RW.2, Lb. Bulus, Kec. Ciputat Tim., Kota Jakarta Selatan', 110000.00, 29, NOW(), NOW(), 3),
	('Orion Tennis Court', 3, 1, 'Jl. M. Saidi Raya No.13-14, RT.1/RW.5, Petukangan Sel., Kec. Pesanggrahan, Kota Jakarta Selatan', 85000.00, 30, NOW(), NOW(), 4),
	('Weedees Tennis Court', 3, 3, 'Jl. Pertanian No.35, Tirtajaya, Kec. Sukmajaya', 150000.00, 31, NOW(), NOW(), 5),
	('Lapangan Tennis Griya Depok', 3, 3, 'Griya Depok Asri Blok A1 Mekarjaya, Sukmajaya Depok', 50000.00, 32, NOW(), NOW(), 2),
	('Griya Tennis Court', 3, 5, 'Jl. Komp. Griya Harapan Permai No.5, RT.004/RW.032, Pejuang, Kecamatan Medan Satria', 80000.00, 33, NOW(), NOW(), 3);

INSERT INTO field_available_hours (field_id, available_hour_id) VALUES
    (1, 1), (1, 2), (1, 3), (1, 4), (1, 5), (1, 6), (1, 7), (1, 8), (1, 9), (1, 10), (1, 11), (1, 12), (1, 13),
    (2, 1), (2, 2), (2, 3), (2, 4), (2, 5), (2, 6), (2, 7), (2, 8), (2, 9), (2, 10), (2, 11), (2, 12), (2, 13),
    (3, 1), (3, 2), (3, 3), (3, 4), (3, 5), (3, 6), (3, 7), (3, 8), (3, 9), (3, 10), (3, 11), (3, 12), (3, 13),
    (4, 1), (4, 2), (4, 3), (4, 4), (4, 5), (4, 6), (4, 7), (4, 8), (4, 9), (4, 10), (4, 11), (4, 12), (4, 13),
    (5, 1), (5, 2), (5, 3), (5, 4), (5, 5), (5, 6), (5, 7), (5, 8), (5, 9), (5, 10), (5, 11), (5, 12), (5, 13),
    (6, 1), (6, 2), (6, 3), (6, 4), (6, 5), (6, 6), (6, 7), (6, 8), (6, 9), (6, 10), (6, 11), (6, 12), (6, 13),
    (7, 1), (7, 2), (7, 3), (7, 4), (7, 5), (7, 6), (7, 7), (7, 8), (7, 9), (7, 10), (7, 11), (7, 12), (7, 13),
    (8, 1), (8, 2), (8, 3), (8, 4), (8, 5), (8, 6), (8, 7), (8, 8), (8, 9), (8, 10), (8, 11), (8, 12), (8, 13),
    (9, 1), (9, 2), (9, 3), (9, 4), (9, 5), (9, 6), (9, 7), (9, 8), (9, 9), (9, 10), (9, 11), (9, 12), (9, 13),
    (10, 1), (10, 2), (10, 3), (10, 4), (10, 5), (10, 6), (10, 7), (10, 8), (10, 9), (10, 10), (10, 11), (10, 12), (10, 13),
    (11, 1), (11, 2), (11, 3), (11, 4), (11, 5), (11, 6), (11, 7), (11, 8), (11, 9), (11, 10), (11, 11), (11, 12), (11, 13),
    (12, 1), (12, 2), (12, 3), (12, 4), (12, 5), (12, 6), (12, 7), (12, 8), (12, 9), (12, 10), (12, 11), (12, 12), (12, 13),
    (13, 1), (13, 2), (13, 3), (13, 4), (13, 5), (13, 6), (13, 7), (13, 8), (13, 9), (13, 10), (13, 11), (13, 12), (13, 13),
    (14, 1), (14, 2), (14, 3), (14, 4), (14, 5), (14, 6), (14, 7), (14, 8), (14, 9), (14, 10), (14, 11), (14, 12), (14, 13),
    (15, 1), (15, 2), (15, 3), (15, 4), (15, 5), (15, 6), (15, 7), (15, 8), (15, 9), (15, 10), (15, 11), (15, 12), (15, 13),
    (16, 1), (16, 2), (16, 3), (16, 4), (16, 5), (16, 6), (16, 7), (16, 8), (16, 9), (16, 10), (16, 11), (16, 12), (16, 13),
    (17, 1), (17, 2), (17, 3), (17, 4), (17, 5), (17, 6), (17, 7), (17, 8), (17, 9), (17, 10), (17, 11), (17, 12), (17, 13),
    (18, 1), (18, 2), (18, 3), (18, 4), (18, 5), (18, 6), (18, 7), (18, 8), (18, 9), (18, 10), (18, 11), (18, 12), (18, 13),
    (19, 1), (19, 2), (19, 3), (19, 4), (19, 5), (19, 6), (19, 7), (19, 8), (19, 9), (19, 10), (19, 11), (19, 12), (19, 13),
    (20, 1), (20, 2), (20, 3), (20, 4), (20, 5), (20, 6), (20, 7), (20, 8), (20, 9), (20, 10), (20, 11), (20, 12), (20, 13),
    (21, 1), (21, 2), (21, 3), (21, 4), (21, 5), (21, 6), (21, 7), (21, 8), (21, 9), (21, 10), (21, 11), (21, 12), (21, 13),
    (22, 1), (22, 2), (22, 3), (22, 4), (22, 5), (22, 6), (22, 7), (22, 8), (22, 9), (22, 10), (22, 11), (22, 12), (22, 13),
    (23, 1), (23, 2), (23, 3), (23, 4), (23, 5), (23, 6), (23, 7), (23, 8), (23, 9), (23, 10), (23, 11), (23, 12), (23, 13),
    (24, 1), (24, 2), (24, 3), (24, 4), (24, 5), (24, 6), (24, 7), (24, 8), (24, 9), (24, 10), (24, 11), (24, 12), (24, 13),
    (25, 1), (25, 2), (25, 3), (25, 4), (25, 5), (25, 6), (25, 7), (25, 8), (25, 9), (25, 10), (25, 11), (25, 12), (25, 13),
    (26, 1), (26, 2), (26, 3), (26, 4), (26, 5), (26, 6), (26, 7), (26, 8), (26, 9), (26, 10), (26, 11), (26, 12), (26, 13),
    (27, 1), (27, 2), (27, 3), (27, 4), (27, 5), (27, 6), (27, 7), (27, 8), (27, 9), (27, 10), (27, 11), (27, 12), (27, 13),
    (28, 1), (28, 2), (28, 3), (28, 4), (28, 5), (28, 6), (28, 7), (28, 8), (28, 9), (28, 10), (28, 11), (28, 12), (28, 13),
    (29, 1), (29, 2), (29, 3), (29, 4), (29, 5), (29, 6), (29, 7), (29, 8), (29, 9), (29, 10), (29, 11), (29, 12), (29, 13),
    (30, 1), (30, 2), (30, 3), (30, 4), (30, 5), (30, 6), (30, 7), (30, 8), (30, 9), (30, 10), (30, 11), (30, 12), (30, 13),
    (31, 1), (31, 2), (31, 3), (31, 4), (31, 5), (31, 6), (31, 7), (31, 8), (31, 9), (31, 10), (31, 11), (31, 12), (31, 13),
    (32, 1), (32, 2), (32, 3), (32, 4), (32, 5), (32, 6), (32, 7), (32, 8), (32, 9), (32, 10), (32, 11), (32, 12), (32, 13),
    (33, 1), (33, 2), (33, 3), (33, 4), (33, 5), (33, 6), (33, 7), (33, 8), (33, 9), (33, 10), (33, 11), (33, 12), (33, 13);


INSERT INTO wallets (user_id, saldo, created_at, updated_at) VALUES
    (1, 1000000.00, NOW(), NOW()),
    (2, 500000.00, NOW(), NOW()),
    (3, 250000.00, NOW(), NOW()),
    (4, 275000.00, NOW(), NOW()),
    (5, 50000.00, NOW(), NOW());

INSERT INTO deposits (amount, wallet_id, created_at, updated_at) VALUES
    (50000.00, 1, NOW(), NOW()),
    (325000.00, 2, NOW(), NOW()),
    (100000.00, 3, NOW(), NOW());

INSERT INTO reservations (user_id, field_id, status, reservation_date, created_at, updated_at) VALUES
	(1, 2, 'confirmed', '2024-08-03', NOW(), NOW());

INSERT INTO payments (reservation_id, amount, payment_date, status, created_at, updated_at) VALUES
	(1, 120000.00, NOW(), 'completed', NOW(), NOW());