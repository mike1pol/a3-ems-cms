package models

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"regexp"
	"time"

	"github.com/go-pg/pg"
)

type Rank struct {
	Id   int64
	Name string
	Next int
	Sort int
}

func (r Rank) IsEqualName(s string) bool {
	return r.Name == s
}

type Personal struct {
	Id            int64
	Name          string
	Status        string
	DismissalDate string
	SteamId       string
	Ranks         []PersonalRank
	Vacations     []PersonalVacation
	Rebuke        []PersonalRebuke
}

func (r Personal) GetCurrentRank() PersonalRank {
	if len(r.Ranks) == 0 {
		return PersonalRank{
			Id:         0,
			PersonalId: r.Id,
			Personal:   r,
			Rank: Rank{
				Id:   0,
				Name: "Not found",
				Sort: 0,
			},
			Date: "1980-01-01",
		}
	}
	return r.Ranks[len(r.Ranks)-1]
}

type PersonalRank struct {
	Id         int64
	PersonalId int64
	Personal   Personal
	RankId     int64
	Rank       Rank
	Date       string
}

type PersonalVacation struct {
	Id         int64
	PersonalId int64
	Personal   Personal
	Start      string
	End        string
}

type PersonalRebuke struct {
	Id          int64
	PersonalId  int64
	Personal    Personal
	Date        string
	Reason      string
	Description string
}

type RankJson struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Next int    `json:"next"`
	Sort int    `json:"sort"`
}

type PersonalRankJson struct {
	Id   int64    `json:"id"`
	Rank RankJson `json:"rank"`
	Date string   `json:"date"`
}

func (r PersonalRankJson) GetDate() string {
	t, e := time.Parse("2006-01-02", r.Date)
	if e != nil {
		return ""
	}
	return t.Format("02.01.2006")
}

func (r PersonalRankJson) GetNextRankDate(vc []PersonalVacationJson) string {
	t, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		return ""
	}
	if r.Rank.Next == 0 {
		return ""
	}
	next := time.Duration(24 * r.Rank.Next)
	nt := t.Add(time.Hour * next)
	if len(vc) > 0 {
		v := vc[len(vc)-1]
		if v.Id > 0 {
			vs, err := time.Parse("2006-01-02", v.Start)
			ve, erre := time.Parse("2006-01-02", v.End)
			if err != nil || erre != nil {
				return ""
			} else {
				tnr := time.Duration(nt.Sub(vs).Hours())
				if vs.Equal(nt) || vs.Before(nt) {
					return ve.Add(time.Hour * tnr).Format("02.01.2006")
				}
			}
		}
	}
	return nt.Format("02.01.2006")
}

