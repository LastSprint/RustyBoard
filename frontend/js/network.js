
const SERVER_URL = "http://127.0.0.1:6644"

function loadAllProjects() {
    var http = new XMLHttpRequest();
    const url= SERVER_URL + "/projects";

    http.open("GET", url, false);
    http.send(null);
    return JSON.parse(http.response)
}