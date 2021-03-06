package gatsby

import "testing"

func BenchmarkCreateAndDelete(b *testing.B) {
	var db = openDB()
	staff := Staff{}
	staff.Name = "John"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Create(db, &staff, DriverPg)
		Delete(db, &staff)
	}
}

func TestCreate(t *testing.T) {
	var db = openDB()
	staff := Staff{}
	staff.Name = "John"
	Create(db, &staff, DriverPg)
	Delete(db, &staff)
}

func TestCreateWithTransactionAndCommit(t *testing.T) {
	var db = openDB()
	SetupConnection(db, DriverPg)

	db.Query("delete from staffs;")

	staff := Staff{}
	staff.Init()

	staff.Name = "Txn Test"
	staff.Gender = "f"
	tx, err := staff.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil {
		t.Fatal("transaction is nil")
	}

	// create with transaction
	res := staff.Create()
	if res.Error != nil {
		t.Fatal(res.Error)
	}

	err = staff.Commit()
	if err != nil {
		t.Fatal(err)
	}
	CloseConnection()
}

func TestCreateWithTransactionAndRollback(t *testing.T) {
	var db = openDB()
	SetupConnection(db, DriverPg)

	staff := Staff{}
	staff.Init()

	staff.Name = "Txn Test With Rollback"
	tx, err := staff.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil {
		t.Fatal("transaction is nil")
	}

	// create with transaction
	res := staff.Create()
	if res.Error != nil {
		t.Fatal(res.Error)
	}
	t.Log(res)
	t.Log(staff)

	err = staff.Rollback()
	if err != nil {
		t.Fatal(err)
	}

	pId := staff.Id
	staff2 := Staff{}
	staff2.Init()

	res = staff2.Load(pId)

	if !res.IsEmpty {
		t.Fatal("Still found the record.")
	}

	CloseConnection()
}
