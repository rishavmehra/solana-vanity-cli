package pkg

import (
	"fmt"
	"strings"
	"time"

	"github.com/gagliardetto/solana-go"
)

var (
	GenerateCount     = 0
	StartTime         = time.Now()
	ShouldStopThreads = false
)

func GenerateWallet(searchString string) {
	for {
		if ShouldStopThreads {
			return
		}
		newWallet := solana.NewWallet()
		if strings.HasPrefix(newWallet.PublicKey().String(), searchString) {
			firstCharAfterSearchTerm := strings.Split(newWallet.PublicKey().String(), searchString)[1][0:1]
			if firstCharAfterSearchTerm == strings.ToUpper(firstCharAfterSearchTerm) {
				fmt.Printf("Success wallet found: %s \n", newWallet.PublicKey())
				fmt.Printf("Secret key: %v\n", newWallet.PrivateKey)
				fmt.Printf("Attempts required: %d, time elapsed: %s\n", GenerateCount+1, time.Since(StartTime))
				ShouldStopThreads = true
			}
		}
		GenerateCount++
		if GenerateCount%1000000 == 0 {
			fmt.Printf("Status: %d wallets generated in %s\n", GenerateCount, time.Since(StartTime))
		}
	}
}
