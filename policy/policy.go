package policy

import (
	"github.com/dropbox/godropbox/container/set"
	"github.com/dropbox/godropbox/errors"
	"github.com/pritunl/pritunl-zero/agent"
	"github.com/pritunl/pritunl-zero/database"
	"github.com/pritunl/pritunl-zero/errortypes"
	"github.com/pritunl/pritunl-zero/service"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Rule struct {
	Type   string   `bson:"type" json:"type"`
	Values []string `bson:"values" json:"values"`
}

type Policy struct {
	Id       bson.ObjectId    `bson:"_id,omitempty" json:"id"`
	Name     string           `bson:"name" json:"name"`
	Services []bson.ObjectId  `bson:"services" json:"services"`
	Roles    []string         `bson:"roles" json:"roles"`
	Rules    map[string]*Rule `bson:"rules" json:"rules"`
}

func (p *Policy) Validate(db *database.Database) (
	errData *errortypes.ErrorData, err error) {

	return
}

func (p *Policy) ValidateUser(r *http.Request) (
	errData *errortypes.ErrorData, err error) {

	agnt := agent.Parse(r)

	for _, rule := range p.Rules {
		switch rule.Type {
		case OperatingSystem:
			match := false
			for _, value := range rule.Values {
				if value == agnt.OperatingSystem {
					match = true
					break
				}
			}

			if !match {
				errData = &errortypes.ErrorData{
					Error:   "operating_system_policy",
					Message: "Operating system not allowed",
				}
				return
			}
		case Browser:
			match := false
			for _, value := range rule.Values {
				if value == agnt.Browser {
					match = true
					break
				}
			}

			if !match {
				errData = &errortypes.ErrorData{
					Error:   "browser_policy",
					Message: "Browser not allowed",
				}
				return
			}
		}
	}

	return
}

func (p *Policy) Commit(db *database.Database) (err error) {
	coll := db.Policies()

	err = coll.Commit(p.Id, p)
	if err != nil {
		return
	}

	return
}

func (p *Policy) CommitFields(db *database.Database, fields set.Set) (
	err error) {

	coll := db.Policies()

	err = coll.CommitFields(p.Id, p, fields)
	if err != nil {
		return
	}

	return
}

func (p *Policy) Insert(db *database.Database) (err error) {
	coll := db.Policies()

	if p.Id != "" {
		err = &errortypes.DatabaseError{
			errors.New("policy: Policy already exists"),
		}
		return
	}

	err = coll.Insert(p)
	if err != nil {
		err = database.ParseError(err)
		return
	}

	return
}
