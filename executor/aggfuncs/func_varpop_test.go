// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aggfuncs_test

import (
	. "github.com/pingcap/check"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/mysql"
	"github.com/pingcap/tidb/executor/aggfuncs"
	"github.com/pingcap/tidb/types"
	"github.com/pingcap/tidb/util/set"
)

func (s *testSuite) TestMergePartialResult4Varpop(c *C) {
	tests := []aggTest{
		buildAggTester(ast.AggFuncVarPop, mysql.TypeDouble, 5, types.NewFloat64Datum(float64(2)), types.NewFloat64Datum(float64(2)/float64(3)), types.NewFloat64Datum(float64(59)/float64(8)-float64(19*19)/float64(8*8))),
	}
	for _, test := range tests {
		s.testMergePartialResult(c, test)
	}
}

func (s *testSuite) TestVarpop(c *C) {
	tests := []aggTest{
		buildAggTester(ast.AggFuncVarPop, mysql.TypeDouble, 5, nil, types.NewFloat64Datum(float64(2))),
	}
	for _, test := range tests {
		s.testAggFunc(c, test)
	}
}

func (s *testSuite) TestMemVarpop(c *C) {
	tests := []aggMemTest{
		buildAggMemTester(ast.AggFuncVarPop, mysql.TypeDouble, 5,
			aggfuncs.DefPartialResult4VarPopFloat64Size, defaultUpdateMemDeltaGens, false),
		buildAggMemTester(ast.AggFuncVarPop, mysql.TypeDouble, 5,
			aggfuncs.DefPartialResult4VarPopDistinctFloat64Size+set.DefFloat64SetBucketMemoryUsage, distinctUpdateMemDeltaGens, true),
	}
	for _, test := range tests {
		s.testAggMemFunc(c, test)
	}
}
