package tri

import (
	"testing"
)

func TestBrief(t *testing.T) {

	// one item only
	tb1 := Brief{
		"item1", "item2",
	}
	e := tb1.Validate()
	if e == nil {
		t.Error("validator allowed more than one")
	}

	// string typed item
	tb2 := Brief{
		1,
	}
	e = tb2.Validate()
	if e == nil {
		t.Error("validator permitted other than a string")
	}

	// string length < 80
	tb3 := Brief{
		"123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
	}
	e = tb3.Validate()
	if e == nil {
		t.Error("validator permitted over 80 characters")
	}

	// no control characters
	tb4 := Brief{
		"this should not have a cr at the end\n",
	}
	e = tb4.Validate()
	if e == nil {
		t.Error("validator permitted over 80 characters")
	}

	tb5 := Brief{
		"this should not have a cr at the end\n",
	}
	e = tb5.Validate()
	if e == nil {
		t.Error("validator permitted over 80 characters")
	}

	tb6 := Brief{
		"this is ok",
	}
	e = tb6.Validate()
	if e != nil {
		t.Error("validator rejected correct input")
	}

}

func TestCommand(t *testing.T) {
	//string in index 0
	tc1 := Command{1}
	if e := tc1.Validate(); e == nil {
		t.Error("validator accepted non-string name")
	}
	//string is valid name (letters only)
	tc2 := Command{"!"}
	if e := tc2.Validate(); e == nil {
		t.Error("validator accepted invalid name")
	}
	//more than one brief not allowed
	tc3 := Command{"name", Brief{""}, Brief{""}}
	if e := tc3.Validate(); e == nil {
		t.Error("validator accepted more than one Brief")
	}
	// brief is invalid
	tc4 := Command{"name", Brief{}}
	if e := tc4.Validate(); e == nil {
		t.Error("validator accepted invalid Brief")
	}
	//more than one handler not allowed
	tc5 := Command{"name", Brief{""}, func(Tri) int { return 0 }, func(Tri) int { return 0 }}
	if e := tc5.Validate(); e == nil {
		t.Error("validator accepted more than one handler")
	}
	//handler not nil
	isnil := func(Tri) int { return 1 }
	_ = isnil
	isnil = nil
	tc6 := Command{"name", isnil}
	if e := tc6.Validate(); e == nil {
		t.Error("validator accepted nil handler")
	}
	// no more than one Short
	tc7 := Command{"name", Short{'a'}, Short{'b'}}
	if e := tc7.Validate(); e == nil {
		t.Error("validator accepted more than one Short")
	}
	// invalid Short
	tc8 := Command{"name", Short{}}
	if e := tc8.Validate(); e == nil {
		t.Error("validator accepted invalid Short")
	}
	// no more than one Usage
	tc9 := Command{"name", Usage{""}, Usage{""}}
	if e := tc9.Validate(); e == nil {
		t.Error("validator accepted more than one Short")
	}
	// invalid Usage
	tc10 := Command{"name", Usage{}}
	if e := tc10.Validate(); e == nil {
		t.Error("validator invalid Usage")
	}
	// no more than one Help
	tc11 := Command{"name", Help{""}, Help{""}}
	if e := tc11.Validate(); e == nil {
		t.Error("validator accepted more than one Short")
	}
	// invalid Help
	tc12 := Command{"name", Help{}}
	if e := tc12.Validate(); e == nil {
		t.Error("validator invalid Usage")
	}
	// no more than one Examples
	tc13 := Command{"name", Examples{"", ""}, Examples{"", ""}}
	if e := tc13.Validate(); e == nil {
		t.Error("validator accepted more than one Short")
	}
	// invalid Examples
	tc14 := Command{"name", Examples{}}
	if e := tc14.Validate(); e == nil {
		t.Error("validator invalid Usage")
	}
	//invalid Var
	tc15 := Command{"name", Var{}}
	if e := tc15.Validate(); e == nil {
		t.Error("validator accepted invalid Var")
	}
	//invalid Trigger
	tc16 := Command{"name", Trigger{}}
	if e := tc16.Validate(); e == nil {
		t.Error("validator accepted invalid Trigger")
	}
	//Brief field present
	tc17 := Command{"name", func(Tri) int { return 1 }}
	if e := tc17.Validate(); e == nil {
		t.Error("validator accepted Command without a Brief")
	}
	//Handler present
	tc18 := Command{"name", Brief{""}}
	if e := tc18.Validate(); e == nil {
		t.Error("validator accepted Command without a handler")
	}
	//invalid typed element
	tc19 := Command{"name", Brief{""}, func(Tri) int { return 1 }, 1}
	if e := tc19.Validate(); e == nil {
		t.Error("validator accepted Command without a Brief")
	}
	// no errors!
	tc20 := Command{"name", Brief{""}, func(Tri) int { return 1 }}
	if e := tc20.Validate(); e != nil {
		t.Error("validator rejected valid Command")
	}
}

