package main

import (
	"fmt"

	tpmclient "github.com/google/go-tpm-tools/client"
	"github.com/google/go-tpm/tpm2"
)

func main() {
	tpm, err := tpm2.OpenTPM()
	if err != nil {
		panic(fmt.Errorf("open tpm: %s", err))
	}

	pcrs, err := tpmclient.ReadAllPCRs(tpm)
	if err != nil {
		panic(fmt.Errorf("failed to read PCRs: %w", err))
	}

	for i := range pcrs {
		for j := 0; j < 24; j++ {
			fmt.Printf("%02d: %X\n", j, pcrs[i].Pcrs[uint32(j)])
		}
	}
}
