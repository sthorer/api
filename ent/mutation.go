// github.com/sthorer/api

package ent

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sthorer/api/ent/file"
	"github.com/sthorer/api/ent/token"
	"github.com/sthorer/api/ent/user"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFile  = "File"
	TypeToken = "Token"
	TypeUser  = "User"
)

// FileMutation represents an operation that mutate the Files
// nodes in the graph.
type FileMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	hash          *string
	size          *int64
	addsize       *int64
	pinned_at     *time.Time
	unpinned_at   *time.Time
	metadata      *map[string]interface{}
	clearedFields map[string]struct{}
	user          *int
	cleareduser   bool
}

var _ ent.Mutation = (*FileMutation)(nil)

// newFileMutation creates new mutation for $n.Name.
func newFileMutation(c config, op Op) *FileMutation {
	return &FileMutation{
		config:        c,
		op:            op,
		typ:           TypeFile,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FileMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FileMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on File creation.
func (m *FileMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *FileMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetHash sets the hash field.
func (m *FileMutation) SetHash(s string) {
	m.hash = &s
}

// Hash returns the hash value in the mutation.
func (m *FileMutation) Hash() (r string, exists bool) {
	v := m.hash
	if v == nil {
		return
	}
	return *v, true
}

// ResetHash reset all changes of the hash field.
func (m *FileMutation) ResetHash() {
	m.hash = nil
}

// SetSize sets the size field.
func (m *FileMutation) SetSize(i int64) {
	m.size = &i
	m.addsize = nil
}

// Size returns the size value in the mutation.
func (m *FileMutation) Size() (r int64, exists bool) {
	v := m.size
	if v == nil {
		return
	}
	return *v, true
}

// AddSize adds i to size.
func (m *FileMutation) AddSize(i int64) {
	if m.addsize != nil {
		*m.addsize += i
	} else {
		m.addsize = &i
	}
}

// AddedSize returns the value that was added to the size field in this mutation.
func (m *FileMutation) AddedSize() (r int64, exists bool) {
	v := m.addsize
	if v == nil {
		return
	}
	return *v, true
}

// ResetSize reset all changes of the size field.
func (m *FileMutation) ResetSize() {
	m.size = nil
	m.addsize = nil
}

// SetPinnedAt sets the pinned_at field.
func (m *FileMutation) SetPinnedAt(t time.Time) {
	m.pinned_at = &t
}

// PinnedAt returns the pinned_at value in the mutation.
func (m *FileMutation) PinnedAt() (r time.Time, exists bool) {
	v := m.pinned_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetPinnedAt reset all changes of the pinned_at field.
func (m *FileMutation) ResetPinnedAt() {
	m.pinned_at = nil
}

// SetUnpinnedAt sets the unpinned_at field.
func (m *FileMutation) SetUnpinnedAt(t time.Time) {
	m.unpinned_at = &t
}

// UnpinnedAt returns the unpinned_at value in the mutation.
func (m *FileMutation) UnpinnedAt() (r time.Time, exists bool) {
	v := m.unpinned_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUnpinnedAt reset all changes of the unpinned_at field.
func (m *FileMutation) ResetUnpinnedAt() {
	m.unpinned_at = nil
}

// SetMetadata sets the metadata field.
func (m *FileMutation) SetMetadata(value map[string]interface{}) {
	m.metadata = &value
}

// Metadata returns the metadata value in the mutation.
func (m *FileMutation) Metadata() (r map[string]interface{}, exists bool) {
	v := m.metadata
	if v == nil {
		return
	}
	return *v, true
}

// ClearMetadata clears the value of metadata.
func (m *FileMutation) ClearMetadata() {
	m.metadata = nil
	m.clearedFields[file.FieldMetadata] = struct{}{}
}

// MetadataCleared returns if the field metadata was cleared in this mutation.
func (m *FileMutation) MetadataCleared() bool {
	_, ok := m.clearedFields[file.FieldMetadata]
	return ok
}

// ResetMetadata reset all changes of the metadata field.
func (m *FileMutation) ResetMetadata() {
	m.metadata = nil
	delete(m.clearedFields, file.FieldMetadata)
}

// SetUserID sets the user edge to User by id.
func (m *FileMutation) SetUserID(id int) {
	m.user = &id
}

// ClearUser clears the user edge to User.
func (m *FileMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared returns if the edge user was cleared.
func (m *FileMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the user id in the mutation.
func (m *FileMutation) UserID() (id int, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the user ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *FileMutation) UserIDs() (ids []int) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser reset all changes of the user edge.
func (m *FileMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Op returns the operation name.
func (m *FileMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (File).
func (m *FileMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *FileMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.hash != nil {
		fields = append(fields, file.FieldHash)
	}
	if m.size != nil {
		fields = append(fields, file.FieldSize)
	}
	if m.pinned_at != nil {
		fields = append(fields, file.FieldPinnedAt)
	}
	if m.unpinned_at != nil {
		fields = append(fields, file.FieldUnpinnedAt)
	}
	if m.metadata != nil {
		fields = append(fields, file.FieldMetadata)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *FileMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case file.FieldHash:
		return m.Hash()
	case file.FieldSize:
		return m.Size()
	case file.FieldPinnedAt:
		return m.PinnedAt()
	case file.FieldUnpinnedAt:
		return m.UnpinnedAt()
	case file.FieldMetadata:
		return m.Metadata()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *FileMutation) SetField(name string, value ent.Value) error {
	switch name {
	case file.FieldHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHash(v)
		return nil
	case file.FieldSize:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSize(v)
		return nil
	case file.FieldPinnedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPinnedAt(v)
		return nil
	case file.FieldUnpinnedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUnpinnedAt(v)
		return nil
	case file.FieldMetadata:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetadata(v)
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *FileMutation) AddedFields() []string {
	var fields []string
	if m.addsize != nil {
		fields = append(fields, file.FieldSize)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *FileMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case file.FieldSize:
		return m.AddedSize()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *FileMutation) AddField(name string, value ent.Value) error {
	switch name {
	case file.FieldSize:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSize(v)
		return nil
	}
	return fmt.Errorf("unknown File numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *FileMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(file.FieldMetadata) {
		fields = append(fields, file.FieldMetadata)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *FileMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *FileMutation) ClearField(name string) error {
	switch name {
	case file.FieldMetadata:
		m.ClearMetadata()
		return nil
	}
	return fmt.Errorf("unknown File nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *FileMutation) ResetField(name string) error {
	switch name {
	case file.FieldHash:
		m.ResetHash()
		return nil
	case file.FieldSize:
		m.ResetSize()
		return nil
	case file.FieldPinnedAt:
		m.ResetPinnedAt()
		return nil
	case file.FieldUnpinnedAt:
		m.ResetUnpinnedAt()
		return nil
	case file.FieldMetadata:
		m.ResetMetadata()
		return nil
	}
	return fmt.Errorf("unknown File field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *FileMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, file.EdgeUser)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *FileMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case file.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *FileMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *FileMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *FileMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, file.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *FileMutation) EdgeCleared(name string) bool {
	switch name {
	case file.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *FileMutation) ClearEdge(name string) error {
	switch name {
	case file.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown File unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *FileMutation) ResetEdge(name string) error {
	switch name {
	case file.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown File edge %s", name)
}

// TokenMutation represents an operation that mutate the Tokens
// nodes in the graph.
type TokenMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	name          *string
	secret        *string
	permissions   *token.Permissions
	created_at    *time.Time
	last_used     *time.Time
	clearedFields map[string]struct{}
	user          *int
	cleareduser   bool
}

var _ ent.Mutation = (*TokenMutation)(nil)

// newTokenMutation creates new mutation for $n.Name.
func newTokenMutation(c config, op Op) *TokenMutation {
	return &TokenMutation{
		config:        c,
		op:            op,
		typ:           TypeToken,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TokenMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TokenMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Token creation.
func (m *TokenMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *TokenMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *TokenMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *TokenMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ResetName reset all changes of the name field.
func (m *TokenMutation) ResetName() {
	m.name = nil
}

// SetSecret sets the secret field.
func (m *TokenMutation) SetSecret(s string) {
	m.secret = &s
}

// Secret returns the secret value in the mutation.
func (m *TokenMutation) Secret() (r string, exists bool) {
	v := m.secret
	if v == nil {
		return
	}
	return *v, true
}

// ResetSecret reset all changes of the secret field.
func (m *TokenMutation) ResetSecret() {
	m.secret = nil
}

// SetPermissions sets the permissions field.
func (m *TokenMutation) SetPermissions(t token.Permissions) {
	m.permissions = &t
}

// Permissions returns the permissions value in the mutation.
func (m *TokenMutation) Permissions() (r token.Permissions, exists bool) {
	v := m.permissions
	if v == nil {
		return
	}
	return *v, true
}

// ResetPermissions reset all changes of the permissions field.
func (m *TokenMutation) ResetPermissions() {
	m.permissions = nil
}

// SetCreatedAt sets the created_at field.
func (m *TokenMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *TokenMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedAt reset all changes of the created_at field.
func (m *TokenMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetLastUsed sets the last_used field.
func (m *TokenMutation) SetLastUsed(t time.Time) {
	m.last_used = &t
}

// LastUsed returns the last_used value in the mutation.
func (m *TokenMutation) LastUsed() (r time.Time, exists bool) {
	v := m.last_used
	if v == nil {
		return
	}
	return *v, true
}

// ClearLastUsed clears the value of last_used.
func (m *TokenMutation) ClearLastUsed() {
	m.last_used = nil
	m.clearedFields[token.FieldLastUsed] = struct{}{}
}

// LastUsedCleared returns if the field last_used was cleared in this mutation.
func (m *TokenMutation) LastUsedCleared() bool {
	_, ok := m.clearedFields[token.FieldLastUsed]
	return ok
}

// ResetLastUsed reset all changes of the last_used field.
func (m *TokenMutation) ResetLastUsed() {
	m.last_used = nil
	delete(m.clearedFields, token.FieldLastUsed)
}

// SetUserID sets the user edge to User by id.
func (m *TokenMutation) SetUserID(id int) {
	m.user = &id
}

// ClearUser clears the user edge to User.
func (m *TokenMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared returns if the edge user was cleared.
func (m *TokenMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the user id in the mutation.
func (m *TokenMutation) UserID() (id int, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the user ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *TokenMutation) UserIDs() (ids []int) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser reset all changes of the user edge.
func (m *TokenMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Op returns the operation name.
func (m *TokenMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Token).
func (m *TokenMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *TokenMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, token.FieldName)
	}
	if m.secret != nil {
		fields = append(fields, token.FieldSecret)
	}
	if m.permissions != nil {
		fields = append(fields, token.FieldPermissions)
	}
	if m.created_at != nil {
		fields = append(fields, token.FieldCreatedAt)
	}
	if m.last_used != nil {
		fields = append(fields, token.FieldLastUsed)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *TokenMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case token.FieldName:
		return m.Name()
	case token.FieldSecret:
		return m.Secret()
	case token.FieldPermissions:
		return m.Permissions()
	case token.FieldCreatedAt:
		return m.CreatedAt()
	case token.FieldLastUsed:
		return m.LastUsed()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TokenMutation) SetField(name string, value ent.Value) error {
	switch name {
	case token.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case token.FieldSecret:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSecret(v)
		return nil
	case token.FieldPermissions:
		v, ok := value.(token.Permissions)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPermissions(v)
		return nil
	case token.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case token.FieldLastUsed:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastUsed(v)
		return nil
	}
	return fmt.Errorf("unknown Token field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *TokenMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *TokenMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *TokenMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Token numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *TokenMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(token.FieldLastUsed) {
		fields = append(fields, token.FieldLastUsed)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *TokenMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *TokenMutation) ClearField(name string) error {
	switch name {
	case token.FieldLastUsed:
		m.ClearLastUsed()
		return nil
	}
	return fmt.Errorf("unknown Token nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *TokenMutation) ResetField(name string) error {
	switch name {
	case token.FieldName:
		m.ResetName()
		return nil
	case token.FieldSecret:
		m.ResetSecret()
		return nil
	case token.FieldPermissions:
		m.ResetPermissions()
		return nil
	case token.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case token.FieldLastUsed:
		m.ResetLastUsed()
		return nil
	}
	return fmt.Errorf("unknown Token field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *TokenMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, token.EdgeUser)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *TokenMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case token.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *TokenMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *TokenMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *TokenMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, token.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *TokenMutation) EdgeCleared(name string) bool {
	switch name {
	case token.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *TokenMutation) ClearEdge(name string) error {
	switch name {
	case token.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown Token unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *TokenMutation) ResetEdge(name string) error {
	switch name {
	case token.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Token edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	email         *string
	password      *string
	active        *bool
	updated_at    *time.Time
	created_at    *time.Time
	plan          *user.Plan
	clearedFields map[string]struct{}
	tokens        map[uuid.UUID]struct{}
	removedtokens map[uuid.UUID]struct{}
	files         map[uuid.UUID]struct{}
	removedfiles  map[uuid.UUID]struct{}
}

var _ ent.Mutation = (*UserMutation)(nil)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op) *UserMutation {
	return &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetEmail sets the email field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the email value in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// ResetEmail reset all changes of the email field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the password field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the password value in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// ResetPassword reset all changes of the password field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetActive sets the active field.
func (m *UserMutation) SetActive(b bool) {
	m.active = &b
}

// Active returns the active value in the mutation.
func (m *UserMutation) Active() (r bool, exists bool) {
	v := m.active
	if v == nil {
		return
	}
	return *v, true
}

// ResetActive reset all changes of the active field.
func (m *UserMutation) ResetActive() {
	m.active = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpdatedAt reset all changes of the updated_at field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetCreatedAt sets the created_at field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedAt reset all changes of the created_at field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetPlan sets the plan field.
func (m *UserMutation) SetPlan(u user.Plan) {
	m.plan = &u
}

// Plan returns the plan value in the mutation.
func (m *UserMutation) Plan() (r user.Plan, exists bool) {
	v := m.plan
	if v == nil {
		return
	}
	return *v, true
}

// ResetPlan reset all changes of the plan field.
func (m *UserMutation) ResetPlan() {
	m.plan = nil
}

// AddTokenIDs adds the tokens edge to Token by ids.
func (m *UserMutation) AddTokenIDs(ids ...uuid.UUID) {
	if m.tokens == nil {
		m.tokens = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.tokens[ids[i]] = struct{}{}
	}
}

// RemoveTokenIDs removes the tokens edge to Token by ids.
func (m *UserMutation) RemoveTokenIDs(ids ...uuid.UUID) {
	if m.removedtokens == nil {
		m.removedtokens = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedtokens[ids[i]] = struct{}{}
	}
}

// RemovedTokens returns the removed ids of tokens.
func (m *UserMutation) RemovedTokensIDs() (ids []uuid.UUID) {
	for id := range m.removedtokens {
		ids = append(ids, id)
	}
	return
}

// TokensIDs returns the tokens ids in the mutation.
func (m *UserMutation) TokensIDs() (ids []uuid.UUID) {
	for id := range m.tokens {
		ids = append(ids, id)
	}
	return
}

// ResetTokens reset all changes of the tokens edge.
func (m *UserMutation) ResetTokens() {
	m.tokens = nil
	m.removedtokens = nil
}

// AddFileIDs adds the files edge to File by ids.
func (m *UserMutation) AddFileIDs(ids ...uuid.UUID) {
	if m.files == nil {
		m.files = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.files[ids[i]] = struct{}{}
	}
}

// RemoveFileIDs removes the files edge to File by ids.
func (m *UserMutation) RemoveFileIDs(ids ...uuid.UUID) {
	if m.removedfiles == nil {
		m.removedfiles = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedfiles[ids[i]] = struct{}{}
	}
}

// RemovedFiles returns the removed ids of files.
func (m *UserMutation) RemovedFilesIDs() (ids []uuid.UUID) {
	for id := range m.removedfiles {
		ids = append(ids, id)
	}
	return
}

// FilesIDs returns the files ids in the mutation.
func (m *UserMutation) FilesIDs() (ids []uuid.UUID) {
	for id := range m.files {
		ids = append(ids, id)
	}
	return
}

// ResetFiles reset all changes of the files edge.
func (m *UserMutation) ResetFiles() {
	m.files = nil
	m.removedfiles = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.active != nil {
		fields = append(fields, user.FieldActive)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.plan != nil {
		fields = append(fields, user.FieldPlan)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	case user.FieldActive:
		return m.Active()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldPlan:
		return m.Plan()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldActive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActive(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldPlan:
		v, ok := value.(user.Plan)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPlan(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldActive:
		m.ResetActive()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldPlan:
		m.ResetPlan()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.tokens != nil {
		edges = append(edges, user.EdgeTokens)
	}
	if m.files != nil {
		edges = append(edges, user.EdgeFiles)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTokens:
		ids := make([]ent.Value, 0, len(m.tokens))
		for id := range m.tokens {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFiles:
		ids := make([]ent.Value, 0, len(m.files))
		for id := range m.files {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedtokens != nil {
		edges = append(edges, user.EdgeTokens)
	}
	if m.removedfiles != nil {
		edges = append(edges, user.EdgeFiles)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTokens:
		ids := make([]ent.Value, 0, len(m.removedtokens))
		for id := range m.removedtokens {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFiles:
		ids := make([]ent.Value, 0, len(m.removedfiles))
		for id := range m.removedfiles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeTokens:
		m.ResetTokens()
		return nil
	case user.EdgeFiles:
		m.ResetFiles()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
