// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package data

import "io"

// DriveRecord contains main properties of cloud server instance
type DriveRecord struct {
	Resource
	Owner  Resource `json:"owner"`
	Status string   `json:"status"`
}

// DriveRecords holds collection of DriverRecord objects
type DriveRecords struct {
	Meta    Meta          `json:"meta"`
	Objects []DriveRecord `json:"objects"`
}

// Drive contains detail properties of cloud server instance
type Drive struct {
	DriveRecord
	Media       string            `json:"media"`
	Meta        map[string]string `json:"meta"`
	Name        string            `json:"name"`
	Size        int64             `json:"size"`
	StorageType string            `json:"storage_type"`
	Jobs        []Resource        `json:"jobs"`
}

// Drives holds collection of Drive objects
type Drives struct {
	Meta    Meta    `json:"meta"`
	Objects []Drive `json:"objects"`
}

// ReadDrives reads and unmarshalls information about cloud drive instances from JSON stream
func ReadDrives(r io.Reader) ([]Drive, error) {
	var drives Drives
	if err := ReadJson(r, &drives); err != nil {
		return nil, err
	}
	return drives.Objects, nil
}

// ReadDrive reads and unmarshalls information about single cloud drive instance from JSON stream
func ReadDrive(r io.Reader) (*Drive, error) {
	var drive Drive
	if err := ReadJson(r, &drive); err != nil {
		return nil, err
	}
	return &drive, nil
}
