// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package qa

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
	"github.com/erda-project/erda/modules/openapi/api/spec"
)

var QA_TESTENV_DELETE = apis.ApiSpec{
	Path:         "/api/testenv/<id>",
	BackendPath:  "/api/testenv/<id>",
	Host:         "qa.marathon.l4lb.thisdcos.directory:3033",
	Scheme:       "http",
	Method:       "DELETE",
	CheckLogin:   true,
	CheckToken:   true,
	IsOpenAPI:    true,
	ResponseType: apistructs.APITestEnvDeleteResponse{},
	Doc:          `summary: 更新项目环境变量信息`,
	Audit: func(ctx *spec.AuditContext) error {
		var resp apistructs.APITestEnvDeleteResponse
		if err := ctx.BindResponseData(&resp); err != nil {
			return err
		}
		project, err := ctx.GetProject(resp.Data.EnvID)
		if err != nil {
			return err
		}
		return ctx.CreateAudit(&apistructs.Audit{
			ScopeType:    apistructs.ProjectScope,
			ScopeID:      project.ID,
			TemplateName: apistructs.QaTestEnvDeleteTemplate,
			Context: map[string]interface{}{
				"projectName": project.Name,
				"testEnvName": resp.Data.Name,
			},
		})
	},
}
