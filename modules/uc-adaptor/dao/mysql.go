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

// Package dbclient 定义数据库操作的方法, orm 等。
package dao

import (
	"github.com/erda-project/erda/modules/uc-adaptor/conf"
	"github.com/erda-project/erda/pkg/dbengine"
)

const BULK_INSERT_CHUNK_SIZE = 3000

type DBClient struct {
	*dbengine.DBEngine
}

func Open() (*DBClient, error) {
	engine, err := dbengine.Open()
	if err != nil {
		return nil, err
	}
	if conf.Debug() {
		engine.LogMode(true)
	}
	db := DBClient{DBEngine: engine}
	// custom init
	if err := db.initOpts(); err != nil {
		return nil, err
	}
	return &db, nil
}

func (db *DBClient) Close() error {
	if db == nil || db.DBEngine == nil {
		return nil
	}
	return db.DBEngine.Close()
}

// TODO: 自定义初始化内容
func (db *DBClient) initOpts() error {
	return nil
}
