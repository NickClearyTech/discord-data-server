// Code generated by ent, DO NOT EDIT.

package message

import (
	"discord-metrics-server/v2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldID, id))
}

// Contents applies equality check predicate on the "contents" field. It's identical to ContentsEQ.
func Contents(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldContents, v))
}

// SentAt applies equality check predicate on the "sent_at" field. It's identical to SentAtEQ.
func SentAt(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSentAt, v))
}

// SenderID applies equality check predicate on the "sender_id" field. It's identical to SenderIDEQ.
func SenderID(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSenderID, v))
}

// MessageID applies equality check predicate on the "message_id" field. It's identical to MessageIDEQ.
func MessageID(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldMessageID, v))
}

// ChannelID applies equality check predicate on the "channel_id" field. It's identical to ChannelIDEQ.
func ChannelID(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldChannelID, v))
}

// InReplyToID applies equality check predicate on the "in_reply_to_id" field. It's identical to InReplyToIDEQ.
func InReplyToID(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldInReplyToID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldUpdatedAt, v))
}

// ContentsEQ applies the EQ predicate on the "contents" field.
func ContentsEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldContents, v))
}

// ContentsNEQ applies the NEQ predicate on the "contents" field.
func ContentsNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldContents, v))
}

// ContentsIn applies the In predicate on the "contents" field.
func ContentsIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldContents, vs...))
}

// ContentsNotIn applies the NotIn predicate on the "contents" field.
func ContentsNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldContents, vs...))
}

// ContentsGT applies the GT predicate on the "contents" field.
func ContentsGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldContents, v))
}

// ContentsGTE applies the GTE predicate on the "contents" field.
func ContentsGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldContents, v))
}

// ContentsLT applies the LT predicate on the "contents" field.
func ContentsLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldContents, v))
}

// ContentsLTE applies the LTE predicate on the "contents" field.
func ContentsLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldContents, v))
}

// ContentsContains applies the Contains predicate on the "contents" field.
func ContentsContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldContents, v))
}

// ContentsHasPrefix applies the HasPrefix predicate on the "contents" field.
func ContentsHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldContents, v))
}

// ContentsHasSuffix applies the HasSuffix predicate on the "contents" field.
func ContentsHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldContents, v))
}

// ContentsEqualFold applies the EqualFold predicate on the "contents" field.
func ContentsEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldContents, v))
}

// ContentsContainsFold applies the ContainsFold predicate on the "contents" field.
func ContentsContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldContents, v))
}

// SentAtEQ applies the EQ predicate on the "sent_at" field.
func SentAtEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSentAt, v))
}

// SentAtNEQ applies the NEQ predicate on the "sent_at" field.
func SentAtNEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldSentAt, v))
}

// SentAtIn applies the In predicate on the "sent_at" field.
func SentAtIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldSentAt, vs...))
}

// SentAtNotIn applies the NotIn predicate on the "sent_at" field.
func SentAtNotIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldSentAt, vs...))
}

// SentAtGT applies the GT predicate on the "sent_at" field.
func SentAtGT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldSentAt, v))
}

// SentAtGTE applies the GTE predicate on the "sent_at" field.
func SentAtGTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldSentAt, v))
}

// SentAtLT applies the LT predicate on the "sent_at" field.
func SentAtLT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldSentAt, v))
}

// SentAtLTE applies the LTE predicate on the "sent_at" field.
func SentAtLTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldSentAt, v))
}

// SenderIDEQ applies the EQ predicate on the "sender_id" field.
func SenderIDEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSenderID, v))
}

// SenderIDNEQ applies the NEQ predicate on the "sender_id" field.
func SenderIDNEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldSenderID, v))
}

// SenderIDIn applies the In predicate on the "sender_id" field.
func SenderIDIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldSenderID, vs...))
}

// SenderIDNotIn applies the NotIn predicate on the "sender_id" field.
func SenderIDNotIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldSenderID, vs...))
}

// MessageIDEQ applies the EQ predicate on the "message_id" field.
func MessageIDEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldMessageID, v))
}

// MessageIDNEQ applies the NEQ predicate on the "message_id" field.
func MessageIDNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldMessageID, v))
}

// MessageIDIn applies the In predicate on the "message_id" field.
func MessageIDIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldMessageID, vs...))
}

// MessageIDNotIn applies the NotIn predicate on the "message_id" field.
func MessageIDNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldMessageID, vs...))
}

// MessageIDGT applies the GT predicate on the "message_id" field.
func MessageIDGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldMessageID, v))
}

// MessageIDGTE applies the GTE predicate on the "message_id" field.
func MessageIDGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldMessageID, v))
}

