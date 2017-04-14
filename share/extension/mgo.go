package extension

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	MGO_HOST     = "127.0.0.1:27017"
	MGO_USER     = "root"
	MGO_PWD      = "root"
	MGO_DATABASE = "blog"
)

type MgoClient struct {
	session *mgo.Session
}

var M *MgoClient = nil

func Init() {
	M = &MgoClient{
		session: nil,
	}

	mconf := fmt.Sprintf("mongodb://%s:%s@%s/%s", MGO_USER, MGO_PWD, MGO_HOST, MGO_DATABASE)
	mgosession, err := mgo.Dial(mconf)

	if err != nil {
		fmt.Printf("Failed to connect MongoDB with hosts: %s, err: %v", mconf, err)
		panic(err)
	} else {
		M.session = mgosession
	}
}

// CopySession get a copy of the Mgo Session; in the case there are some blocks, we need to Copy session to exceute work
// Make sure to call the session.Close()
func (mc *MgoClient) CopySession() *mgo.Session {
	return mc.session.Copy()
}

// Session clone the Mgo session, reuse the same socket
// Make sure to call the session.Close()
func (mc *MgoClient) Session() *mgo.Session {
	return mc.session.Clone()
}

// DB get session's DB
func (mc *MgoClient) DB() *mgo.Database {
	return mc.session.DB(MGO_DATABASE)
}

// C copy session and get the copy-session's DB Collection
func (mc *MgoClient) C(name string) (*mgo.Session, *mgo.Collection) {
	session := mc.CopySession()
	return session, session.DB(MGO_DATABASE).C(name)
}

// SC get session's DB Collection
func (mc *MgoClient) SC(session *mgo.Session, name string) *mgo.Collection {
	return session.DB(MGO_DATABASE).C(name)
}

func handleDatabaseError(err error, funcName string, args ...interface{}) error {
	if err != nil && err != mgo.ErrNotFound {
		fmt.Printf("Error happend during accessing database, error: %v, function: %s, arguments: %v", err, funcName, args)
		panic(err)
	}
	return err
}

// PagingCondition mongo query condition and query pagination
type PagingCondition struct {
	Selector  bson.M
	PageIndex int
	PageSize  int
	Sortor    []string
}

// DatabaseRepository interface of operate mongodb struct
type DatabaseRepository interface {
	NewSession() interface{}
	CloseSession(session interface{})
	FindOne(collectionName string, selector bson.M, result interface{}) error
	FindAll(collectionName string, selector bson.M, sortor []string, limit int, result interface{}) error
	FindByPK(collectionName string, id interface{}, result interface{}) error
	FindByPagination(collectionName string, page PagingCondition, result interface{}) (int, error)
	UpdateOne(collectionName string, selector bson.M, updator bson.M) error
	UpdateAll(collectionName string, selector bson.M, updator bson.M) (int, error)
	Inset(collectionName string, docs ...interface{}) error
	RemoveOne(collectionName string, seletor bson.M) error
	RemoveAll(collectionName string, seletor bson.M) (int, error)

	FindOneSC(session *mgo.Session, collectionName string, selector bson.M, result interface{}) error
	FindAllSC(session *mgo.Session, collectionName string, selector bson.M, sortor []string, limit int, result interface{}) error
	FindByPKSC(session *mgo.Session, collectionName string, id interface{}, result interface{}) error
	FindByPaginationSC(session *mgo.Session, collectionName string, page PagingCondition, result interface{}) (int, error)
	UpdateOneSC(session *mgo.Session, collectionName string, selector bson.M, updator bson.M) error
	UpdateAllSC(session *mgo.Session, collectionName string, selector bson.M, updator bson.M) (int, error)
	InsetSC(session *mgo.Session, collectionName string, docs ...interface{}) error
	RemoveOneSC(session *mgo.Session, collectionName string, seletor bson.M) error
	RemoveAllSC(session *mgo.Session, collectionName string, seletor bson.M) (int, error)
}

type mongoDBRepository struct{}

// MongoDBRepository mongodb repo
var MongoDBRepository DatabaseRepository = &mongoDBRepository{}

func (repo *mongoDBRepository) NewSession() interface{} {
	return M.CopySession()
}

func (repo *mongoDBRepository) CloseSession(session interface{}) {
	session.(*mgo.Session).Close()
}

func (repo *mongoDBRepository) FindOne(collectionName string, selector bson.M, result interface{}) error {
	session := repo.NewSession().(*mgo.Session)
	defer repo.CloseSession(session)

	return repo.FindOneSC(session, collectionName, selector, result)
}

func (repo *mongoDBRepository) FindAll(collectionName string, selector bson.M, sortor []string, limit int, result interface{}) error {
	session := repo.NewSession().(*mgo.Session)
	defer repo.CloseSession(session)

	return repo.FindAllSC(session, collectionName, selector, sortor, limit, result)
}

func (repo *mongoDBRepository) FindByPK(collectionName string, id interface{}, result interface{}) error {
	session := repo.NewSession().(*mgo.Session)
	defer repo.CloseSession(session)

	return repo.FindByPKSC(session, collectionName, id, result)
}

