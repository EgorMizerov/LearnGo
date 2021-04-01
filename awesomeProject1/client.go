package main

import "html/template"

var homeTemplate = template.Must(template.New("").Parse(`
<!doctype html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <title>WebSocket</title>
</head>

<body>
<p id="output"></p>

<script>
    ws = new WebSocket('ws://localhost:8080/ws');

    ws.onopen = function() {
        console.log('Connected')
    };

    ws.onmessage = function(evt) {
        var out = document.getElementById('output');
        console.log(evt);
        out.innerHTML += evt.data + '<br>';
    };
</script>
</body>

</html>
`))
