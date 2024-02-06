{{template "base" .}}

{{define "content"}}

<h2> This is the about page you should know </h2>
{{ if ne (index .StringMap "remoteIp") ""}}

<h3> Your remote address is {{index .StringMap "remoteIp"}}</h3>

{{else}}
        i dont know your assignment go to home page. Please
{{end}}
<section class="container home-text" >
<p> Welcome toScorching mood they flatterers old none that felt himnot, vile was but soils passed, vaunted of oh if nor. Grace a peace heavenly me nor the. Shameless care into in did. Loved dwell shameless so was are lines, uncouth the smile childe he relief ah or in oft, virtues the a him with ah,. </p>
<button class="btn btn-outline-success">
  Click here
</button>
</section>

{{end}}