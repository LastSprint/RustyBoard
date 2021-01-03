function DrawLineChart(chartId, data, labels) {
    var canvas = document.getElementById(chartId);
    var ctx = canvas.getContext("2d");
    var midX = canvas.width/2;
    var midY = canvas.height/2

    Chart.defaults.global.defaultFontFamily = 'Fire Code'

    var dataset = [{
        data: data,
        backgroundColor: [
            'rgba(14, 73, 181, 0.2)',
        ],
        borderColor: [
            'rgba(14, 73, 181, 1)',
        ],
        borderWidth: 1
    }]

    var myChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: labels,
            datasets: dataset,
        },
        options: {
            maintainAspectRatio: false,
            scales: {
                xAxes: [{
                    gridLines: {
                        display: false
                    }
                }],
                yAxes: [{
                    gridLines: {
                        display: false
                    },
                    ticks: {
                        callback: function(label, index, labels) {
                            return secondsToStringView(label)
                        }
                    }
                }],
            },
            // tooltips: {
            //     // enabled: false
            // },
            title: {
                enabled: false
            },
            legend: {
                display: false
            },
        }
    });
}

function secondsToStringView(seconds) {
    if (seconds < 60) {
        return `${seconds}с`
    } else if (seconds < (60 * 60)) {
        return `${(seconds/60).toFixed(0)}м`
    } else {
        return `${(seconds/60/60).toFixed(0)}ч`
    }
}