func TestCommands(t *testing.T) {
	tcc1 := Commands{
		Command{"name", Brief{""}, func(Tri) int { return 1 }, 1},
	}
	if e := tcc1.Validate(); e == nil {
		t.Error("validator accepted Commands with invalid element")
	}
	tcc2 := Commands{Command{"name", Brief{""}, func(Tri) int { return 1 }}}
	if e := tcc2.Validate(); e != nil {
		t.Error("validator rejected valid Commands")
	}
}

func TestDefault(t *testing.T) {

	// only one item
	td1 := Default{
		"item1", "item2",
	}
	e := td1.Validate()
	if e == nil {
		t.Error("validator allowed more than one")
	}

	// item is string
	td2 := Default{
		1,
	}
	e = td2.Validate()
	if e == nil {
		t.Error("validator permitted other than a string")
	}

	// item is a ValidName
	td3 := Default{
		"abc123",
	}
	e = td3.Validate()
	if e == nil {
		t.Error("validator permitted an invalid name")
	}

	// item is a valid
	td4 := Default{
		"abc",
	}
	e = td4.Validate()
	if e != nil {
		t.Error("validator rejected a valid name")
	}

}

func TestDefaultCommand(t *testing.T) {

	// only one item
	tdc1 := DefaultCommand{
		"item1", "item2",
	}
	e := tdc1.Validate()
	if e == nil {
		t.Error("validator allowed more than one")
	}

	// item is string
	tdc2 := DefaultCommand{
		1,
	}
	e = tdc2.Validate()
	if e == nil {
		t.Error("validator permitted other than a string")
	}

	// item is a ValidName
	tdc3 := DefaultCommand{
		"abc123",
	}
	e = tdc3.Validate()
	if e == nil {
		t.Error("validator permitted an invalid name")
	}

	// item is a valid
	tdc4 := DefaultCommand{
		"abc",
	}
	e = tdc4.Validate()
	if e != nil {
		t.Error("validator rejected a valid name")
	}
}

func TestDefaultOn(t *testing.T) {

	// must be empty
	tdo1 := DefaultOn{1}
	e := tdo1.Validate()
	if e == nil {
		t.Error("validator permitted content in DefaultOn")
	}

	// is empty
	tdo2 := DefaultOn{}
	e = tdo2.Validate()
	if e != nil {
		t.Error("validator rejected valid empty DefaultOn")
	}

}

