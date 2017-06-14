package export

import "testing"

func TestChooseExportLocationSuccess(t *testing.T) {
	shareID := "share-id"
	tests := []struct {
		locations    []ExportLocation
		desiredIndex int // index of locations with the correct result
	}{
		{
			[]ExportLocation{
				{"/tmp", shareID, false, "123", true},
			},
			0,
		},
	}
	for _, tt := range tests {
		locations := tt.locations
		want := locations[tt.desiredIndex]
		got, err := chooseExportLocation(locations, shareID)
		if err != nil {
			t.Errorf("chooseExportLocation(%#v, %#v): err = %#v, want nil", locations, shareID, err)
		}
		if got != want {
			t.Errorf("chooseExportLocation(%#v, %#v) = %#v, want %#v", locations, shareID, got, want)
		}
	}
}

func TestChooseExportLocationError(t *testing.T) {
	exportLocations := []ExportLocation{}
	shareID := "share-id"
	if _, err := chooseExportLocation(exportLocations, shareID); err != noExportLocationAvailable {
		t.Errorf("chooseExportLocation(%#v, %#v): err = nil, want %#v", exportLocations, shareID, noExportLocationAvailable)
	}
}
