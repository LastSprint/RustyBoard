function init() {
    loadAllProjects(function (result) {
        runVue(result)
    })
}

function runVue(result) {
    result = result.data

    new Vue({
        el: '#app',
        data: {
            projects: result
        }
    })

    result.forEach(function (item) {
        fillCharts(item)
    })
}

function fillCharts(project) {
    var an = project.wholeWorkAnalytics;
    var data = [an.taskSpent, an.bugSpent, an.serviceSpent, an.testSpent]
    Draw(project.name + 'CANVAS', data)

    var resp = an.workLog.map(function (item) {
        return {
            x: getDateOfISOWeek(item.week, item.year).getTime(),
            y: item.timeSpent
        }
    }).sort(function (a, b) {
        return a.x - b.x
    })

    var labels = resp.map(function (item){
        return new Date(item.x).toLocaleDateString()
    })

    DrawLineChart(project.name + 'CANVAS_TIME', resp, labels)
}

function getDateOfISOWeek(w, y) {
    var simple = new Date(y, 0, 1 + (w - 1) * 7);
    var dow = simple.getDay();
    var ISOweekStart = simple;
    if (dow <= 4)
        ISOweekStart.setDate(simple.getDate() - simple.getDay() + 1);
    else
        ISOweekStart.setDate(simple.getDate() + 8 - simple.getDay());
    return ISOweekStart;
}