<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <meta name="description" content="">
  <meta name="author" content="">
  <title>LEP - Linux Easy Profiling</title>

  <link href="/static/vendors/bootstrap/css/bootstrap.min.css" rel="stylesheet">
  <link href="/static/vendors/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">
  <link href="/static/vendors/sb-admin/css/sb-admin.css" rel="stylesheet">
  <link href="/static/vendors/datatables/media/css/jquery.dataTables.min.css " rel="stylesheet">
  <link href="/static/vendors/c3/c3.min.css" rel="stylesheet">
  <link href="/static/vendors/whhg-font/css/whhg.css" rel="stylesheet">
  <link href="/static/vendors/flameGraph/d4.flameGraph.css" rel="stylesheet">

  <link href="/static/css/lepv.css" rel="stylesheet">

</head>

<body class="fixed-nav sticky-footer bg-dark" id="page-top">
  <!-- Navigation-->
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top" id="mainNav">
    <a class="navbar-brand" href="/">LEP - Linux Easy Profiling</a>
    <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarResponsive">
      <ul class="navbar-nav navbar-sidenav" id="sideBarAccordion">

        <!--<li class="nav-item" data-toggle="tooltip" data-placement="right" title=" Summary">-->
          <!--<a class="nav-link nav-link-collapse collapsed" data-toggle="collapse" href="#collapseSummaryMenu" data-parent="#sideBarAccordion">-->
            <!--<i class="fa fa-fw fa-dashboard"></i>-->
            <!--<span class="nav-link-text">Summary</span>-->
          <!--</a>-->
          <!--<ul class="sidenav-second-level collapse" id="collapseSummaryMenu">-->
            <!--<li>-->
              <!--<a href="#">CPU</a>-->
            <!--</li>-->
            <!--<li>-->
              <!--<a href="#">MEMORY</a>-->
            <!--</li>-->
            <!--<li>-->
              <!--<a href="#">IO</a>-->
            <!--</li>-->
          <!--</ul>-->
        <!--</li>-->

            <!-- Client address input -->
            <form id="client-group" class="form-inline my-2 my-lg-0 mr-lg-2">
              <div class="input-group">

                <input id="txt_server_to_watch" name="txt_server_to_watch" type="text" class="form-control" placeholder="255.255.255.255" value="127.0.0.1"/>
                <span class="input-group-btn">
                  <!--<input type="submit" value="Play" class="btn btn-primary"/>
                      -->
                <button class="btn btn-primary" type="button" onclick="onAddDevices()">
                  <i class="fa fa-play"></i>
                </button>
                </span>
              </div>
            </form>
            <!-- Client address input -->

<!-- Monitor start -->
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title=" Devices">
          <a class="nav-link nav-link-collapse collapsed" data-toggle="collapse" href="#collapseDevicesMenu" data-parent="#sideBarAccordion">
            <i class="fa fa-fw fa-dashboard"></i>
            <span class="nav-link-text">Devices</span>
          </a>


          <ul class="sidenav-second-level collapse" id="collapseDevicesMenu">
            <!--
            <form id="client-group" action="/addclient" method="post" class="form-inline my-2 my-lg-0 mr-lg-2">
             -->

          </ul>
        </li>
