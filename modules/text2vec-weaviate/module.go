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

package modweaviateembed

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/modulecapabilities"
	"github.com/weaviate/weaviate/entities/moduletools"
	"github.com/weaviate/weaviate/modules/text2vec-weaviate/clients"
	"github.com/weaviate/weaviate/modules/text2vec-weaviate/ent"
	"github.com/weaviate/weaviate/usecases/modulecomponents/additional"
	"github.com/weaviate/weaviate/usecases/modulecomponents/batch"
	"github.com/weaviate/weaviate/usecases/modulecomponents/text2vecbase"
)

const Name = "text2vec-weaviate"

var batchSettings = batch.Settings{
	TokenMultiplier:    0,
	MaxTimePerBatch:    float64(10),
	MaxObjectsPerBatch: 200,
	MaxTokensPerBatch:  func(cfg moduletools.ClassConfig) int { return 500000 },
	HasTokenLimit:      false,
	ReturnsRateLimit:   false,
}

type WeaviateEmbedModule struct {
	vectorizer                   text2vecbase.TextVectorizerBatch[[]float32]
	metaProvider                 text2vecbase.MetaProvider
	graphqlProvider              modulecapabilities.GraphQLArguments
	searcher                     modulecapabilities.Searcher[[]float32]
	nearTextTransformer          modulecapabilities.TextTransform
	logger                       logrus.FieldLogger
	additionalPropertiesProvider modulecapabilities.AdditionalProperties
}

func New() *WeaviateEmbedModule {
	return &WeaviateEmbedModule{}
}

func (m *WeaviateEmbedModule) Name() string {
	return Name
}

func (m *WeaviateEmbedModule) Type() modulecapabilities.ModuleType {
	return modulecapabilities.Text2ManyVec
}

func (m *WeaviateEmbedModule) Init(ctx context.Context,
	params moduletools.ModuleInitParams,
) error {
	m.logger = params.GetLogger()

	if err := m.initVectorizer(ctx, params.GetConfig().ModuleHttpClientTimeout, m.logger); err != nil {
		return errors.Wrap(err, "init vectorizer")
	}

	if err := m.initAdditionalPropertiesProvider(); err != nil {
		return errors.Wrap(err, "init additional properties provider")
	}

	return nil
}

func (m *WeaviateEmbedModule) InitExtension(modules []modulecapabilities.Module) error {
	for _, module := range modules {
		if module.Name() == m.Name() {
			continue
		}
		if arg, ok := module.(modulecapabilities.TextTransformers); ok {
			if arg != nil && arg.TextTransformers() != nil {
				m.nearTextTransformer = arg.TextTransformers()["nearText"]
			}
		}
	}

	if err := m.initNearText(); err != nil {
		return errors.Wrap(err, "init graphql provider")
	}
	return nil
}

func (m *WeaviateEmbedModule) initVectorizer(ctx context.Context, timeout time.Duration,
	logger logrus.FieldLogger,
) error {
	apiKey := os.Getenv("WEAVIATE_APIKEY")
	client := clients.New(apiKey, timeout, logger)

	m.vectorizer = text2vecbase.New(client,
		batch.NewBatchVectorizer(client, 50*time.Second, batchSettings, logger, m.Name()),
		batch.ReturnBatchTokenizer(batchSettings.TokenMultiplier, m.Name(), ent.LowerCaseInput),
	)
	m.metaProvider = client

	return nil
}

func (m *WeaviateEmbedModule) initAdditionalPropertiesProvider() error {
	m.additionalPropertiesProvider = additional.NewText2VecProvider()
	return nil
}

func (m *WeaviateEmbedModule) VectorizeObject(ctx context.Context,
	obj *models.Object, cfg moduletools.ClassConfig,
) ([]float32, models.AdditionalProperties, error) {
	return m.vectorizer.Object(ctx, obj, cfg, ent.NewClassSettings(cfg))
}

func (m *WeaviateEmbedModule) VectorizeBatch(ctx context.Context, objs []*models.Object, skipObject []bool, cfg moduletools.ClassConfig) ([][]float32, []models.AdditionalProperties, map[int]error) {
	vecs, errs := m.vectorizer.ObjectBatch(ctx, objs, skipObject, cfg)

	return vecs, nil, errs
}

func (m *WeaviateEmbedModule) MetaInfo() (map[string]interface{}, error) {
	return m.metaProvider.MetaInfo()
}

func (m *WeaviateEmbedModule) VectorizableProperties(cfg moduletools.ClassConfig) (bool, []string, error) {
	return true, nil, nil
}

func (m *WeaviateEmbedModule) VectorizeInput(ctx context.Context,
	input string, cfg moduletools.ClassConfig,
) ([]float32, error) {
	return m.vectorizer.Texts(ctx, []string{input}, cfg)
}

func (m *WeaviateEmbedModule) AdditionalProperties() map[string]modulecapabilities.AdditionalProperty {
	return m.additionalPropertiesProvider.AdditionalProperties()
}

var (
	_ = modulecapabilities.Module(New())
	_ = modulecapabilities.Vectorizer[[]float32](New())
	_ = modulecapabilities.MetaProvider(New())
	_ = modulecapabilities.Searcher[[]float32](New())
	_ = modulecapabilities.GraphQLArguments(New())
	_ = modulecapabilities.InputVectorizer[[]float32](New())
)
