package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// Radicacion describes a radicacion
type Radicacion struct {
	NumeroRadicacion string    `xorm:"varchar(45) primary key not null unique 'numero_radicacion'"`
	FechaRadicacion  string    `xorm:"<-"`
	FechaDocumento   string    `xorm:"<-"`
	Asunto           string    `xorm:"VARCHAR(60) not null 'asunto'"`
	IDRemitente      string    `xorm:"varchar(45) not null 'id_remitente'"`
	PdfidPdf         int16     `xorm:"INT(11) 'pdf_idPDF'"`
	TiempoRespuesta  time.Time `xorm:"TIMESTAMP 'tiempo_respuesta'"`
}

var engine *xorm.Engine

func main() {

	var err error
	engine, err = xorm.NewEngine("mysql", "root:lexco123@tcp(192.168.20.180:3306)/sgdea")
	engine.SetMapper(core.SameMapper{})
	engine.Sync2(new(Radicacion))
	fmt.Print(err)

	//radci := []Radicacion{NumeroRadicacion: "REC540"}
	var radci []Radicacion

	err = engine.Find(&radci)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("radicacionesAll:", radci)

	radicacion := new(Radicacion)
	has, err := engine.Get(radicacion)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nradicacionGET:", has, radicacion)

	var radicacion1 []Radicacion
	err1 := engine.Table("radicacion").Where("numero_radicacion = ?", "REC540").Find(&radicacion1)
	if err1 != nil {
		fmt.Println("errorDB", err1)
		return
	}

	fmt.Println("\nradicacionWhereGET:", radicacion1)
}
