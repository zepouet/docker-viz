{{define "content"}}
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
 <script src="/js/holder.js"></script>
{{end}}