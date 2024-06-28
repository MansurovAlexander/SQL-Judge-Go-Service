package components

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func PrepareUser(db *sqlx.DB, dbName string) error {
	if err := initConfig(); err != nil {
		return err
	}
	studentRole := viper.GetString("db.student_role")
	studentPassword := viper.GetString("db.student_password")
	var count int
	var stRole, stPswd, checkRole strings.Builder
	stRole.WriteString("'")
	stRole.WriteString(studentRole)
	stRole.WriteString("'")
	stPswd.WriteString("'")
	stPswd.WriteString(studentPassword)
	stPswd.WriteString("'")
	checkRole.WriteString("SELECT 1 FROM pg_roles WHERE rolname = ")
	checkRole.WriteString(stRole.String())
	if err := db.QueryRow(checkRole.String()).Scan(&count); err != sql.ErrNoRows { 
		return err
	} else {
		_, err := db.Exec(fmt.Sprintf("CREATE USER %s WITH PASSWORD %s", studentRole, stPswd.String())) //, viper.GetString("student_role"), os.Getenv("STUDENT_PASSWORD"))) //viper.GetString("student_role")
		if err != nil {
			return err
		}
	}
	_, err := db.Exec(fmt.Sprintf("GRANT CONNECT ON DATABASE %s TO %s", dbName, studentRole))
	if err != nil {
		return err
	}
	_, err = db.Exec(fmt.Sprintf("GRANT pg_read_all_data TO %s", studentRole))
	if err != nil {
		return err
	}
	return nil
}

func PrepareDB(dbConn *sqlx.DB, dbFileName string) error {
	if err := initConfig(); err != nil {
		return err
	}
	var tempDbName strings.Builder
	tempDbName.WriteString("'")
	tempDbName.WriteString(dbFileName)
	tempDbName.WriteString("'")
	var dbName string
	row := dbConn.QueryRow("SELECT name FROM databases WHERE file_name = $1 AND name != 'Not restored'", strings.ToLower(tempDbName.String()))
	if err := row.Scan(&dbName); err == sql.ErrNoRows {
		if err := initConfig(); err != nil {
			return err
		}
		var path strings.Builder
		dir, _ := os.Getwd()
		path.WriteString(dir)
		path.WriteString("\\db\\test_db\\")
		path.WriteString(dbFileName)
		path.WriteString(".sql")
		bashPath := strings.Replace(path.String(), "\\", "/", -1)
		err = prepareDumpScript(dbFileName, bashPath)
		if err != nil {
			return err
		}
		err = loadSQLFile(viper.GetString("db.username"), viper.GetString("db.host"), viper.GetString("db.port"), bashPath, viper.GetString("container_name"))
		if err != nil {
			return err
		}
		return nil
	}
	err := PrepareUser(dbConn, dbName)
	if err != nil {
		return err
	}
	return nil
}
func prepareDumpScript(dbName, sqlFile string) error {
	create_func := fmt.Sprintf("DROP DATABASE %s;\nCREATE DATABASE %s;\n\\c %s\n", dbName, dbName, dbName)
	file, err := os.OpenFile(sqlFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	var wrBuff strings.Builder
	wrBuff.WriteString(create_func)

	var wr strings.Builder
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text() + "\n")
	}
	wr.WriteString("\nset statement_timeout = 10000\n")
	wrBuff.WriteString(wr.String())
	_, err = file.WriteAt([]byte(wrBuff.String()), 0)
	if err != nil {
		return err
	}
	return nil
}

func loadSQLFile(dbUserName, host, port, sqlFile, containerName string) error {
	var loadFileWin, loadFileLinux strings.Builder
	loadFileWin.WriteString("psql -U ")
	loadFileWin.WriteString(dbUserName)
	loadFileWin.WriteString(" -h ")
	loadFileWin.WriteString(host)
	loadFileWin.WriteString(" -p ")
	loadFileWin.WriteString(port)
	loadFileWin.WriteString(" -f ")
	loadFileWin.WriteString(sqlFile)

	loadFileLinux.WriteString("sudo cat '")
	loadFileLinux.WriteString(sqlFile)
	loadFileLinux.WriteString("' | sudo docker exec -i ")
	loadFileLinux.WriteString(containerName)
	loadFileLinux.WriteString(" psql -U ")
	loadFileLinux.WriteString(dbUserName)

	runtimeSys := runtime.GOOS

	if runtimeSys == "windows" {
		cmd := exec.Command("cmd", "/C", loadFileWin.String())
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		cmd := exec.Command("sh", "-c", loadFileLinux.String())
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
