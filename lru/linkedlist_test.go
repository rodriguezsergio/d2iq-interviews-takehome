package lru

import "testing"

func TestDoublyLinkedList_AddNode(t *testing.T) {
	tests := map[string]struct {
		Input []string
		Begin string
		End   string
	}{
		"single": {
			Input: []string{"foobar"},
			Begin: "foobar",
			End:   "foobar",
		},
		"triple": {
			Input: []string{"foo", "bar", "baz"},
			Begin: "baz",
			End:   "foo",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ll := &DoublyLinkedList{}
			for _, val := range test.Input {
				ll.AddNode(val)
			}
			if ll.Begin.Data != test.Begin {
				t.Errorf("expected '%v' got '%v'", test.Begin, ll.Begin.Data)
			}
			if ll.End.Data != test.End {
				t.Errorf("expected '%v' got '%v'", test.End, ll.End.Data)
			}
			if ll.Size != len(test.Input) {
				t.Errorf("expected '%v' got '%v'", len(test.Input), ll.Size)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveNode(t *testing.T) {
	tests := map[string]struct {
		Input  []string
		Remove []string
		Begin  interface{}
		End    interface{}
	}{
		"empty set": {
			Input:  []string{"foobar"},
			Remove: []string{"foobar"},
			Begin:  nil,
			End:    nil,
		},
		"remove middle node": {
			Input:  []string{"end", "middle", "start"},
			Remove: []string{"middle"},
			Begin:  "start",
			End:    "end",
		},
		"remove end node": {
			Input:  []string{"end", "middle", "start"},
			Remove: []string{"end"},
			Begin:  "start",
			End:    "middle",
		},
		"remove start node": {
			Input:  []string{"end", "middle", "start"},
			Remove: []string{"start"},
			Begin:  "middle",
			End:    "end",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ll := &DoublyLinkedList{}
			for _, val := range test.Input {
				ll.AddNode(val)
			}
			for _, val := range test.Remove {
				ll.RemoveNode(val)
			}

			if ll.Begin != nil && ll.End != nil {
				if ll.Begin.Data != test.Begin {
					t.Errorf("expected '%v' got '%v'", test.Begin, ll.Begin.Data)
				}
				if ll.End.Data != test.End {
					t.Errorf("expected '%v' got '%v'", test.End, ll.End.Data)
				}
			} else {
				if !(ll.Begin == nil && test.Begin == nil) {
					t.Errorf("expected '%v' got '%v'", test.Begin, ll.Begin)
				}
				if !(ll.End == nil && test.End == nil) {
					t.Errorf("expected '%v' got '%v'", test.End, ll.End)
				}
			}

			inputRemoveDiff := len(test.Input) - len(test.Remove)
			if ll.Size != inputRemoveDiff {
				t.Errorf("expected '%v' got '%v'", inputRemoveDiff, ll.Size)
			}
		})
	}
}
