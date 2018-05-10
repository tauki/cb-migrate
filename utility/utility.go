package utility

import (
	"fmt"
	"github.com/cb-migrate/models"
)

func CheckFlags(data *models.Data, context string) bool {
	check := true

	if data.DBHost == "" {
		check = false
		fmt.Printf("Please input %s address: \n", context)
		fmt.Scanln(&data.DBHost)
		for {
			if checkIfUrl(data.DBHost) {
				break
			}
		}
	}

	if data.DBPort == "" {
		check = false
		fmt.Printf("Please input %s port: \n", context)
		fmt.Scanln(&data.DBPort)
	}

	if data.DBUser == "" {
		check = false
		fmt.Printf("Please input %s DB Username: ", context)
		fmt.Scanln(&data.DBUser)
	}

	if data.DBPassword == "" {
		check = false
		fmt.Printf("Please input %s DB password: ", context)
		fmt.Scanln(&data.DBPassword)
	}

	return check
}

func checkIfUrl(source string) bool {
	return true
}

//
//func GetBucketCreds(name string) *models.Bucket {
//	var bucket models.Bucket
//
//	bucket.BucketName = name
//
//	fmt.Printf("Please input the password for %s\n", name)
//	fmt.Println("leave empty for none")
//	fmt.Scanln(&bucket.BucketPassword)
//
//	return &bucket
//}
