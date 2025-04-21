package controllers

import (
	"crud-employee/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func MethodOverride(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte("gagal parse form"))
			return
		}
		if r.Method == http.MethodPost {
			if method := r.FormValue("_method"); method != "" {
				r.Method = strings.ToUpper(method)
			}
		} else {
			w.Write([]byte("hanya mendukung mehod POST"))
			return
		}
		next(w, r)
	}
}

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

func EditEmployee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//mengambil id
		vars := mux.Vars(r)
		id := vars["id"]

		// mengambil data
		var employee models.Employee
		if result := db.Find(&employee, id); result.Error != nil {
			log.Println("gagal mengambil data")
		}

		//parsing data
		data := struct {
			Employee models.Employee
			Title    string
		}{
			Employee: employee,
			Title:    "Employee List",
		}

		fp := filepath.Join("view", "employee", "edit.html")
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

func UpdateEmployee(db *gorm.DB) http.HandlerFunc {
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

		//mengambil id
		vars := mux.Vars(r)
		id := vars["id"]

		employeeUpdate := models.Employee{
			Name:     name,
			Position: position,
			Email:    email,
			Phone:    phone,
		}

		var employee models.Employee
		// db.First(&employee, id)
		// db.Model(&employee).Updates(employeeUpdate)

		// if result := db.Create(&employee); result.Error != nil {
		// 	w.Write([]byte("gagal menyimpan data"))
		// 	return
		// }

		if result := db.First(&employee, id); result.Error != nil{
			 	w.Write([]byte("gagal mengambil data"))
		}

		if result := db.Model(&employee).Updates(employeeUpdate); result.Error != nil{
			w.Write([]byte("gagal memperbarui data"))
		}

		http.Redirect(w, r, "/employee", http.StatusSeeOther)
	}
}
