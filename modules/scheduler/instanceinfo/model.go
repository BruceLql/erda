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

package instanceinfo

import (
	"time"

	"github.com/erda-project/erda/pkg/dbengine"
	"github.com/erda-project/erda/pkg/strutil"
)

type ServicePhase string
type InstancePhase string
type PodPhase string

const (
	ServicePhaseHealthy   = "Healthy"
	ServicePhaseUnHealthy = "UnHealthy"

	InstancePhaseRunning   = "Running"
	InstancePhaseHealthy   = "Healthy"
	InstancePhaseUnHealthy = "UnHealthy"
	InstancePhaseDead      = "Dead"

	PodPhasePending   = "Pending"
	PodPhaseRunning   = "Running"
	PodPhaseSucceeded = "Succeeded"
	PodPhaseFailed    = "Failed"
	PodPhaseUnknown   = "Unknown"
)

type ServiceInfo struct {
	dbengine.BaseModel
	Cluster   string
	Namespace string `gorm:"type:varchar(64);index"`
	Name      string `gorm:"type:varchar(64);index"`

	// Information obtained from env
	OrgName         string `gorm:"type:varchar(64);index"`
	OrgID           string `gorm:"type:varchar(64);index"`
	ProjectName     string `gorm:"type:varchar(64);index"`
	ProjectID       string `gorm:"type:varchar(64);index"`
	ApplicationName string
	ApplicationID   string
	RuntimeName     string
	RuntimeID       string
	ServiceName     string
	Workspace       string `gorm:"type:varchar(10)"`
	// addon, stateless-service
	ServiceType string `gorm:"type:varchar(64)"`

	Meta string

	Phase      ServicePhase
	Message    string
	StartedAt  time.Time
	FinishedAt *time.Time
}

func (ServiceInfo) TableName() string {
	return "s_service_info"
}

type InstanceInfo struct {
	dbengine.BaseModel
	Cluster   string `gorm:"type:varchar(64);index"`
	Namespace string `gorm:"type:varchar(64);index"`
	Name      string `gorm:"type:varchar(64);index"`

	// Information obtained from env
	OrgName             string `gorm:"type:varchar(64);index"`
	OrgID               string `gorm:"type:varchar(64);index"`
	ProjectName         string `gorm:"type:varchar(64);index"`
	ProjectID           string `gorm:"type:varchar(64);index"`
	ApplicationName     string
	EdgeApplicationName string
	EdgeSite            string
	ApplicationID       string
	RuntimeName         string
	RuntimeID           string
	ServiceName         string
	Workspace           string `gorm:"type:varchar(10)"`
	// addon, stateless-service
	ServiceType string `gorm:"type:varchar(64)"`
	AddonID     string

	Meta string
	// If it is k8s, the value is "k8s"
	TaskID string `gorm:"type:varchar(150);index"`

	Phase       InstancePhase
	Message     string
	ContainerID string `gorm:"type:varchar(100);index"`
	ContainerIP string
	HostIP      string
	StartedAt   time.Time
	FinishedAt  *time.Time
	ExitCode    int
	CpuOrigin   float64
	MemOrigin   int
	CpuRequest  float64
	MemRequest  int
	CpuLimit    float64
	MemLimit    int
	Image       string
}

func (InstanceInfo) TableName() string {
	return "s_instance_info"
}

func (i InstanceInfo) Metadata(k string) (string, bool) {
	kvs := strutil.Split(i.Meta, ",", true)
	for _, kv := range kvs {
		splitted := strutil.Split(kv, "=", true)
		switch len(splitted) {
		case 2:
			if k == strutil.Trim(splitted[0]) {
				return strutil.Trim(splitted[1]), true
			}
		}
	}
	return "", false
}

type PodInfo struct {
	dbengine.BaseModel
	Cluster   string `gorm:"type:varchar(64);index"`
	Namespace string `gorm:"type:varchar(64);index"`
	Name      string `gorm:"type:varchar(64);index"`

	// Information obtained from env
	OrgName         string `gorm:"type:varchar(64);index"`
	OrgID           string `gorm:"type:varchar(64);index"`
	ProjectName     string `gorm:"type:varchar(64);index"`
	ProjectID       string `gorm:"type:varchar(64);index"`
	ApplicationName string
	ApplicationID   string
	RuntimeName     string
	RuntimeID       string
	ServiceName     string
	Workspace       string `gorm:"type:varchar(10)"`
	// addon, stateless-service
	ServiceType string `gorm:"type:varchar(64)"`
	AddonID     string

	Uid          string `gorm:"index"`
	K8sNamespace string `gorm:"index"`
	PodName      string `gorm:"index"`

	Phase     PodPhase
	Message   string
	PodIP     string
	HostIP    string
	StartedAt *time.Time

	MemRequest int
	MemLimit   int
	CpuRequest float64
	CpuLimit   float64
}

func (PodInfo) TableName() string {
	return "s_pod_info"
}