func (repo *mongoDBRepository) FindByPagination(collectionName string, page PagingCondition, result interface{}) (int, error) {
	session := repo.NewSession().(*mgo.Session)
	defer repo.CloseSession(session)

	return repo.FindByPaginationSC(session, collectionName, page, result)
}

func (repo *mongoDBRepository) UpdateOne(collectionName string, selector bson.M, updator bson.M) error {
	session := repo.NewSession().(*mgo.Session)
	defer repo.CloseSession(session)

	return repo.UpdateOneSC(session, collectionName, selector, updator)
}

func (repo *mongoDBRepository) UpdateAll(collectionName string, selector bson.M, updator bson.M) (int, error) {
	session := repo.NewSession().(*mgo.Session)
	defer repo.CloseSession(session)

	return repo.UpdateAllSC(session, collectionName, selector, updator)
}

func (repo *mongoDBRepository) Inset(collectionName string, docs ...interface{}) error {
	session := repo.NewSession().(*mgo.Session)
	defer repo.CloseSession(session)

	return repo.InsetSC(session, collectionName, docs...)
}

func (repo *mongoDBRepository) RemoveOne(collectionName string, seletor bson.M) error {
	session := repo.NewSession().(*mgo.Session)
	defer repo.CloseSession(session)

	return repo.RemoveOneSC(session, collectionName, seletor)
}

func (repo *mongoDBRepository) RemoveAll(collectionName string, seletor bson.M) (int, error) {
	session := repo.NewSession().(*mgo.Session)
	defer repo.CloseSession(session)

	return repo.RemoveAllSC(session, collectionName, seletor)
}

func (repo *mongoDBRepository) FindOneSC(session *mgo.Session, collectionName string, selector bson.M, result interface{}) error {
	collection := M.SC(session, collectionName)
	q := collection.Find(selector)
	err := q.One(result)

	return handleDatabaseError(err, "FindOneSC", collectionName, selector)
}

func (repo *mongoDBRepository) FindAllSC(session *mgo.Session, collectionName string, selector bson.M, sortor []string, limit int, result interface{}) error {
	collection := M.SC(session, collectionName)
	q := collection.Find(selector)

	if len(sortor) > 0 {
		q = q.Sort(sortor...)
	}

	if limit > 0 {
		q = q.Limit(limit)
	}

	err := q.All(result)

	return handleDatabaseError(err, "FindAllSC", collectionName, selector, sortor, limit)
}

func (repo *mongoDBRepository) FindByPKSC(session *mgo.Session, collectionName string, id interface{}, result interface{}) error {
	collection := M.SC(session, collectionName)
	q := collection.FindId(id)

	err := q.One(result)

	return handleDatabaseError(err, "FindByPKSC", collectionName, id)
}

func (repo *mongoDBRepository) FindByPaginationSC(session *mgo.Session, collectionName string, page PagingCondition, result interface{}) (int, error) {
	colletion := M.SC(session, collectionName)

	q := colletion.Find(page.Selector)
	total, err := q.Count()

	if err != nil {
		return total, handleDatabaseError(err, "FindByPaginationSC", collectionName, page)
	}

	if len(page.Sortor) > 0 {
		q = q.Sort(page.Sortor...)
	}

	q = q.Skip((page.PageIndex - 1) * page.PageSize).Limit(page.PageSize)

	err = q.All(result)

	return total, handleDatabaseError(err, "FindByPaginationSC", collectionName, page)
}

func (repo *mongoDBRepository) UpdateOneSC(session *mgo.Session, collectionName string, selector bson.M, updator bson.M) error {
	collection := M.SC(session, collectionName)

	err := collection.Update(selector, updator)

	return handleDatabaseError(err, "UpdateOneSC", collectionName, selector, updator)
}

func (repo *mongoDBRepository) UpdateAllSC(session *mgo.Session, collectionName string, selector bson.M, updator bson.M) (int, error) {
	collection := M.SC(session, collectionName)
	info, err := collection.UpdateAll(selector, updator)

	if err == nil {
		return info.Matched, nil
	}

	return 0, handleDatabaseError(err, "UpdateAllSC", collectionName, selector, updator)
}

func (repo *mongoDBRepository) InsetSC(session *mgo.Session, collectionName string, docs ...interface{}) error {
	collection := M.SC(session, collectionName)

	err := collection.Insert(docs...)

	return handleDatabaseError(err, "InsetSC", collectionName, docs)
}

func (repo *mongoDBRepository) RemoveOneSC(session *mgo.Session, collectionName string, seletor bson.M) error {
	collection := M.SC(session, collectionName)

	err := collection.Remove(seletor)

	return handleDatabaseError(err, "RemoveOneSC", collectionName, seletor)
}

func (repo *mongoDBRepository) RemoveAllSC(session *mgo.Session, collectionName string, seletor bson.M) (int, error) {
	collection := M.SC(session, collectionName)
	info, err := collection.RemoveAll(session)

	if err == nil {
		return info.Removed, nil
	}

	return 0, handleDatabaseError(err, "RemoveAllSC", collectionName, seletor)
}
