package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfig_AddDefaultData(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	ctx := getCtx(req)

	req = req.WithContext(ctx)

	testApp.Session.Put(ctx, "flash", "flash")
	testApp.Session.Put(ctx, "warning", "warning")
	testApp.Session.Put(ctx, "error", "error")

	td := testApp.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Error("failed to get flash value from session")
	}

	if td.Warning != "warning" {
		t.Error("failed to get warning value from session")
	}

	if td.Error != "error" {
		t.Error("failed to get error value from session")
	}
}

func TestConfig_IsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	auth := testApp.isAuthenticated(req)

	if auth {
		t.Error("Returned true for isAuthenticated when no user in session")
	}

	testApp.Session.Put(ctx, "userID", 1)

	auth = testApp.isAuthenticated(req)

	if !auth {
		t.Error("Returned false for isAuthenticated when user in session")
	}
}

func TestConfig_Render(t *testing.T) {
	pathToTemplates = "./templates"

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.render(rr, req, "home.page", &TemplateData{})

	if rr.Code != http.StatusOK {
		t.Error("failed to render template")
	}
}
