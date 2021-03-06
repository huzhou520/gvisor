// Copyright 2018 Google LLC
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

package cmd

import (
	"io/ioutil"
	"testing"

	"gvisor.googlesource.com/gvisor/runsc/boot"
)

func TestNotFound(t *testing.T) {
	ids := []string{"123"}
	dir, err := ioutil.TempDir("", "metadata")
	if err != nil {
		t.Fatalf("error creating dir: %v", err)
	}
	conf := &boot.Config{RootDir: dir}

	d := Delete{}
	if err := d.execute(ids, conf); err == nil {
		t.Error("Deleting non-existent container should have failed")
	}

	d = Delete{force: true}
	if err := d.execute(ids, conf); err != nil {
		t.Errorf("Deleting non-existent container with --force should NOT have failed: %v", err)
	}
}
