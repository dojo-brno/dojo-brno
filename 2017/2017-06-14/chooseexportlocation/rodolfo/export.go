package export

import "errors"

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

var errNoExportLocationAvailable = errors.New("no ExportLocation matches the request")

func chooseExportLocation(locations []ExportLocation, shareID string) (ExportLocation, error) {
	var preferred, others []ExportLocation
	for _, location := range locations {
		if location.ShareInstanceID != shareID || location.IsAdminOnly {
			continue
		}
		if location.Preferred {
			preferred = append(preferred, location)
		} else {
			others = append(others, location)
		}
	}
	for _, location := range preferred {
		return location, nil
	}
	for _, location := range others {
		return location, nil
	}
	return ExportLocation{}, errNoExportLocationAvailable
}
