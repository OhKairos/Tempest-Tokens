package main

import (
	"fmt"
	"log"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/txnbuild"
)

func main() {

	// Use the default pubnet client
	kp, _ := keypair.Parse("QUEST_SERIES_1_SECRET_KEY")
	client := horizonclient.DefaultTestNetClient
	ar := horizonclient.AccountRequest{AccountID: kp.Address()}
	sourceAccount, err := client.AccountDetail(ar)
	// check(err)
	if err != nil {
		fmt.Println("Error 1")
		log.Fatalln(err)
		return
	}

	op := txnbuild.SetOptions{
		// InflationDestination: NewInflationDestination("GCCOBXW2XQNUSL467IEILE6MMCNRR66SSVL4YQADUNYYNUVREF3FIV2Z"),
		// ClearFlags:           []AccountFlag{AuthRevocable},
		// SetFlags:             []AccountFlag{AuthRequired, AuthImmutable},
		// MasterWeight:         NewThreshold(10),
		// LowThreshold:         NewThreshold(1),
		// MediumThreshold:      NewThreshold(2),
		// HighThreshold:        NewThreshold(2),
		// HomeDomain:           NewHomeDomain("LovelyLumensLookLuminous.com"),
		Signer: &txnbuild.Signer{Address: "MASTER_ACCOUNT_PUBLIC_KEY", Weight: txnbuild.Threshold(1)},
	}

	tx, err := txnbuild.NewTransaction(
		txnbuild.TransactionParams{
			SourceAccount:        &sourceAccount,
			IncrementSequenceNum: true,
			Operations:           []txnbuild.Operation{&op},
			BaseFee:              txnbuild.MinBaseFee,
			Timebounds:           txnbuild.NewInfiniteTimeout(), // Use a real timeout in production!
		},
	)
	// check(err)
	if err != nil {
		fmt.Println("Error 2")
		fmt.Println(err)
		return
	}

	tx, err = tx.Sign(network.TestNetworkPassphrase, kp.(*keypair.Full))
	// check(err)
	if err != nil {
		fmt.Println("Error 3")
		fmt.Println(err)
		return
	}

	txe, err := tx.Base64()
	// check(err)
	if err != nil {
		fmt.Println("Error 4")
		fmt.Println(err)
		return
	}
	fmt.Println(txe)

	// submit transaction
	resp, err := client.SubmitTransactionXDR(txe)
	if err != nil {
		fmt.Println("Failed to sent Transactions")
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
