package hashset

import "testing"

func initWithTesting(t *testing.T) *MyHashSet {
	mhs := Constructor()
	t.Cleanup(func() {
		mhs = nil
	})
	return mhs
}

func TestAdd(t *testing.T) {
	mhs := initWithTesting(t)
	st := []struct {
		name string
		key  int
		exp  bool
	}{
		{"add 1", 1, true},
		{"add 2", 2, true},
		{"add 3", 3, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			mhs.Add(tt.key)
			if mhs.Contains(tt.key) != tt.exp {
				t.Fatalf("afte add key:%d, it should exists in the set", tt.key)
			}
			t.Log("pass")
		})
	}
}

func TestContains(t *testing.T) {
	mhs := initWithTesting(t)
	st := []struct {
		name    string
		key     int
		beAdded bool
		exp     bool
	}{
		{"with out add 1", 1, false, false},
		{"after add 1", 1, true, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beAdded {
				mhs.Add(tt.key)
			}
			if mhs.Contains(tt.key) != tt.exp {
				t.Fatalf("check key:%d in set, should %t but got %t", tt.key, tt.exp, !tt.exp)
			}
			t.Log("pass")
		})
	}
}

func TestRemove(t *testing.T) {
	mhs := initWithTesting(t)
	mhs.Add(1)
	mhs.Add(2)
	mhs.Add(3)
	mhs.Add(4)
	st := []struct {
		name string
		key  int
		exp  bool
	}{
		{"after remove 1, calling contains should return false", 1, false},
		{"after remove 3, calling contains should return false", 3, false},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			mhs.Remove(tt.key)
			if mhs.Contains(tt.key) != tt.exp {
				t.Fatalf("%s, %t", tt.name, tt.exp)
			}
			t.Log("pass")
		})
	}

}
