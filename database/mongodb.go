package mongodb

import (
    "context"
    "strings"
    "alumnus/config"
    "alumnus/worker"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type callback func(string,string) string

//Users collection object
var Users *mongo.Collection

// Connect is...
func Connect(dbuser string,dbpass string) (connected bool,err error) {
    uri := func() string {
        if token.IsDev { 
            return "mongodb://localhost:27017/"+config.Config.DB.DBname+"?retryWrites=true&w=majority"
        }
        var str strings.Builder
        parts := [7]string{"mongodb+srv://",dbuser ,":",dbpass,"@realmcluster.njdl8.mongodb.net/",config.Config.DB.DBname,"?retryWrites=true&w=majority"}
        for i := 0; i < len(parts); i++ {
            str.WriteString(parts[i])
        }
        return str.String()
    }()
    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return false,err
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        return false,err
    }
    Users = client.Database(config.Config.DB.DBname).Collection(config.Config.DB.Users)
    return true,nil
}