{{template "base" .}}


{{define "content"}}
{{$res := index .Data "reservation"}}

<div class="container-fluid">
<div class="row">
<div class="col">
<h1 class="mt-4"> Reservation Summary <h1>
</div>
</div>
</div>
<section class="container-fluid">

<table class="table table-striped">
    <tr>
        <td>Full-Name</td>
        <td>{{$res.FirstName}} {{$res.LastName}}</td>
    </tr>    
    <tr>
        <td>Room:</td>
        <td>{{$res.Room.RoomName}}</td>
    </tr>  
     <tr>
        <td>Email:</td>
        <td>{{$res.Email}}</td>
    </tr>   
     <tr>
        <td>Phone:</td>
        <td>{{$res.Phone}}</td>
    </tr>   
     <tr>
        <td>Arrival Date:</td>
        <td>{{index .StringMap "start_date"}}</td>
    </tr>
     <tr>
        <td>Depature Date:</td>
        <td>{{index .StringMap "end_date"}}</td>
    </tr>
</table>
</section>

{{end}}