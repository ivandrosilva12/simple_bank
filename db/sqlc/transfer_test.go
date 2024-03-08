package db

import (
	"context"
	"database/sql"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandonTransfer(t *testing.T) Transfer {

	account1 := CreateRandonAccount(t)
	account2 := CreateRandonAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandonMoney(),
	}

	tranfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tranfer)

	require.Equal(t, arg.FromAccountID, tranfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, tranfer.ToAccountID)
	require.Equal(t, arg.Amount, tranfer.Amount)

	require.NotZero(t, tranfer.ID)
	require.NotZero(t, tranfer.CreatedAt)

	return tranfer
}

func TestCreateTransfer(t *testing.T) {
	createRandonTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandonTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandonTransfer(t)

	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandonMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandonTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestListTransfers(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandonTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	tranfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tranfers, 5)

	for _, tranfer := range tranfers {
		require.NotEmpty(t, tranfer)
	}

}
