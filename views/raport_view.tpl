{{ define "content" }}
<script>
  var raportID = {{ .Raport.ID }};
</script>
<div class="container">
  <div class="row">
    <h1>
      Рапорт {{ .Raport.ID }}
    </h1>
  </div>
  <div class="row">
    <div class="col">
      <h2>Информация:</h2>
      <p>
        <strong>От кого:</strong> {{ .Raport.From.Name }}<br/>
        <strong>На кого:</strong> {{ .Raport.To.Name }}<br/>
        <strong>Статус:</strong> {{ .Raport.GetStatus }}
        <select id="changeStatus">
          <option value="0">Выберите новый статус</option>
          <option value="2">В работе</option>
          <option value="3">Закрыт</option>
        </select>
        <br/>
        <strong>Дата:</strong> {{ .Raport.GetDate }} <br/>
        <strong>Тема:</strong> {{ .Raport.Subject }} <br/>
        <strong>Описание:</strong><br/>
        {{ .Raport.GetBody }} <br/>
      </p>
    </div>
  </div>
  <div class="row">
    <h3>
      Комментарии ({{ .Raport.GetCountComments }}):
    </h3>
  </div>
  <div class="row">
    {{ range .Raport.Comments }}
    <div class="card" style="width: 60%;margin-bottom: 10px;">
      <div class="card-body">
        <h5 class="card-title">{{ .Author.Name }}</h5>
        <h6 class="card-subtitle mb-2 text-muted">{{ .GetDate }}</h6>
        <p class="card-text">{{ .GetMessage }}</p>
      </div>
    </div>
    {{ end }}
  </div>
  <div class="row">
    <div class="col-6">
      <form method="POST" action="/raport/{{ .Raport.ID }}/comment">
        <div class="form-group">
          <label for="pDesc">Комментарий</label>
          <textarea name="message" rows="7" class="form-control" id="pDesc"></textarea>
        </div>
        <button type="submit" class="btn btn-primary">Отправить</button>
      </form>
    </div>
  </div>
</div>
{{ end }}
