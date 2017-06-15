package chooseexportlocation

import "fmt"

// ExportLocation see OpenStack Shares, Export Locations
type ExportLocation struct {
	// The export location path that should be used for mount operation.
	Path string
	// The UUID of the share instance that this export location belongs to.
	ShareInstanceID string
	// Defines purpose of an export location. If set to true, then it is
	// expected to be used for service needs and by administrators only. If it
	// is set to false, then this export location can be used by end users.
	IsAdminOnly bool
	// The share export location UUID.
	ID        string
	Preferred bool
}

// ChooseExportLocation chooses one ExportLocation according to the rules in the README
func ChooseExportLocation(locs []ExportLocation, shareID string) (ExportLocation, error) {
	if len(locs) == 0 {
		return ExportLocation{}, fmt.Errorf("Error: received an empty list of export locations")
	}
	foundMatchingNotPreferred := false
	var matchingNotPreferred ExportLocation
	for _, loc := range locs {
		if !loc.IsAdminOnly && loc.ShareInstanceID == shareID && loc.Preferred {
			return loc, nil
		}
		if !loc.IsAdminOnly && !foundMatchingNotPreferred && loc.ShareInstanceID == shareID {
			matchingNotPreferred = loc
			foundMatchingNotPreferred = true
		}
	}
	if foundMatchingNotPreferred {
		return matchingNotPreferred, nil
	}
	return ExportLocation{}, fmt.Errorf("Error: not found any non-AdminOnly export location")
}
