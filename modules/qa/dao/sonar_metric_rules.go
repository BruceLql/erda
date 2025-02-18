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

package dao

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/qa/dbclient"
)

type QASonarMetricRules struct {
	ID        int64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updatedAt"`

	Description string `gorm:"description" json:"description"`
	ScopeType   string `gorm:"scope_type" json:"scopeType"`
	ScopeID     string `gorm:"scope_id" json:"scopeId"`
	MetricKeyID int64  `gorm:"metric_key_id" json:"metricKeyId"`
	MetricValue string `gorm:"metric_value" json:"metricValue"`
}

func (rule *QASonarMetricRules) ToApi() *apistructs.SonarMetricRuleDto {
	dto := &apistructs.SonarMetricRuleDto{
		ID:          rule.ID,
		CreatedAt:   rule.CreatedAt,
		UpdatedAt:   rule.UpdatedAt,
		Description: rule.Description,
		ScopeType:   rule.ScopeType,
		ScopeID:     rule.ScopeID,
		MetricValue: rule.MetricValue,
		MetricKeyID: rule.MetricKeyID,
	}

	keys := apistructs.SonarMetricKeys[dto.MetricKeyID]

	if keys == nil {
		return dto
	}

	dto.MetricKey = keys.MetricKey
	dto.Operational = apistructs.GetOperationalValue(keys.Operational)
	dto.MetricKeyDesc = keys.MetricKeyDesc
	return dto
}

// TableName QASonar对应的数据库表qa_sonar
func (QASonarMetricRules) TableName() string {
	return "qa_sonar_metric_rules"
}

// PagingTestPlan List test plan
func (client *DBClient) PagingSonarMetricRules(req apistructs.SonarMetricRulesPagingRequest) (*dbclient.Paging, error) {
	var (
		total            int64
		sonarMetricRules []QASonarMetricRules
	)

	sql := client.Where("scope_type = ? and scope_id = ?", req.ScopeType, req.ScopeID)

	if err := sql.Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).
		Order("updated_at desc").Find(&sonarMetricRules).
		Offset(0).Limit(-1).Count(&total).Error; err != nil {
		return nil, err
	}

	return &dbclient.Paging{
		Total: total,
		List:  sonarMetricRules,
	}, nil
}

func (client *DBClient) UpdateSonarMetricRules(updateObj *QASonarMetricRules) (err error) {

	if err = client.Save(updateObj).Error; err != nil {
		return err
	}

	return nil
}

func (client *DBClient) GetSonarMetricRules(ID int64) (*QASonarMetricRules, error) {
	var result QASonarMetricRules
	if err := client.First(&result, "id = ?", ID).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *DBClient) BatchInsertSonarMetricRules(rules []*QASonarMetricRules) (err error) {
	sql := client.DB
	for _, v := range rules {
		sql = sql.Save(v)
	}
	if sql.Error != nil {
		return err
	}

	return nil
}

func (client *DBClient) BatchDeleteSonarMetricRules(rules []QASonarMetricRules) (err error) {
	if rules == nil || len(rules) <= 0 {
		return nil
	}
	var IDs []int64
	for _, v := range rules {
		IDs = append(IDs, v.ID)
	}

	if err := client.Delete(&QASonarMetricRules{}, "scope_type = ? and scope_id = ? and id in (?) ", rules[0].ScopeType, rules[0].ScopeID, IDs).Error; err != nil {
		return err
	}
	return nil
}

func (client *DBClient) ListSonarMetricRules(query *QASonarMetricRules, otherQueryFuncList ...func(sql *gorm.DB) *gorm.DB) (dbRules []QASonarMetricRules, err error) {
	sql := client
	var where = sql.Where("1 = ?", 1)

	if otherQueryFuncList != nil {
		for _, otherQueryFunc := range otherQueryFuncList {
			where = otherQueryFunc(where)
		}
	}
	dbRules = []QASonarMetricRules{}
	if err := where.Find(&dbRules, query).Error; err != nil {
		return nil, err
	}
	return dbRules, nil
}