<!-- Monitor end-->

        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="CPU">
          <a class="nav-link nav-link-collapse collapsed" data-toggle="collapse" href="#collapseCpuMenu" data-parent="#sideBarAccordion">
            <i class="icon-cpu-processor"></i>
            <span class="nav-link-text">CPU</span>
          </a>
          <ul class="sidenav-second-level collapse" id="collapseCpuMenu">
            <li>
              <a href="#container-div-cpu-stat-donut">Stat: Overall</a>
            </li>
            <li>
              <a href="#container-div-cpu-stat-idle">Stat: Idle</a>
            </li>
            <li>
              <a href="#container-div-cpu-stat-user-group">Stat: User Group</a>
            </li>
            <li>
              <a href="#container-div-cpu-stat-irqGroup">Stat: Irq Group</a>
            </li>
            <li>
              <a href="#container-div-cpu-stat-irq">Stat: SoftIRQ</a>
            </li>
            <li>
              <a href="#container-div-cpu-avgload">Average Load</a>
            </li>
            <li>
              <a href="#container-div-cpu-top-table">TOP</a>
            </li>
          </ul>
        </li>



        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Memory">
          <a class="nav-link nav-link-collapse collapsed" data-toggle="collapse" href="#collapseComponents" data-parent="#sideBarAccordion">
            <i class="icon-ram"></i>
            <span class="nav-link-text">Memory</span>
          </a>
          <ul class="sidenav-second-level collapse" id="collapseComponents">
            <li>
              <a href="#container-div-memory-chart">Stat</a>
            </li>
            <li>
              <a href="#container-memory-stat-table">Memory Consumption</a>
            </li>
            <li>
                <a href="#container-div-memory-free-pss-stat">PSS Ratio</a>
            </li>
          </ul>
        </li>


        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="IO">
          <a class="nav-link nav-link-collapse collapsed" data-toggle="collapse" href="#collapseIOMenu" data-parent="#sideBarAccordion">
            <i class="icon-syncalt"></i>
            <span class="nav-link-text">IO</span>
          </a>
          <ul class="sidenav-second-level collapse" id="collapseIOMenu">
              <li>
                <a href="#container-div-IOCharts">Stat</a>
              </li>
              <li>
                <a href="#container-io-top-table">Top</a>
              </li>
              <li>
                <a href="#container-jnet-top-table">Net</a>
              </li>
          </ul>
        </li>


        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Perf">
          <a class="nav-link nav-link-collapse collapsed" data-toggle="collapse" href="#collapsePerfMenu" data-parent="#sideBarAccordion">
            <i class="icon-rawaccesslogs"></i>
            <span class="nav-link-text">Perf</span>
          </a>
          <ul class="sidenav-second-level collapse" id="collapsePerfMenu">
            <li>
              <a href="#container-perf-cpu-table">Perf CPU</a>
            </li>
              <li>
              <a href="#container-perf-flame-graph">Flame Graph</a>
            </li>
          </ul>
        </li>


      </ul>
      <ul class="navbar-nav sidenav-toggler">
        <li class="nav-item">
          <a class="nav-link text-center" id="sidenavToggler">
            <i class="fa fa-fw fa-angle-left"></i>
          </a>
        </li>
      </ul>
      <ul class="navbar-nav ml-auto">
        <li class="nav-item">
            <!--
          <form class="form-inline my-2 my-lg-0 mr-lg-2">
            <div class="input-group">
              <input id='txt_server_to_watch' class="form-control" type="text" placeholder="www.rmlink.cn" value="www.rmlink.cn">
              <span class="input-group-btn">
                <button class="btn btn-primary" type="button" onclick="startWatching()">
                  <i class="fa fa-play"></i>
                </button>
              </span>
            </div>
          </form>
          -->
        </li>
      </ul>
    </div>
  </nav>
  <div class="content-wrapper">
    <div class="container-fluid">


      <!--<div class="col-md-12">-->
        <!--<div id='Summary' class="card mb-3">-->
          <!--<div class="card-header">-->
              <!--<i class="fa fa-fw fa-dashboard"></i> {{ .languages.SystemSummary }}-->
          <!--</div>-->
          <!--<div class="card-body">-->
              <!---->
              <!---->
              <!--<div class="row">-->
                  <!--<div class="col-md-3">-->
                      <!--<div class="card-body">-->
                          <!--<div id="div-cpu-gauge"></div>-->

                          <!--<div class="panel-footer text-center">-->
                              <!--{% if config == 'debug' %}-->
                              <!--处理器 - [GetCmdMpstat]-->
                              <!--{% else %}-->
                              <!--处理器-->
                              <!--{% endif %}-->
                          <!--</div>-->
                      <!--</div>-->
                  <!--</div>-->
                  <!--<div class="col-md-3">-->
                      <!--<div class="card-body">-->
                          <!--<div id="div-memory-gauge"></div>-->
                          <!--<div class="panel-footer text-center">-->
                              <!--{% if config == 'debug' %}-->
                              <!--内存 - [GetProcMeminfo]-->
                              <!--{% else %}-->
                              <!--内存-->
                              <!--{% endif %}-->
                          <!--</div>-->
                      <!--</div>-->
                  <!--</div>-->
                  <!--<div class="col-md-3">-->
                      <!--<div class="card-body">-->
                          <!--<div id="div-io-gauge"></div>-->
                          <!--<div class="panel-footer text-center">-->
                              <!--{% if config == 'debug' %}-->
                              <!--磁盘 - [GetCmdIostat]-->
                              <!--{% else %}-->
                              <!--磁盘-->
                              <!--{% endif %}-->
                          <!--</div>-->
                      <!--</div>-->
                  <!--</div>-->
              <!--</div>-->

              <!--<div class="row">-->
                  <!--<div class="col-md-3">-->
                      <!--{% if config == 'debug' %}-->
                      <!--<div>[GetProcCpuinfo]</div>-->
                      <!--{% endif %}-->
                      <!--<div id="div-cpu-summary" class="panel panel-green">-->
                      <!--</div>-->
                  <!--</div>-->

                  <!--<div class="col-md-3">-->
                      <!--{% if config == 'debug' %}-->
                      <!--<div>[GetProcMeminfo]</div>-->
                      <!--{% endif %}-->
                      <!--<div id="div-memory-summary" class="panel panel-green">-->
                      <!--</div>-->
                  <!--</div>-->

                  <!--<div class="col-md-3">-->
                      <!--{% if config == 'debug' %}-->
                      <!--<div>[GetCmdDf]</div>-->
                      <!--{% endif %}-->
                      <!--<div id="div-io-summary" class="panel panel-green">-->
                      <!--</div>-->
                  <!--</div>-->
              <!--</div>-->
          <!--</div>-->
        <!--</div>-->
      <!--</div>-->



      <!--cpu-stat-donut-->
      <div class="col-md-6">
        <div id="container-div-cpu-stat-donut" class="card mb-3">
          <div class="card-header"><i class="icon-cpu-processor"></i> CPU Stat: Overall</div>
          <div class="card-body"><div class="chart-panel"></div></div>
          <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
        </div>
      </div>


     <!--cpu-stat-idle-->
      <div class="col-md-12">
        <div id="container-div-cpu-stat-idle" class="card mb-3">
          <div class="card-header"><i class="icon-cpu-processor"></i> CPU Stat: Idle</div>
          <div class="card-body"><div class="chart-panel"></div></div>
          <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
        </div>
      </div>


      <!--cpu-stat-user-group-->
      <div class="col-md-12">
        <div id="container-div-cpu-stat-user-group" class="card mb-3">
          <div class="card-header"><i class="icon-cpu-processor"></i> CPU Stat: User + System + Nice</div>
          <div class="card-body"><div class="chart-panel"></div></div>
          <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
        </div>
      </div>


      <!--cpu-stat-irqGroup-->
      <div class="col-md-12">
        <div id="container-div-cpu-stat-irqGroup" class="card mb-3">
            <div class="card-header"><i class="icon-cpu-processor"></i><span class="spanTitle">{{ .languages.cpuIrqGroupChartTitle }}</span></div>
            <div class="card-body"><div class="chart-panel"></div></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
        </div>
      </div>


      <!--cpu-stat-irq-->
      <div class="col-md-12">
          <div id="container-div-cpu-stat-irq" class="card mb-3">
              <div class="card-header"><i class="icon-cpu-processor"></i><span class="spanTitle">{{ .languages.cpuIrqChartTitle }}</span></div>
              <div class="card-body"><div class="chart-panel"></div></div>
              <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>


      <!--cpu-stat-nettxIrq-->
      <div class="col-md-12">
          <div id="container-div-cpu-stat-nettxIrq" class="card mb-3">
            <div class="card-header"><i class="icon-cpu-processor"></i><span class="spanTitle">{{ .languages.cpuNettxIrqChartTitle }}</span></div>
            <div class="card-body"><div class="chart-panel"></div></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>

      <!--cpu-stat-netrxIrq-->
      <div class="col-md-12">
          <div id="container-div-cpu-stat-netrxIrq" class="card mb-3">
            <div class="card-header"><i class="icon-cpu-processor"></i><span class="spanTitle">{{ .languages.cpuNettxIrqChartTitle }}</span></div>
            <div class="card-body"><div class="chart-panel"></div></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>

      <div class="col-md-12">
          <div id="container-div-cpu-stat-taskletIrq" class="card mb-3">
            <div class="card-header">
                <i class="icon-cpu-processor"></i>
                <span class="spanTitle">{{ .languages.cpuNettxIrqChartTitle }}</span>
            </div>
            <div class="card-body"><div class="chart-panel"></div></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>

      <div class="col-md-12">
          <div id="container-div-cpu-stat-hrtimerIrq" class="card mb-3">
            <div class="card-header">
                <i class="icon-cpu-processor"></i>
                <span class="spanTitle">{{ .languages.cpuNettxIrqChartTitle }}</span>
            </div>
            <div class="card-body"><div class="chart-panel"></div></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>



      <div class="col-md-12">
        <div id="container-div-cpu-avgload" class="card mb-3">
          <div class="card-header">
              <i class="icon-cpu-processor"></i>
              <span class="spanTitle" fullTitle="{{ .languages.averageLoadChartTitleFull }}" shortTitle="{{ .languages.averageLoadChartTitle }}">{{ .languages.averageLoadChartTitle }}</span>
          </div>
          <div class="card-body"><div class="chart-panel"></div></div>
          <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
        </div>
      </div>

      <div class="col-md-12">
          <div id="container-div-cpu-top-table" class="card mb-3">
            <div class="card-header">
                <i class="icon-cpu-processor"></i> CPU TOP
            </div>
            <div class="card-body"><table class="display compact chart-panel" cellspacing="0"></table></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>


      <div class="col-md-12">
          <div id="container-div-memory-chart" class="card mb-3">
            <div class="card-header">
                <i class="icon-ram"></i> {{ .languages.ramChartTitle }}
            </div>
            <div class="card-body"><div class="chart-panel"></div></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>

      <div class="col-md-12">
          <div id="container-memory-stat-table" class="card mb-3">
            <div class="card-header">
                <i class="icon-ram"></i> {{ .languages.memoryConsumptionChartTitle }}
            </div>
            <div class="card-body"><table class="display compact chart-panel" cellspacing="0"></table></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>

      <div class="col-md-12">
          <div class="row">

              <div class="col-md-4">
            <div id="container-div-memory-free-pss-stat" class="card mb-3">
              <div class="card-header">
                  <i class="icon-ram"></i> {{ .languages.memoryPssAgainstTotalChartTitle }}
              </div>
              <div class="card-body"><div class="chart-panel"></div></div>
              <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
            </div>
          </div>
          <div class="col-md-8">
            <div id="container-div-memory-pss-stat" class="card mb-3">
              <div class="card-header">
                  <i class="icon-ram"></i> {{ .languages.memoryPssDonutChartTitle }}
              </div>
              <div class="card-body"><div class="chart-panel"></div></div>
              <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
            </div>
          </div>


          </div>

      </div>
      

      <div class="col-md-12">
          <div id='container-div-IOCharts' class="card mb-3">
              <div class="card-header">
                  <i class="icon-syncalt"></i> {{ .languages.ioChartTitle }}
              </div>
              <div class="card-body"><div class="chart-panel"></div></div>
              <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>

      <div class="col-md-12">
          <div id="container-io-top-table" class="card mb-3">
            <div class="card-header">
                <i class="icon-syncalt"></i> I/O Top Table
            </div>
            <div class="card-body"><table class="display compact chart-panel" cellspacing="0"></table></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>

      <div class="col-md-12">
        <div id="container-jnet-top-table" class="card mb-3">
          <div class="card-header">
            <i class="icon-syncalt"></i> Net Top Table
          </div>
          <div class="card-body"><table class="display compact chart-panel" cellspacing="0"></table></div>
          <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
        </div>
      </div>

      <div class="col-md-12">
          <div id="container-perf-cpu-table" class="card mb-3">
            <div class="card-header">
                <i class="glyphicon glyphicon-tasks"></i>
                <span class="spanTitle" fullTitle="{{ .languages.perfTableTitleFull }}" shortTitle="{{ .languages.perfTableTitle }}">{{ .languages.perfTableTitle }}</span>
            </div>
            <div class="card-body"><table class="display compact chart-panel" cellspacing="0"></table></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>


      <!--flame graph-->
      <div class="col-md-12">
          <div id="container-perf-flame-graph" class="card mb-3">
            <div class="card-header">
                <i class="glyphicon glyphicon-tasks"></i>
                <span class="spanTitle" fullTitle="{{ .languages.perfTableTitleFull }}" shortTitle="{{ .languages.perfTableTitle }}">Flame Graph</span>
            </div>
            <div class="card-body"><div class="chart-panel"></div></div>
            <div class="card-footer small text-muted" hidden><i class="fa fa-bell" aria-hidden="true"> </i></div>
          </div>
      </div>

    </div>
    <!-- /.container-fluid-->
    <!-- /.content-wrapper-->
    <footer class="sticky-footer">
      <div class="container">
        <div class="text-center">
          <small>Copyright © LEP Team 2017</small>
        </div>
      </div>
    </footer>
    <!-- Scroll to Top Button-->
    <a class="scroll-to-top rounded" href="#page-top">
      <i class="fa fa-angle-up"></i>
    </a>
    <!-- Logout Modal-->
    <div class="modal fade" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">Ready to Leave?</h5>
            <button class="close" type="button" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">×</span>
            </button>
          </div>
          <div class="modal-body">Select "Logout" below if you are ready to end your current session.</div>
          <div class="modal-footer">
            <button class="btn btn-secondary" type="button" data-dismiss="modal">Cancel</button>
            <a class="btn btn-primary" href="login.html">Logout</a>
          </div>
        </div>
      </div>
    </div>


    <script src="/static/vendors/jquery/jquery.min.js"></script>
    <script src="/static/vendors/popper/popper.min.js"></script>
    <script src="/static/vendors/bootstrap/js/bootstrap.min.js"></script>
    <script src="/static/vendors/jquery-easing/jquery.easing.min.js"></script>
    <script src="/static/vendors/sb-admin/js/sb-admin.min.js"></script>
    <script src="/static/vendors/datatables/media/js/jquery.dataTables.min.js"></script>
    <script src="/static/vendors/c3/d3.min.js" charset="utf-8"></script>
    <script src="/static/vendors/c3/c3.min.js"></script>
    <script src="/static/vendors/socket-io/socket-io-slim.js"></script>
    <script src="/static/vendors/flameGraph/d4.js"></script>
    <script src="/static/vendors/flameGraph/d4.tip.js"></script>
    <script src="/static/vendors/flameGraph/d4.flameGraph.min.js"></script>

    <script src="/static/js/charts/lepvChart.js"></script>
    <script src="/static/js/charts/lepvGaugeChart.js"></script>

    <script src="/static/js/charts/cpu/cpuStatDonutChart.js"></script>
    <script src="/static/js/charts/cpu/cpuStatIdleChart.js"></script>
    <script src="/static/js/charts/cpu/cpuStatIrqGroupChart.js"></script>
    <script src="/static/js/charts/cpu/cpuStatUserGroupChart.js"></script>

    <script src="/static/js/charts/cpu/cpuGauguChart.js"></script>
    <script src="/static/js/charts/cpu/cpuAvgloadChart.js"></script>

    <script src="/static/js/charts/cpu/cpuIrqChart.js"></script>
    <script src="/static/js/charts/cpu/cpuSoftIrqChart.js"></script>
    <script src="/static/js/charts/cpu/cpuTopTable.js"></script>

    <script src="/static/js/charts/memory/memoryGauguChart.js"></script>
    <script src="/static/js/charts/memory/memoryChart.js"></script>
    <script src="/static/js/charts/memory/memoryStatTable.js"></script>
    <script src="/static/js/charts/memory/procrankFreeVsPieChart.js"></script>
    <script src="/static/js/charts/memory/procrankPssPieChart.js"></script>

    <script src="/static/js/charts/io/ioStatChart.js"></script>
    <script src="/static/js/charts/io/ioGauguChart.js"></script>
    <script src="/static/js/charts/io/ioTopTable.js"></script>

    <script src="/static/js/charts/perf/perfCpuTable.js"></script>

    <script src="/static/js/charts/perf/perfFlameGraph.js"></script>

    <!--<script src="/static/ts/charts/tsLepvChart.js"></script>-->

    <!-- Custom scripts for this page-->
    <!-- Toggle between fixed and static navbar-->
    <script>
    $('#toggleNavPosition').click(function() {
      $('body').toggleClass('fixed-nav');
      $('nav').toggleClass('fixed-top static-top');
    });

    </script>
    <!-- Toggle between dark and light navbar-->
    <script>
    $('#toggleNavColor').click(function() {
      $('nav').toggleClass('navbar-dark navbar-light');
      $('nav').toggleClass('bg-dark bg-light');
      $('body').toggleClass('bg-dark bg-light');
    });

    </script>

     <script>

         var socket = null;
         var socketConnected = false;
         var serverToWatch = null;
         var serverList = [];
         var timeout = 10000;

         $( document ).ready(function() {

            $('[data-toggle="popover"]').popover();

            $('.popover-dismiss').popover({
              trigger: 'focus'
            });



        });
     </script>


     <script>
      function addDevicesOnList(device) {
          var li = document.createElement('li');
          var a = document.createElement('a');

          li.id = "DeviceList-" + device;
          li.value = "device"
          a.innerHTML = device;

          a.onclick = function () {
              console.log("start watch:" + device);
              startWatching(device);
          };

          li.appendChild(a);

          //document.getElementById('collapseDevicesList').appendChild(li);
          document.getElementById('collapseDevicesMenu').appendChild(li);
      }
     </script>

    <script>
      // JS with GO

      // Init start
      var server_ws = new WebSocket('ws://' + window.location.host + '/client');

      server_ws.onmessage = function (e) {
          var obj = JSON.parse(e.data);

          clean_client_menu();
          flush_client_list(obj);
      };
      // Init end


      // Clean client menu
      function clean_client_menu() {
          var menu = document.getElementById('collapseDevicesMenu');
          while (menu.hasChildNodes()) {
              menu.removeChild(menu.lastChild);
          }

      }

      // Flush client menu
      function flush_client_list(clients) {
          for (var k in clients) {
              var c = clients[k];

              addDevicesOnList(c);
          }

      }

      // Add client
      function add_client(client) {
          server_ws.send('{"addClient":"'+client+'"}');
      }

      function onAddDevices() {
          var specifiedServer = $("#txt_server_to_watch").val();
          console.log("Add devices:" + specifiedServer);

          add_client(specifiedServer);

      }
    </script>

    <!--triggered when the user specifies the LEPD server and click the start button.-->
    <script>
    var flag = 0;

      function startWatching(specifiedServer) {
          if (flag != 0) {
              return;
          }
          flag = 1;

        if ( !specifiedServer) {
          return;
        }


        serverToWatch = specifiedServer;

          initializeCharts();

      }


    function initializeCharts() {
        var method = [];

        var cpuStatDonutChart = new CpuStatDonutChart("container-div-cpu-stat-donut", ws, serverToWatch);
        var cpuStatIdleChart = new CpuStatIdleChart("container-div-cpu-stat-idle", ws, serverToWatch);
        var cpuStatUserGroupChart = new CpuStatUserGroupChart("container-div-cpu-stat-user-group", ws, serverToWatch);
        var cpuIrqGroupChart = new CpuIrqGroupChart("container-div-cpu-stat-irqGroup", ws, serverToWatch);
        var cpuIrqChart = new CpuIrqChart("container-div-cpu-stat-irq", socket, serverToWatch);

        var cpuNettxIrqChart = new CpuSoftIrqChart("container-div-cpu-stat-nettxIrq", socket, serverToWatch, 'NET_TX');
        var cpuNetrxIrqChart = new CpuSoftIrqChart("container-div-cpu-stat-netrxIrq", socket, serverToWatch, 'NET_RX');
        var cputaskletIrqChart = new CpuSoftIrqChart("container-div-cpu-stat-taskletIrq", socket, serverToWatch, 'TASKLET');
        var cpuhrtimerIrqChart = new CpuSoftIrqChart("container-div-cpu-stat-hrtimerIrq", socket, serverToWatch, 'HRTIMER');

        var cpuAvgloadChart = new CpuAvgLoadChart("container-div-cpu-avgload", socket, serverToWatch);
        var cpuTopTable = new CpuTopTable("container-div-cpu-top-table", socket, serverToWatch);


        var memoryChart = new MemoryChart('container-div-memory-chart', socket, serverToWatch);
        var memoryStatTable = new MemoryStatTable('container-memory-stat-table', socket, serverToWatch);
        var memoryFreePssStatChart = new ProcrankFreeVsPieChart('container-div-memory-free-pss-stat', socket, serverToWatch);
        var memoryPssStatChart = new ProcrankPssPieChart('container-div-memory-pss-stat', socket, serverToWatch);



        var ioStatChart = new IOStatChart('container-div-IOCharts', socket, serverToWatch);
        var ioTopTable = new IoTopTable('container-io-top-table', socket, serverToWatch);

        var jnetTopTable = new IoTopTable('container-jnet-top-table', socket, serverToWatch);

        var perfCpuTable = new PerfCpuTable('container-perf-cpu-table', socket, serverToWatch);
        var perfFlameGraph = new PerfFlameGraph('container-perf-flame-graph',socket, serverToWatch);


        var chartList = {
            cpuStatDonutChart,cpuStatIdleChart,cpuStatUserGroupChart,cpuIrqGroupChart,
            cpuIrqChart,cpuNettxIrqChart,cpuNetrxIrqChart,cputaskletIrqChart,cpuhrtimerIrqChart,
            cpuAvgloadChart,cpuTopTable,memoryChart,memoryStatTable,memoryFreePssStatChart,
            memoryPssStatChart,ioStatChart,ioTopTable,
            jnetTopTable,

            perfCpuTable, perfFlameGraph
        };

        // Add to method
        for (var i in chartList) {
            var chart = chartList[i]
            var key = chart.socket_message_key + "@" + chart.refreshInterval
            method.push(key)
        }

        var ws = new WebSocket('ws://' +
                               window.location.host +
                               '/monitor');

        ws.onmessage = function (e) {
            var result = JSON.parse(e.data);

            for (var i in chartList) {
                var chart = chartList[i]
                var key = chart.socket_message_key + "@" + chart.refreshInterval
                var response = result[key]

                if (response != null) {
                    chart.updateChartData(response);
                }
            }
        };

        var here = this;

        socket_request_id = 0;
        ws.onopen = function() {
            // remove repeat method
            var set = new Set(method);
            var new_method = Array.from(set);

            var req = '{' +
                '"method":"' + new_method + '",' +
                '"server":"' + serverToWatch + '",' +
                '"request_id":"' + socket_request_id + '",' +
                '"request_time":"' + (new Date()).getTime() + '"' +
                '}';

            ws.send(req);
            socket_request_id++;
        };
    }


</script>

  </div>
</body>

</html>
