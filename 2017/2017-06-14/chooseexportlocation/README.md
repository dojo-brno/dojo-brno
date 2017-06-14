# Exercise

Write a function such as:

```go
func chooseExportLocation(locations []ExportLocation, shareID string) (ExportLocation, error)
```

Given:

```go
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
	ID string
	Preferred bool
}
```

Rules for choosing the ExportLocation:

1. shareID == ShareInstanceID
2. IsAdminOnly == false
3. Preferred == true are preferred over Preferred == false
4. Locations with lower slice index are preferred over locations with higher slice index

In case no location complies with the above rules an error is returned.
