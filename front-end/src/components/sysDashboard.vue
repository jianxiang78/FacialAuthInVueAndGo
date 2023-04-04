<template>
    <div id="body" class="nav-md">
        <div class="container body">
            <div class="main_container">
                <div class="col-md-3 left_col">
                    <div class="left_col scroll-view">
                        <div class="navbar nav_title" style="border: 0;">
                            <span class="site_title">
                                <i class="fa fa-user-secret"></i>
                                <span> Demo</span>
                            </span>
                        </div>


                        <!-- menu profile quick info -->
                        <div class="profile clearfix">
                            <div class="profile_info">
                                <h2>Welcome, Guests</h2>
                            </div>
                        </div>
                        <!-- /menu profile quick info -->

                        <br />

                        <!-- sidebar menu -->
                        <div id="sidebar-menu" class="main_menu_side hidden-print main_menu">
                            <div class="menu_section">
                                <h3>Dashboard</h3>
                                <ul class="nav side-menu">
                                    <li>
                                        <a><i class="fa fa-home"></i> Home </a>
                                    </li>
                                </ul>
                            </div>

                        </div>
                        <!-- /sidebar menu -->

                    </div>
                </div>

                <!-- top navigation -->
                <div class="top_nav">
                    <div class="nav_menu">
                        <nav>
                            <ul class="nav navbar-nav navbar-right">
                                <li>
                                    <router-link class="user-profile dropdown-toggle" :to="'/login'"><i class="fa fa-sign-out"></i> Log Out</router-link>
                                </li>
                                <li>
                                    <router-link class="user-profile dropdown-toggle" :to="'#'">Facial Auth In VueJs And Golang</router-link>
                                </li>
                            </ul>
                        </nav>
                    </div>
                </div>
                <!-- /top navigation -->

                <!-- page content -->
                <div class="right_col" role="main">
                    <!-- top tiles -->
                    <div class="row tile_count">
                        <div class="col-md-2 col-sm-4 col-xs-6 tile_stats_count">
                            <span class="count_top"><i class="fa fa-user"></i> Total Users</span>
                            <div class="count">1000</div>
                            <span class="count_bottom"><i class="green">10% </i> From last Week</span>
                        </div>
                        <div class="col-md-2 col-sm-4 col-xs-6 tile_stats_count">
                            <span class="count_top"><i class="fa fa-clock-o"></i> Average Time</span>
                            <div class="count">1000</div>
                            <span class="count_bottom"><i class="green"><i class="fa fa-sort-asc"></i>10% </i> From last Week</span>
                        </div>
                        <div class="col-md-2 col-sm-4 col-xs-6 tile_stats_count">
                            <span class="count_top"><i class="fa fa-user"></i> Total Males</span>
                            <div class="count green">1000</div>
                            <span class="count_bottom"><i class="green"><i class="fa fa-sort-asc"></i>10% </i> From last Week</span>
                        </div>
                        <div class="col-md-2 col-sm-4 col-xs-6 tile_stats_count">
                            <span class="count_top"><i class="fa fa-user"></i> Total Females</span>
                            <div class="count">1000</div>
                            <span class="count_bottom"><i class="red"><i class="fa fa-sort-desc"></i>10% </i> From last Week</span>
                        </div>
                        <div class="col-md-2 col-sm-4 col-xs-6 tile_stats_count">
                            <span class="count_top"><i class="fa fa-user"></i> Total Collections</span>
                            <div class="count">1000</div>
                            <span class="count_bottom"><i class="green"><i class="fa fa-sort-asc"></i>10% </i> From last Week</span>
                        </div>
                        <div class="col-md-2 col-sm-4 col-xs-6 tile_stats_count">
                            <span class="count_top"><i class="fa fa-user"></i> Total Connections</span>
                            <div class="count">1000</div>
                            <span class="count_bottom"><i class="green"><i class="fa fa-sort-asc"></i>10% </i> From last Week</span>
                        </div>
                    </div>
                    <!-- /top tiles -->

                    <div class="row">

                        <div class="col-md-4 col-sm-4 col-xs-12">
                            <div class="x_panel tile fixed_height_320 overflow_hidden">
                                <div class="x_title">
                                    <h2>Login Name</h2>
                                    <ul class="nav navbar-right panel_toolbox">
                                        <li>
                                            <a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                        </li>
                                        <li class="dropdown">
                                            <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>
                                            <ul class="dropdown-menu" role="menu">
                                                <li><a href="#">Settings 1</a>
                                                </li>
                                                <li><a href="#">Settings 2</a>
                                                </li>
                                            </ul>
                                        </li>
                                        <li><a class="close-link"><i class="fa fa-close"></i></a>
                                        </li>
                                    </ul>
                                    <div class="clearfix"></div>
                                </div>
                                <div class="x_content">
                                    <h4>admin</h4>
                                    <div class="dashboard-widget-content">
                                        <ul class="quick-list">
                                            <li>
                                                <i class="fa fa-calendar-o"></i>Default Setting
                                            </li>
                                        </ul>

                                        <div class="sidebar-widget">
                                            <h4>Password</h4>
                                            <div width="150" height="80" id="chart_gauge_01"
                                                 style="width: 160px; height: 100px;"
                                            >
                                                admin
                                            </div>

                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class="col-md-4 col-sm-4 col-xs-12">
                            <div class="x_panel tile fixed_height_320">
                                <div class="x_title">
                                    <h2>Facial Auth</h2>
                                    <ul class="nav navbar-right panel_toolbox">
                                        <li>
                                            <a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                        </li>
                                        <li class="dropdown">
                                            <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>
                                            <ul class="dropdown-menu" role="menu">
                                                <li><a href="#">Settings 1</a>
                                                </li>
                                                <li><a href="#">Settings 2</a>
                                                </li>
                                            </ul>
                                        </li>
                                        <li><a class="close-link"><i class="fa fa-close"></i></a>
                                        </li>
                                    </ul>
                                    <div class="clearfix"></div>
                                </div>
                                <div class="x_content">
                                    <h4 id="myFacialId">Unbound</h4>
                                    <button id="FacialIdReg" type="button" class="btn btn-primary" style="width: 100%" @click="faceReg">Face Recognition Registration</button>
                                </div>
                            </div>
                        </div>

                        <div class="col-md-4 col-sm-4 col-xs-12">
                            <div class="x_panel tile fixed_height_320">
                                <div class="x_title">
                                    <h2>User Name</h2>
                                    <ul class="nav navbar-right panel_toolbox">
                                        <li>
                                            <a class="collapse-link"><i class="fa fa-chevron-up"></i></a>
                                        </li>
                                        <li class="dropdown">
                                            <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><i class="fa fa-wrench"></i></a>
                                            <ul class="dropdown-menu" role="menu">
                                                <li><a href="#">Settings 1</a>
                                                </li>
                                                <li><a href="#">Settings 2</a>
                                                </li>
                                            </ul>
                                        </li>
                                        <li><a class="close-link"><i class="fa fa-close"></i></a>
                                        </li>
                                    </ul>
                                    <div class="clearfix"></div>
                                </div>
                                <div class="x_content">
                                    <h4>Admin</h4>
                                </div>
                            </div>
                        </div>

                    </div>


                </div>
                <!-- /page content -->

                <!-- footer content -->
                <footer>
                    <div class="pull-right">
                        Copyright &copy; Auto17.com 2023
                    </div>
                </footer>
                <!-- /footer content -->
            </div>
        </div>
    </div>

