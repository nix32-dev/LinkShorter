package shorcut

import "errors"

var WrongMethod error = errors.New("Ошибка! Неверный метод подключения!")
var WrongLink error = errors.New("Неверно указана ссылка!")
var WrongTime error = errors.New("Неверно указано время!")
var WrongTimeFormat error = errors.New("Неверно указан формат времени! (доступно: minute, hour, day)")
