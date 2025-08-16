package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CheckFileExistance(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return false, fmt.Errorf("%s does not exist", path)
	}

	return true, nil
}

func LoadEnvironmentFile(envFile string) error {
	if doesFileExist, err := CheckFileExistance(envFile); doesFileExist {
		fmt.Printf("loading environment file: %s\n", envFile)
		if err := godotenv.Load(envFile); err != nil {
			log.Fatal("error loading .env")
			return fmt.Errorf("")
		}
	} else {
		fmt.Println(err)
		return fmt.Errorf("")
	}

	return nil
}

func ReadSQLFile(filename string) (string, error) {
	sqlFileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	return string(sqlFileContents), nil
}
