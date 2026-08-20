package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask28(t *testing.T) {
	now := time.Now()
	r := NewRegistry()
	base := Inspection{ID: "i", StartedAt: now.Add(-time.Hour), CompletedAt: now, Sections: []InspectionSection{{Code: "food", Score: 90, Evidence: []string{"p"}}}}
	require.NoError(t, r.SaveInspection(context.Background(), base, 0))
	base.Version = 1
	require.NoError(t, r.SaveInspection(context.Background(), base, 1))
	base.Version = 2
	s := NewService(r, func() time.Time { return now })
	_, err := s.CompleteInspection(context.Background(), base, 80)
	require.NoError(t, err)
}