type PersonalVacationJson struct {
	Id    int64  `json:"id"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func (v PersonalVacationJson) FormatedDate() string {
	start, e := time.Parse("2006-01-02", v.Start)
	if e != nil {
		return ""
	}
	end, e := time.Parse("2006-01-02", v.End)
	if e != nil {
		return ""
	}
	return fmt.Sprintf("с %s по %s", start.Format("02.01.2006"), end.Format("02.01.2006"))
}

type OdbsPersonalJson struct {
	Date  string  `json:"date"`
	Last  string  `json:"last"`
	First string  `json:"first"`
	Time  float64 `json:"time"`
}

func (o OdbsPersonalJson) GetDate() string {
	t, e := time.Parse("2006-01-02", o.Date)
	if e != nil {
		return ""
	}
	return t.Format("02.01.2006")
}

func (o OdbsPersonalJson) GetTime() string {
	h := math.Floor(o.Time / 60)
	m := math.Floor(o.Time - (h * 60))
	if h > 0 && m > 0 {
		return fmt.Sprintf("%.0fh %.0fm", h, m)
	} else if h > 0 && m <= 0 {
		return fmt.Sprintf("%.0fh", h)
	} else if h <= 0 && m > 0 {
		return fmt.Sprintf("%.0fm", m)
	}
	return "нет"
}

type PersonalReport struct {
	Date  time.Time
	Time  string
	Color string
}

func (r PersonalReport) GetDate(iso bool) string {
	if iso {
		return r.Date.Format("2006-01-02")
	}
	return r.Date.Format("02.01")
}

type PersonalRebukeJson struct {
	Id          int64  `json:"id"`
	Date        string `json:"date"`
	Reason      string `json:"reason"`
	Description string `json:"description"`
}

func (prb PersonalRebukeJson) FormatedDate() string {
	date, e := time.Parse("2006-01-02", prb.Date)
	if e != nil {
		return ""
	}
	return date.Format("02.01.2006")
}

type PersonalList struct {
	Id   int64
	Name string
	Sort int
	List []PersonalJson
}

type PersonalJson struct {
	Id            int64                  `json:"id"`
	Name          string                 `json:"name"`
	Status        string                 `json:"status"`
	DismissalDate string                 `json:"dismissal_date"`
	SteamId       string                 `json:"steam_id"`
	Vacations     []PersonalVacationJson `json:"vacations"`
	Rebukes       []PersonalRebukeJson   `json:"rebukes"`
	Odbs          []OdbsPersonalJson     `json:"odbs"`
	Ranks         []PersonalRankJson     `json:"ranks"`
	Reports       []PersonalReport
}

func (p PersonalJson) GetDismissalDate() string {
	t, e := time.Parse("2006-01-02", p.DismissalDate)
	if e != nil {
		return ""
	}
	return t.Format("02.01.2006")
}

func (p PersonalJson) GetReverseOdbs() []OdbsPersonalJson {
	for i, j := 0, len(p.Odbs)-1; i < j; i, j = i+1, j-1 {
		p.Odbs[i], p.Odbs[j] = p.Odbs[j], p.Odbs[i]
	}
	return p.Odbs
}

func (p PersonalJson) getCurrentVac() (v PersonalVacationJson) {
	if len(p.Vacations) > 0 {
		today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		v = p.GetVacForDate(today)
	}
	return
}

func (p PersonalJson) GetVacForDate(date time.Time) (v PersonalVacationJson) {
	if len(p.Vacations) > 0 {
		for _, vac := range p.Vacations {
			if v.Id > 0 {
				continue
			}
			vs, errS := time.Parse("2006-01-02", vac.Start)
			ve, errE := time.Parse("2006-01-02", vac.End)
			if errS != nil || errE != nil {
				continue
			}
			if (date.Equal(vs) && date.Equal(ve)) || ((date.Equal(vs) || date.After(vs)) && (date.Equal(ve) || date.Before(ve))) {
				v = vac
			}
		}

	}
	return
}

func (p PersonalJson) GetStatus() string {
	if p.Status == "active" {
		vac := p.getCurrentVac()
		if vac.Id > 0 {
			return fmt.Sprintf("Отпуск: %s", vac.FormatedDate())
		}
		return "Активен"
	} else if p.Status == "delete" {
		return "Уволен"
	}
	return p.Status
}

type OType int

const (
	Today OType = 0
	//Week      OType = 1
	Yesterday OType = 2
)

func formatTime(t float64) string {
	h := math.Floor(t / 60)
	m := math.Floor(t - (h * 60))
	return fmt.Sprintf("%.0fh %.0fm", h, m)
}

func findByDate(date string, arr []OdbsPersonalJson) OdbsPersonalJson {
	var r = OdbsPersonalJson{}
	for _, o := range arr {
		if o.Date == date {
			r = o
		}
	}
	return r
}

func (p PersonalJson) GetOnline(t OType) string {
	if len(p.Odbs) == 0 {
		return "0m"
	}
	if t == Today {
		date := findByDate(time.Now().Format("2006-01-02"), p.Odbs)
		if date.Date == "" {
			return "0m"
		}
		return formatTime(date.Time)
	} else if t == Yesterday {
		ty, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		date := findByDate(ty.Add(-time.Hour*24).Format("2006-01-02"), p.Odbs)
		if date.Date == "" {
			return "0m"
		}
		return formatTime(date.Time)
	}
	var tm float64
	for i := range p.Odbs {
		tm = tm + p.Odbs[i].Time
	}
	return formatTime(tm)
}

func (p PersonalJson) GetLastOnline(rev bool) string {
	var idx int
	if rev {
		idx = 0
	} else {
		idx = len(p.Odbs) - 1
	}
	if len(p.Odbs) == 0 || len(p.Odbs[idx].Last) == 0 {
		return "Более 7 дней"
	}
	if len(p.Odbs[idx].First) > 0 {
		return p.Odbs[idx].First + " - " + p.Odbs[idx].Last
	}
	return p.Odbs[idx].Last
}

func (p PersonalJson) GetCurrentRank() PersonalRankJson {
	if len(p.Ranks) == 0 {
		return PersonalRankJson{
			Id: 0,
			Rank: RankJson{
				Id:   0,
				Name: "Not found",
				Next: 0,
				Sort: 0,
			},
			Date: "1980-01-01",
		}

	}
	return p.Ranks[len(p.Ranks)-1]
}

func (p PersonalJson) GetCurrentRankDate() string {
	t, err := time.Parse("2006-01-02", p.GetCurrentRank().Date)
	if err != nil {
		return ""
	}
	return t.Format("02.01.2006")
}

func (p PersonalJson) GetNextRankDate() string {
	t, err := time.Parse("2006-01-02", p.GetCurrentRank().Date)
	if err != nil {
		return ""
	}
	if p.GetCurrentRank().Rank.Next == 0 {
		return ""
	}
	next := time.Duration(24 * p.GetCurrentRank().Rank.Next)
	nt := t.Add(time.Hour * next)
	if len(p.Vacations) > 0 {
		v := p.Vacations[len(p.Vacations)-1]
		if v.Id > 0 {
			vs, err := time.Parse("2006-01-02", v.Start)
			ve, erre := time.Parse("2006-01-02", v.End)
			if err != nil || erre != nil {
				return ""
			} else {
				tnr := time.Duration(nt.Sub(vs).Hours())
				if vs.Equal(nt) || vs.Before(nt) {
					return ve.Add(time.Hour * tnr).Format("02.01.2006")
				}
			}
		}
	}
	return nt.Format("02.01.2006")
}

func (p PersonalJson) GetCurrentRankName() string {
	if len(p.Ranks) == 0 {
		return "Not found"
	}
	return p.Ranks[len(p.Ranks)-1].Rank.Name
}

func (p PersonalJson) GetTimeForDate(date string) string {
	var re = regexp.MustCompile(`(?m)^(\d{2})\.(\d{2})\.(\d{4})`)
	date = re.ReplaceAllString(date, "$3-$2-$1")
	var d OdbsPersonalJson
	for _, o := range p.Odbs {
		if o.Date == date {
			d = o
		}
	}
	return formatTime(d.Time)
}

type PersonalFilter struct {
	Rank   string
	Name   string
	Status string
	Start  time.Time
	End    time.Time
}

func (f PersonalFilter) GetHuman(t int) string {
	if t == 0 {
		return f.Start.Format("02.01.2006")
	}
	return f.End.Format("02.01.2006")
}

func (f PersonalFilter) GetISO(t int) string {
	if t == 0 {
		return f.Start.Format("2006-01-02")
	}
	return f.End.Format("2006-01-02")
}

func GetPersonal(db *pg.DB, filter PersonalFilter) (list []PersonalJson) {
	var jsond []string
	var datesFilter string
	datesFilter = `o.date >= NOW() - INTERVAL '7 DAY'`
	if !filter.Start.IsZero() && !filter.End.IsZero() {
		datesFilter = fmt.Sprintf("o.date >= '%s' AND o.date <= '%s'", filter.Start.Format("2006-01-02"), filter.End.Format("2006-01-02"))
	}
	_, err := db.Query(
		&jsond,
		fmt.Sprintf(`select
  row_to_json(q)
from (
select
  p.*,
  array(
    select row_to_json(rb.*)
    from (
      select prb.id, prb.date, prb.reason, prb.description
      from personal_rebukes prb
      where prb.personal_id = p.id
      order by prb.id asc
    ) rb
  ) rebukes,
  array(
    select row_to_json(v.*)
    from (
      select pv.id, pv.start, pv.end from personal_vacations pv
      where pv.personal_id = p.id
      order by pv.id asc
    ) v
  ) vacations,
  array(
    select
      row_to_json(o.*)
    from (
      select
		o.date,
		to_char(o.first, 'HH24:MI') as first,
		to_char(o.last, 'HH24:MI dd.mm.yyyy') as last,
		o.time
      from odbs o
      where o.name = p.name and %s
      order by o.date asc
    ) o
  ) odbs,
  array(
    select
      row_to_json(prs.*)
    from (
      select
        pr.id,
        pr.date,
        json_build_object('id', r.id, 'name', r.name, 'next', r.next, 'sort', r.sort) rank
      from personal_ranks pr
      inner join ranks r
        on r.id = pr.rank_id where pr.personal_id = p.id
      order by pr.id asc
    ) prs
  ) ranks
  from personals p
) q
where q.status = ? and  q.name like ?`, datesFilter), filter.Status, "%"+filter.Name+"%")
	if err != nil {
		log.Println("Error getting personal ", err)
	}
	for _, j := range jsond {
		var p PersonalJson
		errP := json.Unmarshal([]byte(j), &p)
		if errP != nil {
			log.Println("error parse json: ", errP)
		} else {
			if filter.Rank != "" {
				if p.GetCurrentRankName() == filter.Rank {
					list = append(list, p)
				}
			} else {
				list = append(list, p)
			}
		}
	}

	if !filter.Start.IsZero() && !filter.End.IsZero() {
		for ip := range list {
			diff := int(filter.End.Sub(filter.Start).Hours() / 24)
			for i := 0; i < diff+1; i = i + 1 {
				var r PersonalReport
				if i == 0 {
					r.Date = filter.Start
				} else {
					r.Date = filter.Start.Add(time.Hour * time.Duration(24*i))
				}
				ot := list[ip].getOnlineForDate(r.Date)
				r.Time = ot.GetTime()
				var rank = PersonalRankJson{}
				if len(list[ip].Ranks) > 0 {
					rank = list[ip].Ranks[0]
				} else {
					rank = PersonalRankJson{
						Id: 0,
						Rank: RankJson{
							Id:   0,
							Name: "Nope",
							Next: 0,
							Sort: 0,
						},
						Date: "1980-01-01",
					}
				}
				t, _ := time.Parse("2006-01-02", rank.Date)
				dt, _ := time.Parse("2006-01-02", list[ip].DismissalDate)
				if list[ip].GetVacForDate(r.Date).Id > 0 {
					r.Color = "orange"
				} else if t.After(r.Date) {
					r.Color = "gray"
				} else if ot.Time > 5 {
					r.Color = "green"
				} else {
					r.Color = "red"
				}
				if list[ip].Status == "delete" && (dt.IsZero() || dt.Equal(r.Date) || r.Date.After(dt)) {
					r.Color = "gray"
				}
				list[ip].Reports = append(list[ip].Reports, r)
			}
		}
	}
	return list
}

func (p PersonalJson) getOnlineForDate(date time.Time) (o OdbsPersonalJson) {
	for _, odbs := range p.Odbs {
		if odbs.Date == date.Format("2006-01-02") {
			o = odbs
		}
	}
	return
}

func GetPersonalById(db *pg.DB, id int64) (personal PersonalJson) {
	var jsond []string
	_, err := db.Query(
		&jsond,
		`select
  row_to_json(q)
from (
select
  p.*,
  array(
    select row_to_json(rb.*)
    from (
      select prb.id, prb.date, prb.reason, prb.description
      from personal_rebukes prb
      where prb.personal_id = p.id
      order by prb.id asc
    ) rb
  ) rebukes,
  array(
    select row_to_json(v.*)
    from (
      select pv.id, pv.start, pv.end
      from personal_vacations pv
      where pv.personal_id = p.id
      order by pv.id asc
    ) v
  ) vacations,
  array(
    select
      row_to_json(o.*)
    from (
      select
		o.date,
		to_char(o.first, 'HH24:MI') as first,
		to_char(o.last, 'HH24:MI dd.mm.yyyy') as last,
		o.time
      from odbs o
      where o.name = p.name and o.date >= NOW() - INTERVAL '7 DAY'
      order by o.date asc
    ) o
  ) odbs,
  array(
    select
      row_to_json(prs.*)
    from (
      select
        pr.id,
        pr.date,
        json_build_object('id', r.id, 'name', r.name, 'next', r.next, 'sort', r.sort) rank
      from personal_ranks pr
      inner join ranks r
        on r.id = pr.rank_id where pr.personal_id = p.id
      order by pr.id asc
    ) prs
  ) ranks
  from personals p
) q
where q.id = ?`, id)
	if err != nil {
		log.Println("Error getting personal by id ", err)
	}
	for _, j := range jsond {
		var p PersonalJson
		errP := json.Unmarshal([]byte(j), &p)
		if errP != nil {
			log.Println("error parse json: ", errP)
		} else {
			personal = p
		}
	}
	return
}
