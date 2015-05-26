package kdb

type Index interface {
	NewIndexElement(vals []string) (el *IndexElement, err error)
	Add(vals []string, rpos int64) (el *IndexElement, err error)
	Get(vals []string) (el *IndexElement, err error)
	Find(vals []string) (els []*IndexElement, err error)
	Close() (err error)
}

// Struct representing an element in the index. Here we are maintaining a
// tree structure. So, it's `Values` field only containes in the leaf nodes
// `Children` is  only available in root and intermediate nodes
// All the data elements are on the lowest level, which are leaf nodes
type IndexElement struct {
	Values   []string                 `json:"values"`
	Position int64                    `json:"position"`
	Children map[string]*IndexElement `json:"-"`
}