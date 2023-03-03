package lesson_huesson

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"

	pgx "github.com/jackc/pgx/v4"
)

const (
	host     = "rc1a-y7k98sxnrom6oeg6.mdb.yandexcloud.net"
	port     = 6432
	user     = "dev"
	password = "devdev123"
	dbname   = "dev"
	ca       = "C:\\Users\\foxtrot\\GolandProjects\\dinar\\.postgresql\\root.crt"
)

func Connect() {

	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(ca)
	if err != nil {
		panic(err)
	}

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		panic("Failed to append PEM.")
	}

	connstring := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=verify-full target_session_attrs=read-write",
		host, port, dbname, user, password)

	connConfig, err := pgx.ParseConfig(connstring)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config: %v\n", err)
		os.Exit(1)
	}

	connConfig.TLSConfig = &tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: true,
	}

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	var version string

	err = conn.QueryRow(context.Background(), "select version()").Scan(&version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(version)
}
