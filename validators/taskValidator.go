package validators

import (
	"errors"
	"practice/server/store"
	"time"
)

func validateText(text string) (bool, error) {
	if text == "" {
		return false, errors.New("Text field should not be empty")
	}
	return true, nil
}

func validateTags(tags []string) (bool, error) {
	if len(tags) == 0 {
		return false, errors.New("tags should not be empty")
	}
	return true, nil
}

func validateDue(due time.Time) (bool, error) {
	// if due == nil {
	// 	return false, errors.New("Due should not be empty")
	// }

	return true, nil
}

func ValidateTask(task store.Task) error {
	var isValid bool
	isValid, err := validateText(task.Text)
	if !isValid {
		return err
	}

	isValid, err = validateTags(task.Tags)
	if !isValid {
		return err
	}

	isValid, err = validateDue(task.Due)
	if !isValid {
		return err
	}

	return nil
}
