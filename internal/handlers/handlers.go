package handlers

import (
	"go1fl-sprint6-final-tpl/internal/service"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Создаём стартовую страницу index.html для загрузки файлов
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, "Ошибка чтения index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//Ограничение размера
	r.ParseMultipartForm(10 << 20)

	//Если метод не POST, то выдать ошибку
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Получаем файл
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	//Читаем файл
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//Конвертируем данные из файла в морзе или текст
	input := string(data)
	output, err := service.Replace(input)
	if err != nil {
		log.Printf("Ошибка конвертирования: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//Создаём и записываем в файл
	filename := time.Now().UTC().Format("2006-01-02-150405") + filepath.Ext(handler.Filename)
	err = os.Chdir("../")
	if err != nil {
		http.Error(w, "Ошибка пути сохранения", http.StatusInternalServerError)
		return
	}
	err = os.WriteFile(filename, []byte(output), 0755)
	if err != nil {
		http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
		return
	}

	//Устанавливаем заголовок, как результат изменений
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(output))
}
