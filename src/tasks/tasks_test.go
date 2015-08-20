package tasks

import (
	"testing"
)

func TestHasRequiredFields(t *testing.T) {
	jsonTask := TaskJSON {
		Title: "",
		Description: "",
		IsCompleted: false,
		Priority: 1,
	}

	if hasRequiredFields(jsonTask) == nil {
		t.Errorf("Missing title!")
	}

	jsonTask.Title = "test"
	if hasRequiredFields(jsonTask) == nil {
		t.Errorf("Missing description!")
	}

	jsonTask.Description = "test"
	if hasRequiredFields(jsonTask) != nil {
		t.Errorf("Nothing should be missing now!")
	}
}

func TestSetPriority(t *testing.T) {
	jsonTask := TaskJSON {
		Title: "",
		Description: "",
		IsCompleted: false,
		Priority: "3",
	}

	if priority, err := setPriority(jsonTask); err != nil || priority != 3 {
		t.Errorf("Priority shall be 3")
	}

	jsonTask.Priority = ""
	if priority, err := setPriority(jsonTask); err != nil || priority != 1 {
		t.Errorf("Priority shall be 1")
	}

	jsonTask.Priority = "sada"
	if _, err := setPriority(jsonTask); err == nil {
		t.Errorf("Priority should return error")
	}

}