func TestExamples(t *testing.T) {

	// must not be empty
	te1 := Examples{}
	e := te1.Validate()
	if e == nil {
		t.Error("validator invalid empty Examples")
	}

	// must have pairs of elements
	te2 := Examples{"1", 2, 3.0}
	e = te2.Validate()
	if e == nil {
		t.Error("validator permitted odd number of items in Examples")
	}

	// elements must be strings
	te3 := Examples{"1", 2}
	e = te3.Validate()
	if e == nil {
		t.Error("validator permitted non-string in Examples")
	}

	// string length < 80
	te4 := Examples{
		"--max=0", "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
	}
	e = te4.Validate()
	if e == nil {
		t.Error("validator permitted over 80 characters")
	}

	// string length < 80
	te5 := Examples{
		"a1234567890123456789012345678901234567890", "123",
	}
	e = te5.Validate()
	if e == nil {
		t.Error("validator permitted over 80 characters")
	}

	// even numbered (first in pair) elements have no control characters
	te6 := Examples{
		"--max=0", "aaa\n",
	}
	e = te6.Validate()
	if e == nil {
		t.Error("validator permitted control character in explainer")
	}

	// no error
	te7 := Examples{
		"--max=0", "aaaaaaaa",
	}
	e = te7.Validate()
	if e != nil {
		t.Error("validator rejected valid example")
	}

}

func TestGroup(t *testing.T) {

	// contains only one element
	tg1 := Group{1, 1}
	tg2 := Group{}
	e := tg1.Validate()
	if e == nil {
		t.Error("validator accepted more than one")
	}
	e = tg2.Validate()
	if e == nil {
		t.Error("validator accepted no elements")
	}

	// element is a string
	tg3 := Group{1}
	e = tg3.Validate()
	if e == nil {
		t.Error("validator accepted non string elemeent")
	}

	// string is a ValidName
	tg4 := Group{"abc123"}
	e = tg4.Validate()
	if e == nil {
		t.Error("validator accepted invalid name")
	}

	// no error!
	tg5 := Group{"abc"}
	e = tg5.Validate()
	if e != nil {
		t.Error("validator rejected valid name")
	}

}

func TestHelp(t *testing.T) {

	// contains only one element
	th1 := Help{1, 1}
	e := th1.Validate()
	if e == nil {
		t.Error("validator accepted more than one")
	}
	th2 := Help{}
	e = th2.Validate()
	if e == nil {
		t.Error("validator accepted no elements")
	}

	// element is a string
	th3 := Help{1.0}
	e = th3.Validate()
	if e == nil {
		t.Error("validator accepted non-string")
	}

	// no error!
	th4 := Help{"this is a test with cr\nand other things"}
	e = th4.Validate()
	if e != nil {
		t.Error("validator rejected valid element")
	}

}

func TestRunAfter(t *testing.T) {

	// may not contain anything
	tra1 := RunAfter{
		"",
	}
	e := tra1.Validate()
	if e == nil {
		t.Error("validator accepted content in RunAfter")
	}
	// no error
	tra2 := RunAfter{}
	e = tra2.Validate()
	if e != nil {
		t.Error("validator rejected valid RunAfter")
	}

}

func TestShort(t *testing.T) {

	// contains only one element
	ts1 := Short{'1', 2}
	e := ts1.Validate()
	if e == nil {
		t.Error("validator accepted more than one item")
	}

	// element is a rune (single character/unicode point)
	ts2 := Short{1}
	e = ts2.Validate()
	if e == nil {
		t.Error("validator accepted a non-rune element")
	}

	// element is a letter or number
	ts3 := Short{'!'}
	e = ts3.Validate()
	if e == nil {
		t.Error("validator accepted non alphanumeric element")
	}

	// no error!
	ts4 := Short{'a'}
	e = ts4.Validate()
	if e != nil {
		t.Error("validator rejected a valid short element")
	}

}

func TestSlot(t *testing.T) {

	// slots are all the same type (pointer to said type)
	a := 1
	b := "string"
	ts1 := Slot{&a, &b}
	e := ts1.Validate()
	if e == nil {
		t.Error("validator accepted heteregenous types")
	}

	// slots are all pointers
	c := 2
	ts2 := Slot{a, c}
	e = ts2.Validate()
	if e == nil {
		t.Error("validator accepted heteregenous types")
	}

	// no error!
	ts3 := Slot{&a, &c}
	e = ts3.Validate()
	if e != nil {
		t.Error("validator rejected valid contents")
	}

}

