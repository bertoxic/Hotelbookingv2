{{define "admin"}}

<!DOCTYPE html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>Administration</title>
    <!-- plugins:css -->
    <link rel="stylesheet" href="/static/admin/assets/vendors/mdi/css/materialdesignicons.min.css">
    <link rel="stylesheet" href="/static/admin/assets/vendors/css/vendor.bundle.base.css">
    <!-- endinject -->
    <!-- Plugin css for this page -->
    <!-- End plugin css for this page -->
    <!-- inject:css -->
    <!-- endinject -->
    <!-- Layout styles -->
    <link rel="stylesheet" href="/static/admin/assets/css/style.css">
    <!-- End layout styles -->
    <link rel="shortcut icon" href="/static/admin/assets/images/favicon.ico" />
    <style>
 .notie-container {
    z-index:10000;
    }

  thead { 
      vertical-align: inherit;
      border-color: red;
      background: black;
      color: white;
  }
</style>
    {{block "css" .}}
    {{end}}
  </head>

   <body>
    {{/* <div class="container-scroller">
      <div class="row p-0 m-0 proBanner" id="proBanner">
        <div class="col-md-12 p-0 m-0">
          <div class="card-body card-body-padding d-flex align-items-center justify-content-between">
            <div class="ps-lg-1">
              <div class="d-flex align-items-center justify-content-between">
               
                {{/* <a href="https://www.bootstrapdash.com/product/purple-bootstrap-admin-template/?utm_source=organic&utm_medium=banner&utm_campaign=buynow_demo" target="_blank" class="btn me-2 buy-now-btn border-0">Get Pro</a> */}}
              </div>   
            </div>
            <div class="d-flex align-items-center justify-content-between">
              <a href="https://www.bootstrapdash.com/product/purple-bootstrap-admin-template/"><i class="mdi mdi-home me-3 text-white"></i></a>
              <button id="bannerClose" class="btn border-0 p-0">
                <i class="mdi mdi-close text-white me-0"></i>
              </button>
            </div>
          </div>
        </div>
      </div> */}}
      <!-- partial:partials/_navbar.html -->
      <nav class="navbar default-layout-navbar col-lg-12 col-12 p-0 fixed-top d-flex flex-row">
        <div class="text-center navbar-brand-wrapper d-flex align-items-center justify-content-center">
          <a class="navbar-brand brand-logo" href="/admin/dashboard"><img src="/static/admin/assets/images/logo.svg" alt="logo" /></a>
          <a class="navbar-brand brand-logo-mini" href="/admin/dashboard"><img src="/static/admin/assets/images/logo-mini.svg" alt="logo" /></a>
        </div>
        <div class="navbar-menu-wrapper d-flex align-items-stretch">
          <button class="navbar-toggler navbar-toggler align-self-center" type="button" data-toggle="minimize">
            <span class="mdi mdi-menu"></span>
          </button>
          {{/* <div class="search-field d-none d-md-block">
            <form class="d-flex align-items-center h-100" action="#">
              <div class="input-group">
                <div class="input-group-prepend bg-transparent">
                  <i class="input-group-text border-0 mdi mdi-magnify"></i>
                </div>
                <input type="text" class="form-control bg-transparent border-0" placeholder="Search projects">
              </div>
            </form>
          </div> */}}
          <ul class="navbar-nav navbar-nav-right">
            <li class="nav-item nav-profile dropdown">
              <a class="nav-link dropdown-toggle" id="profileDropdown" href="#" data-bs-toggle="dropdown" aria-expanded="false">
                <div class="nav-profile-img">
                  <img src="/static/admin/assets/images/faces/face1.jpg" alt="image">
                  <span class="availability-status online"></span>
                </div>
                <div class="nav-profile-text">
                  <p class="mb-1 text-black">Bertoxic</p>
                </div>
              </a>
              <div class="dropdown-menu navbar-dropdown" aria-labelledby="profileDropdown">
                <a class="dropdown-item" href="/">
                  <i class="mdi mdi-cached me-2 text-success"></i> Public site </a>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item" href="/user-logout">
                  <i class="mdi mdi-logout me-2 text-primary"></i> Signout </a>
              </div>
            </li>
            <li class="nav-item d-none d-lg-block full-screen-link">
              <a class="nav-link">
                <i class="mdi mdi-fullscreen" id="fullscreen-button"></i>
              </a>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link count-indicator dropdown-toggle" id="messageDropdown" href="#" data-bs-toggle="dropdown" aria-expanded="false">
                <i class="mdi mdi-email-outline"></i>
                <span class="count-symbol bg-warning"></span>
              </a>
              <div class="dropdown-menu dropdown-menu-right navbar-dropdown preview-list" aria-labelledby="messageDropdown">
                <h6 class="p-3 mb-0">Messages</h6>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item preview-item">
                  <div class="preview-thumbnail">
                    <img src="/static/admin/assets/images/faces/face4.jpg" alt="image" class="profile-pic">
                  </div>
                  <div class="preview-item-content d-flex align-items-start flex-column justify-content-center">
                    <h6 class="preview-subject ellipsis mb-1 font-weight-normal">Mark send you a message</h6>
                    <p class="text-gray mb-0"> 1 Minutes ago </p>
                  </div>
                </a>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item preview-item">
                  <div class="preview-thumbnail">
                    <img src="/static/admin/assets/images/faces/face2.jpg" alt="image" class="profile-pic">
                  </div>
                  <div class="preview-item-content d-flex align-items-start flex-column justify-content-center">
                    <h6 class="preview-subject ellipsis mb-1 font-weight-normal">Cregh send you a message</h6>
                    <p class="text-gray mb-0"> 15 Minutes ago </p>
                  </div>
                </a>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item preview-item">
                  <div class="preview-thumbnail">
                    <img src="/static/admin/assets/images/faces/face3.jpg" alt="image" class="profile-pic">
                  </div>
                  <div class="preview-item-content d-flex align-items-start flex-column justify-content-center">
                    <h6 class="preview-subject ellipsis mb-1 font-weight-normal">Profile picture updated</h6>
                    <p class="text-gray mb-0"> 18 Minutes ago </p>
                  </div>
                </a>
                <div class="dropdown-divider"></div>
                <h6 class="p-3 mb-0 text-center">4 new messages</h6>
              </div>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link count-indicator dropdown-toggle" id="notificationDropdown" href="#" data-bs-toggle="dropdown">
                <i class="mdi mdi-bell-outline"></i>
                <span class="count-symbol bg-danger"></span>
              </a>
              <div class="dropdown-menu dropdown-menu-right navbar-dropdown preview-list" aria-labelledby="notificationDropdown">
                <h6 class="p-3 mb-0">Notifications</h6>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item preview-item">
                  <div class="preview-thumbnail">
                    <div class="preview-icon bg-success">
                      <i class="mdi mdi-calendar"></i>
                    </div>
                  </div>
                  <div class="preview-item-content d-flex align-items-start flex-column justify-content-center">
                    <h6 class="preview-subject font-weight-normal mb-1">Event today</h6>
                    <p class="text-gray ellipsis mb-0"> Just a reminder that you have an event today </p>
                  </div>
                </a>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item preview-item">
                  <div class="preview-thumbnail">
                    <div class="preview-icon bg-warning">
                      <i class="mdi mdi-settings"></i>
                    </div>
                  </div>
                  <div class="preview-item-content d-flex align-items-start flex-column justify-content-center">
                    <h6 class="preview-subject font-weight-normal mb-1">Settings</h6>
                    <p class="text-gray ellipsis mb-0"> Update dashboard </p>
                  </div>
                </a>
                <div class="dropdown-divider"></div>
                <a class="dropdown-item preview-item">
                  <div class="preview-thumbnail">
                    <div class="preview-icon bg-info">
                      <i class="mdi mdi-link-variant"></i>
                    </div>
                  </div>
                  <div class="preview-item-content d-flex align-items-start flex-column justify-content-center">
                    <h6 class="preview-subject font-weight-normal mb-1">Launch Admin</h6>
                    <p class="text-gray ellipsis mb-0"> New admin wow! </p>
                  </div>
                </a>
                <div class="dropdown-divider"></div>
                <h6 class="p-3 mb-0 text-center">See all notifications</h6>
              </div>
            </li>
            <li class="nav-item nav-logout d-none d-lg-block">
              <a class="nav-link" href="#">
                <i class="mdi mdi-power"></i>
              </a>
            </li>
            <li class="nav-item nav-settings d-none d-lg-block">
              <a class="nav-link" href="#">
                <i class="mdi mdi-format-line-spacing"></i>
              </a>
            </li>
          </ul>
          <button class="navbar-toggler navbar-toggler-right d-lg-none align-self-center" type="button" data-toggle="offcanvas">
            <span class="mdi mdi-menu"></span>
          </button>
        </div>
      </nav>
            <div class="container-fluid page-body-wrapper">
        <!-- partial:partials/_sidebar.html -->
        <nav class="sidebar sidebar-offcanvas" id="sidebar">
          <ul class="nav">
            <li class="nav-item nav-profile">
              <a href="#" class="nav-link">
                <div class="nav-profile-image">
                  <img src="/static/admin/assets/images/faces/face1.jpg" alt="profile">
                  <span class="login-status online"></span>
                  <!--change to offline or busy as needed-->
                </div>
                <div class="nav-profile-text d-flex flex-column">
                  <span class="font-weight-bold mb-2">Bertoxic</span>
                  <span class="text-secondary text-small">Project Manager</span>
                </div>
                <i class="mdi mdi-bookmark-check text-success nav-profile-badge"></i>
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="/admin/dashboard">
                <span class="menu-title">Dashboard</span>
                <i class="mdi mdi-home menu-icon"></i>
              </a>
            </li>
            <li class="nav-item">
              <a class="nav-link" data-bs-toggle="collapse" href="#ui-basic" aria-expanded="false" aria-controls="ui-basic">
                <span class="menu-title">Reservations</span>
                <i class="menu-arrow"></i>
                <i class="mdi mdi-crosshairs-gps menu-icon"></i>
              </a>
              <div class="collapse" id="ui-basic">
                <ul class="nav flex-column sub-menu">
                  <li class="nav-item"> <a class="nav-link" href="/admin/reservations-new">New Reservations</a></li>
                  <li class="nav-item"> <a class="nav-link" href="/admin/reservations-all">All Reservations</a></li>
                </ul>
              </div>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="/admin/reservation-calender">
                <span class="menu-title">Reservation calendar</span>
                <i class="mdi mdi-contacts menu-icon"></i>
              </a>
            </li>
             <li class="nav-item">
              <a class="nav-link" href="/">
                <span class="menu-title">make reservation</span>
                <i class="mdi mdi-format-list-bulleted menu-icon"></i>
              </a>
            </li>
           
        
            <li class="nav-item sidebar-actions">
              <span class="nav-link">
                <div class="border-bottom">
                  <h6 class="font-weight-normal mb-3">Regular</h6>
                </div>
                <button href="/" class="btn btn-block btn-lg btn-gradient-primary mt-4">Booking page</button>
              
              </span>
            </li>
          </ul>
        </nav>
         <div class="main-panel">
          <div class="content-wrapper">
            <div class="page-header">
              <h3 class="page-title">
                <span class="page-title-icon bg-gradient-primary text-white me-2">
                  <i class="mdi mdi-home"></i>
                </span> {{block "page-title" .}}
                        {{end}}
              </h3>
              <nav aria-label="breadcrumb">
              <nav aria-label="breadcrumb">
                <ul class="breadcrumb">
                  <li class="breadcrumb-item active" aria-current="page">
                    <span></span>Overview <i class="mdi mdi-alert-circle-outline icon-sm text-primary align-middle"></i>
                  </li>
                </ul>
              </nav>
               <div class="row">
              <div class="col-md-4 stretch-card grid-margin">
              </div>
              </div>
              </div>
              



       <hr>
          <div class="container-fluid row">
              {{block "content" .}}

              {{end}}
                  
            </div>
        
        
       
        

   <footer class="footer">
            <div class="container-fluid d-flex justify-content-between">
            
            </div>
          </footer>


 <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js"
            integrity="sha384-w76AqPfDkMBDXo30jS1Sgez6pr3x5MlQ1ZAGC+nuZB+EYdgRZgiwxhTBTkF7CXvN"
            crossorigin="anonymous"></script>
    <script src="https://unpkg.com/notie"></script>
    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
    <script src="/static/admin/assets/vendors/js/vendor.bundle.base.js"></script>
    <!-- endinject -->
    <!-- Plugin js for this page -->
    <script src="/static/admin/assets/vendors/chart.js/Chart.min.js"></script>
    <script src="/static/admin/assets/js/jquery.cookie.js" type="text/javascript"></script>
    <!-- End plugin js for this page -->
    <!-- inject:js -->
    <script src="/static/admin/assets/js/off-canvas.js"></script>
    <script src="/static/admin/assets/js/hoverable-collapse.js"></script>
    <script src="/static/admin/assets/js/misc.js"></script>
    <!-- endinject -->
    <!-- Custom js for this page -->
    <script src="/static/admin/assets/js/dashboard.js"></script>
    <script src="/static/admin/assets/js/todolist.js"></script>
    
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
                        {{/* preConfirm: () => {
                            return [
                                document.getElementById('start').value,
                                document.getElementById('end').value
                            ]
                        }, */}}
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
