package fbx_test

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/dgraph/fb"
	"github.com/dgraph-io/dgraph/fbx"
	"github.com/dgraph-io/dgraph/protos/pb"
	"github.com/stretchr/testify/require"
)

func TestPosting(t *testing.T) {
	uid := uint64(1)
	value := []byte("value")
	valueType := pb.Posting_BINARY
	langTag := []byte("langTag")
	label := "label"
	facets := make([]*fb.Facet, 0)
	for i := 0; i < 5; i++ {
		facet := fbx.NewFacet().
			SetKey(fmt.Sprintf("facet%d", i)).
			Build()
		facets = append(facets, facet)
	}
	op := fb.DirectedEdgeOpDEL
	startTs := uint64(2)
	commitTs := uint64(3)

	builder := fbx.NewPosting().
		SetUid(uid).
		SetValue(value).
		SetValueType(valueType).
		SetLangTag(langTag).
		SetLabel(label).
		SetOp(op).
		SetStartTs(startTs).
		SetCommitTs(commitTs)

	for _, facet := range facets {
		builder.StartFacet().
			SetKey(fbx.BytesToString(facet.Key())).
			EndFacet()
	}

	p := builder.Build()
	require.Equal(t, p.Uid(), uid)
	require.Equal(t, p.ValueBytes(), value)
	require.Equal(t, pb.Posting_ValType(p.ValueType()), valueType)
	require.Equal(t, p.LangTagBytes(), langTag)
	require.Equal(t, fbx.BytesToString(p.Label()), label)
	require.Equal(t, p.FacetsLength(), len(facets))
	for i, expFacet := range facets {
		var gotFacet fb.Facet
		require.True(t, p.Facets(&gotFacet, i))
		require.Equal(t, gotFacet.Key(), expFacet.Key())
	}
	require.Equal(t, p.Op(), op)
	require.Equal(t, p.StartTs(), startTs)
	require.Equal(t, p.CommitTs(), commitTs)
}