package database

var connection string

func init() {
	connection = "MySQL"
} // init function akan dijalankan pertama kali saat package ini diakses

func GetDatabase() string {
	return connection
}
