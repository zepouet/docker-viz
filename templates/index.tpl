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
      <img data-src="holder.js/200x200/auto/sky" class="img-responsive" alt="">
      <h4>Label</h4>
      <span class="text-muted">Docker Images</span>
    </div>
    <div class="col-xs-6 col-sm-3 placeholder">
      <img data-src="holder.js/200x200/auto/vine" class="img-responsive" alt="">
      <h4>Label</h4>
      <span class="text-muted">Docker Container</span>
    </div>
</div>
<p id="status">Search Docker engine</p>
 <script src="/js/holder.js"></script>
{{end}}