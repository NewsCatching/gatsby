package sqlutils
import "testing"

type Staff struct {
	Id        int64 `json:"id" field:",primary,serial"`
	Name      string `json:"name" field:",required"`
	Gender    string `json:"gender"`
	StaffType string `json:"staff_type"` // valid types: doctor, nurse, ...etc
	Phone     string `json:"phone"`
}

// Implement the GetPkId interface
func (self *Staff) GetPkId() int64 {
	return self.Id
}

func TestBuildSelectColumns(t * testing.T) {
	str := BuildSelectColumnClause( &fooRecord{Id:4, Name: "John"} )
	if len(str) == 0 {
		t.Fail()
	}
	if str != "id,name,type" {
		t.Fatal(str)
	}
}

func TestBuildSelectClause(t * testing.T) {
	staff := Staff{Id:4, Name: "John", Gender: "m", Phone: "0975277696"}
	sql := BuildSelectClause(&staff)
	if sql != "SELECT id,name,gender,staff_type,phone FROM staffs" {
		t.Fatal(sql)
	}
}

