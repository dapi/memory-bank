package agentinstructions

import (
	"bytes"
	"testing"
)

func TestBuildPlan(t *testing.T) {
	oldBlock := []byte(StartMarker + "\nold routing\n" + EndMarker + "\n")
	tests := []struct {
		name     string
		input    []byte
		status   Status
		wantData []byte
	}{
		{name: "empty file", status: Missing, wantData: CurrentBlock},
		{name: "existing content", input: []byte("user before\nuser after"), status: Missing, wantData: append([]byte("user before\nuser after\n\n"), CurrentBlock...)},
		{name: "current block", input: append([]byte(nil), CurrentBlock...), status: Current, wantData: CurrentBlock},
		{name: "outdated block", input: append(append([]byte("before\x00\n"), oldBlock...), []byte("after\n")...), status: Outdated, wantData: append(append([]byte("before\x00\n"), CurrentBlock...), []byte("after\n")...)},
		{name: "duplicate markers", input: []byte(StartMarker + "\n" + EndMarker + "\n" + StartMarker + "\n" + EndMarker), status: Ambiguous},
		{name: "missing end marker", input: []byte(StartMarker + "\ntext"), status: Ambiguous},
		{name: "damaged start marker", input: []byte("<!-- MEMORY BANK START v0 -->\ntext\n" + EndMarker), status: Ambiguous},
		{name: "lowercase damaged markers", input: []byte("<!-- memory bank start -->\nold\n<!-- memory bank end -->\n"), status: Ambiguous},
		{name: "mixed case damaged markers", input: []byte("<!-- Memory Bank Start -->\nold\n<!-- Memory Bank End -->\n"), status: Ambiguous},
		{name: "inline marker references", input: []byte("Document " + StartMarker + " and later " + EndMarker + " in prose.\n"), status: Ambiguous},
		{name: "indented markers", input: []byte("  " + StartMarker + "\ntext\n  " + EndMarker + "\n"), status: Ambiguous},
		{name: "trailing marker whitespace", input: []byte(StartMarker + " \ntext\n" + EndMarker + "\n"), status: Ambiguous},
		{name: "outdated CRLF block", input: []byte(StartMarker + "\r\nold\r\n" + EndMarker + "\r\nafter\r\n"), status: Outdated, wantData: append(append([]byte(nil), CurrentBlock...), []byte("after\r\n")...)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			plan := BuildPlan(test.input)
			if plan.Status != test.status || !bytes.Equal(plan.Data, test.wantData) {
				t.Fatalf("unexpected plan: status=%s data=%q", plan.Status, plan.Data)
			}
			if (test.status == Missing || test.status == Outdated) && plan.Diff == "" {
				t.Fatal("planned mutation has no diff")
			}
		})
	}
}
