// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import "fmt"

// Meta describes properties of dataset
type Meta struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"total_count"`
}

// Resource describes properties of linked resource
type Resource struct {
	URI  string `json:"resource_uri"`
	UUID string `json:"uuid"`
}

// MakeResource returns Resource structure from given type and UUID
func MakeResource(t, uuid string) Resource {
	return Resource{
		URI:  fmt.Sprintf("/api/2.0/%s/%s/", t, uuid),
		UUID: uuid,
	}
}