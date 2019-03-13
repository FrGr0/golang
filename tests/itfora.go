package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "database/sql"
    "odbc"
    "fmt"
    . "trycatch"
)

func main() {
    var inTE, outTE *walk.TextEdit
    db, sc, err := oraConnect( "GLDDF008", "******", "******" )
    if ( sc!=0 ){
        //
    }
    if ( err!= nil ) {
        //
    }

    MainWindow{
        Title:   "GOLANG ORACLE TEST",
        MinSize: Size{600, 400},
        Layout:  VBox{},
        Children: []Widget{
            VSplitter{
                Children: []Widget{
                    TextEdit{AssignTo: &inTE},
                    TextEdit{AssignTo: &outTE, ReadOnly: true},
                },
            },
            PushButton{
                Text: "run",
                OnClicked: func() {
                    Block{
                        Try: func() {
                            outTE.SetText( "" )
                            qry := fmt.Sprintf("%s", inTE.Text())
                        
                            rows, _ := db.Query(qry)
                            cols, _ := rows.Columns()
                            for rows.Next() {
                                columns := make([]interface{}, len(cols))
                                columnPointers := make([]interface{}, len(cols))
                                for i, _ := range columns {
                                    columnPointers[i] = &columns[i]
                                }
                                
                                if err := rows.Scan(columnPointers...); err != nil {
                                    //panic(err)
                                }
                            
                                m := make(map[string]interface{})
                                for i, colName := range cols {
                                    val := columnPointers[i].(*interface{})
                                    m[colName] = *val
                                }

                                outTE.SetText(fmt.Sprintf( "%s%s\r\n", outTE.Text(), m ))
                            }
                        },
                        Catch: func(e Exception) {
                            outTE.SetText(fmt.Sprintf( "%s\r\n", e ))
                        },
                    
                        Finally: func() {
                            //
                        },
                    }.Do()
                },
            },
        },
    }.Run()
    db.Close()
}

func oraConnect(dbhost string, dbuid string, dbpass string) (db *sql.DB, stmtCount int, err error) {
    conn := "driver={Oracle dans OraClient11g_home1};DBQ="+dbhost+";UID="+dbuid+";PWD="+dbpass+";"
    db, err = sql.Open("odbc", conn)
    if err != nil {
        return nil, 0, err
    }
    stats := db.Driver().(*odbc.Driver).Stats
    return db, stats.StmtCount, nil
}