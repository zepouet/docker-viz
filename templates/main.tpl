{{define "base"}}
<!DOCTYPE html>
<html lang="fr">
    <head>
        <meta charset="utf-8" />
        <!--[if lt IE 9]><script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script><![endif]-->
        <link rel="stylesheet" href="/css/bootstrap.min.css">
        <link type="text/css" rel="stylesheet" href="/css/stylesheet.css" />
        <link rel="icon" type="image/x-icon" href="/images/docker.ico" />
        <title>Docker-Viz :: {{ .title }}</title>
        <script src="/js/d3.v3.min.js"></script>
        <script src="/js/jquery-1.11.2.min.js"></script>
    </head>
    <body>
        <nav class="navbar navbar-inverse navbar-fixed-top">
              <div class="container-fluid">
                <div class="navbar-header">
                  <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                  </button>
                  <a class="navbar-brand" href="/">Docker-Viz</a>
                </div>
                <div id="navbar" class="navbar-collapse collapse">
                  <ul class="nav navbar-nav navbar-right">
                    <li><a href="/">Dashboard</a></li>
                    <li><a href="#">Help</a></li>
                  </ul>
                </div>
              </div>
            </nav>

            <div class="container-fluid">
              <div class="row">
                <div class="col-sm-3 col-md-2 sidebar">
                  <ul class="nav nav-sidebar">
                    <li><a href="/">Dashboard</a></li>
                  </ul>
                  <ul class="nav nav-sidebar">
                    <li><a href="/dendrogam">Images Dendrogam</a></li>
                    <li><a href="/bubble/images">Images Bubbles</a></li>
                  </ul>
                  <ul class="nav nav-sidebar">
                    <li><a href="/bubble/containers">Containers Bubbles</a></li>
                    <li><a href="/miserables">Containers Miserables</a></li>
                  </ul>
                </div>
                <div class="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
                  <h1 class="page-header">{{ .title }}</h1>
                  {{ template "content" .}}
                </div>
              </div>
            </div>

        <script src="/js/bootstrap.min.js"></script>
    </body>
</html>
{{end}}