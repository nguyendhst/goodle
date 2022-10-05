package goodle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSiteInfo(t *testing.T) {
	m := NewMoodleAPI(BASE, TOKEN)
	a, err := m.GetSiteInfo()
	require.Empty(t, err)
	t.Log(a)
}

func TestGetUnreadConversationsCount(t *testing.T) {
	m := NewMoodleAPI(BASE, TOKEN)
	a, err := m.GetUnreadConversationsCount()
	require.Empty(t, err)
	t.Log(a)
}
