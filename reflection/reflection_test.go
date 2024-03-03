package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Luke"},
			[]string{"Luke"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Luke", "Waterloo"},
			[]string{"Luke", "Waterloo"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Luke", 28},
			[]string{"Luke"},
		},
		{
			"struct with nested fields",
			Person{
				"Luke",
				Profile{28, "Waterloo"},
			},
			[]string{"Luke", "Waterloo"},
		},
		{
			"struct with pointers to things",
			&Person{
				"Luke",
				Profile{28, "Waterloo"},
			},
			[]string{"Luke", "Waterloo"},
		},
		{
			"test slices",
			[]Profile{
				{28, "Waterloo"},
				{26, "Markham"},
			},
			[]string{"Waterloo", "Markham"},
		},
		{
			"test arrays",
			[2]Profile{
				{28, "Waterloo"},
				{26, "Markham"},
			},
			[]string{"Waterloo", "Markham"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {

			var got []string

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}

		})
	}
	t.Run("test with Maps", func(t *testing.T) {
		testingMap := map[string]string{
			"Pig":  "Oink",
			"Fish": "Bloop",
		}

		var got []string
		Walk(testingMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Oink")
		assertContains(t, got, "Bloop")
	})
	t.Run("test with channels", func(t *testing.T) {
		testingChannel := make(chan Profile)

		go func() {
			testingChannel <- Profile{26, "Berlin"}
			testingChannel <- Profile{32, "Frankfurt"}
			close(testingChannel)
		}()

		var got []string
		want := []string{"Berlin", "Frankfurt"}

		Walk(testingChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("test with functions", func(t *testing.T) {
		testingFunction := func() (Profile, Profile) {
			return Profile{26, "Berlin"}, Profile{32, "Frankfurt"}
		}

		var got []string
		want := []string{"Berlin", "Frankfurt"}

		Walk(testingFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("Expected %v to contain %q but it was not found", haystack, needle)
	}
}
