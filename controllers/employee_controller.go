package controllers

import (
	"crud-employee/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"gorm.io/gorm"
)

func IndexEmployee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// mengambil data
		var employee []models.Employee
		if result := db.Find(&employee); result.Error != nil {
			log.Println("gagal mengambil data")
		}

		//parsing data
		data := struct {
			Employee []models.Employee
			Title    string
		}{
			Employee: employee,
			Title:    "Employee List",
		}

		fp := filepath.Join("view", "employee", "index.html")
		t, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}

func CreateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fp := filepath.Join("view", "employee", "create.html")
		t, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func StoreEmployee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte("gagal parse form"))
			return
		}
		name := r.FormValue("name")
		position := r.FormValue("position")
		email := r.FormValue("email")
		phone := r.FormValue("phone")

		if name == "" {
			w.Write([]byte("nama tidak bileh kosong"))
			return
		}
		if position == "" {
			w.Write([]byte("position tidak bileh kosong"))
			return
		}
		if email == "" {
			w.Write([]byte("email tidak bileh kosong"))
			return
		}
		if phone == "" {
			w.Write([]byte("phone tidak bileh kosong"))
			return
		}

		employee := models.Employee{
			Name:     name,
			Position: position,
			Email:    email,
			Phone:    phone,
		}

		if result := db.Create(&employee); result.Error != nil {
			w.Write([]byte("gagal menyimpan data"))
			return
		}

		http.Redirect(w, r, "/employee", http.StatusSeeOther)
	}
}
