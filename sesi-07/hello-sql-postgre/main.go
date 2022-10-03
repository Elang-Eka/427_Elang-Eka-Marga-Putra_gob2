package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// config db
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "elangeka"
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database")

	Pilihan()
}

func resetId() {
	sqlIncrement := `
	ALTER SQUENCE employees_id_seq RESTART WITH 1
	`
	_, err = db.Exec(sqlIncrement)
}

func createEmployee(fullname string, email string, age int, division string) {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *
	`

	err = db.QueryRow(sqlStatement, fullname, email, age, division).
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetEmployees() {
	var result = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		result = append(result, employee)
	}

	fmt.Println("Employee datas: ", result)
}

func UpdateEmployees(id int, fullname string, email string, age int, division string) {
	sqlStatement := `
	UPDATE  employees
	SET full_name = $2, email = $3, division = $4, age =$5
	WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, id, fullname, email, division, age)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Update data amount:", count)
}

func DeleteEmployees(id int) {
	sqlStatement := `DELETE from employees
	WHERE id = $1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete Data:", count)
}

func Pilihan() {
	var pilih int
	var loop = 1
	for i := 0; i <= loop; i++ {
		if loop == 1 {
			fmt.Println("==============================")
			fmt.Println("1. Create Employee")
			fmt.Println("2. Get All Employee")
			fmt.Println("3. Update Employee")
			fmt.Println("4. Delete Employee")
			fmt.Println("5. Exit")
			fmt.Println("==============================")
			fmt.Print("Masukkan Pilihan : ")
			fmt.Scanln(&pilih)
			if pilih == 1 || pilih == 2 || pilih == 3 || pilih == 4 || pilih == 5 {
				switch pilih {
				case 1:
					var email, division string
					var age int
					resetId()
					consoleReader := bufio.NewReader(os.Stdin)
					fmt.Println("==============================")
					fmt.Print("Masukkan nama lengkap : ")
					full_name, _ := consoleReader.ReadString('\n')
					fmt.Println("==============================")
					fmt.Print("Masukkan email : ")
					fmt.Scanln(&email)
					fmt.Println("==============================")
					fmt.Print("Masukkan umur : ")
					fmt.Scanln(&age)
					fmt.Println("==============================")
					fmt.Print("Masukkan divisi : ")
					fmt.Scanln(&division)
					fmt.Println("==============================")
					createEmployee(full_name, email, age, division)
					i = 0
					loop = 1
				case 2:
					fmt.Println("==============================")
					GetEmployees()
					i = 0
					loop = 1
				case 3:
					consoleReader := bufio.NewReader(os.Stdin)
					var email, division string
					var id, age int
					fmt.Println("==============================")
					fmt.Print("Masukkan id yang akan diganti : ")
					fmt.Scanln(&id)
					fmt.Println("==============================")
					fmt.Print("Masukkan nama lengkap : ")
					full_name, _ := consoleReader.ReadString('\n')
					fmt.Println("==============================")
					fmt.Print("Masukkan email : ")
					fmt.Scanln(&email)
					fmt.Println("==============================")
					fmt.Print("Masukkan umur : ")
					fmt.Scanln(&age)
					fmt.Println("==============================")
					fmt.Print("Masukkan division : ")
					fmt.Scanln(&division)
					fmt.Println("==============================")
					UpdateEmployees(id, full_name, email, age, division)
					i = 0
					loop = 1
				case 4:
					var id int
					fmt.Println("==============================")
					fmt.Print("Masukkan id yang akan dihapus : ")
					fmt.Scanln(&id)
					fmt.Println("==============================")
					DeleteEmployees(id)
					i = 0
					loop = 1
				case 5:
					i = 1
					loop = 0
					fmt.Println("==============================")
					fmt.Println("Terima Kasih")
					fmt.Println("==============================")
				}
			} else {
				i = 0
				loop = 1
				fmt.Println("Pilihan tidak ada menu hanya 1 hingga 5")
			}
		}
	}
}
