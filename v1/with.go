package bluzelledbgo

import (
	"context"

	pb "github.com/cpurta/bluzelle-db-go/types"
)

var (
	// compile time interface checking for TransactionOperation implementations
	_ TransactionOperation = &transactionCreate{}
	_ TransactionOperation = &transactionDelete{}
	_ TransactionOperation = &transactionDeleteAll{}
	_ TransactionOperation = &transactionMultiUpdate{}
	_ TransactionOperation = &transactionRenewLeasesAll{}
	_ TransactionOperation = &transactionRenewLease{}
	_ TransactionOperation = &transactionRename{}
	_ TransactionOperation = &transactionUpdate{}
	_ TransactionOperation = &transactionUpsert{}
)

//
//  -- CREATE
//

// NewTransactionCreate provides a wrapper around a *MsgCreate so that it can be
// used in the WithTransactions batch processing function.
func (client *defaultTransactionClient) NewTransactionCreate(create *pb.MsgCreate) TransactionOperation {
	return &transactionCreate{
		create: create,
		client: client,
	}
}

type transactionCreate struct {
	create *pb.MsgCreate
	client *defaultTransactionClient
}

// PerformOperation will perform the underlying create transaction operation.
func (txnCreate *transactionCreate) PerformOperation(ctx context.Context) error {
	_, err := txnCreate.client.Create(ctx, txnCreate.create)

	return err
}

//
//  -- DELETE
//

// NewTransactionDelete provides a wrapper around a *MsgDelete so that it can be
// used in the WithTransactions batch processing function.
func (client *defaultTransactionClient) NewTransactionDelete(delete *pb.MsgDelete) TransactionOperation {
	return &transactionDelete{
		delete: delete,
		client: client,
	}
}

type transactionDelete struct {
	delete *pb.MsgDelete
	client *defaultTransactionClient
}

// PerformOperation will perform the underlying delete transaction operation.
func (txnCreate *transactionDelete) PerformOperation(ctx context.Context) error {
	_, err := txnCreate.client.Delete(ctx, txnCreate.delete)

	return err
}

//
//  -- DELETE ALL
//

// NewTransactionDeleteAll provides a wrapper around a *MsgDeleteAll so that it can be
// used in the WithTransactions batch processing function.
func (client *defaultTransactionClient) NewTransactionDeleteAll(deleteAll *pb.MsgDeleteAll) TransactionOperation {
	return &transactionDeleteAll{
		deleteAll: deleteAll,
		client:    client,
	}
}

type transactionDeleteAll struct {
	deleteAll *pb.MsgDeleteAll
	client    *defaultTransactionClient
}

// PerformOperation will perform the underlying delete all transaction operation.
func (txnCreate *transactionDeleteAll) PerformOperation(ctx context.Context) error {
	_, err := txnCreate.client.DeleteAll(ctx, txnCreate.deleteAll)

	return err
}

//
//  -- MULTI UPDATE
//

// NewTransactionMultiUpdate provides a wrapper around a *MsgMultiUpdate so that it can be
// used in the WithTransactions batch processing function.
func (client *defaultTransactionClient) NewTransactionMultiUpdate(multiUpdate *pb.MsgMultiUpdate) TransactionOperation {
	return &transactionMultiUpdate{
		multiUpdate: multiUpdate,
		client:      client,
	}
}

type transactionMultiUpdate struct {
	multiUpdate *pb.MsgMultiUpdate
	client      *defaultTransactionClient
}

// PerformOperation will perform the underlying multi update transaction operation.
func (txnCreate *transactionMultiUpdate) PerformOperation(ctx context.Context) error {
	_, err := txnCreate.client.MultiUpdate(ctx, txnCreate.multiUpdate)

	return err
}

//
//  -- RENEW LEASES ALL
//

// NewTransactionRenewLeasesAll provides a wrapper around a *MsgRenewLeasesAll so that it can be
// used in the WithTransactions batch processing function.
func (client *defaultTransactionClient) NewTransactionRenewLeasesAll(renewLeasesAll *pb.MsgRenewLeasesAll) TransactionOperation {
	return &transactionRenewLeasesAll{
		renewLeasesAll: renewLeasesAll,
		client:         client,
	}
}

type transactionRenewLeasesAll struct {
	renewLeasesAll *pb.MsgRenewLeasesAll
	client         *defaultTransactionClient
}

// PerformOperation will perform the underlying renew leases all transaction operation.
func (txnCreate *transactionRenewLeasesAll) PerformOperation(ctx context.Context) error {
	_, err := txnCreate.client.RenewLeasesAll(ctx, txnCreate.renewLeasesAll)

	return err
}

//
//  -- RENEW LEASE
//

// NewTransactionRenewLease provides a wrapper around a *MsgRenewLease so that it can be
// used in the WithTransactions batch processing function.
func (client *defaultTransactionClient) NewTransactionRenewLease(renewLease *pb.MsgRenewLease) TransactionOperation {
	return &transactionRenewLease{
		renewLease: renewLease,
		client:     client,
	}
}

type transactionRenewLease struct {
	renewLease *pb.MsgRenewLease
	client     *defaultTransactionClient
}

// PerformOperation will perform the underlying renew leases all transaction operation.
func (txnCreate *transactionRenewLease) PerformOperation(ctx context.Context) error {
	_, err := txnCreate.client.RenewLease(ctx, txnCreate.renewLease)

	return err
}

//
//  -- RENAME
//

// NewTransactionRename provides a wrapper around a *MsgRename so that it can be
// used in the WithTransactions batch processing function.
func (client *defaultTransactionClient) NewTransactionRename(rename *pb.MsgRename) TransactionOperation {
	return &transactionRename{
		rename: rename,
		client: client,
	}
}

type transactionRename struct {
	rename *pb.MsgRename
	client *defaultTransactionClient
}

// PerformOperation will perform the underlying rename transaction operation.
func (txnCreate *transactionRename) PerformOperation(ctx context.Context) error {
	_, err := txnCreate.client.Rename(ctx, txnCreate.rename)

	return err
}

//
//  -- UPDATE
//

// NewTransactionUpdate provides a wrapper around a *MsgUpdate so that it can be
// used in the WithTransactions batch processing function.
func (client *defaultTransactionClient) NewTransactionUpdate(update *pb.MsgUpdate) TransactionOperation {
	return &transactionUpdate{
		update: update,
		client: client,
	}
}

type transactionUpdate struct {
	update *pb.MsgUpdate
	client *defaultTransactionClient
}

// PerformOperation will perform the underlying rename transaction operation.
func (txnCreate *transactionUpdate) PerformOperation(ctx context.Context) error {
	_, err := txnCreate.client.Update(ctx, txnCreate.update)

	return err
}

//
//  -- UPSERT
//

// NewTransactionUpsert provides a wrapper around a *MsgUpsert so that it can be
// used in the WithTransactions batch processing function.
func (client *defaultTransactionClient) NewTransactionUpsert(upsert *pb.MsgUpsert) TransactionOperation {
	return &transactionUpsert{
		upsert: upsert,
		client: client,
	}
}

type transactionUpsert struct {
	upsert *pb.MsgUpsert
	client *defaultTransactionClient
}

// PerformOperation will perform the underlying rename transaction operation.
func (txnCreate *transactionUpsert) PerformOperation(ctx context.Context) error {
	_, err := txnCreate.client.Upsert(ctx, txnCreate.upsert)

	return err
}
