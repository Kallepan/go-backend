/*
* This package serves as a domain layer for the application.
* It contains models which are not replicated in the database.
 */

package dco

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}
