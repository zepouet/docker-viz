{{define "content"}}
<script>
    $.ajax({
        url: 'docker',
        type: 'GET',
        success: function(data) {
            if((data == "false")) {
                $('#status').text("Docker Engine not found");
            } else {
                $('#status').text("");
            }
        }
    });
</script>
<p id="status">Search Docker engine</p>
<p><a href="./dendrogam">Images Dendrogam</a></p>
<p><a href="./bubble/images">Images Bubbles</a></p>
<p><a href="./bubble/containers">Containers Bubbles</a></p>
<p><a href="./miserables">Containers Miserables</a></p>
{{end}}