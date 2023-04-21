{{template "base" .}}

{{define "content"}}

{{$res := index .Data "reservation"}}
<p class="container"><strong>Reservation Details</strong> <br>

Arrival:{{ index .StringMap "startDate"}} <br>
Depature: {{ index .StringMap "endDate"}} <br>
Room Name: {{$res.Room.RoomName}} <br> 
</P>


<form method="POST" action="/make-reservation" class="needs-validation container" novalidate>
<input type="text" name="csrf_token" value="{{.CSRFToken}}">
<input type="hidden" name="room_id" value="{{$res.RoomID}}">
<input type="hidden" name="start_date" value="{{ index .StringMap "startDate"}}">
<input type="hidden" name="end_date" value="{{index .StringMap "endDate" }}">
      <div class="form-group mt-3">
        <label for="first_name">First Name</label>
        {{with .Form.Errors.Get "first_name"}}
        <label class="text-danger">{{.}}</label>
        {{end}}
        <input name="first_name" value="{{$res.FirstName}}" id="first_name" type="text" placeholder="Enter your first name" class="form-control" required
          autocomplete="off" type="text" >
      </div>
      <div class="form-group mt-3">
        <label for="last_name">Last Name</label>
         {{with .Form.Errors.Get "last_name"}}
        <label class="text-danger">{{.}}</label>
        {{end}}
        <input name="last_name" value="{{$res.LastName}}" id="last_name" class="form-control" required autocomplete="off" type="text"
          placeholder="Enter your last name">
      </div>
    
      <div class="form-group mt-3">
        <label for="email">Email Address </label>
         {{with .Form.Errors.Get "email"}}
        <label class="text-danger">{{.}}</label>
        {{end}}
        <input name="email" value="{{$res.Email}}" id="email" class="form-control" required autocomplete="off" type="text"
          placeholder="Enter your email">
      </div>

      <div class="form-group mt-3">
        <label for="phone">Phone Number </label>
         {{with .Form.Errors.Get "phone"}}
        <label class="text-danger">{{.}}</label>
        {{end}}
        <input name="phone" value="{{$res.Phone}}" id="Phone" class="form-control" required autocomplete="off" type="text"
          placeholder="Enter your Phone number">
          <hr>
         
      <button type="submit"class="btn btn-success" value="make reservation">Make reservation</button>
    </form>
    
    <br>


    <section class="container home-text">
      <p>Welcome toScorching mood they flatterers old none that felt himnot, vile was but soils passed, vaunted of oh if
        nor. Grace a peace heavenly me nor the. Shameless care into in did. Loved dwell shameless so was are lines,
        uncouth the smile childe he relief ah or in oft, virtues the a him with ah,. </p>
      <button class="btn btn-outline-success">
        Click here
      </button>
    </section>
{{end}}

