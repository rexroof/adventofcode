package passport

import "strconv"
import "regexp"
import "fmt"

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

func (p passport) validNumberRange(_in string, _min int, _max int) bool {

	if len(_in) < 1 {
		return false
	}

	i, err := strconv.Atoi(_in)
	if err != nil {
		fmt.Printf("%s %d %d %s\n", _in, _min, _max, err)
		return false
	}
	if i < _min {
		return false
	}
	if i > _max {
		return false
	}

	return true
}

func (p passport) CheckValidity() bool {
	valid := true
	if !p.validNumberRange(p.byr, 1920, 2002) {
		valid = false
	}
	if !p.validNumberRange(p.iyr, 2010, 2020) {
		valid = false
	}
	if !p.validNumberRange(p.eyr, 2020, 2030) {
		valid = false
	}

	// 150-193cm or 59-76in
	if len(p.hgt) > 1 {
		re := regexp.MustCompile(`^([0-9]{2,3})(in|cm)$`)
		if re.MatchString(p.hgt) {
			matches := re.FindStringSubmatch(p.hgt)
			if matches[2] == "in" {
				if !p.validNumberRange(matches[1], 59, 76) {
					valid = false
					fmt.Println("Did not pass")
					fmt.Println(matches)
				}
			} else if matches[2] == "cm" {
				if !p.validNumberRange(matches[1], 150, 193) {
					valid = false
					fmt.Println("Did not pass")
					fmt.Println(matches)
				}
			}
		} else {
			fmt.Printf("`%s` did not pass `^([0-9]{2,3})(in|cm)$`\n", p.hgt)
			valid = false
		}
	} else {
		valid = false
	}

	// #[0-9a-f]{6}
	if len(p.hcl) > 1 {
		re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		if !(re.MatchString(p.hcl)) {
			fmt.Printf("`%s` did not pass `^#[0-9a-f]{6}$`\n", p.hcl)
			valid = false
		}
	} else {
		valid = false
	}

	// amb|blu|brn|gry|grn|hzl|oth
	if len(p.ecl) > 1 {
		re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		if !(re.MatchString(p.ecl)) {
			fmt.Printf("`%s` did not pass `^(amb|blu|brn|gry|grn|hzl|oth)$`\n", p.ecl)
			valid = false
		}
	} else {
		valid = false
	}

	// [0-9]{9}
	if len(p.pid) > 1 {
		re := regexp.MustCompile(`^[0-9]{9}$`)
		if !(re.MatchString(p.pid)) {
			fmt.Printf("`%s` did not pass `^[0-9]{9}$`\n", p.pid)
			valid = false
		}
	} else {
		valid = false
	}

	return valid
}
