package text

import "github.com/microcosm-cc/bluemonday"

// html cleanup

func HTMLSanitize(text string) (string, error) {

	// Do this once for each unique policy, and use the policy for the life of the program
	// Policy creation/editing is not safe to use in multiple goroutines
	p := bluemonday.UGCPolicy()

	// The policy can then be used to sanitize lots of input and it is safe to use the policy in multiple goroutines
	html := p.Sanitize(text)

	return html, nil
}
