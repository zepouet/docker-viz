{{define "base"}}
<!DOCTYPE html>
<html lang="fr">
    <head>
        <meta charset="utf-8" />
        <!--[if lt IE 9]><script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script><![endif]-->
        <link type="text/css" rel="stylesheet" href="/css/stylesheet.css" />
        <link rel="icon" type="image/x-icon" href="/images/docker.ico" />
        <title>Docker-Viz :: {{ .title }}</title>
        <script src="/js/d3.v3.min.js"></script>
        <script src="/js/jquery-1.11.2.min.js"></script>
    </head>
    <body>
        {{template "content"}}
    </body>
</html>
{{end}}