package feedback

import (
	"bytes"
	"context"
	"encoding/csv"
	"strconv"
)

type Service struct {
	Repo *Repository
}

func (s *Service) Create(ctx context.Context, feedback *Feedback) error {
	return s.Repo.Create(ctx, feedback)
}

func (s *Service) CSV(ctx context.Context) ([]byte, error) {
	feedbacks, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	records := make([][]string, 0, len(feedbacks))

	for _, f := range feedbacks {
		record := []string{strconv.Itoa(f.Stars), f.Comment}
		records = append(records, record)
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.WriteAll(records)

	return buf.Bytes(), nil
}
