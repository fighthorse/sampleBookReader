// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/fighthorse/sampleBookReader/domain/models/model"
)

func newGoadminSession(db *gorm.DB, opts ...gen.DOOption) goadminSession {
	_goadminSession := goadminSession{}

	_goadminSession.goadminSessionDo.UseDB(db, opts...)
	_goadminSession.goadminSessionDo.UseModel(&model.GoadminSession{})

	tableName := _goadminSession.goadminSessionDo.TableName()
	_goadminSession.ALL = field.NewAsterisk(tableName)
	_goadminSession.ID = field.NewInt32(tableName, "id")
	_goadminSession.Sid = field.NewString(tableName, "sid")
	_goadminSession.Values = field.NewString(tableName, "values")
	_goadminSession.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminSession.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminSession.fillFieldMap()

	return _goadminSession
}

type goadminSession struct {
	goadminSessionDo goadminSessionDo

	ALL       field.Asterisk
	ID        field.Int32
	Sid       field.String
	Values    field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (g goadminSession) Table(newTableName string) *goadminSession {
	g.goadminSessionDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g goadminSession) As(alias string) *goadminSession {
	g.goadminSessionDo.DO = *(g.goadminSessionDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *goadminSession) updateTableName(table string) *goadminSession {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.Sid = field.NewString(table, "sid")
	g.Values = field.NewString(table, "values")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")

	g.fillFieldMap()

	return g
}

func (g *goadminSession) WithContext(ctx context.Context) *goadminSessionDo {
	return g.goadminSessionDo.WithContext(ctx)
}

func (g goadminSession) TableName() string { return g.goadminSessionDo.TableName() }

func (g goadminSession) Alias() string { return g.goadminSessionDo.Alias() }

func (g goadminSession) Columns(cols ...field.Expr) gen.Columns {
	return g.goadminSessionDo.Columns(cols...)
}

func (g *goadminSession) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *goadminSession) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 5)
	g.fieldMap["id"] = g.ID
	g.fieldMap["sid"] = g.Sid
	g.fieldMap["values"] = g.Values
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminSession) clone(db *gorm.DB) goadminSession {
	g.goadminSessionDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g goadminSession) replaceDB(db *gorm.DB) goadminSession {
	g.goadminSessionDo.ReplaceDB(db)
	return g
}

type goadminSessionDo struct{ gen.DO }

func (g goadminSessionDo) Debug() *goadminSessionDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminSessionDo) WithContext(ctx context.Context) *goadminSessionDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminSessionDo) ReadDB() *goadminSessionDo {
	return g.Clauses(dbresolver.Read)
}

func (g goadminSessionDo) WriteDB() *goadminSessionDo {
	return g.Clauses(dbresolver.Write)
}

func (g goadminSessionDo) Session(config *gorm.Session) *goadminSessionDo {
	return g.withDO(g.DO.Session(config))
}

func (g goadminSessionDo) Clauses(conds ...clause.Expression) *goadminSessionDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminSessionDo) Returning(value interface{}, columns ...string) *goadminSessionDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g goadminSessionDo) Not(conds ...gen.Condition) *goadminSessionDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminSessionDo) Or(conds ...gen.Condition) *goadminSessionDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminSessionDo) Select(conds ...field.Expr) *goadminSessionDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminSessionDo) Where(conds ...gen.Condition) *goadminSessionDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminSessionDo) Order(conds ...field.Expr) *goadminSessionDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminSessionDo) Distinct(cols ...field.Expr) *goadminSessionDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminSessionDo) Omit(cols ...field.Expr) *goadminSessionDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminSessionDo) Join(table schema.Tabler, on ...field.Expr) *goadminSessionDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminSessionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *goadminSessionDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminSessionDo) RightJoin(table schema.Tabler, on ...field.Expr) *goadminSessionDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminSessionDo) Group(cols ...field.Expr) *goadminSessionDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminSessionDo) Having(conds ...gen.Condition) *goadminSessionDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminSessionDo) Limit(limit int) *goadminSessionDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminSessionDo) Offset(offset int) *goadminSessionDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminSessionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *goadminSessionDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminSessionDo) Unscoped() *goadminSessionDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminSessionDo) Create(values ...*model.GoadminSession) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminSessionDo) CreateInBatches(values []*model.GoadminSession, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminSessionDo) Save(values ...*model.GoadminSession) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminSessionDo) First() (*model.GoadminSession, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSession), nil
	}
}

func (g goadminSessionDo) Take() (*model.GoadminSession, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSession), nil
	}
}

func (g goadminSessionDo) Last() (*model.GoadminSession, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSession), nil
	}
}

func (g goadminSessionDo) Find() ([]*model.GoadminSession, error) {
	result, err := g.DO.Find()
	return result.([]*model.GoadminSession), err
}

func (g goadminSessionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoadminSession, err error) {
	buf := make([]*model.GoadminSession, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminSessionDo) FindInBatches(result *[]*model.GoadminSession, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminSessionDo) Attrs(attrs ...field.AssignExpr) *goadminSessionDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminSessionDo) Assign(attrs ...field.AssignExpr) *goadminSessionDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminSessionDo) Joins(fields ...field.RelationField) *goadminSessionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g goadminSessionDo) Preload(fields ...field.RelationField) *goadminSessionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g goadminSessionDo) FirstOrInit() (*model.GoadminSession, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSession), nil
	}
}

func (g goadminSessionDo) FirstOrCreate() (*model.GoadminSession, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSession), nil
	}
}

func (g goadminSessionDo) FindByPage(offset int, limit int) (result []*model.GoadminSession, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g goadminSessionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g goadminSessionDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g goadminSessionDo) Delete(models ...*model.GoadminSession) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *goadminSessionDo) withDO(do gen.Dao) *goadminSessionDo {
	g.DO = *do.(*gen.DO)
	return g
}
