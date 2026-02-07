package shorcut

import (
	"math/rand"
	"time"
)

type Links struct {
	Description string    `json:description` // описание
	Link        string    `json:link`        // входная ссылка
	NewLink     string    `json:newlink`     // новая ссылка
	ExpTime     time.Time `json:exptime`     // время когда ссылка будет недействительна
	Time        int       `json: time`       // входное время удаления
	InpTimForm  string    `json: TimeFormat` // тип входного времени
}

var LinksList map[string]Links = make(map[string]Links)

func Create(desc string, link string, timeInp int, timeFormat string) (Links, error) {
	// Проверка на пустоту строки description
	if desc == "" {
		desc = "empty"
	}

	// Проверка на пустоту строки link
	if link == "" {
		return Links{}, WrongLink
	}

	// Проверка на пустоту строки time
	if timeInp <= 0 {
		return Links{}, WrongTime
	}

	// Проверка формата времени и его подсчет
	var expTm time.Duration
	switch {
	case timeFormat == "minute":
		expTm = time.Duration(timeInp) * time.Minute
	case timeFormat == "hour":
		expTm = time.Duration(timeInp) * time.Hour
	case timeFormat == "day":
		expTm = time.Duration(timeInp*24) * time.Hour
	default:
		return Links{}, WrongTimeFormat
	}

	// Создаем рандомный набор букв и цифр длиною 6 символов
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	nLink := string(b)
	// Создаем инфо о ссылке
	new := Links{
		Description: desc,
		Link:        link,
		NewLink:     nLink,
		ExpTime:     time.Now().Add(expTm),
		Time:        timeInp,
		InpTimForm:  timeFormat,
	}

	// Добавляем значения в мапу
	LinksList[nLink] = new

	// Добавляем авто-удаление после истечения времени
	time.AfterFunc(expTm, func() { delete(LinksList, string(b)) })

	// Возвращаем значения в запрос
	return LinksList[string(b)], nil
}
