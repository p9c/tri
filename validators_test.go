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

	//string is valid name (letters only)

	//more than one brief not allowed

	//more than one handler not allowed

	//handler not nil

	//Brief field present

	//Handler present

}

func TestCommands(t *testing.T) {

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

	// contains only one elemennt

	// element is a rune (single character/unicode point)

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
	if e == nil {
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

	// element is string

	// string is no more than 80 chars long

	// string contains no control characters

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

	// has no more than 4 fields

	// has at least 3 fields

	// first three are integers

	// integers are under 100

	// field 4 is a string

	// string contains only letters and numbers

}

func TestVersion(t *testing.T) {

	//contains only letters

}
