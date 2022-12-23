package entity

import "testing"

func TestValidate(t *testing.T) {
	task1 := Task{
		Name:        "test",
		Description: "test",
		Type:        "http",
		Target:      "http://localhost",
		HttpHost:    "localhost",
		SslVerify:   true,
		SslExpire:   true,
		Interval:    60,
		Timeout:     10000,
	}
	err := Validate(&task1)
	if err != nil {
		t.Error(err)
	}
	t.Log("task1 is passed")

	task2 := Task{
		Name:        "test",
		Description: "test",
		Type:        "blabla",
		Target:      "http://localhost",
		HttpHost:    "localhost",
		SslVerify:   true,
		SslExpire:   true,
		Interval:    60,
		Timeout:     10000,
	}
	err = Validate(&task2)
	if err == nil {
		t.Error("task2 is invalid but no validate error")
	}
	t.Log("task2 is passed, err=", err)

	task3 := Task{
		Name:        "test",
		Description: "test",
		Type:        "http",
		Target:      "http://localhost",
		HttpHost:    "localhost",
		SslVerify:   true,
		SslExpire:   true,
		Interval:    -1,
		Timeout:     10000,
	}

	err = Validate(&task3)
	if err == nil {
		t.Error("task3 is invalid but no validate error")
	}
	t.Log("task3 is passed, err=", err)
}
