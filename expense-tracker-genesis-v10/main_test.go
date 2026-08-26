package main

import (
	"testing"
	"time"

	"expense-tracker/internal/service"
	"expense-tracker/internal/store"
)

func TestServiceAndStore(t *testing.T) {
	st, err := store.NewSQLiteStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer st.Close()

	svc := service.NewService(st)

	// Test Add and List
	_, err = svc.AddExpense(25.0, "Coffee", "Food", time.Now())
	if err != nil {
		t.Error(err)
	}

	expenses, err := svc.ListExpenses()
	if err != nil || len(expenses) == 0 {
		t.Error("expected at least one expense")
	}

	// Test Summary
	_, total, err := svc.GetSummary()
	if err != nil || total != 25.0 {
		t.Errorf("expected total 25.0, got %f", total)
	}

	// Test Categories
	err = svc.AddCategory("Travel")
	if err != nil {
		t.Error(err)
	}

	cats, err := svc.GetCategories()
	if err != nil || len(cats) == 0 {
		t.Error("expected categories")
	}

	// Test Update and Delete
	exp := expenses[0]
	_, err = svc.UpdateExpense(exp.ID, 30.0, "Updated Coffee", "Food", time.Now())
	if err != nil {
		t.Error(err)
	}

	err = svc.DeleteExpense(exp.ID)
	if err != nil {
		t.Error(err)
	}

	t.Log("All service and store tests passed")
}
