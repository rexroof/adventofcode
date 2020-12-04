package passport

/*
The expected fields are as follows:

byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)
*/

//north pole id or passport
type passport struct {
	byr string // (Birth Year)
	iyr string // (Issue Year)
	eyr string // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid string // (Country ID) (optional)
}

func New(byr string, iyr string, eyr string, hgt string, hcl string, ecl string, pid string, cid string) passport {
	p := passport{byr, iyr, eyr, hgt, hcl, ecl, pid, cid}
	return p
}

func (p passport) CheckValidity() bool {
  valid := true
  if len(p.byr) < 1 { valid = false }
  if len(p.iyr) < 1 { valid = false }
  if len(p.eyr) < 1 { valid = false }
  if len(p.hgt) < 1 { valid = false }
  if len(p.hcl) < 1 { valid = false }
  if len(p.ecl) < 1 { valid = false }
  if len(p.pid) < 1 { valid = false }

	return valid
}
