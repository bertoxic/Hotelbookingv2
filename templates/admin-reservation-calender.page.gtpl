{{template "admin" .}}

{{define "page-title"}}
    Reservation Calender
{{end}}


{{define "content"}}
{{$now:= index .Data "now"}}
{{$rooms := index .Data "rooms"}} 
{{$dim:= index .IntMap "days_in_month"}}
{{$currMonth:= index .StringMap "this_month"}}
{{$currYear:= index .StringMap "this_month_year"}}
 
<div>
<h3 class="text-center">
    {{formatDate $now "January"}} {{formatDate $now "2006"}}
</h3>
</div> 

<div >
<div class="float-start">
<a class="btn btn-sm btn-outline-danger"
href="/admin/reservation-calender?y={{index .StringMap "last_month_year"}}&m={{index .StringMap "last_month"}}" >&lt&lt</a>
</div>

<div class="float-end">
<a class="btn btn-sm btn-outline-info"
href="/admin/reservation-calender?y={{index .StringMap "next_month_year"}}&m={{index .StringMap "next_month"}}" >&gt&gt</a>
</div>
</div>
<div class="clearfix"></div>
<div>
<form method="post" action="/admin/reservation-calender">
 <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
 <input type="hidden" name="m" value="{{index .StringMap "this_month"}}">
 <input type="hidden" name="y" value="{{index .StringMap "this_month_year"}}">
    {{range $rooms}}
    {{$roomID:= .ID}}
    {{$blocks:= index $.Data (printf "block_map_%d" .ID)}}
    {{$reservations:= index $.Data (printf "reservation_map_%d" .ID)}}
    <h4 class ="mt-4">{{.RoomName}}</h4>
    <div>

 
       <div class="table-responsive">
         <table class= "table  table-bordered table-sm">
            <tr class = "table-info">
                {{range $index := iterate $dim}}
                <td class="text-center">
                {{add $index 1}}</td>
             {{end}}
            </tr>
             <tr>
             {{range $index :=iterate $dim}}
                <td class="text-center">
                {{if gt (index $reservations (printf "%s-%s-%d" $currYear $currMonth (add $index 1))) 0}}
                <a href="/admin/reservation/cal/{{index $reservations (printf "%s-%s-%d" $currYear $currMonth (add  $index 1))}}/show?y={{$currYear}}&m={{$currMonth}}">
                <span class="text-danger">R</span>
               
                {{else}}
                 <input 
                {{ if gt (index $blocks (printf "%s-%s-%d" $currYear $currMonth (add $index 1))) 0}}
                    checked 
                    name ="remove_block_{{$roomID}}_{{printf "%s-%s-%d" $currYear $currMonth (add $index 1)}}"
                    value= "{{index $blocks (printf "%s-%s-%d" $currYear $currMonth (add $index 1))}}"
                {{else}}
                    name = "add_block_{{$roomID}}_{{printf "%s-%s-%d" $currYear $currMonth (add $index 1)}}"
                {{end}}
                type="checkbox">
                
                {{end}}
                </td>
                 {{end}}
                </tr>
       </table>
       </div>
       {{end}}
       </div>
  {{end}}

  <input type="submit" class="btn btn-primary" value="Save Changes">
</form>
</div>



