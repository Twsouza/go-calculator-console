package calculator

import "testing"

func Test_Do_Multiplication(t *testing.T) {
	op := "*"
	equation := "6*16+2*8"
	expectedOutput := "96.000000+16.000000"

	calc := New()
	output := calc.do(op, ops[op], equation)
	if output != expectedOutput {
		t.Fatalf("Expected to return %s but got %s", expectedOutput, output)
	}
}

func Test_Do_Multiplication_Float(t *testing.T) {
	op := "*"
	equation := "6*16+2*8+1.7*2"
	expectedOutput := "96.000000+16.000000+3.400000"

	calc := New()
	output := calc.do(op, ops[op], equation)
	if output != expectedOutput {
		t.Fatalf("Expected to return %s but got %s", expectedOutput, output)
	}
}

func Test_Do_Division(t *testing.T) {
	op := "/"
	equation := "36/6+40/8"
	expectedOutput := "6.000000+5.000000"

	calc := New()
	output := calc.do(op, ops[op], equation)
	if output != expectedOutput {
		t.Fatalf("Expected to return %s but got %s", expectedOutput, output)
	}
}

func Test_Do_Division_Float(t *testing.T) {
	op := "/"
	equation := "36/6+40/8-3.4/2"
	expectedOutput := "6.000000+5.000000-1.700000"

	calc := New()
	output := calc.do(op, ops[op], equation)
	if output != expectedOutput {
		t.Fatalf("Expected to return %s but got %s", expectedOutput, output)
	}
}

func Test_Do_Addition(t *testing.T) {
	op := "+"
	equation := "6+16+2+0.5-8"
	expectedOutput := "24.500000-8"

	calc := New()
	output := calc.do(op, ops[op], equation)
	if output != expectedOutput {
		t.Fatalf("Expected to return %s but got %s", expectedOutput, output)
	}
}

func Test_Do_SubSubtraction(t *testing.T) {
	op := "-"
	equation := "6+16.5-2-8+1"
	expectedOutput := "6+6.500000+1"

	calc := New()
	output := calc.do(op, ops[op], equation)
	if output != expectedOutput {
		t.Fatalf("Expected to return %s but got %s", expectedOutput, output)
	}
}

func Test_Calculate_FirstEquation(t *testing.T) {
	equation := "07+8*0-2"
	expected := "5"
	calc := New()
	result := calc.Calculate(equation)
	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func Test_Calculate_SecondEquation(t *testing.T) {
	equation := "0*02+2+0-2*0/2"
	expected := "2"
	calc := New()
	result := calc.Calculate(equation)
	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}

func Test_Calculate_EquationWithFloat(t *testing.T) {
	equation := "016.5+2.1*2-1.3"
	expected := "19.4"
	calc := New()
	result := calc.Calculate(equation)
	if result != expected {
		t.Fatalf("Expected %s but got %s", expected, result)
	}
}
