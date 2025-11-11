package reverse

import "testing"

func TestReverseASCII(t *testing.T) {
	if got := Reverse("hello"); got != "olleh" {
		t.Fatalf("want %q, got %q", "olleh", got)
	}
}

func TestReverseEmoji(t *testing.T) {
	if got := Reverse("Go🙂"); got != "🙂oG" {
		t.Fatalf("want %q, got %q", "🙂oG", got)
	}
}

func TestReverseCJK(t *testing.T) {
	if got := Reverse("日本"); got != "本日" {
		t.Fatalf("want %q, got %q", "本日", got)
	}
}
