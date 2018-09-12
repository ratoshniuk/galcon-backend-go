package cassandra

import (
    "fmt"
    "github.com/gocql/gocql"
    "log"
)

type CassandraContext struct {
    Cluster *gocql.ClusterConfig
    Session *gocql.Session
    KeyspaceSession *gocql.Session
}

func (ctx *CassandraContext) Init(host string, keyspace string) {
    ctx.Cluster = gocql.NewCluster(host)

    ctx.Session, _ = ctx.Cluster.CreateSession()
    var err error
    err = ctx.Session.Query(fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }", keyspace)).Exec()
    if err != nil {
        log.Fatal(err)
    }

    ctx.Cluster.Keyspace = keyspace
    ctx.KeyspaceSession, _ = ctx.Cluster.CreateSession()
}

func (ctx *CassandraContext) PerformDDL(ddls ...func(keyspace string) *string) {

    fmt.Println("running ddls...")
    for _, ddl := range ddls {
        res := ddl(ctx.Cluster.Keyspace)
        if res != nil {
            fmt.Printf("upping %s\n", *res)
            err := ctx.Session.Query(*res).Exec()
            if err != nil {
                log.Fatal(err)
            }
        }

    }
}

func (ctx *CassandraContext) Stop(keyspaceCleanup bool) {
    if keyspaceCleanup {
        ctx.Session.Query(fmt.Sprintf("delete keyspace %s", ctx.Cluster.Keyspace))
    }
}
