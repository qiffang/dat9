package parser

import (
	"context"
	"io"
)

// Request describes a source object that may be parsed into smaller derived
// artifacts. SourcePath is the logical dat9 path of the original L2 file.
type Request struct {
	SourcePath string
	MediaType  string
	Size       int64
	Reader     io.Reader
	Metadata   map[string]string
}

// Section is a parsed unit that may later be written back as a small file.
// PathHint is relative to the destination tree chosen by the TreeBuilder.
type Section struct {
	PathHint  string
	Title     string
	MediaType string
	Content   []byte
	Metadata  map[string]string
}

// Relation is the advisory edge schema stored in .relations.json.
type Relation struct {
	Target      string `json:"target"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

// Result captures parser output without deciding where it should be stored.
// Placement is owned by the TreeBuilder layer.
type Result struct {
	Sections  []Section
	Relations []Relation
}

// Parser transforms one source object into zero or more small derived
// artifacts. Implementations may leave Result empty when a format should be
// preserved without section-level expansion.
type Parser interface {
	Parse(ctx context.Context, req Request) (*Result, error)
}
