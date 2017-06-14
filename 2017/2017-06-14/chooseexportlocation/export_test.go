package export

import "testing"

func TestChooseExportLocationSuccess(t *testing.T) {
	shareID := "share-id"
	tests := []struct {
		should       string // document what behavior is being tested
		locations    []ExportLocation
		desiredIndex int // index of locations with the correct result
	}{
		{
			should: "return the first location possible",
			locations: []ExportLocation{
				{"/tmp", shareID, false, "id1", true},
			},
			desiredIndex: 0,
		},
		{
			should: "return a location even if not preferred",
			locations: []ExportLocation{
				{"/tmp", shareID, false, "id1", false},
			},
			desiredIndex: 0,
		},
		{
			should: "filter out invalid shareIDs",
			locations: []ExportLocation{
				{"/tmp", "a-different-share-id", false, "id1", true},
				{"/any/path", shareID, false, "id2", true},
			},
			desiredIndex: 1,
		},
		{
			should: "filter out admin locations",
			locations: []ExportLocation{
				{"/tmp", shareID, true, "id1", true},
				{"/any/path", shareID, false, "id2", true},
			},
			desiredIndex: 1,
		},
		{
			should: "prefer preferred locations",
			locations: []ExportLocation{
				{"/tmp", shareID, false, "id1", false},
				{"/any/path", shareID, false, "id2", false},
				{"/yet/another/path", shareID, false, "id3", true},
			},
			desiredIndex: 2,
		},
	}
	for _, tt := range tests {
		locations := tt.locations
		want := locations[tt.desiredIndex]
		got, err := chooseExportLocation(locations, shareID)
		if err != nil {
			t.Errorf("chooseExportLocation should %v\nerr = %v, want nil", tt.should, err)
		}
		if got != want {
			t.Errorf("chooseExportLocation should %v\ngot %+v, want %+v", tt.should, got, want)
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
