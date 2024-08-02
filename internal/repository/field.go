package repository

import (
	"database/sql"
	"sportix-cli/internal/entity"
)

type FieldRepo interface {
	FindAllFields() ([]entity.Field, error)
	FindAllHoursByFieldID(fieldID uint) ([]entity.FieldAvailableHour, error)
	FindFieldById(fieldID int) (*entity.Field, error)
	EditField(field *entity.Field) error
}

type fieldRepo struct {
	db *sql.DB
}

func NewFieldRepo(db *sql.DB) FieldRepo {
	return &fieldRepo{db}
}

func (f *fieldRepo) FindAllFields() ([]entity.Field, error) {
	query := `SELECT f.field_id, f.name AS field_name, f.price, c.name AS category_name, l.name AS location_name, fac.bathroom,
				fac.cafeteria, fac.vehicle_park, fac.prayer_room, fac.changing_room, fac.cctv, f.address, u.username AS created_by_username
			FROM fields f
			JOIN categories c ON f.category_id = c.category_id
			JOIN locations l ON f.location_id = l.location_id
			JOIN facilities fac ON f.facility_id = fac.facility_id
			LEFT JOIN users u ON f.created_by = u.user_id
			ORDER BY f.name;`
	rows, err := f.db.Query(query)
	if err != nil {
		return nil, err
	}

	var fields []entity.Field
	for rows.Next() {
		var field entity.Field
		if err := rows.Scan(&field.FieldID, &field.Name, &field.Price, &field.Category.Name, &field.Location.Name, &field.Facility.Bathroom, &field.Facility.Cafeteria, &field.Facility.VehiclePark, &field.Facility.PrayerRoom, &field.Facility.ChangingRoom, &field.Facility.CCTV, &field.Address, &field.CreatedBy.Username); err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}

	return fields, nil
}

func (f *fieldRepo) FindAllHoursByFieldID(fieldID uint) ([]entity.FieldAvailableHour, error) {
	query := `SELECT fa.field_available_hour_id, a.start_time, a.end_time, fa.status
			FROM field_available_hours fa
			JOIN available_hours a ON fa.available_hour_id = a.available_hour_id
			WHERE fa.field_id = ?;`
	rows, err := f.db.Query(query, fieldID)
	if err != nil {
		return nil, err
	}

	var fieldAvailableHours []entity.FieldAvailableHour
	for rows.Next() {
		var availableHours entity.FieldAvailableHour
		if err := rows.Scan(&availableHours.FieldAvailableHourID, &availableHours.AvailableHourID.StartTime, &availableHours.AvailableHourID.EndTime, &availableHours.Status); err != nil {
			return nil, err
		}
		fieldAvailableHours = append(fieldAvailableHours, availableHours)

	}

	return fieldAvailableHours, nil
}

func (f *fieldRepo) FindFieldById(fieldID int) (*entity.Field, error) {
	var field entity.Field

	query := `SELECT 
    f.field_id,
    f.name AS field_name,
    f.price,
    c.name AS category_name,
    l.name AS location_name,
    fac.bathroom,
    fac.cafeteria,
    fac.vehicle_park,
    fac.prayer_room,
    fac.changing_room,
    fac.cctv,
    u.username AS created_by_username
FROM 
    fields f
JOIN 
    categories c ON f.category_id = c.category_id
JOIN 
    locations l ON f.location_id = l.location_id
JOIN 
    facilities fac ON f.facility_id = fac.facility_id
LEFT JOIN 
    users u ON f.created_by = u.user_id
WHERE field_id=?`
	row := f.db.QueryRow(query, fieldID)

	if err := row.Scan(&field.FieldID, &field.Name, &field.Price, &field.Category.Name, &field.Location.Name, &field.Facility.Bathroom, &field.Facility.Cafeteria, &field.Facility.VehiclePark, &field.Facility.PrayerRoom, &field.Facility.ChangingRoom, &field.Facility.CCTV, &field.CreatedBy.Username); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &field, nil
}

func (f *fieldRepo) EditField(field *entity.Field) error {
	editFieldQuery := `UPDATE fields SET name=?, price=?, address=? WHERE field_id=?;`
	_, err := f.db.Exec(editFieldQuery, field.Name, field.Price, field.Address, "", field.FieldID)
	return err
}
