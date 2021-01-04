
const SERVER_URL = "http://127.0.0.1:6644"

function loadAllProjects() {

    var http = new XMLHttpRequest();
    const url= SERVER_URL + "/projects";

    console.log(url)

    http.open("GET", url, false);
    http.send(null);

    console.log(JSON.parse(http.response))

    return http.response
}