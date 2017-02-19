package controllers

import (
	"github.com/go-gorp/gorp"
	r "github.com/revel/revel"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/modules/db/app"
	"portfolio/app/models"
	"fmt"
	"database/sql"

)

var Dbm *gorp.DbMap

func InitDB(){
	fmt.Println("DB initialization")
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	setColumnsizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.Project{}).SetKeys(true, "Id")
	setColumnsizes(t, map[string]int{
		"Id":		20,
		"Title":	50,
		"Git":		200,
		"Images":	100,
		"Video":	200,
		"Deploy":	200,
		"Details":	300,
		"Back":		100,
	})

	Dbm.TraceOn("[gorp]", r.INFO)
	err := Dbm.CreateTablesIfNotExists()
	if err != nil{
		fmt.Println("Error on Create Tables")
	}

	fmt.Println("DB TEST")
	Dbm.Exec("select * from project", nil)

	projects := []*models.Project{
		// id
		// title
		// git address
		// language
		// images
		// video
		// deploy
		// details
		// back
		&models.Project{1,
				"golang-chat-server",
				"https://github.com/jayleec/golang-mobile-server",
				"Go",
				"1-1.png",
				"",
				"http://13.112.157.192/chat",
				"채팅서버, oAuth2연동, AWS deploy",
				"lab.jpg"},
		&models.Project{2,
				"erlang-chat-server",
				"https://github.com/jayleec/erlang-heroku.git",
				"Erlang",
				"2-1.png",
				"",
				"https://erlang-heroku.herokuapp.com",
				"웹소켓 채팅 서버 Heroku",
				"lab.jpg"},
		&models.Project{3,
				"mobile-chat-client",
				"https://github.com/jayleec/simple-chat-client.git",
				"AngularJs (ionic2 framework)",
				"3-1.png",
				"",
				"https://ionic-chat-heroku.herokuapp.com",
				"모바일 채팅 클라이언트, Heroku",
				"lab.jpg"},
		&models.Project{4,
				"software-visualization",
				"https://github.com/sohn126/torch",
				"Python, Javascript",
				"4-1.png",
				"https://www.youtube.com/embed/Og5AEVFfZzw",
				"",
				"소프트웨어 품질 지표 시각화",
				"lab.jpg"},
		&models.Project{5,
				"java-galaga-game",
				"",
				"Java",
				"5-1.png",
				"https://www.youtube.com/embed/krw6liVOkl8",
				"",
				"자바 갈라가 게임",
				"lab.jpg"},
		&models.Project{6,
				"reverse-camera",
				"https://github.com/jayleec/reverseCamera",
				"Objective-c, Swift",
				"6-1.png",
				"https://www.youtube.com/embed/n_EhY7EHevA",
				"",
				"아이폰 비디오 어플리케이션",
				"lab.jpg"},
		&models.Project{7,
				"mean-stack-web-app",
				"https://github.com/jayleec/mean_ionic",
				"AngularJs, Javascript",
				"7-1.png",
				"https://www.youtube.com/embed/iN7daRAPLNg",
				"",
				"MEAN Stack 기반의 웹앱 쇼핑몰 구현",
				"lab.jpg"},

	}
	for _, project := range projects {
		if err := Dbm.Insert(project); err != nil{
			panic(err)
		}
	}
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
