function Draw(chartId, data) {
    var canvas = document.getElementById(chartId);
    var ctx = canvas.getContext("2d");
    var midX = canvas.width/2;
    var midY = canvas.height/2

    Chart.defaults.global.defaultFontFamily = 'Fire Code'

    var dataset = [{
        data: data,
        backgroundColor: [
            'rgba(14, 73, 181, 0.4)',
            'rgba(202, 62, 71, 0.4)',
            `rgba(221, 221, 221, 0.4)`,
            'rgba(22, 199, 154, 0.4)'
        ],
        borderColor: [
            'rgba(14, 73, 181, 1)',
            'rgba(202, 62, 71, 1)',
            `rgba(221, 221, 221, 1)`,
            'rgba(22, 199, 154, 1)'
        ],
        borderWidth: 0.8
    }]

    var myChart = new Chart(ctx, {
        type: 'horizontalBar',
        data: {
            labels: ['На таски', 'На баги', 'На севрис-таски', 'На тесты'],
            datasets: dataset,
        },
        options: {
            maintainAspectRatio: false,
            animation: {
                onComplete: function () {

                    var chartInstance = this.chart;
                    var ctx = chartInstance.ctx;

                    var ctx = this.chart.ctx;
                    var height = chartInstance.controller.boxes[0].bottom;
                    ctx.fontFamily = "Fire Code";
                    ctx.fontSize = 16;
                    ctx.fillStyle = "#f5f3f1";
                    ctx.textAlign = "center";
                    ctx.textBaseline = "bottom";

                    this.data.datasets.forEach(function (dataset, i) {
                        var meta = chartInstance.controller.getDatasetMeta(i);
                        meta.data.forEach(function (bar, index){
                            var width = bar._model.x;
                            var height = bar._model.height;

                            var xPos = 0 + 24
                            var yPos = bar._model.y + height/4

                            ctx.fillText(secondsToStringView(dataset.data[index]), xPos, yPos);
                        })
                    })
                }
            },
            scales: {
                xAxes: [{
                    display: false
                }],
                yAxes: [{
                    display: false
                }],
            },
            tooltips: {
                enabled: true,
                callbacks: {
                    label: function(tooltipItem, data) {
                        return secondsToStringView(tooltipItem.value)
                    }
                }
            },
            title: {
                enabled: false
            },
            legend: {
                display: false,
            },
        }
    });
}

function secondsToStringView(seconds) {
    if (seconds < 60) {
        return `${seconds}с`
    } else if (seconds < (60 * 60)) {
        return `${(seconds/60).toFixed(2)}м`
    } else {
        return `${(seconds/60/60).toFixed(2)}ч`
    }
}