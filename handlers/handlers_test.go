package handlers

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers_Home(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/")
	if err != nil {
		//t.Log(err)
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("for home page, expected status 200 but got %d", resp.StatusCode)
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bodyText), "Welcome") {
		bend.TakeScreenshot(ts.URL+"/", "Hometest", 1500, 1000)
		t.Error("did not find submarine")
	}
}

//func TestHandlers_Clicker(t *testing.T) {
//	routes := getRoutes()
//	ts := httptest.NewTLSServer(routes)
//	defer ts.Close()
//
//	page := bend.FetchPage(ts.URL + "/tester")
//	outputElement := bend.SelectElementByID(page, "output")
//	button := bend.SelectElementByID(page, "clicker")
//
//	testHtml, err := outputElement.HTML()
//	if err != nil {
//		t.Fatal(err)
//	}
//	if strings.Contains(testHtml, "Clicked the button") {
//		t.Errorf("found text that should not be there")
//	}
//
//	button.MustClick()
//	testHtml, err = outputElement.HTML()
//	if err != nil {
//		t.Fatal(err)
//	}
//	if !strings.Contains(testHtml, "Clicked the button") {
//		t.Errorf("did not find text that should be there")
//	}
//}
//
//func TestHandlers_Home2(t *testing.T) {
//	req, _ := http.NewRequest("GET", "/", nil)
//	ctx := getCtx(req)
//	req = req.WithContext(ctx)
//
//	rr := httptest.NewRecorder()
//	bend.Session.Put(ctx, "test_key", "Hello, world.")
//
//	h := http.HandlerFunc(testHandlers.Home)
//	h.ServeHTTP(rr, req)
//
//	if rr.Code != 200 {
//		t.Errorf("returned wrond response code; got %d but expected 200", rr.Code)
//	}
//
//	if bend.Session.GetString(ctx, "test_key") != "Hello, world." {
//		t.Errorf("did not get correctvalue from session")
//	}
//}
