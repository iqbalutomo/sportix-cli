

-- Trigger untuk memastikan reservation_date bukan tanggal kemarin
DELIMITER $$
DROP TRIGGER IF EXISTS before_insert_reservations$$
CREATE TRIGGER before_insert_reservations
    BEFORE INSERT ON reservations
    FOR EACH ROW
BEGIN
    IF NEW.reservation_date < CURDATE() THEN
    SIGNAL SQLSTATE '45000'
    SET MESSAGE_TEXT = 'Reservation date cannot be in the past';
END IF;
END$$