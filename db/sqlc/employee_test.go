package db

import (
	"context"
	"example-crud/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateEmployee(t *testing.T) {
	CreateEmployee(t)
}

func TestGetEmployee(t *testing.T) {
	employee1 := CreateEmployee(t)

	employee2, err := testQueries.GetEmployee(context.Background(), employee1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, employee2)
	require.Equal(t, employee1.Code, employee2.Code)
	require.Equal(t, employee1.Name, employee2.Name)
	require.Equal(t, employee1.Email, employee2.Email)
	require.Equal(t, employee1.PhoneNumber, employee2.PhoneNumber)
	require.WithinDuration(t, employee1.CreatedAt, employee2.CreatedAt, time.Second)
}

func CreateEmployee(t *testing.T) Employee {
	params := CreateParams()
	employee, err := testQueries.CreateEmployee(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, employee)
	require.Equal(t, params.Code, employee.Code)
	require.Equal(t, params.Name, employee.Name)
	require.Equal(t, params.Email, employee.Email)
	require.Equal(t, params.PhoneNumber, employee.PhoneNumber)
	require.NotZero(t, employee.CreatedAt)

	return employee
}

func CreateParams() CreateEmployeeParams {
	return CreateEmployeeParams{
		Code:        util.RandomCode(),
		Name:        util.RandomName(),
		Email:       util.RandomEmail(),
		PhoneNumber: util.RandomPhoneNumber(),
	}
}
