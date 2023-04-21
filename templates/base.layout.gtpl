{{define "base"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
     <link rel="stylesheet" href="/static/css/style.css" type="text/css">
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet"
            integrity="sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD" crossorigin="anonymous">
        <link rel="stylesheet"
            href="https://cdn.jsdelivr.net/npm/vanillajs-datepicker@1.3.1/dist/css/datepicker.min.css">
        <link rel="stylesheet" type="text/css" href="https://unpkg.com/notie/dist/notie.min.css">
  

{{/* {{block "css" .}}

{{end}} */}}

</head>
 <style>
        .swal2-actions {
     position: relative;
    z-index: 0 !important;
    background-color: #e0d6d604;
}
    </style>

<body>
<header>
  {{/* <nav class="navbar navbar-expand-lg bg-body-secondary"> */}}
  <nav id ="Nav" class="navbar navbar-expand-lg">
            <div class="container-fluid">
                <a class="navbar-brand" href="#">Tulith</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                    data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false"
                    aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarNavDropdown">
                    <ul class="navbar-nav">
                        <li class="nav-item">
                            <a class="nav-link active" aria-current="page" href="/">Home</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="/make-reservation">Blog</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="/about">About</a>
                        </li>
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown"
                                aria-expanded="false">
                                Rooms
                            </a>
                            <ul class="dropdown-menu">
                                <li><a class="dropdown-item" href="/general-quaters">Generals</a></li>
                                <li><a class="dropdown-item" href="/majors-suite">Majors</a></li>
                                <li><a class="dropdown-item" href="/search-availability">Reservations</a></li>
                            </ul>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="/contact" target="_blank">Contact</a>
                        </li>  
                       
                        <li class="nav-item">
                         {{if eq .IsAuthenticated 1}}
                                 <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown"
                                aria-expanded="false">
                                Admin
                            </a>
                            <ul class="dropdown-menu">
                                <li><a class="dropdown-item" href="/admin/dashboard">Dashboard</a></li>
                                <li><a class="dropdown-item" href="/user-logout">logout</a></li>
                                <li><a class="dropdown-item" href="/search-availability">Reservations</a></li>
                            </ul>
                        </li>
                        {{else}}
                            <a class="nav-link" href="/user-login" target="_blank">Login</a>
                             {{end}}
                        </li>
                       
                    </ul>
                </div>
            </div>
            </div>
        </nav>
        {{block "nav" .}}

        {{end}}
        </header>




{{block "content" .}}

{{end}}


		<footer class="footer-distributed mt-5">

			<div class="footer-right">

				<a href="#"><i class="fa fa-facebook"></i></a>
				<a href="#"><i class="fa fa-twitter"></i></a>
				<a href="#"><i class="fa fa-linkedin"></i></a>
				<a href="#"><i class="fa fa-github"></i></a>

			</div>

			<div class="footer-left">

				<p class="footer-links">
					<a class="link-1" href="#">Home</a>

					<a href="#">Blog</a>

					<a href="/">Pricing</a>

					<a href="/about">About</a>

					<a href="/">Faq</a>

					<a href="#">Contact</a>
				</p>

				<p>Tulith Hotels &copy; 2019</p>
			</div>

		</footer>

        <script src="https://cdn.jsdelivr.net/npm/vanillajs-datepicker@1.3.1/dist/js/datepicker-full.min.js"></script>
        <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js"
            integrity="sha384-w76AqPfDkMBDXo30jS1Sgez6pr3x5MlQ1ZAGC+nuZB+EYdgRZgiwxhTBTkF7CXvN"
            crossorigin="anonymous"></script>
        <script src="https://unpkg.com/notie"></script>
        <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>


        {{block "js" .}}


        {{end}}
        <script>
      
let attention = Prompt()
            {{/* const elem = document.getElementById('Reservation-date');
            const rangepicker = new DateRangePicker(elem, {
                // ...options
            }); */}}

            function notify(msg, msgType) {
                // notie.alert({
                //     type: "success",
                //     text: msg,
                // })
                Swal.fire({
                    title: msgType,
                    text: msg,
                    icon: msgType,
                    confirmButtonText: 'Cancel'
                })
            }

            function notifyModal(title, text, icon, confirmButtonText) {


                Swal.fire({
                    title: title,
                    text: text,
                    icon: icon,
                    confirmButtonText: confirmButtonText
                })
            }

              {{with .Error}}
        notify("{{.}}","error")
        
        {{end}}    
        
           {{with .Flash}}
        attention.toast({msg:"{{.}}",icon:"success"})
        
        {{end}}

            function Prompt() {
                let toast = function (c) {
                    const {
                        msg = "",
                        position = "top-end",
                        icon = "success",

                    } = c;
                    const Toast = Swal.mixin({
                        toast: true,
                        title: msg,
                        icon: icon,
                        position: position,
                        showConfirmButton: false,
                        timer: 3000,
                        timerProgressBar: true,
                        didOpen: (toast) => {
                            toast.addEventListener('mouseenter', Swal.stopTimer)
                            toast.addEventListener('mouseleave', Swal.resumeTimer)
                        }
                    })

                    Toast.fire({

                    })
                }

                let success = function (c) {
                    const {
                        msg = "success",
                        title = ""

                    } = c
                    Swal.fire({
                        title: msg,
                        text: '',
                        icon: 'success',
                    })
                    Toast.fire({})
                }

                let Error = function (c) {
                    const {
                        msg = "",
                        title = ""

                    } = c
                    Swal.fire({
                        title: msg,
                        text: '',
                        icon: 'error',
                    })
                    Toast.fire({})

                }
                async function custom(c){
                    const {
                        msg = "",
                        title = "",
                        icon = "",
                        showConfirmButton = true,
                    } = c;
                    const { value: result } = await Swal.fire({
                        icon:icon,
                        title: title,
                        html:msg,
                        focusConfirm: false,
                        backdrop:true,
                        showCancelButton: true,
                        showConfirmButton:showConfirmButton,
                        willOpen:()=>{
                            if (c.willOpen !== undefined) {
                                c.willOpen();
                            }
                        } ,
                        preConfirm: () => {
                            return [
                                document.getElementById('start').value,
                                document.getElementById('end').value
                            ]
                        },
                        didOpen:()=>{
                           if(c.didOpen !==undefined){
                            c.didOpen();
                           }
                        }
                    })
                            
                        if(result){
                            if (result.dismiss !== Swal.DismissReason.cancel){
                                if (result.value !==""){
                                    if (c.callback !== undefined){
                                        c.callback(result);
                                        console.log("xxxxxx",c.speedy);
                                        
                                    }
                                }else {
                                    c.callback(false);
                                }
                            }else {
                                c.callback(false);
                            }
                        }
                    {{/* if (result) {
                        Swal.fire(JSON.stringify(result))
                    } */}}

                }

                return {
                    toast: toast,
                    success: success,
                    error: Error,
                    custom:custom,
                }
            }


        </script>

    </body>

</html>

{{end}}