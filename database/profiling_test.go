/*
 * Copyright 2018 Xiaomi, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package database

import (
	"testing"

	"github.com/BruceDu521/soar/common"

	"github.com/kr/pretty"
)

func TestProfiling(t *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	rows, err := connTest.Profiling("select 1")
	if err != nil {
		t.Error(err)
	}
	pretty.Println(rows)
	_, err = connTest.Profiling("delete from film")
	if err == nil {
		t.Error(err)
	}
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}

func TestFormatProfiling(t *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	res, err := connTest.Profiling("select 1")
	if err != nil {
		t.Error(err)
	}
	pretty.Println(FormatProfiling(res))
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}
