// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import "strconv"

type FacetValueType int32

const (
	FacetValueTypeSTRING   FacetValueType = 0
	FacetValueTypeINT      FacetValueType = 1
	FacetValueTypeFLOAT    FacetValueType = 2
	FacetValueTypeBOOL     FacetValueType = 3
	FacetValueTypeDATETIME FacetValueType = 4
)

var EnumNamesFacetValueType = map[FacetValueType]string{
	FacetValueTypeSTRING:   "STRING",
	FacetValueTypeINT:      "INT",
	FacetValueTypeFLOAT:    "FLOAT",
	FacetValueTypeBOOL:     "BOOL",
	FacetValueTypeDATETIME: "DATETIME",
}

var EnumValuesFacetValueType = map[string]FacetValueType{
	"STRING":   FacetValueTypeSTRING,
	"INT":      FacetValueTypeINT,
	"FLOAT":    FacetValueTypeFLOAT,
	"BOOL":     FacetValueTypeBOOL,
	"DATETIME": FacetValueTypeDATETIME,
}

func (v FacetValueType) String() string {
	if s, ok := EnumNamesFacetValueType[v]; ok {
		return s
	}
	return "FacetValueType(" + strconv.FormatInt(int64(v), 10) + ")"
}