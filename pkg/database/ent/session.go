// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/willie-lin/cloud-terminal/pkg/database/ent/session"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Protocol holds the value of the "protocol" field.
	Protocol string `json:"protocol,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// Port holds the value of the "port" field.
	Port int `json:"port,omitempty"`
	// ConnectionID holds the value of the "connection_id" field.
	ConnectionID string `json:"connection_id,omitempty"`
	// AssetID holds the value of the "asset_id" field.
	AssetID string `json:"asset_id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator string `json:"creator,omitempty"`
	// ClientIP holds the value of the "client_ip" field.
	ClientIP string `json:"client_ip,omitempty"`
	// Width holds the value of the "width" field.
	Width int `json:"width,omitempty"`
	// Height holds the value of the "height" field.
	Height int `json:"height,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Recording holds the value of the "recording" field.
	Recording string `json:"recording,omitempty"`
	// PrivateKey holds the value of the "private_key" field.
	PrivateKey string `json:"private_key,omitempty"`
	// Passphrase holds the value of the "passphrase" field.
	Passphrase string `json:"passphrase,omitempty"`
	// Code holds the value of the "code" field.
	Code int `json:"code,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Connected holds the value of the "connected" field.
	Connected time.Time `json:"connected,omitempty"`
	// Disconnected holds the value of the "disconnected" field.
	Disconnected time.Time `json:"disconnected,omitempty"`
	// Mode holds the value of the "mode" field.
	Mode string `json:"mode,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges SessionEdges `json:"edges"`
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// Assets holds the value of the assets edge.
	Assets []*Asset `json:"assets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AssetsOrErr returns the Assets value or an error if the edge
// was not loaded in eager-loading.
func (e SessionEdges) AssetsOrErr() ([]*Asset, error) {
	if e.loadedTypes[0] {
		return e.Assets, nil
	}
	return nil, &NotLoadedError{edge: "assets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldPort, session.FieldWidth, session.FieldHeight, session.FieldCode:
			values[i] = new(sql.NullInt64)
		case session.FieldID, session.FieldProtocol, session.FieldIP, session.FieldConnectionID, session.FieldAssetID, session.FieldUsername, session.FieldPassword, session.FieldCreator, session.FieldClientIP, session.FieldStatus, session.FieldRecording, session.FieldPrivateKey, session.FieldPassphrase, session.FieldMessage, session.FieldMode:
			values[i] = new(sql.NullString)
		case session.FieldCreatedAt, session.FieldUpdatedAt, session.FieldConnected, session.FieldDisconnected:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case session.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case session.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case session.FieldProtocol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field protocol", values[i])
			} else if value.Valid {
				s.Protocol = value.String
			}
		case session.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				s.IP = value.String
			}
		case session.FieldPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				s.Port = int(value.Int64)
			}
		case session.FieldConnectionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field connection_id", values[i])
			} else if value.Valid {
				s.ConnectionID = value.String
			}
		case session.FieldAssetID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field asset_id", values[i])
			} else if value.Valid {
				s.AssetID = value.String
			}
		case session.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				s.Username = value.String
			}
		case session.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				s.Password = value.String
			}
		case session.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				s.Creator = value.String
			}
		case session.FieldClientIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_ip", values[i])
			} else if value.Valid {
				s.ClientIP = value.String
			}
		case session.FieldWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				s.Width = int(value.Int64)
			}
		case session.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				s.Height = int(value.Int64)
			}
		case session.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = value.String
			}
		case session.FieldRecording:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recording", values[i])
			} else if value.Valid {
				s.Recording = value.String
			}
		case session.FieldPrivateKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_key", values[i])
			} else if value.Valid {
				s.PrivateKey = value.String
			}
		case session.FieldPassphrase:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field passphrase", values[i])
			} else if value.Valid {
				s.Passphrase = value.String
			}
		case session.FieldCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = int(value.Int64)
			}
		case session.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				s.Message = value.String
			}
		case session.FieldConnected:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field connected", values[i])
			} else if value.Valid {
				s.Connected = value.Time
			}
		case session.FieldDisconnected:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field disconnected", values[i])
			} else if value.Valid {
				s.Disconnected = value.Time
			}
		case session.FieldMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mode", values[i])
			} else if value.Valid {
				s.Mode = value.String
			}
		}
	}
	return nil
}

// QueryAssets queries the "assets" edge of the Session entity.
func (s *Session) QueryAssets() *AssetQuery {
	return (&SessionClient{config: s.config}).QueryAssets(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", protocol=")
	builder.WriteString(s.Protocol)
	builder.WriteString(", ip=")
	builder.WriteString(s.IP)
	builder.WriteString(", port=")
	builder.WriteString(fmt.Sprintf("%v", s.Port))
	builder.WriteString(", connection_id=")
	builder.WriteString(s.ConnectionID)
	builder.WriteString(", asset_id=")
	builder.WriteString(s.AssetID)
	builder.WriteString(", username=")
	builder.WriteString(s.Username)
	builder.WriteString(", password=")
	builder.WriteString(s.Password)
	builder.WriteString(", creator=")
	builder.WriteString(s.Creator)
	builder.WriteString(", client_ip=")
	builder.WriteString(s.ClientIP)
	builder.WriteString(", width=")
	builder.WriteString(fmt.Sprintf("%v", s.Width))
	builder.WriteString(", height=")
	builder.WriteString(fmt.Sprintf("%v", s.Height))
	builder.WriteString(", status=")
	builder.WriteString(s.Status)
	builder.WriteString(", recording=")
	builder.WriteString(s.Recording)
	builder.WriteString(", private_key=")
	builder.WriteString(s.PrivateKey)
	builder.WriteString(", passphrase=")
	builder.WriteString(s.Passphrase)
	builder.WriteString(", code=")
	builder.WriteString(fmt.Sprintf("%v", s.Code))
	builder.WriteString(", message=")
	builder.WriteString(s.Message)
	builder.WriteString(", connected=")
	builder.WriteString(s.Connected.Format(time.ANSIC))
	builder.WriteString(", disconnected=")
	builder.WriteString(s.Disconnected.Format(time.ANSIC))
	builder.WriteString(", mode=")
	builder.WriteString(s.Mode)
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
