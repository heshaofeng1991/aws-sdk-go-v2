// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApplyMethod string

// Values returns all known values for ApplyMethod. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ApplyMethod) Values() []ApplyMethod {
	return []ApplyMethod{
		"immediate",
		"pending-reboot",
	}
}

type SourceType string

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"db-instance",
		"db-parameter-group",
		"db-security-group",
		"db-snapshot",
		"db-cluster",
		"db-cluster-snapshot",
	}
}