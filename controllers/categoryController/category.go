package categorycontroller

import (
	"go-init/entities"
	categorymodel "go-init/models/categoryModel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}

	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}
		category := categorymodel.Detail(id)
		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

}
