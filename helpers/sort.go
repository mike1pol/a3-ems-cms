package helpers

import (
	"sort"
	"time"

	m "github.com/mike1pol/rms/models"
)

// OnlineSort online sort
func OnlineSort(s []m.OnlineMed) []m.OnlineMed {
	sort.Slice(s, func(i int, j int) bool {
		if s[i].Personal.Id == 0 && s[j].Personal.Id == 0 {
			return false
		}
		if s[i].Personal.Id == 0 && s[j].Personal.Id != 0 {
			return false
		}
		if s[i].Personal.Id != 0 && s[j].Personal.Id == 0 {
			return true
		}
		return s[i].Personal.GetCurrentRank().Rank.Sort < s[j].Personal.GetCurrentRank().Rank.Sort
	})
	return s
}

// OdbSortByDate online db sort by date
func OdbSortByDate(s []m.OdbMed, srt bool) []m.OdbMed {
	if !srt {
		return s
	}
	sort.Slice(s, func(i int, j int) bool {
		return s[i].Odb.Last.Before(s[j].Odb.Last)
	})
	return s
}

// OdbSort online db sort
func OdbSort(s []m.OdbMed) []m.OdbMed {
	sort.Slice(s, func(i int, j int) bool {
		if s[i].Personal.Id == 0 && s[j].Personal.Id == 0 {
			return false
		}
		if s[i].Personal.Id == 0 && s[j].Personal.Id != 0 {
			return false
		}
		if s[i].Personal.Id != 0 && s[j].Personal.Id == 0 {
			return true
		}
		return s[i].Personal.GetCurrentRank().Rank.Sort < s[j].Personal.GetCurrentRank().Rank.Sort
	})
	return s
}

func rankDateSort(list []m.PersonalJson) []m.PersonalJson {
	sort.Slice(list, func(i int, j int) bool {
		t1, _ := time.Parse("2006-01-02", list[i].GetCurrentRank().Date)
		t2, _ := time.Parse("2006-01-02", list[j].GetCurrentRank().Date)
		return t1.Before(t2)
	})
	return list
}

// RankDateSort Sort by rank and date
func RankDateSort(list []m.PersonalJson) (r []m.PersonalJson) {
	var l1 []m.PersonalJson
	var l2 []m.PersonalJson
	var l3 []m.PersonalJson
	var l4 []m.PersonalJson
	var l5 []m.PersonalJson
	var l6 []m.PersonalJson
	var l7 []m.PersonalJson
	for i := range list {
		if list[i].GetCurrentRank().Id == 0 {
			l7 = append(l7, list[i])
		} else if list[i].GetCurrentRank().Rank.Sort == 1 {
			l1 = append(l1, list[i])
		} else if list[i].GetCurrentRank().Rank.Sort == 2 {
			l2 = append(l2, list[i])
		} else if list[i].GetCurrentRank().Rank.Sort == 3 {
			l3 = append(l3, list[i])
		} else if list[i].GetCurrentRank().Rank.Sort == 4 {
			l4 = append(l4, list[i])
		} else if list[i].GetCurrentRank().Rank.Sort == 5 {
			l5 = append(l5, list[i])
		} else if list[i].GetCurrentRank().Rank.Sort == 6 {
			l6 = append(l6, list[i])
		} else if list[i].GetCurrentRank().Rank.Sort == 7 {
			l7 = append(l7, list[i])
		}
	}
	r = append(r, rankDateSort(l1)...)
	r = append(r, rankDateSort(l2)...)
	r = append(r, rankDateSort(l3)...)
	r = append(r, rankDateSort(l4)...)
	r = append(r, rankDateSort(l5)...)
	r = append(r, rankDateSort(l6)...)
	r = append(r, rankDateSort(l7)...)
	return
}

func findInRList(id int64, list []m.PersonalList) int {
	for i := range list {
		if list[i].Id == id {
			return i
		}
	}
	return -1
}

// DutySortByRank Sort duty by rank
func DutySortByRank(list []m.Duty) []m.Duty {
	sort.Slice(list, func(i int, j int) bool {
		return list[i].Personal.GetCurrentRank().Rank.Sort < list[j].Personal.GetCurrentRank().Rank.Sort
	})
	return list
}

// RankDateSortByRank Sort rank by date
func RankDateSortByRank(list []m.PersonalJson) (rList []m.PersonalList) {
	for i := range list {
		indexRL := findInRList(list[i].GetCurrentRank().Rank.Id, rList)
		if indexRL == -1 {
			rList = append(rList, m.PersonalList{
				Id:   list[i].GetCurrentRank().Rank.Id,
				Name: list[i].GetCurrentRank().Rank.Name,
				Sort: list[i].GetCurrentRank().Rank.Sort,
				List: []m.PersonalJson{list[i]},
			})
		} else {
			rList[indexRL].List = append(rList[indexRL].List, list[i])
		}
	}
	for i := range rList {
		rList[i].List = rankDateSort(rList[i].List)
	}
	sort.Slice(rList, func(i int, j int) bool {
		return rList[i].Sort < rList[j].Sort
	})
	return
}

// RankSort rank sort
func RankSort(list []m.Rank, desc bool) []m.Rank {
	if desc {
		sort.Slice(list, func(i int, j int) bool {
			return list[i].Sort < list[j].Sort
		})
		return list
	}

	sort.Slice(list, func(i int, j int) bool {
		return list[i].Sort > list[j].Sort
	})
	return list
}
