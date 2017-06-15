package chooseexportlocation

import (
	"reflect"
	"testing"
)

func TestChooseExportLocationSuccess(t *testing.T) {
	tests := []struct {
		locs    []ExportLocation
		shareID string
		want    ExportLocation
	}{
		{
			locs: []ExportLocation{
				{
					Path:            "ip://directory",
					ShareInstanceID: "123456",
					IsAdminOnly:     false,
					ID:              "123456-1",
					Preferred:       false,
				},
			},
			shareID: "123456",
			want: ExportLocation{
				Path:            "ip://directory",
				ShareInstanceID: "123456",
				IsAdminOnly:     false,
				ID:              "123456-1",
				Preferred:       false,
			},
		},
		{
			locs: []ExportLocation{
				{
					Path:            "ip://directory",
					ShareInstanceID: "765321",
					IsAdminOnly:     false,
					ID:              "7654321-1",
					Preferred:       false,
				},
				{
					Path:            "ip://directory",
					ShareInstanceID: "1234567",
					IsAdminOnly:     false,
					ID:              "1234567-1",
					Preferred:       false,
				},
			},
			shareID: "1234567",
			want: ExportLocation{
				Path:            "ip://directory",
				ShareInstanceID: "1234567",
				IsAdminOnly:     false,
				ID:              "1234567-1",
				Preferred:       false,
			},
		},
		{
			locs: []ExportLocation{
				{
					Path:            "ip://directory",
					ShareInstanceID: "765321",
					IsAdminOnly:     false,
					ID:              "7654321-1",
					Preferred:       false,
				},
				{
					Path:            "ip://directory",
					ShareInstanceID: "1234567",
					IsAdminOnly:     false,
					ID:              "1234567-1",
					Preferred:       false,
				},
				{
					Path:            "ip://preferred/directory",
					ShareInstanceID: "1234567",
					IsAdminOnly:     false,
					ID:              "1234567-2",
					Preferred:       true,
				},
			},
			shareID: "1234567",
			want: ExportLocation{
				Path:            "ip://preferred/directory",
				ShareInstanceID: "1234567",
				IsAdminOnly:     false,
				ID:              "1234567-2",
				Preferred:       true,
			},
		},
		{
			locs: []ExportLocation{
				{
					Path:            "ip://directory",
					ShareInstanceID: "765321",
					IsAdminOnly:     false,
					ID:              "7654321-1",
					Preferred:       false,
				},
				{
					Path:            "ip://directory",
					ShareInstanceID: "1234567",
					IsAdminOnly:     false,
					ID:              "1234567-1",
					Preferred:       false,
				},
				{
					Path:            "ip://preferred/directory",
					ShareInstanceID: "1234567",
					IsAdminOnly:     false,
					ID:              "1234567-2",
					Preferred:       false,
				},
			},
			shareID: "1234567",
			want: ExportLocation{
				Path:            "ip://directory",
				ShareInstanceID: "1234567",
				IsAdminOnly:     false,
				ID:              "1234567-1",
				Preferred:       false,
			},
		},
	}

	for _, tt := range tests {
		if got, err := ChooseExportLocation(tt.locs, tt.shareID); err != nil {
			t.Errorf("ChooseExportLocation(%v, %q) = (%v, %q) want (%v, nil)", tt.locs, tt.shareID, got, err.Error(), tt.want)
		} else if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("ChooseExportLocation(%v, %q) = (%v, nil) want (%v, nil)", tt.locs, tt.shareID, got, tt.want)
		}
	}
}

func TestChooseExportLocationNotFound(t *testing.T) {
	tests := []struct {
		locs    []ExportLocation
		shareID string
	}{
		{
			locs:    []ExportLocation{},
			shareID: "123456",
		},
		{
			locs: []ExportLocation{
				{
					Path:            "ip://directory",
					ShareInstanceID: "123456",
					IsAdminOnly:     false,
					ID:              "123456-1",
					Preferred:       false,
				},
			},
			shareID: "654321",
		},
		{
			locs: []ExportLocation{
				{
					Path:            "ip://directory",
					ShareInstanceID: "123456",
					IsAdminOnly:     true,
					ID:              "123456-1",
					Preferred:       false,
				},
			},
			shareID: "123456",
		},
		{
			locs: []ExportLocation{
				{
					Path:            "ip://directory",
					ShareInstanceID: "123456",
					IsAdminOnly:     true,
					ID:              "123456-1",
					Preferred:       true,
				},
			},
			shareID: "123456",
		},
		{
			locs: []ExportLocation{
				{
					Path:            "ip://directory",
					ShareInstanceID: "123456",
					IsAdminOnly:     false,
					ID:              "123456-1",
					Preferred:       true,
				},
			},
			shareID: "654321",
		},
	}
	for _, tt := range tests {
		if got, err := ChooseExportLocation(tt.locs, tt.shareID); err == nil {
			t.Errorf("ChooseExportLocation(%v, %q) = (%v, nil) want (\"N/A\", \"an error\")", tt.locs, tt.shareID, got)
		}
	}
}
