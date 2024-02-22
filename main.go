package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	// Defina o horário em que deseja executar o backup (22h)
	backupTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 22, 00, 0, 0, time.Local)

	// Inicia a goroutine para agendar a execução do backup
	go scheduleBackup(backupTime)

	// Mantém o programa em execução para permitir que a goroutine agende a execução
	select {}
}

func scheduleBackup(backupTime time.Time) {
	for {
		// Obtém a data atual
		now := time.Now()

		// Calcula a duração até o próximo horário de backup
		durationUntilNextBackup := backupTime.Sub(now)

		// Aguarda até o próximo horário de backup
		time.Sleep(durationUntilNextBackup)

		// Quando chegar o horário de backup, executa a função de backup
		go performBackup()

		// Agenda o próximo backup para o próximo dia
		backupTime = backupTime.Add(24 * time.Hour)
	}
}

func performBackup() {
	// Parâmetros de conexão ao banco de dados PostgreSQL
	dbHost := os.Getenv("HOST_DATABASE")
	dbPort := os.Getenv("PORT_DATABASE")
	dbUser := os.Getenv("USER_DATABASE")
	dbPassword := os.Getenv("PASS_DATABASE")

	// Nome da pasta que será criada para os Backups
	backupFolder := fmt.Sprintf("bkp_db_arval_%s", time.Now().Format("2006-01-02"))

	// Verifica se a pasta já existe
	if _, err := os.Stat(backupFolder); os.IsNotExist(err) {
		// Permissão de leitura e escrita
		permissoes := os.FileMode(0755)

		// Cria a pasta
		err := os.Mkdir(backupFolder, permissoes)
		if err != nil {
			fmt.Println("Erro ao criar a pasta:", err)
			return
		}
	}

	// Obtém o timestamp atual para usar no nome do arquivo de backup
	timestamp := time.Now().Format("2006-01-02_15.04.05")

	// Backup DBs playground
	playgroundDBs := []string{os.Getenv("DBNAME_PERU_PLAYGROUND"), os.Getenv("DBNAME_CHILE_PLAYGROUND"), os.Getenv("DBNAME_MARROCOS_PLAYGROUND"), os.Getenv("DBNAME_RUSSIA_PLAYGROUND")}

	for _, dbName := range playgroundDBs {

		fmt.Println("Iniciando o backup do banco: ", dbName)
		fmt.Println("-------------------------------------")

		// Nome do arquivo de backup
		backupFileName := fmt.Sprintf("%s/%s_%s_bkp.sql", backupFolder, dbName, timestamp)

		// Comando pg_dump para criar o backup
		cmd := exec.Command("pg_dump", "-h", dbHost, "-p", dbPort, "-U", dbUser, "-d", dbName, "-f", backupFileName, "-F", "c")

		// Define a senha do usuário do banco de dados
		cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", dbPassword))

		// Executa o comando pg_dump
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Erro ao criar o Backup para o banco de dados %s: %v\n", dbName, err)
		}

		fmt.Printf("Backup criado com sucesso para o banco de dados %s: %s\n", dbName, backupFileName)
		fmt.Println("-------------------------------------")
	}

	// Backup DBs produção
	productionDBs := []string{os.Getenv("DBNAME_PERU_PRODUCTION"), os.Getenv("DBNAME_CHILE_PRODUCTION"), os.Getenv("DBNAME_MARROCOS_PRODUCTION")}

	for _, dbName := range productionDBs {

		fmt.Println("Iniciando o backup do banco: ", dbName)
		fmt.Println("-------------------------------------")

		// Nome do arquivo de backup
		backupFileName := fmt.Sprintf("%s/%s_%s_bkp.sql", backupFolder, dbName, timestamp)

		// Comando pg_dump para criar o backup
		cmd := exec.Command("pg_dump", "-h", dbHost, "-p", dbPort, "-U", dbUser, "-d", dbName, "-f", backupFileName, "-F", "c")

		// Define a senha do usuário do banco de dados
		cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", dbPassword))

		// Executa o comando pg_dump
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Erro ao criar o Backup para o banco de dados %s: %v\n", dbName, err)
		}

		fmt.Printf("Backup criado com sucesso para o banco de dados %s: %s\n", dbName, backupFileName)
		fmt.Println("-------------------------------------")
	}

	fmt.Println("Backups finalizados com sucesso!")
}
