package connect_to_db

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

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

type Row struct {
	Id           int       `json:"Id"`
	CreatedAt    time.Time `json:"Created_at"`
	ClientName   string    `json:"Client_name"`
	Amount       *int      `json:"Amount"`
	OrderComment *string   `json:"Order_comment"`
}

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

	rows, err := conn.Query(context.Background(), "SELECT * FROM public.table_dinar")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var rowsSlice []Row
	for rows.Next() {
		var r Row
		err := rows.Scan(&r.Id, &r.CreatedAt, &r.ClientName, &r.Amount, &r.OrderComment)
		if err != nil {
			log.Fatal(err)
		}
		rowsSlice = append(rowsSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, v := range rowsSlice {
		fmt.Println(v.Id, v.CreatedAt, v.ClientName, *v.Amount, *v.OrderComment)
	}
}
