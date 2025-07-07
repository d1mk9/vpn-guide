package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// PageData — данные, доступные в любом шаблоне
type PageData struct {
	Title string
	Now   time.Time
}

// newPage формирует PageData с актуальной датой/временем
func newPage(title string) PageData {
	return PageData{
		Title: title,
		Now:   time.Now(),
	}
}

// render выводит нужный HTML-шаблон
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

	if err := t.ExecuteTemplate(w, "layout", data); err != nil {
		log.Printf("Ошибка рендера: %v", err)
		http.Error(w, "Ошибка рендера", http.StatusInternalServerError)
	}
}

func main() {
	// Отдаём статику
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Основные маршруты
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "index.html", newPage("VPN для безопасности и приватности"))
	})
	http.HandleFunc("/vpn-dlya-raboty", func(w http.ResponseWriter, r *http.Request) {
		render(w, "vpn-dlya-raboty.html", newPage("VPN для удалённой работы"))
	})
	http.HandleFunc("/vpn-dlya-wifi", func(w http.ResponseWriter, r *http.Request) {
		render(w, "vpn-dlya-wifi.html", newPage("VPN для домашнего Wi-Fi"))
	})
	http.HandleFunc("/privacy-policy", func(w http.ResponseWriter, r *http.Request) {
		render(w, "privacy-policy.html", newPage("Политика конфиденциальности"))
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		render(w, "about.html", newPage("О проекте"))
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		render(w, "contact.html", newPage("Контакты"))
	})

	// robots.txt и sitemap.xml
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/robots.txt")
	})
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		http.ServeFile(w, r, "static/sitemap.xml")
	})

	// Запуск сервера
	const port = ":8081"
	log.Printf("✅ VPN-сайт работает на http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
