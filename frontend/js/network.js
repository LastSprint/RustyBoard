
const SERVER_URL = "https://lastsprint.dev"

function loadAllProjects() {

    var http = new XMLHttpRequest();
    const url= SERVER_URL + "/projects";

    console.log(url)

    http.open("GET", url, false);
    http.send(null);

    console.log(JSON.parse(http.response))

    return http.response
}