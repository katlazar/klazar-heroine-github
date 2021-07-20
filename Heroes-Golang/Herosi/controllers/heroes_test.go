//controllers/heroes_test.go

package controllers

import (
	"bytes"
	"fmt"
	"herosi/models"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error
var apiPath = "/api/heroes"

func testInit(router *gin.Engine) {
	db, err = gorm.Open("sqlite3", "hero_test.db")
	db.AutoMigrate(&models.Hero{})
	hero := models.Hero{Name: "Szymon"}
	db.Create(&hero)
	hero = models.Hero{Name: "Jakub"}
	db.Create(&hero)
	hero = models.Hero{Name: "Jan"}
	db.Create(&hero)

	if err != nil {
		panic("Failed to connect to database!")
	}
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
}

func testEnd() {
	_ = db.Close()
	os.Remove("hero_test.db")
}

func TestGetHeroes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	testInit(router)

	router.GET(apiPath, GetHeroes)

	req, err := http.NewRequest("GET", apiPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	expected := `[{"id":1,"name":"Szymon"},{"id":2,"name":"Jakub"},{"id":3,"name":"Jan"}]`
	if resp.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
	testEnd()
}

func TestAddHero(t *testing.T) {
	var jsonStr = []byte(`{"name":"Andrzej"}`)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	testInit(router)
	req, err := http.NewRequest("POST", apiPath, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	router.POST(apiPath, AddHero)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":4,"name":"Andrzej"}`
	if resp.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
	testEnd()
}

func TestGetHero(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	testInit(router)
	req, err := http.NewRequest("GET", apiPath+"/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	router.GET(apiPath+"/:id", GetHero)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":3,"name":"Jan"}`
	if resp.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
	testEnd()
}

func TestGetNoHero(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	testInit(router)
	req, err := http.NewRequest("GET", apiPath+"/10", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	router.GET(apiPath+"/:id", GetHero)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if status := resp.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	expected := `{"error":"Hero with id: 10 not found!"}`
	if resp.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
	testEnd()
}

func TestPutHero(t *testing.T) {
	var jsonStr = []byte(`{"name":"Piotr"}`)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	testInit(router)
	req, err := http.NewRequest("PUT", apiPath+"/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	router.PUT(apiPath+"/:id", PutHero)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"name":"Piotr"}`
	if resp.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
	testEnd()
}

func TestDeleteHero(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	testInit(router)
	req, err := http.NewRequest("DELETE", apiPath+"/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	router.DELETE(apiPath+"/:id", DeleteHero)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "true"
	if resp.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
	testEnd()
}

func TestDeleteNoHero(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	testInit(router)
	req, err := http.NewRequest("DELETE", apiPath+"/100", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	router.DELETE(apiPath+"/:id", DeleteHero)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if status := resp.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	expected := `{"error":"Hero with id: 100 not found!"}`
	if resp.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
	testEnd()
}

