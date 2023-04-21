{{template "base" .}}

{{define "css"}}
    <style>
        .swal2-actions {
     position: relative;
    z-index: 0 !important;
    background-color: #e0d6d604;
}
    </style>
{{end}}

{{define "content"}}
<section class="container">
            <h1>Check availability</h1>
            <div>
              <img src="/static/resources/city/image10.jpg" class="gallery__img" alt="Image 3">
            </div>
            <form method="post" action="/search-availability" novalidate class="needs-validation mt-3 col-6">
            <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
            <div class="row">
            
                <div class="col">
                    <div class="row" id="Reservation-date">
                        <div class="col">
                        <label for="start"> Arrival:</label>
                            <input type="text" class="form-control" placeholder="arrival date" name="start">
                        </div>
                        <div class="col">
                         <label for="end"> Depature:</label>
                            <input type="text" class="form-control" placeholder="depature-date" name="end">
                        </div>
                    </div>

                </div>


            </div>
          
            <button id="Reservation-date"  type="submit" method="post" class="btn btn-primary mt-3"> Search Availability</button>
            </form>
        </section>

        <section class="container  mt-3">
            <p>Welcome toScorching mood they flatterers old none that felt himnot, vile was but soils passed, vaunted of
                oh if nor. Grace a peace heavenly me nor the. Shameless care into in did. Loved dwell shameless so was
                are lines, uncouth the smile childe he relief ah or in oft, virtues the a him . </p>
           
        </section>




{{end}}

{{define "js"}}
<script>
 

                 const elem = document.getElementById('Reservation-date');
            const rangepicker = new DateRangePicker(elem, {
               
                      format: "yyyy-mm-dd",
                      minDate: new Date(),
      
                // ...options
            });
</script>

{{end}}