func TestTerminates(t *testing.T) {
	// must be empty
	tt1 := Terminates{1}
	e := tt1.Validate()
	if e == nil {
		t.Error("validator permitted content")
	}

	// is empty
	tt2 := Terminates{}
	e = tt2.Validate()
	if e != nil {
		t.Error("validator rejected valid empty Terminates")
	}

}

func TestTri(t *testing.T) {

	// contains at least 4 elements

	// first element is a string

	// string is a ValidName

	// contains (only) one Brief

	// contains (only) one Version

	// contains (only) one Commands

	// contains no more than one DefaultCommand

	// only contains element from set of possible elements

	// Brief is missing

	// Version is missing

	// Commands is missing

}

func TestTrigger(t *testing.T) {

	// must contain name, Brief and Handler

	// first element is name

	// name is a ValidName

	// only one Brief

	// only one Handler

	// rest of items only (one of) Short, Usage, Help, DefaultOn, Terminates and RunAfter

}

func TestUsage(t *testing.T) {

	// only one element
	tu1 := Usage{}
	tu2 := Usage{1, 2}
	e := tu1.Validate()
	if e == nil {
		t.Error("validator permitted empty Usage")
	}
	e = tu2.Validate()
	if e == nil {
		t.Error("validator permitted more than one element in Usage")
	}

	// element is string
	tu3 := Usage{0.1}
	e = tu3.Validate()
	if e == nil {
		t.Error("validator permitted element that is not a string")
	}

	// string is no more than 80 chars long
	tu4 := Usage{
		"123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
	}
	e = tu4.Validate()
	if e == nil {
		t.Error("validator permitted over 80 characters")
	}

	// string contains no control characters
	tu5 := Usage{
		"aaa\n",
	}
	e = tu5.Validate()
	if e == nil {
		t.Error("validator permitted control character in explainer")
	}

	// no error!
	tu6 := Usage{
		"aaaaaaa",
	}
	e = tu6.Validate()
	if e != nil {
		t.Error("validator rejected valid input")
	}

}

func TestValidName(t *testing.T) {

	// name is 3 or more characters long
	tvn1 := "ab"
	e := ValidName(tvn1)
	if e == nil {
		t.Error("validator accepted string under 3 characters length")
	}

	// name is only composed of letters
	tvn2 := "ab3"
	e = ValidName(tvn2)
	if e == nil {
		t.Error("validator accepted non-letter characters")
	}

	// no error!
	tvn3 := "proper"
	e = ValidName(tvn3)
	if e != nil {
		t.Error("validator rejected valid namme")
	}

}

