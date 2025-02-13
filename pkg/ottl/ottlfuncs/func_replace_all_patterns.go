// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"
	"fmt"
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const (
	modeKey   = "key"
	modeValue = "value"
)

func ReplaceAllPatterns[K any](target ottl.PMapGetter[K], mode string, regexPattern string, replacement string) (ottl.ExprFunc[K], error) {
	compiledPattern, err := regexp.Compile(regexPattern)
	if err != nil {
		return nil, fmt.Errorf("the regex pattern supplied to replace_all_patterns is not a valid pattern: %w", err)
	}
	if mode != modeValue && mode != modeKey {
		return nil, fmt.Errorf("invalid mode %v, must be either 'key' or 'value'", mode)
	}

	return func(ctx context.Context, tCtx K) (interface{}, error) {
		val, err := target.Get(ctx, tCtx)
		if err != nil {
			return nil, err
		}
		updated := pcommon.NewMap()
		updated.EnsureCapacity(val.Len())
		val.Range(func(key string, originalValue pcommon.Value) bool {
			switch mode {
			case modeValue:
				if compiledPattern.MatchString(originalValue.Str()) {
					updatedString := compiledPattern.ReplaceAllString(originalValue.Str(), replacement)
					updated.PutStr(key, updatedString)
				} else {
					originalValue.CopyTo(updated.PutEmpty(key))
				}
			case modeKey:
				if compiledPattern.MatchString(key) {
					updatedKey := compiledPattern.ReplaceAllLiteralString(key, replacement)
					originalValue.CopyTo(updated.PutEmpty(updatedKey))
				} else {
					originalValue.CopyTo(updated.PutEmpty(key))
				}
			}
			return true
		})
		updated.CopyTo(val)
		return nil, nil
	}, nil
}
