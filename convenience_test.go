package xmlpath

import (
	"strings"
	"testing"
)

var testXML = `
<?xml version="1.0"?>
<library>
  <!-- Great book. -->
  <book id="b0836217462" available="true">
    <isbn>0836217462</isbn>
    <title lang="en">Being a Dog Is a Full-Time Job</title>
    <quote>I'd dog paddle the deepest ocean.</quote>
    <author id="CMS">
      <?echo "go rocks"?>
      <name>Charles M Schulz</name>
      <born>1922-11-26</born>
      <dead>2000-02-12</dead>
    </author>
    <character id="PP">
      <name>Peppermint Patty</name>
      <born>1966-08-22</born>
      <qualification>bold, brash and tomboyish</qualification>
    </character>
    <character id="Snoopy">
      <name>Snoopy</name>
      <born>1950-10-04</born>
      <qualification>extroverted beagle</qualification>
    </character>
    <character id="Schroeder">
      <name>Schroeder</name>
      <born>1951-05-30</born>
      <qualification>brought classical music to the Peanuts strip</qualification>
    </character>
    <character id="Lucy">
      <name>Lucy</name>
      <born>1952-03-03</born>
      <qualification>bossy, crabby and selfish</qualification>
    </character>
  </book>
  <!-- Another great book. -->
  <book id="b0883556316" available="true">
    <isbn>0883556316</isbn>
    <title lang="en">Barney <i>Google</i> and Snuffy Smith</title>
    <author id="CMS">
      <name>Charles M Schulz</name>
      <born>1922-11-26</born>
      <dead>2000-02-12</dead>
    </author>
    <character id="Barney">
      <name>Barney Google</name>
      <born>1919-01-01</born>
      <qualification>goggle-eyed, moustached, gloved and top-hatted, bulbous-nosed, cigar-chomping shrimp</qualification>
    </character>
    <character id="Spark">
      <name>Spark Plug</name>
      <born>1922-07-17</born>
      <qualification>brown-eyed, bow-legged nag, seldom races, patched blanket</qualification>
    </character>
    <character id="Snuffy">
      <name>Snuffy Smith</name>
      <born>1934-01-01</born>
      <qualification>volatile and diminutive moonshiner, ornery little cuss, sawed-off and shiftless</qualification>
    </character>
  </book>
</library>
`

func TestFindAllString(t *testing.T) {
	func() {
		result, err := FindAllString(testXML, "//book/title")
		if err != nil {
			t.Errorf("FindAllString() failed with error: %s\n", err.Error())
		}
		expected := "Being a Dog Is a Full-Time Job|Barney Google and Snuffy Smith"
		got := strings.Join(result, "|")
		if ok := (expected == got); !ok {
			t.Errorf("FindAllString() failed, expected: %s, got: %s\n", expected, got)
		}
	}()
	func() {
		result, err := FindAllString(testXML, "//book/character/name")
		if err != nil {
			t.Errorf("FindAllString() failed with error: %s\n", err.Error())
		}
		expected := "Peppermint Patty|Snoopy|Schroeder|Lucy|Barney Google|Spark Plug|Snuffy Smith"
		got := strings.Join(result, "|")
		if ok := (expected == got); !ok {
			t.Errorf("FindAllString() failed, expected: %s, got: %s\n", expected, got)
		}
	}()
	func() {
		result, err := FindAllString(testXML, "//book/character[@id='Snuffy']/qualification")
		if err != nil {
			t.Errorf("FindAllString() failed with error: %s\n", err.Error())
		}
		expected := "volatile and diminutive moonshiner, ornery little cuss, sawed-off and shiftless"
		got := result[0]
		if ok := (expected == got && len(result) == 1); !ok {
			t.Errorf("FindAllString() failed, expected: %s, got: %s\n", expected, got)
		}
	}()
	func() {
		result, err := FindAllString(testXML, "//book/foo")
		if err != nil {
			t.Errorf("FindAllString() failed with error: %s\n", err.Error())
		}
		if ok := (len(result) == 0); !ok {
			t.Errorf("FindAllString() failed, expected len(result) to be 0, result was %d", len(result))
		}
	}()
}

func TestFindString(t *testing.T) {
	func() {
		result, err := FindString(testXML, "//book/character[@id='Snuffy']/qualification")
		if err != nil {
			t.Errorf("FindString() failed with error: %s\n", err.Error())
		}
		expected := "volatile and diminutive moonshiner, ornery little cuss, sawed-off and shiftless"
		got := result
		if ok := (expected == got); !ok {
			t.Errorf("FindString() failed, expected: %s, got: %s\n", expected, got)
		}
	}()
	func() {
		result, err := FindString(testXML, "//book/character/name")
		if err != nil {
			t.Errorf("FindString() failed with error: %s\n", err.Error())
		}
		expected := "Peppermint Patty"
		if ok := (expected == result); !ok {
			t.Errorf("FindString() failed, expected: %s, got: %s\n", expected, result)
		}
	}()
	func() {
		_, err := FindString(testXML, "//book/foo")
		if ok := (err.Error() == "No Strings found for given xpath"); !ok {
			t.Errorf("FindString() failed, expected error to render, instead error was: %s", err.Error())
		}
	}()
}
