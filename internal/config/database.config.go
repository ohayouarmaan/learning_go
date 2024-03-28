package config

import "gopkg.in/mgo.v2"

var session *mgo.Database = nil

func CreateDBInstance(uri string) *mgo.Database {
	if session != nil {
		return session
	} else {
		_session, err := mgo.Dial(uri)
		if err != nil {
			panic(err)
		}
		session = _session.DB("blocks")
		return session
	}
}
