// go runコマンドで実行するので、パッケージはmain
// $ docker-compose exec app go run tools/migrate.go -exec up
// $ docker-compose exec app go run tools/migrate.go -exec down
package main

import (
	"flag"
	"log"
	"os"
	"simple_pm_api/pkg/core"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	command  = flag.String("exec", "", "-execの引数にはup または downを指定してください")
	force    = flag.Bool("f", false, "force exec fixed sql")
	database = ""
	source   = ""
)

func main() {
	flag.Parse()
	core.LoadEnv(".env")
	if len(*command) < 1 {
		log.Fatal("-exec および その引数を指定してください")
	}

	if *command == "up" {
		source = "file:tools//.//up"
	} else {
		source = "file:tools//.//down"
	}

	migrate, err := migrate.New(source, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("マイグレートインスタンスの生成に失敗しました", err)
	}

	log.Println("Execute command: ", *command)
	version, dirty, err := migrate.Version()
	run(migrate, version, dirty)
}

// run マイグレート実行
func run(migrate *migrate.Migrate, version uint, dirty bool) {
	if dirty && *force {
		log.Println("force=true: force execute current version sql")
		migrate.Force(int(version))
	}

	var err error
	switch *command {
	case "up":
		err = migrate.Up()
	case "down":
		err = migrate.Down()
	case "version":
		return
	default:
		log.Fatal("invalid command: ", *command)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully: ", *command)
	log.Println("Updated version info")
	showResult(migrate.Version())
}

// showResult マイグレート実行結果
func showResult(version uint, dirty bool, err error) {
	log.Println("-------------------")
	log.Printf("version	:%v\n", version)
	log.Printf("dirty	:%v\n", dirty)
	log.Printf("err		:%v\n", err)
	log.Println("-------------------")
}
