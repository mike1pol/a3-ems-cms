{{ define "content" }}
<div class="container">
    <div class="row">
        <h1>Кто в сети</h1>
    </div>
    <div class="row">
    {{ range .Servers }}
        <div class="col">
            <h4>
            {{ if .Status }}️✳️{{ else }}🔴{{ end }} {{ .Name }}<br>
            {{ if .Status }}В сети: ({{ .Online }}/{{ .OnlineMed }}/{{ .MaxPlayers }})<br>{{ end }}
            {{ .GetLastUpdate }}
            </h4>
        {{ if .Status }}
            <ul class="list-unstyled">
            {{ range .Players }}
                <li>
                {{ if ne .Personal.Id 0 }}
                    <img src="/static/img/med.png"/>
                    [{{ .Personal.GetCurrentRank.Rank.Name }}]
                {{ end }}
                {{ .Online.Name }}
                </li>
            {{ end }}
            </ul>
        {{ end }}
        </div>
    {{ end }}
    </div>
</div>
<script>
    setTimeout(function () {
        location.reload();
    }, 1000 * 60 * 2);
</script>
{{ end }}
