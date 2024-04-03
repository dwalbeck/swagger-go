package api

import (
	"encoding/json"
	"net/http"
)

type DemoApi struct{}

//var _ ServerInterface = (*DemoApi)(nil)

func NewDemoApi() *DemoApi {
	return &DemoApi{}
}

func demoApiError(w http.ResponseWriter, code int, message string) {
	demoErr := Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(demoErr)
}

func (f *DemoApi) FooList(w http.ResponseWriter, r *http.Request, params FooListParams) {
	var result []Foo

	// magically populate the result
	// ...

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

func (f *DemoApi) FooNew(w http.ResponseWriter, r *http.Request) {
	var newFoo Foo
	if err := json.NewDecoder(r.Body).Decode(&newFoo); err != nil {
		demoApiError(w, http.StatusBadRequest, "Invalid format for Foo")
		return
	}

	// magically handle the post data
	// ...

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Success")
}

func (f *DemoApi) FooGet(w http.ResponseWriter, r *http.Request, fooId int32) {
	clr := FooFooColor("purple")
	qty := int(2)
	result := Foo{
		FooColor:    &clr,
		FooQuantity: &qty,
		FooUsState:  "UT",
		PriKey:      456,
	}

	// magically populate the result
	// ...

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}
