package service

import (
	"bytes"
	"encoding/json"

	"github.com/hawkkiller/timely/internal/model"
	"github.com/pkg/errors"
)

// InsertSchedule implements ScheduleService
func (s *scheduleService) InsertSchedule(link string) error {
	resp, err := s.client.Get(link)
	if err != nil {
		return errors.Wrap(err, "can't get schedule from link")
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	_, err = resp.Body.Read(buf.Bytes())
	if err != nil {
		return err
	}

	var schedule []model.Schedule

	err = json.Unmarshal(buf.Bytes(), &schedule)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal schedule")
	}

	return nil
}
