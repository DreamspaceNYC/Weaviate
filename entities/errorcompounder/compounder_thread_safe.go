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

package errorcompounder

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type SafeErrorCompounder struct {
	sync.Mutex
	errors []error
}

func NewSafe() *SafeErrorCompounder {
	return &SafeErrorCompounder{}
}

func (ec *SafeErrorCompounder) Add(err error) {
	ec.Lock()
	defer ec.Unlock()
	if err != nil {
		ec.errors = append(ec.errors, err)
	}
}

func (ec *SafeErrorCompounder) Addf(msg string, args ...interface{}) {
	ec.Lock()
	defer ec.Unlock()
	ec.errors = append(ec.errors, fmt.Errorf(msg, args...))
}

func (ec *SafeErrorCompounder) ToError() error {
	ec.Lock()
	defer ec.Unlock()
	if len(ec.errors) == 0 {
		return nil
	}

	var msg strings.Builder
	for i, err := range ec.errors {
		if i != 0 {
			msg.WriteString(", ")
		}

		msg.WriteString(err.Error())
	}

	return errors.New(msg.String())
}

func (ec *SafeErrorCompounder) First() error {
	if len(ec.errors) == 0 {
		return nil
	}
	return ec.errors[0]
}
