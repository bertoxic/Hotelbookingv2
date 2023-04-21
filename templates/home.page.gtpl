{{template "base" .}}



<style>
header{
    background-image: linear-gradient(
          rgba(0, 0, 10, 0.5), 
          rgba(0, 0,10, 0.4)
        ), url("/static/resources/city/image8.jpg");
    height: 800px;
    background-size:cover;
    background-position:center center;
    background-blend-mode: darken;

}

#Nav{
  background-color:none !important;
  color:white;
}
</style>
{{define "nav"}}
<h1 style="color:azure">Tulith Hotels</h1>
<section class="container home-text mt-5" >
<div class="paragraph-effect" style="--d: 0.5s; color:azure">
<h3> Tulith passed, vaunted beauty</div>
</div>
</div>
<div class="paragraph-effect"style="--d: 1.0s; color:azure">
    <span >Grace a peace heavenly me nor the.ce heavenly me nor the. Shameless care Shameless care into in did. Loved dwell shameless so was are</span>
</div>
<div class="paragraph-effect" style="--d:1.5s; color:azure">
    <span >shameless so was are lines, uncouth the smile childe he relief ah or in oft, virtues the a him with ah.</span>
</div>

  <a class="btn btn-outline-info mt-4" href="/search-availability">Make Reservation Now</a>

</section>
{{end}}
{{define "content"}}

{{/* <div class="hero-page"> 
<div class="container"> 
<h2>
kilopele
</h2>
</div>
</div> */}}


<section class="container home-text" >
<div class="paragraph-effect" style="--d: 0.5s">
<h3> Tulith's best vaunted beauty</div>
</div>
</div>
<div class="paragraph-effect"style="--d: 1.0s">
    <span >Grace a peace heavenly me nor the. Shameless care into in did. Loved dwell shameless so was are</span>
</div>
<div class="paragraph-effect" style="--d:1.5s">
    <span >shameless so was are lines, uncouth the smile childe he relief ah or in oft, virtues the a him with ah.</span>
</div>

  <a class="btn btn-outline-success mt-4" href="/search-availability">Make Reservation Now</a>

</section>
{{end}}

