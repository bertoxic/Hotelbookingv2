
{{template "base" .}}

{{define "content"}}

<h2> PLease Choose a room</h2>
{{ if ne (index .StringMap "remoteIp") ""}}

<h3> Your remote address is {{index .StringMap "remoteIp"}}</h3>

{{else}}
        i dont know your Location go to home page. <hr>
{{end}}
 
{{$rooms := index .Data "rooms"}}
<ul>
{{range $rooms}}
<li> <a href="/choose-room/{{.ID}}"> {{.RoomName}} </a>
    </li>
    <hr>
{{end}}
</ul>
<section class="container home-text" >
<p style="font-size:14px"> Welcome toScorching mood they flatterers old none that felt himnot, vile was but soils passed, vaunted of oh if nor. Grace a peace heavenly me nor the. Shameless care into in did. Loved dwell shameless so was are lines, uncouth the smile childe he relief ah or in oft, virtues the a him with ah,. </p>
<button class="btn btn-outline-success">
  Click here
</button>
</section>

{{end}}