package gittag

import (
	"fmt"
	"strings"

	"github.com/updatecli/updatecli/pkg/plugins/version"
)

// Spec defines a specification for a "gittag" resource
// parsed from an updatecli manifest file
type Spec struct {
	Path          string         // Path contains the git repository path
	VersionFilter version.Filter // VersionFilter provides parameters to specify version pattern and its type like regex, semver, or just latest.
	Message       string         // Message associated to the git tag
}

// GitTag defines a resource of kind "gittag"
type GitTag struct {
	spec         Spec
	foundVersion version.Version // Holds both parsed version and original version (to allow retrieving metadata such as changelog)
}

// New returns a reference to a newly initialized GitTag object from a Spec
// or an error if the provided Filespec triggers a validation error.
func New(spec Spec) (*GitTag, error) {
	return &GitTag{
		spec: spec,
	}, nil
}

// Validate tests that tag struct is correctly configured
func (gt *GitTag) Validate() error {
	validationErrors := []string{}
	if gt.spec.Path == "" {
		validationErrors = append(validationErrors, "Git working directory path is empty while it must be specified. Did you specify an `scmID` or a `spec.path`?")
	}

	err := gt.spec.VersionFilter.Validate()
	if err != nil {
		validationErrors = append(validationErrors, err.Error())
	}

	// Return all the validation errors if found any
	if len(validationErrors) > 0 {
		return fmt.Errorf("Validation error: the provided manifest configuration has the following validation errors:\n%s", strings.Join(validationErrors, "\n\n"))
	}

	return nil
}