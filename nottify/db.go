package nottify

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
)

type Db struct {
	hostname string
	username string
	database string
	password string
	conn     *sql.DB
	conf     *Config
}

// NewDb provisions an instance of Db
func NewDb(c *Config) *Db {
	d := new(Db)

	d.hostname, _ = c.Get([]string{"database", "hostname"})
	d.username, _ = c.Get([]string{"database", "user"})
	d.password, _ = c.Get([]string{"database", "password"})
	d.database, _ = c.Get([]string{"database", "database"})

	if err := d.Connect(); err != nil {
		// needs elegance
		panic(fmt.Sprintf("Unable to connect to Database [%s]", err.Error()))
	}

	return d
}

// Connect will open the connection to the server
func (d *Db) Connect() error {
	d.conn, err = sql.Open("postgres", fmt.Sprintf("user='%s' dbname='%s' password='%s' host='%s' sslmode=require", d.username, d.database, d.password, d.hostname))
	if err != nil {
		return err
	}

	return nil
}

// RunFile will use Postgres functionality to execute the full file
func (d *Db) RunFile(f string) error {
	// put the full path together
	sqlFile := path.Join(d.conf.GetBaseDir(), "schema", f)
	if _, err := os.Stat(sqlFile); err != nil {
		return errors.New("Cannot find SQL file")
	}

	// run the file
	data, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	reg := regexp.MustCompile(`(?m)^-- .+$`)
	tbls := strings.Split(string(data), ";--end")
	for i := 0; i < len(tbls); i++ {
		q := strings.TrimSpace(reg.ReplaceAllString(tbls[i], ""))

		if len(q) < 4 {
			continue
		}

		_, err = d.conn.Exec(q)
		if err != nil {
			return err
		}
	}

	return nil
}
