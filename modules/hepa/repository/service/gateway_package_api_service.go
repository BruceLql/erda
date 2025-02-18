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

package service

import (
	"strings"

	"github.com/erda-project/erda/modules/hepa/common"
	. "github.com/erda-project/erda/modules/hepa/common/vars"
	"github.com/erda-project/erda/modules/hepa/repository/orm"
	"github.com/pkg/errors"
	"github.com/xormplus/xorm"
)

type GatewayPackageApiServiceImpl struct {
	engine *orm.OrmEngine
	*SessionHelper
	executor xorm.Interface
}

func NewGatewayPackageApiServiceImpl() (*GatewayPackageApiServiceImpl, error) {
	engine, error := orm.GetSingleton()
	if error != nil {
		return nil, errors.Wrap(error, "new GatewayPackageApiServiceImpl failed")
	}
	return &GatewayPackageApiServiceImpl{
		engine:   engine,
		executor: engine,
	}, nil
}

func (impl *GatewayPackageApiServiceImpl) NewSession(helper ...*SessionHelper) (GatewayPackageApiService, error) {
	var session *SessionHelper
	var err error
	if len(helper) == 0 {
		session, err = NewSessionHelper()
		if err != nil {
			return nil, err
		}
	} else if helper[0] == nil {
		return &GatewayPackageApiServiceImpl{
			engine:   impl.engine,
			executor: impl.engine,
		}, nil
	} else {
		session = helper[0]
	}
	return &GatewayPackageApiServiceImpl{
		engine:        impl.engine,
		executor:      session.session,
		SessionHelper: session,
	}, nil
}

func (impl *GatewayPackageApiServiceImpl) Update(dao *orm.GatewayPackageApi, columns ...string) error {
	if dao == nil {
		return errors.New(ERR_INVALID_ARG)
	}
	_, err := orm.Update(impl.executor, dao, columns...)
	if err != nil {
		return errors.Wrap(err, ERR_SQL_FAIL)
	}
	return nil
}

func (impl *GatewayPackageApiServiceImpl) Insert(dao *orm.GatewayPackageApi) error {
	if dao == nil {
		return errors.New(ERR_INVALID_ARG)
	}
	_, err := orm.Insert(impl.executor, dao)
	if err != nil {
		return errors.Wrap(err, ERR_SQL_FAIL)
	}
	return nil
}

func (impl *GatewayPackageApiServiceImpl) GetByAny(cond *orm.GatewayPackageApi) (*orm.GatewayPackageApi, error) {
	if cond == nil {
		return nil, errors.New(ERR_INVALID_ARG)
	}
	dao := &orm.GatewayPackageApi{}
	bCond, err := orm.BuildConds(impl.engine, cond, cond.GetMustCondCols())
	if err != nil {
		return nil, errors.Wrap(err, "buildConds failed")
	}
	succ, err := orm.GetByAnyI(impl.executor, bCond, dao)
	if err != nil {
		return nil, errors.Wrap(err, ERR_SQL_FAIL)
	}
	if !succ {
		return nil, nil
	}
	return dao, nil
}

func (impl *GatewayPackageApiServiceImpl) GetRawByAny(cond *orm.GatewayPackageApi) (*orm.GatewayPackageApi, error) {
	if cond == nil {
		return nil, errors.New(ERR_INVALID_ARG)
	}
	dao := &orm.GatewayPackageApi{}
	bCond, err := orm.BuildConds(impl.engine, cond, cond.GetMustCondCols())
	if err != nil {
		return nil, errors.Wrap(err, "buildConds failed")
	}
	succ, err := orm.GetRawByAnyI(impl.executor, bCond, dao)
	if err != nil {
		return nil, errors.Wrap(err, ERR_SQL_FAIL)
	}
	if !succ {
		return nil, nil
	}
	return dao, nil
}

func (impl *GatewayPackageApiServiceImpl) SelectByAny(cond *orm.GatewayPackageApi) ([]orm.GatewayPackageApi, error) {
	var result []orm.GatewayPackageApi
	if cond == nil {
		return result, errors.New(ERR_INVALID_ARG)
	}
	bCond, err := orm.BuildConds(impl.engine, cond, cond.GetMustCondCols())
	if err != nil {
		return result, errors.Wrap(err, ERR_SQL_FAIL)
	}
	err = orm.SelectByAnyI(impl.executor, bCond, &result)
	if err != nil {
		return result, errors.Wrap(err, ERR_SQL_FAIL)
	}
	return result, nil
}

