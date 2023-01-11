// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/live"
)

// Live is the model entity for the Live schema.
type Live struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Watch live streams
	WatchLive bool `json:"watch_live,omitempty"`
	// Watch new VODs
	WatchVod bool `json:"watch_vod,omitempty"`
	// Download archives
	DownloadArchives bool `json:"download_archives,omitempty"`
	// Download highlights
	DownloadHighlights bool `json:"download_highlights,omitempty"`
	// Download uploads
	DownloadUploads bool `json:"download_uploads,omitempty"`
	// Whether the channel is currently live.
	IsLive bool `json:"is_live,omitempty"`
	// Whether the chat archive is enabled.
	ArchiveChat bool `json:"archive_chat,omitempty"`
	// Resolution holds the value of the "resolution" field.
	Resolution string `json:"resolution,omitempty"`
	// The time the channel last went live.
	LastLive time.Time `json:"last_live,omitempty"`
	// Whether the chat should be rendered.
	RenderChat bool `json:"render_chat,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LiveQuery when eager-loading is set.
	Edges        LiveEdges `json:"edges"`
	channel_live *uuid.UUID
}

// LiveEdges holds the relations/edges for other nodes in the graph.
type LiveEdges struct {
	// Channel holds the value of the channel edge.
	Channel *Channel `json:"channel,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ChannelOrErr returns the Channel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LiveEdges) ChannelOrErr() (*Channel, error) {
	if e.loadedTypes[0] {
		if e.Channel == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: channel.Label}
		}
		return e.Channel, nil
	}
	return nil, &NotLoadedError{edge: "channel"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Live) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case live.FieldWatchLive, live.FieldWatchVod, live.FieldDownloadArchives, live.FieldDownloadHighlights, live.FieldDownloadUploads, live.FieldIsLive, live.FieldArchiveChat, live.FieldRenderChat:
			values[i] = new(sql.NullBool)
		case live.FieldResolution:
			values[i] = new(sql.NullString)
		case live.FieldLastLive, live.FieldUpdatedAt, live.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case live.FieldID:
			values[i] = new(uuid.UUID)
		case live.ForeignKeys[0]: // channel_live
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Live", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Live fields.
func (l *Live) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case live.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case live.FieldWatchLive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field watch_live", values[i])
			} else if value.Valid {
				l.WatchLive = value.Bool
			}
		case live.FieldWatchVod:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field watch_vod", values[i])
			} else if value.Valid {
				l.WatchVod = value.Bool
			}
		case live.FieldDownloadArchives:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field download_archives", values[i])
			} else if value.Valid {
				l.DownloadArchives = value.Bool
			}
		case live.FieldDownloadHighlights:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field download_highlights", values[i])
			} else if value.Valid {
				l.DownloadHighlights = value.Bool
			}
		case live.FieldDownloadUploads:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field download_uploads", values[i])
			} else if value.Valid {
				l.DownloadUploads = value.Bool
			}
		case live.FieldIsLive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_live", values[i])
			} else if value.Valid {
				l.IsLive = value.Bool
			}
		case live.FieldArchiveChat:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field archive_chat", values[i])
			} else if value.Valid {
				l.ArchiveChat = value.Bool
			}
		case live.FieldResolution:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resolution", values[i])
			} else if value.Valid {
				l.Resolution = value.String
			}
		case live.FieldLastLive:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_live", values[i])
			} else if value.Valid {
				l.LastLive = value.Time
			}
		case live.FieldRenderChat:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field render_chat", values[i])
			} else if value.Valid {
				l.RenderChat = value.Bool
			}
		case live.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case live.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case live.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field channel_live", values[i])
			} else if value.Valid {
				l.channel_live = new(uuid.UUID)
				*l.channel_live = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryChannel queries the "channel" edge of the Live entity.
func (l *Live) QueryChannel() *ChannelQuery {
	return (&LiveClient{config: l.config}).QueryChannel(l)
}

// Update returns a builder for updating this Live.
// Note that you need to call Live.Unwrap() before calling this method if this Live
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Live) Update() *LiveUpdateOne {
	return (&LiveClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Live entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Live) Unwrap() *Live {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Live is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Live) String() string {
	var builder strings.Builder
	builder.WriteString("Live(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("watch_live=")
	builder.WriteString(fmt.Sprintf("%v", l.WatchLive))
	builder.WriteString(", ")
	builder.WriteString("watch_vod=")
	builder.WriteString(fmt.Sprintf("%v", l.WatchVod))
	builder.WriteString(", ")
	builder.WriteString("download_archives=")
	builder.WriteString(fmt.Sprintf("%v", l.DownloadArchives))
	builder.WriteString(", ")
	builder.WriteString("download_highlights=")
	builder.WriteString(fmt.Sprintf("%v", l.DownloadHighlights))
	builder.WriteString(", ")
	builder.WriteString("download_uploads=")
	builder.WriteString(fmt.Sprintf("%v", l.DownloadUploads))
	builder.WriteString(", ")
	builder.WriteString("is_live=")
	builder.WriteString(fmt.Sprintf("%v", l.IsLive))
	builder.WriteString(", ")
	builder.WriteString("archive_chat=")
	builder.WriteString(fmt.Sprintf("%v", l.ArchiveChat))
	builder.WriteString(", ")
	builder.WriteString("resolution=")
	builder.WriteString(l.Resolution)
	builder.WriteString(", ")
	builder.WriteString("last_live=")
	builder.WriteString(l.LastLive.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("render_chat=")
	builder.WriteString(fmt.Sprintf("%v", l.RenderChat))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Lives is a parsable slice of Live.
type Lives []*Live

func (l Lives) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
