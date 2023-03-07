package dtos

import "time"

// Meta ...
type Meta struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Required     bool   `json:"required"`
	Order        int    `json:"order"`
	Enabled      bool   `json:"enabled"`
	DefaultValue string `json:"defaultValue"`
	ID           int    `json:"id"`
}

// Item ...
type Item struct {
	CreatedAt        time.Time         `json:"createdAt"`
	UpdatedAt        time.Time         `json:"updatedAt"`
	DeletedAt        interface{}       `json:"deletedAt"`
	Version          int               `json:"version"`
	ID               int               `json:"id"`
	Label            string            `json:"label"`
	Order            int               `json:"order"`
	Meta             map[string]string `json:"meta"`
	ParentID         interface{}       `json:"parentId"`
	MenuID           int               `json:"menuId"`
	Enabled          bool              `json:"enabled"`
	StartPublication interface{}       `json:"startPublication"`
	EndPublication   interface{}       `json:"endPublication"`
	Template         interface{}       `json:"template"`
	TemplateFormat   interface{}       `json:"templateFormat"`
}

// Menu ...
type Menu struct {
	Name           string `json:"name"`
	Meta           []Meta `json:"meta"`
	Template       string `json:"template"`
	TemplateFormat string `json:"templateFormat"`
	Items          []Item `json:"items"`
}
