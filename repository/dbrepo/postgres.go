package dbrepo

import (
	"context"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/bertoxic/bert/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var newID int
	stmt := `insert into reservations (first_name, last_name , email , phone, start_date,
        end_date , room_id, created_at , updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

//inserts a room restriction into the database

func (m *postgresDBRepo) InsertRoomRestriction(md models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into room_restrictions (start_date, end_date ,room_id, reservation_id ,
        restriction_id, created_at, updated_at) values ($1, $2, $3, $4, $5,$6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		md.StartDate,
		md.EndDate,
		md.RoomID,
		md.ReservationID,
		md.RestrictionID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil

}

// returns true if availability for roomID but false if not

func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var numRows int

	stmt := `
select
   count(id)
from
   room_restrictions rr
where  
    room_id = $1
and $2 < end_date and $3 >= start_date; 
`

	row := m.DB.QueryRowContext(ctx, stmt, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		log.Println("...........>>>>>>>>.>>>>>>>>>>>>>>>>>>>>cbccbcbcbc")
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

// SearchAvailabilityForAllRooms returns room id and name fo all rooms available
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
		select
			r.id, r.room_name
		from
			rooms r
		where r.id not in 
		(select room_id from room_restrictions rr where $1 < rr.end_date and $2 >= rr.start_date);		`

	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}
	defer rows.Close()
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

func (m *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room

	query := `
		select id, room_name, created_at, updated_at from rooms where id = $1
`
	//log.Println("error in postgres.go vvvvvvvvvvvvvvvx", id)
	row := m.DB.QueryRowContext(ctx, query, id)
	//log.Println("error in postgres.go vvvvvvvvvvvvvvvx", room.RoomName)

	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		log.Println("error in postgres.go vvvvvvvvvvvvvvv", room.RoomName)
		return room, err
	}
	if err = row.Err(); err != nil {
		return room, err
	}

	return room, nil
}

//returns a user by ID

func (m *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u models.User
	query := `--sql
	select id, first_name, last_name, email, password, access_level, created_at, updated_at from users where id = $1 
	`

	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (m *postgresDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `--sql
	update users set first_name =$1, last_name = $2, email = $3, access_level=$4, updated_at = $5
	`
	_, err := m.DB.ExecContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.AccessLevel,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

//Authentucates a user
func (m *postgresDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var hashedPassword string

	row := m.DB.QueryRowContext(ctx, "select id, password from users where email = $1", email)
	err := row.Scan(&id, &hashedPassword)

	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return id, "", errors.New("incorrect password")
	} else if err != nil {
		return id, "", err
	}
	return id, hashedPassword, nil
}

func (m *postgresDBRepo) AllReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `--sql
	select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at, r.processed, rm.id, rm.room_name from reservations r 
	left join rooms rm on (r.room_id = rm.id) 
	order by r.start_date asc
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()
	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Processed,
			&i.Room.ID,
			&i.Room.RoomName,
		)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}

	if err = rows.Err(); err != nil {

		return reservations, err
	}
	return reservations, nil
}

func (m *postgresDBRepo) AllNewReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `--sql
		select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at, rm.id, rm.room_name from reservations r 
		left join rooms rm on (r.room_id = rm.id)
		where processed = 0 
		order by r.start_date asc
		`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()
	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Room.ID,
			&i.Room.RoomName,
		)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}

	if err = rows.Err(); err != nil {

		return reservations, err
	}
	return reservations, nil
}

// returns one reservation by id
func (m *postgresDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var res models.Reservation

	// stmt :=` update set reservations processed =1 where id=$1
	// `

	query := `--sql
	select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date, r.end_date, r.room_id, r.created_at, r.processed, r.updated_at, rm.id, rm.room_name from reservations r 
	left join rooms rm on (r.room_id = rm.id) 
	where r.id = $1
	`
	row := m.DB.QueryRowContext(ctx, query,id)
	err := row.Scan(
			&res.ID,
			&res.FirstName,
			&res.LastName,
			&res.Email,
			&res.Phone,
			&res.StartDate,
			&res.EndDate,
			&res.RoomID,
			&res.CreatedAt,
			&res.Processed,
			&res.UpdatedAt,
			&res.Room.ID,
			&res.Room.RoomName,	
)
if err != nil {
	return res , err
}
return res , nil
}


func (m *postgresDBRepo) UpdateReservation (u models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `--sql
	update reservations set first_name =$1, last_name = $2, email = $3, phone=$4,  updated_at = $5 
	where id = $6
	`
	_, err := m.DB.ExecContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Phone,
		time.Now(),
		u.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) DeleteReservation(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `--sql
	delete from reservations where id = $1
	`
	_, err := m.DB.ExecContext(ctx, query,
	id,
	)
	if err != nil {
		return err
	}
	return nil
}


//UpdateReservationProcessed updates processed for reservation
func (m *postgresDBRepo) UpdateReservationProcessed(id,processed int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `--sql
	update reservations set processed = $1 where id = $2
	`
	_, err := m.DB.ExecContext(ctx, query,
		processed,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo)AllRooms ()([]models.Room, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	var rooms []models.Room
	query := ` select id, room_name, created_at, updated_at from rooms order by room_name
	`
	rows, err:=m.DB.QueryContext(ctx, query)
	if err != nil {
		return rooms, err
	}
 defer rows.Close()

  for rows.Next(){
	var rm models.Room

	err := rows.Scan(
		 &rm.ID,
		 &rm.RoomName,
		 &rm.CreatedAt,
		 &rm.UpdatedAt,
		)
		
		rooms = append(rooms,rm )
		if err =rows.Err(); err != nil {
			return rooms, err
		}
	}
	return rooms , err
	
}

// returns restrictions for a room by date range
func (m *postgresDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction,error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var restrictions []models.RoomRestriction
	query := `--sql
	select id, coalesce(reservation_id, 0), restriction_id, room_id, start_date, end_date from room_restrictions where $1 < end_date and $2 >= start_date and room_id =$3
	`
	rows, err := m.DB.QueryContext(ctx, query, start,end,roomID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var r models.RoomRestriction
			err:= rows.Scan(
				&r.ID,
				&r.ReservationID,
				&r.RestrictionID,
				&r.RoomID,
				&r.StartDate,
				&r.EndDate,
			
			)
			if err != nil {
				return nil, err
			}
			restrictions = append(restrictions, r)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
	return restrictions, nil
}

// InsertBlockForRoom inserts room blodk
func (m *postgresDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `--sql
		insert into room_restrictions (start_date, end_date, room_id, restriction_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)
	`
	_, err := m.DB.ExecContext( ctx, query, startDate, startDate.AddDate(0, 0, 1), id, 2, time.Now(), time.Now())

		if err != nil {
			log.Println(err)
			return err
		}
		return nil

}

func (m *postgresDBRepo) DeleteBlockBYID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `--sql
		delete from room_restrictions where id = $1
	`
	_, err := m.DB.ExecContext( ctx, query, id)

		if err != nil {
			log.Println(err)
			return err
		}
		return nil

}