</template>

<script>
    export default {
        name: 'sysDashboard',
        activated() {
            this.AjaxPost("dashboard","",initData);
            function initData(result) {
                if(result.Code===1){
                    let fid=result.Result.FacialId
                    if(fid.length>4){
                        document.getElementById("myFacialId").innerHTML=fid
                        document.getElementById("FacialIdReg").disabled=true;
                    }
                }
            }
        },
        methods: {
            faceReg: function() {
                const that=this
                window.myFaceIO.enroll({
                    "locale": "auto"
                }).then(userInfo => {
                    addFacialAuth(userInfo);
                }).catch(errCode => {
                    console.log(errCode);
                    if(errCode===6){
                        window.myFaceIO.authenticate({
                            "locale": "auto"
                        }).then(userInfo => {
                            addFacialAuth(userInfo);
                        }).catch(errCode => {
                            console.log(errCode);
                        })
                    }
                })
                function addFacialAuth(userInfo) {
                    let obj = {};
                    obj.FacialId = userInfo.facialId;
                    that.AjaxPost("userEdit",obj,addFacialAuthEnd);
                }
                function addFacialAuthEnd(result) {
                    if(result.Code===1){
                        document.getElementById("myFacialId").innerHTML=result.Result.FacialId
                        document.getElementById("FacialIdReg").disabled=true;
                    }else {
                        alert("add Facial Auth fail")
                    }
                }
            }
        }
    }
</script>

<style scoped src="../assets/build/css/custom.css"></style>