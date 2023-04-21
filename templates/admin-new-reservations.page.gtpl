{{template "admin" .}}

{{define "css"}}
<link href="https://cdn.jsdelivr.net/npm/simple-datatables@latest/dist/style.css" rel="stylesheet" type="text/css">
<style>

.content-wrapper {
    background-color:white;
}
.footer {
    background-color:white;
    color: color(dark);
    border-top: 1px solid #e7dee9;
    padding: 30px 1rem;
     font-size: calc(0.875rem - 0.05rem);
    font-family: "ubuntu-regular", sans-serif;
}
.datatable-table {
    max-width: 100%;
    width: 100%;  
    border-spacing: 2px !important;
}

</style>
{{end}}
{{define "page-title"}}
    New Reservations
{{end}}


{{define "content"}}
   {{$res := index .Data "reservations"}}
    <table class="table table-striped table-hover table-color" id="new-res">
    <thead>
    <tr class="table-dark">
    <th>ID</th>
    <th>Name</th>
    <th>Room</th>
    <th>Arrival</th>
    <th>Depature</th>
    </tr>
    </thead>
    <tbody>
    {{range $res}}
        <tr>
        <td class="table-secondary">{{.ID}}</td>
        <td class="table-primary">
        <a href="/admin/reservation/new/{{.ID}}/show">
        {{.FirstName}} {{.LastName}}</a>
        </td>
        <td class="table-success">{{.Room.RoomName}}</td>
        <td class="table-danger">{{humanDate .StartDate}}</td>
        <td class="table-info">{{humanDate .EndDate}}</td>
          <tr>  
    {{end}}
    </tbody>
    </table>
{{end}}

{{define "js"}}
<script src="https://cdn.jsdelivr.net/npm/simple-datatables@latest" type="text/javascript"></script>
<script>
document.addEventListener("DOMContentLoaded", function() {
const dataTable = new simpleDatatables.DataTable("#new-res", {
    select: 3, sort: "desc"
})
})

</script>
{{end}}

