package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	//"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bertoxic/bert/drivers"
	"github.com/bertoxic/bert/helpers"
	"github.com/bertoxic/bert/internal/config"
	"github.com/bertoxic/bert/internal/forms"
	"github.com/bertoxic/bert/internal/render"
	"github.com/bertoxic/bert/models"
	"github.com/bertoxic/bert/repository"
	"github.com/bertoxic/bert/repository/dbrepo"
	"github.com/go-chi/chi"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig, db *drivers.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewTestingRepo(a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Sessions.Put(r.Context(), "remoteIp", remoteIp)

	render.Template(w, r, "home.page.gtpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteI := m.App.Sessions.GetString(r.Context(), "remoteIp")
	stringMp := make(map[string]string)
	stringMp["remoteIp"] = remoteI

	render.Template(w, r, "about.page.gtpl", &models.TemplateData{
		StringMap: stringMp,
	})
}
func (m *Repository) Example(w http.ResponseWriter, r *http.Request) {
	stringmap := make(map[string]string)
	stringmap["any"] = "anything in krips!"

	render.Template(w, r, "example.page.gtpl", &models.TemplateData{
		StringMap: stringmap,
	})
}
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "generals.page.gtpl", &models.TemplateData{})
}
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "majors.page.gtpl", &models.TemplateData{})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "search-availability.page.gtpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	start := r.Form.Get("start")
	end := r.Form.Get("end")
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "can't parse start date!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "can't parse end date!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	rooms, err := m.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "can't get availability for rooms")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	println(len(rooms), "the lennnnnnnnnnnn")
	if len(rooms) == 0 {
		m.App.Sessions.Put(r.Context(), "error", "Sorry this room has already been booked at that date")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		return
	}
	data := make(map[string]interface{})
	data["rooms"] = rooms
	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}
	m.App.Sessions.Put(r.Context(), "reservation", res)
	render.Template(w, r, "choose-room.page.gtpl", &models.TemplateData{
		Data: data,
	})

	//w.Write([]byte(fmt.Sprintf("your startdate is %s and end depature is %s", start, end)))

	for _, i := range rooms {
		m.App.InfoLog.Println("rooms avail:", i.ID, i.RoomName, i.CreatedAt)
	}

}

