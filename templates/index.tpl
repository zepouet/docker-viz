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
<div class="row placeholders">
    <div class="col-xs-6 col-sm-3 placeholder">
      <img data-src="holder.js/200x200?theme=sky&amp;text={{ .countImages}}" class="img-responsive" alt="{{ .countImages}}">
      <h4>Docker Images</h4>
    </div>
    <div class="col-xs-6 col-sm-3 placeholder">
      <img data-src="holder.js/200x200?theme=vine&amp;text={{ .countContainers}}" class="img-responsive" alt="{{ .countContainers}}">
      <h4>Docker Container</h4>
    </div>
</div>
<p id="status">Search Docker engine</p>
 <script src="/js/holder.js"></script>
{{end}}