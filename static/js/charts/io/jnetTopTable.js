/*
 * Open source under the GPLv2 License or later.
 * Copyright (c) 2016, Mac Xu <shinyxxn@hotmail.com>.
 */

var JnetTopTable = function(rootDivName, socket, server) {

    LepvChart.call(this, rootDivName, socket, server);

    this.rootDivName = rootDivName;
    this.socket = socket;
    this.serverToWatch = server;

    this.locateUIElements();

    this.socket_message_key = 'io.jnet';
    
    this.setTableDivName(rootDivName);
    
    this.chartTitle = "Net Top Table";
    this.chartHeaderColor = 'yellow';
    
    this.maxDataCount = 25;
    this.refreshInterval = 20;

    this.initializeChart();
    this.setupSocketIO();
};

JnetTopTable.prototype = Object.create(LepvChart.prototype);
JnetTopTable.prototype.constructor = JnetTopTable;

JnetTopTable.prototype.initializeChart = function() {
    this.table = $('#' + this.mainDivName).DataTable( {
        destroy: true,
        paging: false,
        info: false,
        searching: true,
        columns: [
            {
                title: "Src",
                orderable: false
            },
            {
                title: "Src port",
                orderable: false
            },
            {
                title: "Src bytes",
                orderable: false
            },
            {
                title: "Dst",
                orderable: true
            },
            {
                title: "Dst port",
                orderable: true
            },
            {
                title: "Dst bytes",
                orderable: false
            },
            {
                title: "Proto",
                orderable: false
            },
            {
                title: "Total bytes",
                orderable: false
            },
            {
                title: "Filter data",
                orderable: false
            }
        ],
        order: [[4, "desc"], [5, "desc"]]
    });
};

JnetTopTable.prototype.updateChartData = function(response) {
    data = response['data']
    // console.log(data)
    var thisChart = this;

    this.table.rows().remove().draw( true );
    if (data != null) {
        $.each( data, function( itemIndex, ioppData ) {

            if (itemIndex >= thisChart.maxDataCount) {
                return;
            }

            thisChart.table.row.add([
                ioppData['Src'],
                ioppData['Src port'],
                ioppData['Src bytes'],
                ioppData['Dst'],
                ioppData['Dst port'],
                ioppData['Dst bytes'],
                ioppData['Proto'],
                ioppData['Total bytes'],
                ioppData['Filter data']
            ]);
            index = index + 1;
        });
    } else {
        var index = 0;
        while(index < thisChart.maxDataCount) {
            thisChart.table.row.add([
                "--",
                "--",
                "--",
                "--",
                "--",
                "--",
                "--",
                "--",
                "--"
            ]);
            index = index + 1;
        }
    }
    this.table.draw(true);

    // this.requestData();
};
