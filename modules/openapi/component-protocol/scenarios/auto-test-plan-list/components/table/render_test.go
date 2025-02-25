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

package table

import (
	"context"
	"reflect"
	"testing"

	"bou.ke/monkey"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	protocol "github.com/erda-project/erda/modules/openapi/component-protocol"
	"github.com/erda-project/erda/pkg/i18n"
)

func TestTestPlanManageTable_Render(t *testing.T) {
	type args struct {
		ctx      context.Context
		c        *apistructs.Component
		scenario apistructs.ComponentProtocolScenario
		event    apistructs.ComponentEvent
		gs       *apistructs.GlobalStateData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_default_render",
			args: args{
				ctx:      context.Background(),
				c:        &apistructs.Component{},
				scenario: apistructs.ComponentProtocolScenario{},
				event: apistructs.ComponentEvent{
					Operation: apistructs.InitializeOperation,
				},
				gs: &apistructs.GlobalStateData{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bdl = protocol.ContextBundle{Locale: "zh"}
			dl := bundle.New(bundle.WithI18nLoader(&i18n.LocaleResourceLoader{}))
			m := monkey.PatchInstanceMethod(reflect.TypeOf(dl), "GetLocale",
				func(bdl *bundle.Bundle, local ...string) *i18n.LocaleResource {
					return &i18n.LocaleResource{}
				})
			defer m.Unpatch()
			monkey.PatchInstanceMethod(reflect.TypeOf(dl), "PagingTestPlansV2", func(b *bundle.Bundle, req apistructs.TestPlanV2PagingRequest) (*apistructs.TestPlanV2PagingResponseData, error) {
				return &apistructs.TestPlanV2PagingResponseData{}, nil
			})
			bdl.Bdl = dl
			bdl.InParams = map[string]interface{}{"projectId": 1}
			tt.args.ctx = context.WithValue(tt.args.ctx, protocol.GlobalInnerKeyCtxBundle.String(), bdl)

			tpmt := &TestPlanManageTable{}
			if err := tpmt.Render(tt.args.ctx, tt.args.c, tt.args.scenario, tt.args.event, tt.args.gs); (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetIterations(t *testing.T) {
	tt := []struct {
		state map[string]interface{}
		want  []uint64
	}{
		{
			state: nil, want: nil,
		},
		{
			state: map[string]interface{}{"foo": "bar"}, want: nil,
		},
		{
			state: map[string]interface{}{"iteration": "bar"}, want: nil,
		},
		{
			state: map[string]interface{}{"iteration": []string{"1", "2"}}, want: nil,
		},
		{
			state: map[string]interface{}{"iteration": []uint64{1, 2}}, want: []uint64{1, 2},
		},
		{
			state: map[string]interface{}{"iteration": []interface{}{float64(1), float64(2)}}, want: []uint64{1, 2},
		},
	}
	for _, v := range tt {
		if !reflect.DeepEqual(v.want, getIterations(v.state)) {
			t.Error("fail")
		}
	}
}