func TestVar(t *testing.T) {
	// contains at least 3 elements
	tv1 := Var{1, 1}
	if e := tv1.Validate(); e == nil {
		t.Error("Var must contain at least 3 elements")
	}
	// first is string
	tv2 := Var{1, 1, 1}
	if e := tv2.Validate(); e == nil {
		t.Error("first element must be a string")
	}
	// name is ValidName
	tv3 := Var{"a ", 1, 1}
	if e := tv3.Validate(); e == nil {
		t.Error("validator accepted invalid name")
	}
	// has only one Brief
	tv4 := Var{"aaaa", Brief{""}, Brief{""}}
	if e := tv4.Validate(); e == nil {
		t.Error("validator accepted more than one Brief")
	}
	// has only one Short
	tv5 := Var{"aaaa", Short{'a'}, Short{'a'}}
	if e := tv5.Validate(); e == nil {
		t.Error("validator allowed more than one Short")
	}
	// has only one Usage
	tv6 := Var{"aaaa", Usage{""}, Usage{""}}
	if e := tv6.Validate(); e == nil {
		t.Error("validator allowed more than one Usage")
	}
	// has only one Help
	tv7 := Var{"aaaa", Help{""}, Help{""}}
	if e := tv7.Validate(); e == nil {
		t.Error("validator allowed more than one Help")
	}
	// has only one Default
	tv8 := Var{"aaaa", Default{"aaa"}, Default{"aaa"}}
	if e := tv8.Validate(); e == nil {
		t.Error("validator allowed more than one Default")
	}
	// has only one Slot
	tstring := "valid string"
	tv9 := Var{"aaaa", Slot{&tstring}, Slot{&tstring}}
	if e := tv9.Validate(); e == nil {
		t.Error("validator allowed more than one Slot")
	}
	// has invalid Brief
	tv10 := Var{"aaaa", Brief{}, Short{""}}
	if e := tv10.Validate(); e == nil {
		t.Error("validator accepted invalid Brief")
	}
	// has invalid Short
	tv11 := Var{"aaaa", Brief{"aaaa"}, Short{1}}
	if e := tv11.Validate(); e == nil {
		t.Error("validator allowed invalid Short")
	}
	// has invalid Usage
	tv12 := Var{"aaaa", Brief{"aaaa"}, Usage{0.1}}
	if e := tv12.Validate(); e == nil {
		t.Error("validator allowed invalid Usage")
	}
	// has invalid Help
	tv13 := Var{"aaaa", Brief{""}, Help{"aaa", 1}}
	if e := tv13.Validate(); e == nil {
		t.Error("validator allowed invalid Help")
	}
	// has invalid Default
	tv14 := Var{"aaaa", Brief{"aaa"}, Default{1, 3}}
	if e := tv14.Validate(); e == nil {
		t.Error("validator allowed invalid Default")
	}
	// has invalid Slot
	tv15 := Var{"aaaa", Brief{tstring}, Slot{tstring}}
	if e := tv15.Validate(); e == nil {
		t.Error("validator allowed invalid Slot")
	}
	// has one each of Brief and Slot
	tv16 := Var{"aaaa", Brief{"aaa"}, Default{"aa"}}
	t.Log(tv16.Validate())
	if e := tv16.Validate(); e == nil {
		t.Error("validator allowed absence of Brief or Slot")
	}

	// has no other type than those foregoing
	tv17 := Var{"aaaa", Brief{tstring}, Slot{&tstring}, 1}
	if e := tv17.Validate(); e == nil {
		t.Error("validator rejected valid Var")
	}

	// no error!}
	tv18 := Var{"aaaa", Brief{tstring}, Slot{&tstring}}
	if e := tv18.Validate(); e == nil {
		t.Error("validator rejected valid Var")
	}

}

func TestVersion(t *testing.T) {

	// has no more than 4 fields
	tv1 := Version{1, 2, 3, 4, 5}
	e := tv1.Validate()
	if e == nil {
		t.Error("validator accepted more than three items")
	}

	// has at least 3 fields
	tv2 := Version{4, 5}
	e = tv2.Validate()
	if e == nil {
		t.Error("validator accepted less than three items")
	}

	// first three are integers
	tv3 := Version{1.1, 2, 3, 4}
	e = tv3.Validate()
	if e == nil {
		t.Error("validator accepted non-integer version numbers")
	}

	// integers are under 100
	tv4 := Version{100, 2, 3, 4}
	e = tv4.Validate()
	if e == nil {
		t.Error("validator accepted a version number over 100")
	}

	// 4th field is a string
	tv5 := Version{10, 2, 3, 4}
	e = tv5.Validate()
	if e == nil {
		t.Error("validator accepted a 4th field that is not a string")
	}

	// string contains only letters and numbers
	tv6 := Version{10, 2, 3, "alpha3! "}
	e = tv6.Validate()
	if e == nil {
		t.Error(
			"validator accepted a 4th field that contains other than letters and numbers")
	}

	// no error!
	tv7 := Version{10, 2, 3, "alpha3"}
	e = tv7.Validate()
	if e != nil {
		t.Error(
			"validator accepted a 4th field that contains other than letters and numbers")
	}

}
