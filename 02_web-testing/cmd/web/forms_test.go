package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)
	fmt.Println("_form.data__", form.Data)
	fmt.Println("_form.error__", form.Errors)
	has := form.Has("test")
	if has {
		t.Error("form shows has field when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("a", "b")
	postedData.Add("b", "a")
	form = NewForm(postedData)
	fmt.Println("_form.data__", form.Data)
	has = form.Has("b")

	if !has {
		t.Error("form says it does not have the field when it should")
	}

}

func Test_TableHas(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("a", "b")
	postedData.Add("b", "a")
	var testHas = []struct {
		name           string
		input          string
		Formvalue      url.Values
		expectedOutput bool
	}{
		{name: "Not Has", Formvalue: url.Values{}, input: "test", expectedOutput: false},
		{name: "Has", Formvalue: postedData, input: "a", expectedOutput: true},
		{name: "Has1", Formvalue: postedData, input: "c", expectedOutput: false},
	}

	for _, test := range testHas {
		form := NewForm(test.Formvalue)
		has := form.Has(test.input)
		if has != test.expectedOutput {
			t.Errorf("expected %v, but has %v", test.expectedOutput, has)
		}
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/login", nil)
	form := NewForm(r.PostForm)
	form.Required("a", "b", "c")
	fmt.Println("Form Valid", form.Data, form.Valid())
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("a", "b")
	postedData.Add("a", "c")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r = httptest.NewRequest("POST", "/login", nil)
	r.PostForm = postedData
	form = NewForm(r.PostForm)

	form.Required("a", "b", "c")
	fmt.Println("Form Valid2", form.Data)
	if !form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}
}

func Test_TableRequired(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("a", "b")
	postedData.Add("a", "c")
	postedData.Add("b", "a")
	postedData.Add("c", "a")
	var testRequired = []struct {
		name           string
		request        *http.Request
		input          []string
		Formvalue      url.Values
		expectedOutput bool
	}{
		{name: "Value Not Available", request: httptest.NewRequest("POST", "/login", nil), input: []string{"a", "b", "c"}, Formvalue: url.Values{}, expectedOutput: false},
		{name: "Value Availble", request: httptest.NewRequest("POST", "/login", nil), input: []string{"a", "b", "c"}, Formvalue: postedData, expectedOutput: true},
	}

	for _, test := range testRequired {
		test.request.PostForm = test.Formvalue
		form := NewForm(test.request.PostForm)
		form.Required(test.input...)
		if form.Valid() != test.expectedOutput {
			t.Errorf("form shows valid when required fields are missing")
		}
	}

}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is mandatory")
	if form.Valid() {
		t.Error("valid() returns false, and it should be true when calling Check()")
	}
}

func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	s := form.Errors.Get("password")
	fmt.Println(len(s), s)
	if len(s) == 0 {
		t.Error("should have an error returned from Get, but do not")
	}

	s = form.Errors.Get("whatever")
	fmt.Println(len(s), s)
	if len(s) != 0 {
		t.Error("should not have an error, but got one")
	}
}

func Test_tableErrorGet(t *testing.T) {
	testError := []struct {
		name           string
		formData       Form
		input          string
		expectedOutput string
	}{
		{name: "Has key", formData: Form{Errors: make(errors)}, input: "password", expectedOutput: "password is required"},
		{name: "Not Has key", formData: Form{Errors: make(errors)}, input: "whatever", expectedOutput: ""},
	}

	for _, test := range testError {
		test.formData.Check(false, "password", "password is required")
		s := test.formData.Errors.Get(test.input)
		if s != test.expectedOutput {
			t.Error("should not have an error, but got one")
		}
	}
}