// MessageIDLT applies the LT predicate on the "message_id" field.
func MessageIDLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldMessageID, v))
}

// MessageIDLTE applies the LTE predicate on the "message_id" field.
func MessageIDLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldMessageID, v))
}

// MessageIDContains applies the Contains predicate on the "message_id" field.
func MessageIDContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldMessageID, v))
}

// MessageIDHasPrefix applies the HasPrefix predicate on the "message_id" field.
func MessageIDHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldMessageID, v))
}

// MessageIDHasSuffix applies the HasSuffix predicate on the "message_id" field.
func MessageIDHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldMessageID, v))
}

// MessageIDEqualFold applies the EqualFold predicate on the "message_id" field.
func MessageIDEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldMessageID, v))
}

// MessageIDContainsFold applies the ContainsFold predicate on the "message_id" field.
func MessageIDContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldMessageID, v))
}

// ChannelIDEQ applies the EQ predicate on the "channel_id" field.
func ChannelIDEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldChannelID, v))
}

// ChannelIDNEQ applies the NEQ predicate on the "channel_id" field.
func ChannelIDNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldChannelID, v))
}

// ChannelIDIn applies the In predicate on the "channel_id" field.
func ChannelIDIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldChannelID, vs...))
}

// ChannelIDNotIn applies the NotIn predicate on the "channel_id" field.
func ChannelIDNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldChannelID, vs...))
}

// ChannelIDGT applies the GT predicate on the "channel_id" field.
func ChannelIDGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldChannelID, v))
}

// ChannelIDGTE applies the GTE predicate on the "channel_id" field.
func ChannelIDGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldChannelID, v))
}

// ChannelIDLT applies the LT predicate on the "channel_id" field.
func ChannelIDLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldChannelID, v))
}

// ChannelIDLTE applies the LTE predicate on the "channel_id" field.
func ChannelIDLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldChannelID, v))
}

// ChannelIDContains applies the Contains predicate on the "channel_id" field.
func ChannelIDContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldChannelID, v))
}

// ChannelIDHasPrefix applies the HasPrefix predicate on the "channel_id" field.
func ChannelIDHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldChannelID, v))
}

// ChannelIDHasSuffix applies the HasSuffix predicate on the "channel_id" field.
func ChannelIDHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldChannelID, v))
}

// ChannelIDEqualFold applies the EqualFold predicate on the "channel_id" field.
func ChannelIDEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldChannelID, v))
}

// ChannelIDContainsFold applies the ContainsFold predicate on the "channel_id" field.
func ChannelIDContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldChannelID, v))
}

// InReplyToIDEQ applies the EQ predicate on the "in_reply_to_id" field.
func InReplyToIDEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldInReplyToID, v))
}

// InReplyToIDNEQ applies the NEQ predicate on the "in_reply_to_id" field.
func InReplyToIDNEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldInReplyToID, v))
}

// InReplyToIDIn applies the In predicate on the "in_reply_to_id" field.
func InReplyToIDIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldInReplyToID, vs...))
}

// InReplyToIDNotIn applies the NotIn predicate on the "in_reply_to_id" field.
func InReplyToIDNotIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldInReplyToID, vs...))
}

// InReplyToIDIsNil applies the IsNil predicate on the "in_reply_to_id" field.
func InReplyToIDIsNil() predicate.Message {
	return predicate.Message(sql.FieldIsNull(FieldInReplyToID))
}

// InReplyToIDNotNil applies the NotNil predicate on the "in_reply_to_id" field.
func InReplyToIDNotNil() predicate.Message {
	return predicate.Message(sql.FieldNotNull(FieldInReplyToID))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasSender applies the HasEdge predicate on the "sender" edge.
func HasSender() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SenderTable, SenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSenderWith applies the HasEdge predicate on the "sender" edge with a given conditions (other predicates).
func HasSenderWith(preds ...predicate.User) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newSenderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInReplyTo applies the HasEdge predicate on the "in_reply_to" edge.
func HasInReplyTo() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InReplyToTable, InReplyToColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInReplyToWith applies the HasEdge predicate on the "in_reply_to" edge with a given conditions (other predicates).
func HasInReplyToWith(preds ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newInReplyToStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResponders applies the HasEdge predicate on the "responders" edge.
func HasResponders() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RespondersTable, RespondersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRespondersWith applies the HasEdge predicate on the "responders" edge with a given conditions (other predicates).
func HasRespondersWith(preds ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newRespondersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMentions applies the HasEdge predicate on the "mentions" edge.
func HasMentions() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MentionsTable, MentionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMentionsWith applies the HasEdge predicate on the "mentions" edge with a given conditions (other predicates).
func HasMentionsWith(preds ...predicate.User) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newMentionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(sql.NotPredicates(p))
}
