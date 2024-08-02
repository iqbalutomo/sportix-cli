package repository

import (
	"database/sql"
	"sportix-cli/internal/entity"
)

type FacilityRepo interface {
	FindFacilityById(facilitID int) (*entity.Facility, error)
	EditFacility(facility *entity.Facility) error
}

type facilityRepo struct {
	db *sql.DB
}

func (fa *facilityRepo) FindFacilityById(facilityID int) (*entity.Facility, error) {
	var facility entity.Facility

	query := `SELECT`
	row := fa.db.QueryRow(query, facilityID)

	if err := row.Scan(); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &facility, nil
}

func (fa *facilityRepo) EditFacility(facility *entity.Facility) error {
	editFacility := `UPDATE facilities SET bathroom=?, cafetaria=?, vehicle_park=?, prayer_room=?, changing_room=?, cctv=? WHERE facility_id=?`
	_, err := fa.db.Exec(editFacility, facility.Bathroom, facility.Cafeteria, facility.VehiclePark, facility.PrayerRoom, facility.ChangingRoom, facility.CCTV, facility.FacilityID)
	return err
}
