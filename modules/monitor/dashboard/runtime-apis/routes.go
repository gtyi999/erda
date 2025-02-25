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

package runtimeapis

import (
	"github.com/erda-project/erda-infra/providers/httpserver"
	"github.com/erda-project/erda/modules/monitor/common"
	"github.com/erda-project/erda/modules/monitor/common/permission"
)

func (p *provider) intRoutes(routes httpserver.Router) error {
	// metric for runtime view
	checkApplication := permission.QueryValue("filter_application_id")
	routes.GET("/api/runtime/metrics/:scope", p.metricq.HandleV1, permission.Intercepter(
		permission.ScopeApp, checkApplication,
		common.ResourceRuntime, permission.ActionGet,
	))
	routes.GET("/api/runtime/metrics/:scope/:aggregate", p.metricq.HandleV1, permission.Intercepter(
		permission.ScopeApp, checkApplication,
		common.ResourceRuntime, permission.ActionGet,
	))
	routes.GET("/api/runtime/metrics/query", p.metricq.Handle, permission.Intercepter(
		permission.ScopeApp, checkApplication,
		common.ResourceRuntime, permission.ActionGet,
	))
	routes.POST("/api/runtime/metrics/query", p.metricq.Handle, permission.Intercepter(
		permission.ScopeApp, checkApplication,
		common.ResourceRuntime, permission.ActionGet,
	))
	return nil
}
