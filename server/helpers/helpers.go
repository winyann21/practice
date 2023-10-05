package helpers

import "regexp"

func IsValidEmail(email string) bool {
    // A more permissive regular expression pattern for email validation
    // You can customize this pattern according to your requirements
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
    return regexp.MustCompile(pattern).MatchString(email)
}