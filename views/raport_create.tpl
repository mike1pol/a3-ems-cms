{{ define "content" }}
<div class="container">
    <div class="row">
        <h1>
        Подать рапорт
        </h1>
    </div>
    <div class="row">
      <div class="col-6">
        <form method="POST">
          <div class="form-group">
            <label for="toID">Сотрудник</label>
            <select name="toID" class="form-control" id="toID">
              <option value="">Выберите сотрудника</option>
              {{ range .Personal }}
              <option value="{{ .Id }}">{{ .Name }}</option>
              {{ end }}
            </select>

          </div>
          <div class="form-group">
            <label for="pDate">Дата</label>
            <input name="date" type="date" data-toggle="datepicker" class="form-control" id="pDate">
          </div>
          <div class="form-group">
            <label for="bSubj">Тема</label>
            <input name="subject" type="text" class="form-control" id="bSubj">
          </div>
          <div class="form-group">
            <label for="pDesc">Описание</label>
            <textarea name="body" rows="7" class="form-control" id="pDesc"></textarea>
          </div>
          <button type="submit" class="btn btn-primary">Отправить</button>
        </form>
      </div>
    </div>
</div>
{{ end }}
