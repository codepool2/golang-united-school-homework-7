package coverage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPerson2DateOfBirthIsElderThanPerson1DateOfBirth(t *testing.T) {

	//Arrange & Act
	result := generateTestData().Less(1, 0)

	//Assert
	assert.True(t, result)
}

func TestPerson2DateOfBirthSameAsPerson1DateOfBirthShouldSortOnFirstName(t *testing.T) {

	//Arrange
	people := generateTestData()
	per := &people[0]
	per1 := &people[1]
	per.birthDay = per1.birthDay

	//Act
	result := people.Less(0, 1)

	assert.True(t, result)
}

func TestPerson2AndPerson1withSameDOBAndFirstNameShouldgetSortedBasedOnLastName(t *testing.T) {

	//Arrange
	people := generateTestData()
	per := &people[0]
	per1 := &people[1]
	per.birthDay = per1.birthDay
	per.firstName = per1.firstName

	//Act
	result := people.Less(1, 0)

	//Assert
	assert.True(t, result)
}

func TestShouldSwapTomWithTomBrother(t *testing.T) {

	//Arrange
	people := generateTestData()

	//Act
	people.Swap(0, 1)

	//Assert
	assert.Equal(t, people[0].firstName, "tom_elder_brother")
	assert.Equal(t, people[1].firstName, "tom")

}

func TestShouldReturnLengthOf2(t *testing.T) {

	//Arrange & Act & Assert
	assert.Equal(t, generateTestData().Len(), 2)
}

func TestShouldReturnMatrix(t *testing.T) {

	matrix, _ := New(generateStrings())

	assert.NotNil(t, matrix)
	assert.Equal(t, len(matrix.data), matrix.cols*matrix.rows)
}

func TestShouldReturnNilForUnEvenStringSizes(t *testing.T) {

	s := generateStrings() + " 13"
	actualResult, _ := New(s)

	assert.Nil(t, actualResult)
}

func TestShouldReturnErrorForInvalidChar(t *testing.T) {

	//Arrange
	s := generateStrings() + "\n13 14 15 a"

	//Act
	_, err := New(s)

	//Assert
	_, ok := err.(*strconv.NumError)
	assert.True(t, ok)
}

func TestShouldReturnOnlyRows(t *testing.T) {

	matrix, _ := New(generateStrings())

	actualResult := matrix.Rows()

	assert.Equal(t, 3, len(actualResult))
	assert.Equal(t, 4, len(actualResult[0]))
	assert.Equal(t, 1, actualResult[0][0])
	assert.Equal(t, 2, actualResult[0][1])
	assert.Equal(t, 3, actualResult[0][2])
	assert.Equal(t, 4, actualResult[0][3])

}

func TestShouldReturnOnlyColums(t *testing.T) {

	matrix, _ := New(generateStrings())

	actualResult := matrix.Cols()

	assert.Equal(t, 4, len(actualResult))
	assert.Equal(t, 3, len(actualResult[0]))
	assert.Equal(t, 1, actualResult[0][0])
	assert.Equal(t, 5, actualResult[0][1])
	assert.Equal(t, 9, actualResult[0][2])

}

func TestSetShouldReturnFalseForNegativeRows(t *testing.T) {

	matrix, _ := New(generateStrings())

	assert.False(t, matrix.Set(-1, 0, 3))
}

func TestSetShouldReturnFalseForNegativeColumnValue(t *testing.T) {

	matrix, _ := New(generateStrings())

	assert.False(t, matrix.Set(0, -1, 3))
}

func TestSetShouldReturnFalseRowSizeGreaterThanMatrixRowSize(t *testing.T) {

	matrix, _ := New(generateStrings())

	assert.False(t, matrix.Set(4, 0, 3))
}

func TestSetShouldReturnFalseForColumnSizeGreaterThanMatrixColumnSize(t *testing.T) {

	matrix, _ := New(generateStrings())

	assert.False(t, matrix.Set(0, 5, 3))
}

func TestSetShouldReturnTrueAfterSettingValueAtGivenPosition(t *testing.T) {

	matrix, _ := New(generateStrings())
	assert.True(t, matrix.Set(1, 2, 70))
	assert.Equal(t, matrix.data[6], 70)
}
func generateStrings() string {

	return "1 2 3 4\n5 6 7 8\n9 10 11 12"
}

func generateTestData() People {

	person1 :=
		Person{
			"tom",
			"cruise",
			time.Date(1994, time.March, 01, 0, 0, 0, 0, time.UTC)}

	person2 :=
		Person{
			"tom_elder_brother",
			"Cruise",
			time.Date(1994, time.April, 01, 0, 0, 0, 0, time.UTC)}

	twoPersonsWithDiffDateOfBirth := []Person{person1, person2}

	return twoPersonsWithDiffDateOfBirth
}
