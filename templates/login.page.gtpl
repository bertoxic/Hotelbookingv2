{{template "base" .}}

{{define "content"}}

<p class="container"><strong>LOGIN PAGE</strong> <br>
<form method="POST" action="/user-login" class="needs-validation container" novalidate>
<input type="text" name="csrf_token" value="{{.CSRFToken}}">
<div class="row">
   <div class=" form-group mt-3">
        <label for="email">Email</label>
        {{with .Form.Errors.Get "email"}}
        <label class="text-danger">{{.}}</label>
        {{end}}
        <input name="email" value="" id="email" type="text" placeholder="Enter your email" class="form-control" required
          autocomplete="off" type="text" >
      </div> 
      </div> 
      <div class="col form-group ">
        <label for="password">Password</label>
        {{with .Form.Errors.Get "password"}}
        <label class="text-danger">{{.}}</label>
        {{end}}
        <input name="password" value="" id="password" type="text" placeholder="Enter your password" class="form-control" required
          autocomplete="off" type="text" >
          <hr>
          <input type="submit" class="btn btn-primary" value="submit"> 
      </div>
</div>

      </form>

{{end}}