type jsonResponse struct {
	OK        bool   `json:"ok"`
	Message   string `json:"Available"`
	RoomID    string `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		resp := jsonResponse{
			OK:      false,
			Message: "Internal server error",
		}
		out, _ := json.MarshalIndent(resp, "", "   ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	sd := r.Form.Get("start")
	ed := r.Form.Get("end")
	roomID, _ := strconv.Atoi(r.Form.Get("room_id"))
	layout := "2006-01-02"

	startDate, _ := time.Parse(layout, sd)

	endDate, _ := time.Parse(layout, ed)

	available, err := m.DB.SearchAvailabilityByDatesByRoomID(startDate, endDate, roomID)
	if err != nil {
		resp := jsonResponse{
			OK:      false,
			Message: "unable to connect to the database",
		}
		out, _ := json.MarshalIndent(resp, "", "   ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	resp := jsonResponse{
		OK:        available,
		Message:   "we the best GODDID",
		RoomID:    strconv.Itoa(roomID),
		StartDate: sd,
		EndDate:   ed,
	}

	out, _ := json.MarshalIndent(resp, "", "   ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	//return
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	res, ok := m.App.Sessions.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Sessions.Put(r.Context(), "error", "can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	room, err := m.DB.GetRoomByID(res.RoomID)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "can't find room!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	res.Room.RoomName = room.RoomName

	m.App.Sessions.Put(r.Context(), "reservation", res)

	sd := res.StartDate.Format("2006-01-02")
	ed := res.EndDate.Format("2006-01-02")

	stringMap := make(map[string]string)
	stringMap["startDate"] = sd
	stringMap["endDate"] = ed

	data := make(map[string]interface{})
	data["reservation"] = res

	render.Template(w, r, "make-reservation.page.gtpl", &models.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

	// reservation, ok := m.App.Sessions.Get(r.Context(), "reservation").(models.Reservation)
	// if !ok {
	// 	helpers.ServerError(w, errors.New("reservation not found in postreservation "))
	// 	return
	// }
	err := r.ParseForm()
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "cannot parse form")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	//2023 01 02 -- 01/01 03:04:05PM '06 -0700
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, sd)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "can't parse startdate")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "can't get parse enddate")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "invalid data")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	room, err := m.DB.GetRoomByID(roomID)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "invalid data")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// reservation.FirstName = r.Form.Get("first_name")
	// reservation.LastName = r.Form.Get("last_name")
	// reservation.Phone = r.Form.Get("phone")
	// reservation.Email = r.Form.Get("email")

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomID,
		Room:      room,
	}
	rx := m.App.Sessions.Get(r.Context(), "reservation").(models.Reservation)
	 reservation.Room.RoomName = rx.Room.RoomName
	form := forms.New(r.PostForm)
	// form.Has("first_name",r)
	form.MinLength("first_name", 3)
	form.Required("first_name", "last_name", "email", "phone")
	if !form.Valid() {
		data := make(map[string]interface{})
		stringMap := make(map[string]string)
		stringMap["startDate"] = sd
		stringMap["endDate"] = sd
		data["reservation"] = reservation
		//http.Error(w, "my own error message", http.StatusSeeOther)
		render.Template(w, r, "make-reservation.page.gtpl", &models.TemplateData{
			Form: form,
			Data: data,
			StringMap: stringMap,
		})

		return
	}

	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "cannot insert reservation into Database")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	restriction := models.RoomRestriction{
		StartDate:     startDate,
		EndDate:       endDate,
		RoomID:        roomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "can't insert room-restriction")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	htmlMessages := fmt.Sprintf("Dear %s you have made a reservation for %s from %s to %s thank you", reservation.FirstName, reservation.Room.RoomName, reservation.StartDate.Format("2006-01-02"), reservation.EndDate.Format("2006-01-02"))
	//send Notifications
	msg := models.MailData{
		To:       reservation.Email,
		From:     "ok@fmail.com",
		Subject:  "Heazo's Lodge Reservation",
		Content:  htmlMessages,
		Template: "basic.html",
	}
	//	m.App.MailChan <- msg

	htmlMessages = fmt.Sprintf("<strong>Bertoxic Vents</strong> <br>Dear Admin a reservation has been made for <strong>%s</strong> from %s to %s. thank you", reservation.Room.RoomName, reservation.StartDate.Format("2006-01-02"), reservation.EndDate.Format("2006-01-02"))
	//send Notifications
	msg = models.MailData{
		To:      "Admin@mail.com",
		From:    "me@here.com",
		Subject: "Bertoxic's Lodge Reservation",
		Content: htmlMessages,
	}
	m.App.MailChan <- msg

	m.App.Sessions.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Sessions.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot obtain item from our session")
		m.App.Sessions.Put(r.Context(), "error", "Cannot find in reservation summary")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data := make(map[string]interface{})
	data["reservation"] = reservation
	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")
	stringmap := make(map[string]string)
	stringmap["start_date"] = sd
	stringmap["end_date"] = ed
	render.Template(w, r, "reservation-summary.page.gtpl", &models.TemplateData{
		Data:      data,
		StringMap: stringmap,
	})
}

func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")
	//roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	roomID, err := strconv.Atoi(exploded[2])
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "missing url parameter")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	res, ok := m.App.Sessions.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Sessions.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	res.RoomID = roomID
	m.App.Sessions.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)

}

func (m *Repository) BookRoom(w http.ResponseWriter, r *http.Request) {
	roomID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	sd := r.URL.Query().Get("s")
	ed := r.URL.Query().Get("e")

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, sd)
	endDate, _ := time.Parse(layout, ed)

	var res models.Reservation

	room, err := m.DB.GetRoomByID(roomID)
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "Can't get room from db!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	res.Room.RoomName = room.RoomName
	res.RoomID = roomID
	res.StartDate = startDate
	res.EndDate = endDate

	m.App.Sessions.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}

func (m *Repository) ShowLogin(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "login.page.gtpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository) PostShowLogin(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Sessions.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	forms := forms.New(r.PostForm)
	forms.MinLength("email", 3)
	forms.Required("email", "password")
	if !forms.Valid() {
		render.Template(w, r, "login.page.gtpl", &models.TemplateData{
			Form: forms,
		})
		return
	}
	id, n, err := m.DB.Authenticate(email, password)
	if err != nil {
		log.Println(err, n)
		m.App.Sessions.Put(r.Context(), "error", "invalid login details")
		http.Redirect(w, r, "/user-login", http.StatusSeeOther)
		return
	}
	m.App.Sessions.Put(r.Context(), "user_id", id)
	m.App.Sessions.Put(r.Context(), "flash", "Logged in successfully")
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

//logs out user from session
func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	m.App.Sessions.Destroy(r.Context())
	m.App.Sessions.RenewToken(r.Context())
	http.Redirect(w, r, "/user-login", http.StatusSeeOther)

}

//navigates to the admindashboard
func (m *Repository) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "admin-dashboard.page.gtpl", &models.TemplateData{})
}

//Shows all new Reservartions in admin page
func (m *Repository) AdminNewReservations(w http.ResponseWriter, r *http.Request) {
	reservations, err := m.DB.AllNewReservations()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	data := make(map[string]interface{})
	data["reservations"] = reservations
	render.Template(w, r, "admin-new-reservations.page.gtpl", &models.TemplateData{
		Data: data,
	})

}

//Shows all reservations in admin page
func (m *Repository) AdminAllReservations(w http.ResponseWriter, r *http.Request) {
	reservations, err := m.DB.AllReservations()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	data := make(map[string]interface{})
	data["reservations"] = reservations
	render.Template(w, r, "admin-all-reservations.page.gtpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) AdminShowReservation(w http.ResponseWriter, r *http.Request) {
	explode := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(explode[4])
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	src := explode[3]

	reservation, err := m.DB.GetReservationByID(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	data := make(map[string]interface{})
	stringMap := make(map[string]string)
	stringMap["src"] = src

	year := r.URL.Query().Get("y")
	month := r.URL.Query().Get("m")
	stringMap["year"] = year
	stringMap["month"] = month

	data["reservation"] = reservation
	render.Template(w, r, "admin-reservation-show.page.gtpl", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
		Form:      forms.New(nil),
	})
}

func (m *Repository) AdminModifyReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		m.App.Sessions.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	explode := strings.Split(r.RequestURI, "/")
	fid := explode[4]
	id, err := strconv.Atoi(fid)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	src := explode[3]

	res, err := m.DB.GetReservationByID(id)
	res.FirstName = r.Form.Get("first_name")
	res.LastName = r.Form.Get("last_name")
	res.Email = r.Form.Get("email")
	res.Phone = r.Form.Get("phone")

	println("previous:", res.FirstName, res.LastName)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	err = m.DB.UpdateReservation(res)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	month := r.Form.Get("month")
	year := r.Form.Get("year")

	m.App.Sessions.Put(r.Context(), "flash", "Changes made")
	if year == "" {
		http.Redirect(w, r, fmt.Sprintf("/admin/reservations-%s", src), http.StatusSeeOther)

	} else {
		http.Redirect(w, r, fmt.Sprintf("/admin/reservation-calender?y=%s&m=%s", year, month), http.StatusSeeOther)

	}

}

func (m *Repository) AdminReservationCalender(w http.ResponseWriter, r *http.Request) {
	//case were there are no mpnth and year specified

	now := time.Now()
	if r.URL.Query().Get("y") != "" {
		year, _ := strconv.Atoi(r.URL.Query().Get("y"))
		month, _ := strconv.Atoi(r.URL.Query().Get("m"))
		now = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	}
	data := make(map[string]interface{})
	data["now"] = now
	next := now.AddDate(0, 1, 0)
	last := now.AddDate(0, -1, 0)

	nextMonth := next.Format("01")
	nextMonthYear := next.Format("2006")
	lastMonth := last.Format("01")
	lastMonthYear := last.Format("2006")

	stringMap := make(map[string]string)
	stringMap["next_month"] = nextMonth
	stringMap["next_month_year"] = nextMonthYear
	stringMap["last_month_year"] = lastMonthYear
	stringMap["last_month"] = lastMonth

	stringMap["this_month"] = now.Format("01")
	stringMap["this_month_year"] = now.Format("2006")

	currentYear, currrentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currrentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	intMap := make(map[string]int)
	intMap["days_in_month"] = lastOfMonth.Day()
	rooms, err := m.DB.AllRooms()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	data["rooms"] = rooms
	for _, v := range rooms {
		reservationMap := make(map[string]int)
		blockMap := make(map[string]int)

		for d := firstOfMonth; d.After(lastOfMonth) == false; d = d.AddDate(0, 0, 1) {
			reservationMap[d.Format("2006-01-02")] = 0
			blockMap[d.Format("2006-01-02")] = 0
		}
		// getting all the restructions fotr the current row

		restrictions, err := m.DB.GetRestrictionsForRoomByDate(v.ID, firstOfMonth, lastOfMonth)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}

		for _, y := range restrictions {
			if y.ReservationID > 0 {

				//it's a reservation
				for d := y.StartDate; d.After(y.EndDate) == false; d = d.AddDate(0, 0, 1) {

					reservationMap[d.Format("2006-01-02")] = y.ReservationID
				}

			} else {
				// it is a block
				blockMap[y.StartDate.Format("2006-01-02")] = y.ID
			}
		}
		data[fmt.Sprintf("reservation_map_%d", v.ID)] = reservationMap
		data[fmt.Sprintf("block_map_%d", v.ID)] = blockMap

		m.App.Sessions.Put(r.Context(), fmt.Sprintf("block_map_%d", v.ID), blockMap)

	}

	render.Template(w, r, "admin-reservation-calender.page.gtpl", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
		IntMap:    intMap,
	})

}

func (m *Repository) AdminProcessReservation(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	src := chi.URLParam(r, "src")
	year := r.URL.Query().Get("y")
	month := r.URL.Query().Get("m")
	err = m.DB.UpdateReservationProcessed(id, 1)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	m.App.Sessions.Put(r.Context(), "flash", "Reservation marked as processed")
	if year == "" {
		http.Redirect(w, r, fmt.Sprintf("/admin/reservations-%s", src), http.StatusSeeOther)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/admin/reservation-calender?y=%s&m=%s", year, month), http.StatusSeeOther)
	}

}

//deletes a reservation
func (m *Repository) AdminDeleteReservation(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	src := chi.URLParam(r, "src")
	year := r.URL.Query().Get("y")
	month := r.URL.Query().Get("m")
	err = m.DB.DeleteReservation(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	if year == "" {
		http.Redirect(w, r, fmt.Sprintf("/admin/reservations-%s", src), http.StatusSeeOther)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/admin/reservation-calender?y=%s&m=%s", year, month), http.StatusSeeOther)
	}
}

func (m *Repository) AdminPostReservationCalender(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	year, _ := strconv.Atoi(r.Form.Get("y"))
	month, _ := strconv.Atoi(r.Form.Get("m"))

	rooms, err := m.DB.AllRooms()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	form := forms.New(r.PostForm)
	for _, room := range rooms {

		curMap := m.App.Sessions.Get(r.Context(), fmt.Sprintf("block_map_%d", room.ID)).(map[string]int)

		for name, value := range curMap {

			if val, ok := curMap[name]; ok {
				if val > 0 {
					if !form.Has(fmt.Sprintf("remove_block_%d_%s", room.ID, name)) {
						err := m.DB.DeleteBlockBYID(value)
						if err != nil {
							helpers.ServerError(w, err)
							return
						}
					}
				}
			}
		}
	}

	// now handle blocks

	for name, _ := range r.PostForm {
		if strings.HasPrefix(name, "add_block") {
			exploded := strings.Split(name, "_")
			roomID, _ := strconv.Atoi(exploded[2])
			t, _ := time.Parse("2006-01-2", exploded[3])

			err := m.DB.InsertBlockForRoom(roomID, t)
			if err != nil {
				helpers.ServerError(w, err)
				return
			}
			log.Println("would insert block into room ", roomID, "for date", exploded[3])
		}
	}

	m.App.Sessions.Put(r.Context(), "flash", "changes are saved")
	http.Redirect(w, r, fmt.Sprintf("/admin/reservation-calender?y=%d&m=%d", year, month), http.StatusSeeOther)

}
