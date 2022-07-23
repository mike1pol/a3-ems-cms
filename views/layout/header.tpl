{{ define "header" }}
<header style="margin-bottom:10px;">
  <nav class="navbar navbar-expand-md navbar-dark bg-dark">
    <a class="navbar-brand" href="/">МЧС</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarCollapse"
            aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarCollapse">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item{{ if (eq .Header.Page "home") }} active{{ end }}">
          <a class="nav-link" href="/">В сети</a>
        </li>
        <li class="nav-item{{ if (eq .Header.Page "duty") }} active{{ end }}">
          <a class="nav-link" href="/duty">Зоны дежурства</a>
        </li>
        <li class="nav-item{{ if (eq .Header.Page "blacklist") }} active{{ end }}">
          <a class="nav-link" href="/blacklist">Черный список</a>
        </li>
        <li
          class="nav-item dropdown{{ if or (eq .Header.Page "refDB") (eq .Header.Page "odb") (eq .Header.Page "psy") }} active{{ end }}">
          <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink" data-toggle="dropdown"
             aria-haspopup="true" aria-expanded="false">
            База данных
          </a>
          <div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
            <a class="dropdown-item{{ if (eq .Header.Page "refDB") }} active{{ end }}"
               href="/refdb">База
              справок</a>
          {{ if .Header.User.IsUser }}
            <a class="dropdown-item{{ if (eq .Header.Page "psy") }} active{{ end }}"
               href="/psy">Психический больных</a>
            <a class="dropdown-item{{ if (eq .Header.Page "odb") }} active{{ end }}" href="/odb">База
              онлана</a>
          {{ end }}
          </div>
        </li>
      {{ if .Header.User.IsUser }}
        <li class="nav-item{{ if (eq .Header.Page "personal") }} active{{ end }}">
          <a class="nav-link" href="/personal">Личный состав</a>
        </li>
      {{ if .Header.User.IsAdmin }}
        <li class="nav-item{{ if (eq .Header.Page "report") }} active{{ end }}">
          <a class="nav-link" href="/report">Отчеты</a>
        </li>
        <li
          class="nav-item dropdown{{ if or (eq .Header.Page "raport") (eq .Header.Page "raport_list") }} active{{ end }}">
          <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuRLink" data-toggle="dropdown"
             aria-haspopup="true" aria-expanded="false">
            Рапорта
          </a>
          <div class="dropdown-menu" aria-labelledby="navbarDropdownMenuRLink">
            <a class="dropdown-item{{ if (eq .Header.Page "raport") }} active{{ end }}"
               href="/raport">Подать</a>
            <a class="dropdown-item{{ if (eq .Header.Page "raport_list") }} active{{ end }}"
               href="/raport/list">Посмотреть</a>
          </div>
        </li>
      {{ else }}
        <li class="nav-item{{ if (eq .Header.Page "raport") }} active{{ end }}">
          <a class="nav-link" href="/raport">Подать рапорт</a>
        </li>
      {{ end }}
      {{ end }}
        <li class="nav-item">
          <a class="nav-link" target="_blank" href="http://rimasrp.ru">Стенгазета</a>
        </li>
      </ul>
    {{ if gt .Header.User.Id 0 }}
      <span class="navbar-text">
        <img style="height:21px;margin-right:10px;" src="{{ .Header.User.Avatar }}"><a target="_blank"
                                                                                       href="{{ .Header.User.ProfileUrl }}">{{ .Header.User.Name }}</a>
      </span>
    {{ end }}
      <ul class="navbar-nav">
      {{ if .Header.User.IsAdmin }}
        <li class="nav-item dropdown{{ if (eq .Header.Page "admin") }} active{{ end }}">
          <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink" data-toggle="dropdown"
             aria-haspopup="true" aria-expanded="false">
            Admin
          </a>
          <div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
            <a class="dropdown-item" href="/server">Сервера</a>
            <a class="dropdown-item" href="/user">Пользователи</a>
            <a class="dropdown-item" href="/rank">Ранги</a>
          </div>
        </li>
      {{ end }}
        <li class="navbar-item">
        {{ if gt .Header.User.Id 0 }}
          <a class="nav-link" href="/logout">Выйти</a>
        {{ else }}
          <a class="nav-link" href="/login">Вход</a>
        {{ end }}
        </li>
      </ul>
    </div>
  </nav>
</header>
{{ end }}
