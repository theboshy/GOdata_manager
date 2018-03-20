package main

import (
	"github.com/go-xorm/xorm"
	"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"

)

var engine *xorm.Engine

func main() {
	type radicacion struct {
		numeroRadicacion string `xorm:"varchar(45) not null unique 'numero_radicacion'"`
		fechaRadicacion  time.Time `xorm:"TIMESTAMP 'fecha_radicacion'"`
		fechaDocumento string `xorm:"<-"`
		asunto string `xorm:"VARCHAR(60) 'asunto'"`
		idRemitente string `xorm:"varchar(45) 'id_remitente'"`
		pdfidPdf int16 `xorm:"INT(11) 'pdf_idPDF'"`
		tiempoRespuesta time.Time `xorm:"TIMESTAMP 'tiempo_respuesta'"`
	}

	var errGe error
	engine, errGe = xorm.NewEngine("mysql", "root:lexco123@tcp(127.0.0.1:3306)/sgdea")
	engine.SetMapper(core.SameMapper{})
	engine.Sync2(new(radicacion))
	fmt.Print(errGe)

	result,errrrr := engine.QueryString("select *  from radicacion");

//ff := engine.Table("radicacion").Select("radicacion.*");

for i:=0;i<len(result);i++{
	fmt.Print("posicion ",i,"[");
	for k, v := range result[i] {
		fmt.Printf("nombreColumna[%s] valor[%s]\n", k, v )
	}
	fmt.Println("]");
}


	fmt.Print(errrrr)


	/*session := engine.NewSession()
	defer session.Close()

	err := session.Begin()

	user1 := radicacion{numeroRadicacion: "REC591"}
	has, err  := engine.Get(&user1)
fmt.Print(has)
	err = session.Commit()
	if err != nil {
		panic(err)
		return
	}*/

}