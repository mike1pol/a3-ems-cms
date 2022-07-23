{{ define "base" }}
<html>
<head>
    <title>RimasRP {{ if ne .Header.Title ""}}- {{ .Header.Title }}{{ end }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="/static/css/styles.css"/>
    <link rel="stylesheet" href="/static/css/datepicker.min.css"/>
    <link rel="stylesheet" href="/static/css/jquery.typeahead.min.css"/>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css"
          integrity="sha384-WskhaSGFgHYWDcbwN70/dfYBj47jz9qbsMId/iRN3ewGhXQFZCSftd1LZCfmhktB" crossorigin="anonymous">
</head>
<body>
{{ template "header" . }}
{{ template "content" . }}
{{ template "footer" . }}
</body>
</html>
{{ end }}
