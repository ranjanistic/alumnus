package config

type db struct {
    Username string
    Pass string
    DBname string
    Users string
    Dpass string
    Cpass string
}

type config struct {
    Appname string
    DB *db
}

//Config is
var Config = config{
    Appname: "Alum",
    DB: &db{
        DBname:"alumnusDB",
        Users:"0users",
    },
}