func (impl *GatewayPackageApiServiceImpl) Get(id string) (*orm.GatewayPackageApi, error) {
	if len(id) == 0 {
		return nil, errors.New(ERR_INVALID_ARG)
	}
	dao := &orm.GatewayPackageApi{}
	succ, err := orm.Get(impl.executor, dao, "id = ?", id)
	if err != nil {
		return nil, errors.Wrap(err, ERR_SQL_FAIL)
	}
	if !succ {
		return nil, nil
	}
	return dao, nil
}

func (impl *GatewayPackageApiServiceImpl) Delete(id string) error {
	if len(id) == 0 {
		return errors.New(ERR_INVALID_ARG)
	}
	_, err := orm.Delete(impl.executor, &orm.GatewayPackageApi{}, "id = ?", id)
	if err != nil {
		return errors.Wrap(err, ERR_SQL_FAIL)
	}
	return nil
}

func (impl *GatewayPackageApiServiceImpl) DeleteByPackageDiceApi(packageId, diceApiId string) error {
	if packageId == "" || diceApiId == "" {
		return errors.New(ERR_INVALID_ARG)
	}
	_, err := orm.Delete(impl.executor, &orm.GatewayPackageApi{}, "package_id = ? and dice_api_id",
		packageId, diceApiId)
	if err != nil {
		return errors.Wrap(err, ERR_SQL_FAIL)
	}
	return nil
}

func (impl *GatewayPackageApiServiceImpl) DeleteByPackageId(packageId string) error {
	if packageId == "" {
		return errors.New(ERR_INVALID_ARG)
	}
	_, err := orm.Delete(impl.executor, &orm.GatewayPackageApi{}, "package_id = ?", packageId)
	if err != nil {
		return errors.Wrap(err, ERR_SQL_FAIL)
	}
	return nil
}

func (impl *GatewayPackageApiServiceImpl) GetPage(options []orm.SelectOption, page *common.Page) (*common.PageQuery, error) {
	total, err := impl.Count(options)
	if err != nil {
		return nil, errors.Wrap(err, "get total count failed")
	}
	page.SetTotalNum(total)
	if total == 0 {
		return &common.PageQuery{Result: []orm.GatewayPackageApi{}, Page: page}, nil
	}
	var result []orm.GatewayPackageApi
	err = orm.SelectPageWithOption(options, impl.executor, &result, page)
	if err != nil {
		return nil, errors.Wrap(err, ERR_SQL_FAIL)
	}
	return &common.PageQuery{Result: result, Page: page}, nil
}

func (impl *GatewayPackageApiServiceImpl) Count(options []orm.SelectOption) (int64, error) {
	count, err := orm.CountWithOption(options, impl.executor, &orm.GatewayPackageApi{})
	if err != nil {
		return 0, errors.Wrap(err, ERR_SQL_FAIL)
	}
	return count, nil
}

func (impl *GatewayPackageApiServiceImpl) CheckUnique(dao *orm.GatewayPackageApi) (bool, error) {
	if dao == nil {
		return false, errors.New(ERR_INVALID_ARG)
	}
	shadowExist := &orm.GatewayPackageApi{}
	succ, err := orm.Get(impl.executor, shadowExist, "package_id = ? and origin = ? and api_path like ?",
		dao.PackageId, "shadow", strings.ReplaceAll(dao.ApiPath, `_`, `\_`)+"%")
	if err != nil {
		return false, errors.Wrap(err, ERR_SQL_FAIL)
	}
	if succ {
		return false, nil
	}
	cond := &orm.GatewayPackageApi{
		PackageId: dao.PackageId,
		ApiPath:   dao.ApiPath,
		Method:    dao.Method,
	}
	cond.SetMustCondCols("package_id", "api_path", "method")
	exist, err := impl.GetByAny(cond)
	if err != nil {
		return false, errors.Wrap(err, ERR_SQL_FAIL)
	}
	unique := (exist == nil) || (exist.Id == dao.Id)
	if !unique {
		return false, nil
	}
	exist, err = impl.GetByAny(&orm.GatewayPackageApi{
		PackageId: dao.PackageId,
		ApiPath:   dao.ApiPath,
	})
	if err != nil {
		return false, errors.WithStack(err)
	}
	if exist == nil || exist.Id == dao.Id {
		return true, nil
	}
	if exist.RedirectType == "service" || dao.RedirectType == "service" {
		return false, nil
	}
	return true, nil
}
