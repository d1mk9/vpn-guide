package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// PageData — структура для передачи данных в шаблон
type PageData struct {
	Title string
}

// render — универсальный рендер шаблонов
func render(w http.ResponseWriter, tmpl string, data PageData) {
	files := []string{
		filepath.Join("templates", "layout.html"),
		filepath.Join("templates", tmpl),
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Ошибка шаблона: %v", err)
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Ошибка рендера: %v", err)
		http.Error(w, "Ошибка рендера", http.StatusInternalServerError)
	}
}

func main() {
	// Статика
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Маршруты
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "index.html", PageData{Title: "VPN для безопасности и приватности"})
	})

	http.HandleFunc("/vpn-dlya-raboty", func(w http.ResponseWriter, r *http.Request) {
		render(w, "vpn-dlya-raboty.html", PageData{Title: "VPN для удалённой работы"})
	})

	http.HandleFunc("/vpn-dlya-wifi", func(w http.ResponseWriter, r *http.Request) {
		render(w, "vpn-dlya-wifi.html", PageData{Title: "VPN для защиты домашнего Wi-Fi"})
	})

	// Запуск сервера
	port := ":8081"
	log.Printf("✅ VPN сайт работает на http://localhost%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
