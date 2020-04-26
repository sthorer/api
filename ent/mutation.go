// github.com/sthorer/api

package ent

import (
	"fmt"
	"time"

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
	TypeToken = "Token"
	TypeUser  = "User"
)

// TokenMutation represents an operation that mutate the Tokens
// nodes in the graph.
type TokenMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	name          *string
	token         *string
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
func (m *TokenMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *TokenMutation) ID() (id int64, exists bool) {
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

// SetToken sets the token field.
func (m *TokenMutation) SetToken(s string) {
	m.token = &s
}

// Token returns the token value in the mutation.
func (m *TokenMutation) Token() (r string, exists bool) {
	v := m.token
	if v == nil {
		return
	}
	return *v, true
}

// ResetToken reset all changes of the token field.
func (m *TokenMutation) ResetToken() {
	m.token = nil
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

// ResetLastUsed reset all changes of the last_used field.
func (m *TokenMutation) ResetLastUsed() {
	m.last_used = nil
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
	if m.token != nil {
		fields = append(fields, token.FieldToken)
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
	case token.FieldToken:
		return m.Token()
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
	case token.FieldToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetToken(v)
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
	return nil
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
	case token.FieldToken:
		m.ResetToken()
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
	tokens        map[int64]struct{}
	removedtokens map[int64]struct{}
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
func (m *UserMutation) AddTokenIDs(ids ...int64) {
	if m.tokens == nil {
		m.tokens = make(map[int64]struct{})
	}
	for i := range ids {
		m.tokens[ids[i]] = struct{}{}
	}
}

// RemoveTokenIDs removes the tokens edge to Token by ids.
func (m *UserMutation) RemoveTokenIDs(ids ...int64) {
	if m.removedtokens == nil {
		m.removedtokens = make(map[int64]struct{})
	}
	for i := range ids {
		m.removedtokens[ids[i]] = struct{}{}
	}
}

// RemovedTokens returns the removed ids of tokens.
func (m *UserMutation) RemovedTokensIDs() (ids []int64) {
	for id := range m.removedtokens {
		ids = append(ids, id)
	}
	return
}

// TokensIDs returns the tokens ids in the mutation.
func (m *UserMutation) TokensIDs() (ids []int64) {
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
	edges := make([]string, 0, 1)
	if m.tokens != nil {
		edges = append(edges, user.EdgeTokens)
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
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtokens != nil {
		edges = append(edges, user.EdgeTokens)
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
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
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
	}
	return fmt.Errorf("unknown User edge %s", name)
}
