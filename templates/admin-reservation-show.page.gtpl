{{template "admin" .}}

{{define "page-title"}}
    reservation show
{{end}}


{{define "content"}}

   {{$res := index .Data "reservation"}}
   {{$src := index .StringMap "src"}}

  <p>Name:<strong> {{$res.FirstName}} {{$res.LastName}} </strong></p>
   <br>
   <br>

   <p>
   Arrival Date: <strong>{{humanDate $res.StartDate}}</strong> <br>
   Depature Date: <strong>{{humanDate $res.EndDate}}</strong><br>
   Room Name: <strong>{{ $res.Room.RoomName}}</strong><br>
   </p>
{{/* admin/reservation/{{$src}}/{{$res.ID}} */}}
   <form method="POST" action="/admin/reservation/{{$src}}/{{$res.ID}}" class="" novalidate>
    <input type="text" name="csrf_token" value="{{.CSRFToken}}">
    <input type="hidden" name="year" value="{{index .StringMap "year"}}">
    <input type="hidden" name="month" value="{{index .StringMap "month"}}">

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
         <div class="float-start">
      <input type="submit"class="btn btn-success" value="Modify Reservation">
      {{if eq $src "cal"}}
       <a href="#!" onclick="window.history.go(-1)" class ="btn btn-light">Cancel</a>
      {{else}}
      <a href="/admin/reservations-{{$src}}" class ="btn btn-light">Cancel</a>
      {{end}}
      {{if eq $res.Processed 0}}
      <a href="#!" onclick="processRes({{$res.ID}})" class ="btn btn-dark">Mark as processed</a>
      {{end}}
      </div>
      <div class="float-end">
         <a href="#!" onclick="deleteRes({{$res.ID}})" class ="btn btn-danger">delete</a>
      </div>
      <div class="clearfix"></div>
    </form>
    
{{end}}

{{define "js"}}
     {{$src := index .StringMap "src"}}
<script>
function processRes(id) {
    attention.custom({
        icon: 'warning',
        msg: 'Are you sure?',
        callback: function(result) {
            console.log(result)
            if (result!==false) {
                        window.location.href = "/admin/process-reservations/{{$src}}/" 
                        + id
                        +"/do?y={{index .StringMap "year"}}&m={{index .StringMap "month"}}"
                        ;
            }
        }

    })
}

function deleteRes(id){
   attention.custom({
        icon: 'warning',
        msg: 'Are you sure?',
        callback: function(result) {
            console.log(result)
            if (result!==false) {
                        window.location.href = "/admin/delete-reservations/{{$src}}/" 
                        + id
                        +"/do?y={{index .StringMap "year"}}&m={{index .StringMap "month"}}"
                        ;
            }
        }

    })
}
</script>
{{end}}