package valueobjects

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Timestamps struct {
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewTimestamps() Timestamps {
	ts := Timestamps{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return ts
}

func (ts *Timestamps) Update() {
	ts.UpdatedAt = time.Now()
}

type AuditTrail struct {
	Timestamps

	CreatedBy uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedBy uuid.UUID `db:"updated_by" json:"updated_by"`
}

func NewAuditTrail(userID uuid.UUID) AuditTrail {
	ts := AuditTrail{
		Timestamps: NewTimestamps(),
		CreatedBy:  userID,
		UpdatedBy:  uuid.Nil,
	}
	return ts
}

func (ts *AuditTrail) Update(userID uuid.UUID) {
	ts.Timestamps.Update()
	ts.UpdatedBy = userID
}
