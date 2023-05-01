package validation

type ValidationResult struct {
	Valid  bool     `json:"valid"`
	Errors []string `json:"errors,omitempty"`
}

type Validatable interface {
	Validate() error
}

func Valid() ValidationResult {
	return ValidationResult{Valid: true}
}

func Invalid(errors ...string) ValidationResult {
	return ValidationResult{Valid: false, Errors: errors}
}
