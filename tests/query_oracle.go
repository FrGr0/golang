//test go ORACLE SQL
package main

import (
    "fmt"
    "time"
    "strftime"
    "io/ioutil"
    "database/sql"
    "odbc"
    "strconv"
    . "trycatch"
)

func main() {

    Block{
        Try: func() {
            
            Qry := readfile( "./qry/select_test.sql" )

            // *** CONNEXION ORACLE ODBC ***
            db, sc, err := oraConnect( "GLDDF008", "TPXBAT", "TPXBAT" )
            fmt.Println( "**************************************************" )
            fmt.Printf(  "[%s] connexion à la BDD Oracle --> OK\n", Now() )
            fmt.Printf(  "[%s] DBStat  : %d\n", Now(), sc )
            fmt.Printf(  "[%s] DBError : %v\n", Now(), err )
            fmt.Println( "**************************************************" )

            rows, _ := db.Query(Qry)

            defer rows.Close()
            
            //variables retour sql
            xr  := ""
            xvl := 0

            for rows.Next() {
                if err := rows.Scan(&xr, &xvl); err != nil {
                    panic( err )
                }
                fmt.Printf( "r: %s, vl: %s\n", xr, 
                                               Lpad( strconv.Itoa( xvl ), 3, "0" ) )
            } 

            db.Close()
        },
        Catch: func(e Exception) {
            fmt.Printf("Erreur: %v\n", e)
        },

        Finally: func() {
            fmt.Println( "**************************************************" )
            fmt.Printf(  "[%s] Fermeture du programme.\n", Now() )
            fmt.Println( "**************************************************" )
        },
    }.Do()

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

func readfile(file string) (string){
    dat, err := ioutil.ReadFile(file)
    if (err!=nil){
        panic( err )
    }

    return string(dat)
}

func Now() (string) {
    return strftime.Format("%d-%m-%Y %H:%M:%S", time.Now() )
}

func Lpad(s string, l int, comp string)(string) {
    for len(s) < l {
        s = comp+s
    }
    return s
}