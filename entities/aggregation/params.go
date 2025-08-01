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

package aggregation

import (
	"encoding/json"
	"fmt"

	"github.com/weaviate/weaviate/entities/filters"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
	"github.com/weaviate/weaviate/entities/searchparams"
)

type Params struct {
	Filters          *filters.LocalFilter       `json:"filters"`
	ClassName        schema.ClassName           `json:"className"`
	Properties       []ParamProperty            `json:"properties"`
	GroupBy          *filters.Path              `json:"groupBy"`
	IncludeMetaCount bool                       `json:"includeMetaCount"`
	Limit            *int                       `json:"limit"`
	ObjectLimit      *int                       `json:"objectLimit"`
	SearchVector     models.Vector              `json:"searchVector"`
	TargetVector     string                     `json:"targetVector"`
	Certainty        float64                    `json:"certainty"`
	Tenant           string                     `json:"tenant"`
	ModuleParams     map[string]interface{}     `json:"moduleParams"`
	NearVector       *searchparams.NearVector   `json:"nearVector"`
	NearObject       *searchparams.NearObject   `json:"nearObject"`
	Hybrid           *searchparams.HybridSearch `json:"hybrid"`
}

func (p *Params) UnmarshalJSON(data []byte) error {
	type alias Params
	aux := &struct {
		SearchVector json.RawMessage `json:"searchVector"`
		*alias
	}{
		alias: (*alias)(p),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// SearchVector is nil
	if aux.SearchVector == nil {
		return nil
	}

	// Try unmarshaling as []float32
	var vector []float32
	if err := json.Unmarshal(aux.SearchVector, &vector); err == nil {
		if len(vector) > 0 {
			p.SearchVector = vector
		}
		return nil
	}

	// Try unmarshaling as [][]float32
	var multiVector [][]float32
	if err := json.Unmarshal(aux.SearchVector, &multiVector); err == nil {
		if len(multiVector) > 0 {
			p.SearchVector = multiVector
		}
		return nil
	}
	return fmt.Errorf("searchVector: cannot unmarshal into either []float32 or [][]float32: %v", aux.SearchVector)
}

type ParamProperty struct {
	Name        schema.PropertyName `json:"name"`
	Aggregators []Aggregator        `json:"aggregators"`
}

type Aggregator struct {
	Type  string `json:"type"`
	Limit *int   `json:"limit"` // used on TopOccurrence Agg
}

func (a Aggregator) String() string {
	return a.Type
}

// Aggregators used in every prop
var (
	CountAggregator = Aggregator{Type: "count"}
	TypeAggregator  = Aggregator{Type: "type"}
)

// Aggregators used in numerical props
var (
	SumAggregator     = Aggregator{Type: "sum"}
	MeanAggregator    = Aggregator{Type: "mean"}
	ModeAggregator    = Aggregator{Type: "mode"}
	MedianAggregator  = Aggregator{Type: "median"}
	MaximumAggregator = Aggregator{Type: "maximum"}
	MinimumAggregator = Aggregator{Type: "minimum"}
)

// Aggregators used in boolean props
var (
	TotalTrueAggregator       = Aggregator{Type: "totalTrue"}
	PercentageTrueAggregator  = Aggregator{Type: "percentageTrue"}
	TotalFalseAggregator      = Aggregator{Type: "totalFalse"}
	PercentageFalseAggregator = Aggregator{Type: "percentageFalse"}
)

const TopOccurrencesType = "topOccurrences"

// NewTopOccurrencesAggregator creates a TopOccurrencesAggregator, we cannot
// use a singleton for this as the desired limit can be different each time
func NewTopOccurrencesAggregator(limit *int) Aggregator {
	return Aggregator{Type: TopOccurrencesType, Limit: limit}
}

// Aggregators used in ref props
var (
	PointingToAggregator = Aggregator{Type: "pointingTo"}
)

func ParseAggregatorProp(name string) (Aggregator, error) {
	switch name {
	// common
	case CountAggregator.String():
		return CountAggregator, nil
	case TypeAggregator.String():
		return TypeAggregator, nil

	// numerical
	case MeanAggregator.String():
		return MeanAggregator, nil
	case MedianAggregator.String():
		return MedianAggregator, nil
	case ModeAggregator.String():
		return ModeAggregator, nil
	case MaximumAggregator.String():
		return MaximumAggregator, nil
	case MinimumAggregator.String():
		return MinimumAggregator, nil
	case SumAggregator.String():
		return SumAggregator, nil

	// boolean
	case TotalTrueAggregator.String():
		return TotalTrueAggregator, nil
	case TotalFalseAggregator.String():
		return TotalFalseAggregator, nil
	case PercentageTrueAggregator.String():
		return PercentageTrueAggregator, nil
	case PercentageFalseAggregator.String():
		return PercentageFalseAggregator, nil

	// string/text
	case TopOccurrencesType:
		return NewTopOccurrencesAggregator(ptInt(5)), nil // default to limit 5, can be overwritten

	// ref
	case PointingToAggregator.String():
		return PointingToAggregator, nil

	default:
		return Aggregator{}, fmt.Errorf("unrecognized aggregator prop '%s'", name)
	}
}

func ptInt(in int) *int {
	return &in
}
