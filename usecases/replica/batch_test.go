//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package replica_test

import (
	"sort"
	"strconv"
	"testing"

	"github.com/weaviate/weaviate/usecases/replica"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/weaviate/weaviate/cluster/router/types"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/storobj"
)

// createBatch creates IndexedBatch from xs
func createBatch(xs []*storobj.Object) replica.IndexedBatch {
	var bi replica.IndexedBatch
	bi.Data = xs
	bi.Index = make([]int, len(xs))
	for i := 0; i < len(xs); i++ {
		bi.Index[i] = i
	}
	return bi
}

// cluster data object by shard
func cluster(bi replica.IndexedBatch) []replica.ShardPart {
	index := bi.Index
	data := bi.Data
	sort.Slice(index, func(i, j int) bool {
		return data[index[i]].BelongsToShard < data[index[j]].BelongsToShard
	})
	clusters := make([]replica.ShardPart, 0, 16)
	// partition
	cur := data[index[0]]
	j := 0
	for i := 1; i < len(index); i++ {
		if data[index[i]].BelongsToShard == cur.BelongsToShard {
			continue
		}
		clusters = append(clusters, replica.ShardPart{
			Shard: cur.BelongsToShard,
			Node:  cur.BelongsToNode, Data: data,
			Index: index[j:i],
		})
		j = i
		cur = data[index[j]]

	}
	clusters = append(clusters, replica.ShardPart{
		Shard: cur.BelongsToShard,
		Node:  cur.BelongsToNode, Data: data,
		Index: index[j:],
	})
	return clusters
}

func TestBatchInput(t *testing.T) {
	var (
		N    = 9
		ids  = make([]strfmt.UUID, N)
		data = make([]*storobj.Object, N)
	)
	for i := 0; i < N; i++ {
		uuid := strfmt.UUID(strconv.Itoa(i))
		ids[i] = uuid
		data[i] = objectEx(uuid, 1, "S1", "N1")
	}
	parts := cluster(createBatch(data))
	assert.Len(t, parts, 1)
	assert.Equal(t, parts[0], replica.ShardPart{
		Shard: "S1",
		Node:  "N1",
		Data:  data,
		Index: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
	})
	assert.Equal(t, parts[0].ObjectIDs(), ids)

	data[0].BelongsToShard = "S2"
	data[0].BelongsToNode = "N2"
	data[2].BelongsToShard = "S2"
	data[2].BelongsToNode = "N2"
	data[3].BelongsToShard = "S2"
	data[4].BelongsToNode = "N2"
	data[5].BelongsToShard = "S2"
	data[5].BelongsToNode = "N2"

	parts = cluster(createBatch(data))
	sort.Slice(parts, func(i, j int) bool { return len(parts[i].Index) < len(parts[j].Index) })
	assert.Len(t, parts, 2)
	assert.ElementsMatch(t, parts[0].ObjectIDs(), []strfmt.UUID{ids[0], ids[2], ids[3], ids[5]})
	assert.Equal(t, parts[0].Shard, "S2")
	assert.Equal(t, parts[0].Node, "N2")

	assert.ElementsMatch(t, parts[1].ObjectIDs(), []strfmt.UUID{ids[1], ids[4], ids[6], ids[7], ids[8]})
	assert.Equal(t, parts[1].Shard, "S1")
	assert.Equal(t, parts[1].Node, "N1")
}

func genInputs(node, shard string, updateTime int64, ids []strfmt.UUID) ([]*storobj.Object, []types.RepairResponse) {
	xs := make([]*storobj.Object, len(ids))
	digestR := make([]types.RepairResponse, len(ids))
	for i, id := range ids {
		xs[i] = &storobj.Object{
			Object: models.Object{
				ID:                 id,
				LastUpdateTimeUnix: updateTime,
			},
			BelongsToShard: shard,
			BelongsToNode:  node,
		}
		digestR[i] = types.RepairResponse{ID: ids[i].String(), UpdateTime: updateTime}
	}
	return xs, digestR
}

func setObjectsConsistency(xs []*storobj.Object, isConsistent bool) []*storobj.Object {
	want := make([]*storobj.Object, len(xs))
	for i, x := range xs {
		cp := *x
		cp.IsConsistent = isConsistent
		want[i] = &cp
	}
	return want
}

func objectEx(id strfmt.UUID, lastTime int64, shard, node string) *storobj.Object {
	return &storobj.Object{
		Object: models.Object{
			ID:                 id,
			LastUpdateTimeUnix: lastTime,
		},
		BelongsToShard: shard,
		BelongsToNode:  node,
	}
}
