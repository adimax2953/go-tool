package argtool

import (
	"encoding/json"
	"math"
	"testing"
)

func TestAssert(t *testing.T) {
	{
		var arg string = ""
		err := Assert(
			NonEmptyString(arg, "arg"),
		)
		if err == nil {
			t.Errorf("should get error")
		}
	}
}

func TestNonEmptyString(t *testing.T) {
	{
		var arg string = "1234"
		err := NonEmptyString(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg string = ""
		err := NonEmptyString(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
}

func TestNonNanNorInf(t *testing.T) {
	{
		var arg float64 = 1.0
		err := NonNanNorInf(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg float64 = math.Inf(1)
		err := NonNanNorInf(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg float64 = math.Inf(-1)
		err := NonNanNorInf(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg float64 = math.NaN()
		err := NonNanNorInf(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
}

func TestNonNegativeInteger(t *testing.T) {
	{
		var arg int64 = 1
		err := NonNegativeInteger(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = 0
		err := NonNegativeInteger(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg int64 = -1
		err := NonNegativeInteger(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
}

func TestJsonInteger(t *testing.T) {
	{
		var arg json.Number = "1"
		err := JsonInteger(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "1.5"
		err := JsonInteger(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg json.Number = "aaa"
		err := JsonInteger(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
}

func TestJsonFloatNumber(t *testing.T) {
	{
		var arg json.Number = "1"
		err := JsonNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "1.5"
		err := JsonNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "nan"
		err := JsonNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "inf"
		err := JsonNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "-inf"
		err := JsonNumber(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "aaa"
		err := JsonNumber(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
	}
}

func TestJsonFloatNumberWithNonNanNorInf(t *testing.T) {
	{
		var arg json.Number = "1"
		err := JsonNumber(arg, "arg", NonNanNorInf)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "1.5"
		err := JsonNumber(arg, "arg", NonNanNorInf)
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "nan"
		err := JsonNumber(arg, "arg", NonNanNorInf)
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg json.Number = "inf"
		err := JsonNumber(arg, "arg", NonNanNorInf)
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg json.Number = "-inf"
		err := JsonNumber(arg, "arg", NonNanNorInf)
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg json.Number = "aaa"
		err := JsonNumber(arg, "arg", NonNanNorInf)
		if err == nil {
			t.Errorf("should get error")
		}
	}
}

func TestJsonFloatNumberWithNonNanNorInf_RangeBetweenFloat(t *testing.T) {
	{
		var arg json.Number = "1"
		err := JsonNumber(arg, "arg",
			NonNanNorInf, RangeBetweenFloat(0, 1.49))
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "1.5"
		err := JsonNumber(arg, "arg",
			NonNanNorInf, RangeBetweenFloat(0, 1.49))
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg json.Number = "nan"
		err := JsonNumber(arg, "arg",
			NonNanNorInf, RangeBetweenFloat(0, 1.49))
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg json.Number = "inf"
		err := JsonNumber(arg, "arg",
			NonNanNorInf, RangeBetweenFloat(0, 1.49))
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg json.Number = "-inf"
		err := JsonNumber(arg, "arg",
			NonNanNorInf, RangeBetweenFloat(0, 1.49))
		if err == nil {
			t.Errorf("should get error")
		}
	}
	{
		var arg json.Number = "aaa"
		err := JsonNumber(arg, "arg",
			NonNanNorInf, RangeBetweenFloat(0, 1.49))
		if err == nil {
			t.Errorf("should get error")
		}
	}
}

func TestIntegerNotIn(t *testing.T) {
	{
		var arg json.Number = "1"
		err := JsonInteger(arg, "arg",
			IntegerNotIn(0, 7, 8))
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg json.Number = "0"
		err := JsonInteger(arg, "arg",
			IntegerNotIn(0, 7, 8))
		if err == nil {
			t.Errorf("should get error")
		}
	}
}
