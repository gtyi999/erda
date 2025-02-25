// Copyright (c) 2021 Terminus, Inc.
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

package snapshot_test

import (
	"os"
	"testing"

	"github.com/erda-project/erda/pkg/database/sqlparser/snapshot"
)

func TestSampling(t *testing.T) {
	snapshot.Sampling()
}

func TestMaxSampling(t *testing.T) {
	if n := snapshot.MaxSampling(); n != 300 {
		t.Fatal("error")
	}
	os.Setenv("PIPELINE_MIGRATION_SAMPLING_SIZE", "100")
	if n := snapshot.MaxSampling(); n != 100 {
		t.Error("error")
	}
}
