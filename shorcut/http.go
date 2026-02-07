package shorcut

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	checkLink := path.Base(r.URL.Path)
	// Проверяем, существует ли что-либо на указаной ссылке
	if _, ok := LinksList[checkLink]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(WrongLink.Error()))
		return
	}
	// переводим пользователя по ссылке
	http.Redirect(w, r, LinksList[checkLink].Link, http.StatusFound)
}

func CreateLink(w http.ResponseWriter, r *http.Request) {
	// Нам нужен метод POST для создания ссылки, проверяем
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(WrongMethod.Error()))
		return
	}

	// Создаем переменную для записи значений из запроса и проверяем ее на ошибки
	var input Links
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println("Произошла ошибка: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Произошла ошибка: " + string(err.Error())))
		return
	}

	// Создаем новую укороченную ссылку через функцию описанную в другом файле
	newLink, err := Create(input.Description, input.Link, input.Time, input.InpTimForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Произошла ошибка: " + string(err.Error())))
		return
	}

	// Записываем данные о новой ссылке в byte и возвращаем пользователю
	output, err := json.Marshal(newLink)
	if err != nil {
		fmt.Println("Произошла ошибка: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Произошла ошибка: " + string(err.Error())))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	// Нам нужен метод POST для создания ссылки, проверяем
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(WrongMethod.Error()))
		return
	}

	// Проверяем есть ли у нас querry параметры, чтобы узнать, делаем ли мы запрос о существующей информации
	checkQuerry := r.URL.Query().Get("info")
	// Проверяем, существует ли что-либо на указаному параметру, или мы хотим получить всю информацию
	if checkQuerry == "ALL" || checkQuerry == "all" || checkQuerry == "" {
		output, err := json.Marshal(LinksList)
		if err != nil {
			fmt.Println("Произошла ошибка: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Произошла ошибка: " + string(err.Error())))
			return
		}
		w.Write(output)
		return
	}

	if _, ok := LinksList[checkQuerry]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(WrongLink.Error()))
		return
	}

	output, err := json.Marshal(LinksList[checkQuerry])
	if err != nil {
		fmt.Println("Произошла ошибка: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Произошла ошибка: " + string(err.Error())))
		return
	}
	w.Write(output)
}
