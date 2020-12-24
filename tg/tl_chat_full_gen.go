// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// ChatFull represents TL type `chatFull#dc8c181`.
// Detailed chat info
//
// See https://core.telegram.org/constructor/chatFull for reference.
type ChatFull struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Can we change the username of this chat
	CanSetUsername bool
	// Whether scheduled messages¹ are available
	//
	// Links:
	//  1) https://core.telegram.org/api/scheduled-messages
	HasScheduled bool
	// ID of the chat
	ID int
	// About string for this chat
	About string
	// Participant list
	Participants ChatParticipantsClass
	// Chat photo
	//
	// Use SetChatPhoto and GetChatPhoto helpers.
	ChatPhoto PhotoClass
	// Notification settings
	NotifySettings PeerNotifySettings
	// Chat invite
	ExportedInvite ExportedChatInviteClass
	// Info about bots that are in this chat
	//
	// Use SetBotInfo and GetBotInfo helpers.
	BotInfo []BotInfo
	// Message ID of the last pinned message¹
	//
	// Links:
	//  1) https://core.telegram.org/api/pin
	//
	// Use SetPinnedMsgID and GetPinnedMsgID helpers.
	PinnedMsgID int
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// Call field of ChatFull.
	//
	// Use SetCall and GetCall helpers.
	Call InputGroupCall
}

// ChatFullTypeID is TL type id of ChatFull.
const ChatFullTypeID = 0xdc8c181

// String implements fmt.Stringer.
func (c *ChatFull) String() string {
	if c == nil {
		return "ChatFull(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatFull")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(c.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(c.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAbout: ")
	sb.WriteString(fmt.Sprint(c.About))
	sb.WriteString(",\n")
	sb.WriteString("\tParticipants: ")
	sb.WriteString(c.Participants.String())
	sb.WriteString(",\n")
	if c.Flags.Has(2) {
		sb.WriteString("\tChatPhoto: ")
		sb.WriteString(c.ChatPhoto.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("\tNotifySettings: ")
	sb.WriteString(c.NotifySettings.String())
	sb.WriteString(",\n")
	sb.WriteString("\tExportedInvite: ")
	sb.WriteString(c.ExportedInvite.String())
	sb.WriteString(",\n")
	if c.Flags.Has(3) {
		sb.WriteByte('[')
		for _, v := range c.BotInfo {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	if c.Flags.Has(6) {
		sb.WriteString("\tPinnedMsgID: ")
		sb.WriteString(fmt.Sprint(c.PinnedMsgID))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(11) {
		sb.WriteString("\tFolderID: ")
		sb.WriteString(fmt.Sprint(c.FolderID))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(12) {
		sb.WriteString("\tCall: ")
		sb.WriteString(c.Call.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChatFull) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatFull#dc8c181 as nil")
	}
	b.PutID(ChatFullTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatFull#dc8c181: field flags: %w", err)
	}
	b.PutInt(c.ID)
	b.PutString(c.About)
	if c.Participants == nil {
		return fmt.Errorf("unable to encode chatFull#dc8c181: field participants is nil")
	}
	if err := c.Participants.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatFull#dc8c181: field participants: %w", err)
	}
	if c.Flags.Has(2) {
		if c.ChatPhoto == nil {
			return fmt.Errorf("unable to encode chatFull#dc8c181: field chat_photo is nil")
		}
		if err := c.ChatPhoto.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatFull#dc8c181: field chat_photo: %w", err)
		}
	}
	if err := c.NotifySettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatFull#dc8c181: field notify_settings: %w", err)
	}
	if c.ExportedInvite == nil {
		return fmt.Errorf("unable to encode chatFull#dc8c181: field exported_invite is nil")
	}
	if err := c.ExportedInvite.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatFull#dc8c181: field exported_invite: %w", err)
	}
	if c.Flags.Has(3) {
		b.PutVectorHeader(len(c.BotInfo))
		for idx, v := range c.BotInfo {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode chatFull#dc8c181: field bot_info element with index %d: %w", idx, err)
			}
		}
	}
	if c.Flags.Has(6) {
		b.PutInt(c.PinnedMsgID)
	}
	if c.Flags.Has(11) {
		b.PutInt(c.FolderID)
	}
	if c.Flags.Has(12) {
		if err := c.Call.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatFull#dc8c181: field call: %w", err)
		}
	}
	return nil
}

// SetCanSetUsername sets value of CanSetUsername conditional field.
func (c *ChatFull) SetCanSetUsername(value bool) {
	if value {
		c.Flags.Set(7)
		c.CanSetUsername = true
	} else {
		c.Flags.Unset(7)
		c.CanSetUsername = false
	}
}

// SetHasScheduled sets value of HasScheduled conditional field.
func (c *ChatFull) SetHasScheduled(value bool) {
	if value {
		c.Flags.Set(8)
		c.HasScheduled = true
	} else {
		c.Flags.Unset(8)
		c.HasScheduled = false
	}
}

// SetChatPhoto sets value of ChatPhoto conditional field.
func (c *ChatFull) SetChatPhoto(value PhotoClass) {
	c.Flags.Set(2)
	c.ChatPhoto = value
}

// GetChatPhoto returns value of ChatPhoto conditional field and
// boolean which is true if field was set.
func (c *ChatFull) GetChatPhoto() (value PhotoClass, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.ChatPhoto, true
}

// SetBotInfo sets value of BotInfo conditional field.
func (c *ChatFull) SetBotInfo(value []BotInfo) {
	c.Flags.Set(3)
	c.BotInfo = value
}

// GetBotInfo returns value of BotInfo conditional field and
// boolean which is true if field was set.
func (c *ChatFull) GetBotInfo() (value []BotInfo, ok bool) {
	if !c.Flags.Has(3) {
		return value, false
	}
	return c.BotInfo, true
}

// SetPinnedMsgID sets value of PinnedMsgID conditional field.
func (c *ChatFull) SetPinnedMsgID(value int) {
	c.Flags.Set(6)
	c.PinnedMsgID = value
}

// GetPinnedMsgID returns value of PinnedMsgID conditional field and
// boolean which is true if field was set.
func (c *ChatFull) GetPinnedMsgID() (value int, ok bool) {
	if !c.Flags.Has(6) {
		return value, false
	}
	return c.PinnedMsgID, true
}

// SetFolderID sets value of FolderID conditional field.
func (c *ChatFull) SetFolderID(value int) {
	c.Flags.Set(11)
	c.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (c *ChatFull) GetFolderID() (value int, ok bool) {
	if !c.Flags.Has(11) {
		return value, false
	}
	return c.FolderID, true
}

// SetCall sets value of Call conditional field.
func (c *ChatFull) SetCall(value InputGroupCall) {
	c.Flags.Set(12)
	c.Call = value
}

// GetCall returns value of Call conditional field and
// boolean which is true if field was set.
func (c *ChatFull) GetCall() (value InputGroupCall, ok bool) {
	if !c.Flags.Has(12) {
		return value, false
	}
	return c.Call, true
}

// Decode implements bin.Decoder.
func (c *ChatFull) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatFull#dc8c181 to nil")
	}
	if err := b.ConsumeID(ChatFullTypeID); err != nil {
		return fmt.Errorf("unable to decode chatFull#dc8c181: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field flags: %w", err)
		}
	}
	c.CanSetUsername = c.Flags.Has(7)
	c.HasScheduled = c.Flags.Has(8)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field about: %w", err)
		}
		c.About = value
	}
	{
		value, err := DecodeChatParticipants(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field participants: %w", err)
		}
		c.Participants = value
	}
	if c.Flags.Has(2) {
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field chat_photo: %w", err)
		}
		c.ChatPhoto = value
	}
	{
		if err := c.NotifySettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field notify_settings: %w", err)
		}
	}
	{
		value, err := DecodeExportedChatInvite(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field exported_invite: %w", err)
		}
		c.ExportedInvite = value
	}
	if c.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field bot_info: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotInfo
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode chatFull#dc8c181: field bot_info: %w", err)
			}
			c.BotInfo = append(c.BotInfo, value)
		}
	}
	if c.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field pinned_msg_id: %w", err)
		}
		c.PinnedMsgID = value
	}
	if c.Flags.Has(11) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field folder_id: %w", err)
		}
		c.FolderID = value
	}
	if c.Flags.Has(12) {
		if err := c.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatFull#dc8c181: field call: %w", err)
		}
	}
	return nil
}

// construct implements constructor of ChatFullClass.
func (c ChatFull) construct() ChatFullClass { return &c }

// Ensuring interfaces in compile-time for ChatFull.
var (
	_ bin.Encoder = &ChatFull{}
	_ bin.Decoder = &ChatFull{}

	_ ChatFullClass = &ChatFull{}
)

// ChannelFull represents TL type `channelFull#ef3a6acd`.
// Full info about a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/constructor/channelFull for reference.
type ChannelFull struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Can we vew the participant list?
	CanViewParticipants bool
	// Can we set the channel's username?
	CanSetUsername bool
	// Can we associate¹ a stickerpack to the supergroup?
	//
	// Links:
	//  1) https://core.telegram.org/method/channels.setStickers
	CanSetStickers bool
	// Is the history before we joined hidden to us?
	HiddenPrehistory bool
	// Can we set the geolocation of this group (for geogroups)
	CanSetLocation bool
	// Whether scheduled messages are available
	HasScheduled bool
	// Can the user view channel/supergroup statistics¹
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	CanViewStats bool
	// Whether any anonymous admin of this supergroup was blocked: if set, you won't receive messages from anonymous group admins in discussion replies via @replies¹
	//
	// Links:
	//  1) https://core.telegram.org/api/discussion
	Blocked bool
	// ID of the channel
	ID int
	// Info about the channel
	About string
	// Number of participants of the channel
	//
	// Use SetParticipantsCount and GetParticipantsCount helpers.
	ParticipantsCount int
	// Number of channel admins
	//
	// Use SetAdminsCount and GetAdminsCount helpers.
	AdminsCount int
	// Number of users kicked¹ from the channel
	//
	// Links:
	//  1) https://core.telegram.org/api/rights
	//
	// Use SetKickedCount and GetKickedCount helpers.
	KickedCount int
	// Number of users banned¹ from the channel
	//
	// Links:
	//  1) https://core.telegram.org/api/rights
	//
	// Use SetBannedCount and GetBannedCount helpers.
	BannedCount int
	// Number of users currently online
	//
	// Use SetOnlineCount and GetOnlineCount helpers.
	OnlineCount int
	// Position up to which all incoming messages are read.
	ReadInboxMaxID int
	// Position up to which all outgoing messages are read.
	ReadOutboxMaxID int
	// Count of unread messages
	UnreadCount int
	// Channel picture
	ChatPhoto PhotoClass
	// Notification settings
	NotifySettings PeerNotifySettings
	// Invite link
	ExportedInvite ExportedChatInviteClass
	// Info about bots in the channel/supergrup
	BotInfo []BotInfo
	// The chat ID from which this group was migrated¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	//
	// Use SetMigratedFromChatID and GetMigratedFromChatID helpers.
	MigratedFromChatID int
	// The message ID in the original chat at which this group was migrated¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	//
	// Use SetMigratedFromMaxID and GetMigratedFromMaxID helpers.
	MigratedFromMaxID int
	// Message ID of the last pinned message¹
	//
	// Links:
	//  1) https://core.telegram.org/api/pin
	//
	// Use SetPinnedMsgID and GetPinnedMsgID helpers.
	PinnedMsgID int
	// Associated stickerset
	//
	// Use SetStickerset and GetStickerset helpers.
	Stickerset StickerSet
	// Identifier of a maximum unavailable message in a channel due to hidden history.
	//
	// Use SetAvailableMinID and GetAvailableMinID helpers.
	AvailableMinID int
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// ID of the linked discussion chat¹ for channels
	//
	// Links:
	//  1) https://core.telegram.org/api/discussion
	//
	// Use SetLinkedChatID and GetLinkedChatID helpers.
	LinkedChatID int
	// Location of the geogroup
	//
	// Use SetLocation and GetLocation helpers.
	Location ChannelLocationClass
	// If specified, users in supergroups will only be able to send one message every slowmode_seconds seconds
	//
	// Use SetSlowmodeSeconds and GetSlowmodeSeconds helpers.
	SlowmodeSeconds int
	// Indicates when the user will be allowed to send another message in the supergroup (unixdate)
	//
	// Use SetSlowmodeNextSendDate and GetSlowmodeNextSendDate helpers.
	SlowmodeNextSendDate int
	// If set, specifies the DC to use for fetching channel statistics
	//
	// Use SetStatsDC and GetStatsDC helpers.
	StatsDC int
	// Latest PTS¹ for this channel
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int
	// Call field of ChannelFull.
	//
	// Use SetCall and GetCall helpers.
	Call InputGroupCall
}

// ChannelFullTypeID is TL type id of ChannelFull.
const ChannelFullTypeID = 0xef3a6acd

// String implements fmt.Stringer.
func (c *ChannelFull) String() string {
	if c == nil {
		return "ChannelFull(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelFull")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(c.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(c.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAbout: ")
	sb.WriteString(fmt.Sprint(c.About))
	sb.WriteString(",\n")
	if c.Flags.Has(0) {
		sb.WriteString("\tParticipantsCount: ")
		sb.WriteString(fmt.Sprint(c.ParticipantsCount))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(1) {
		sb.WriteString("\tAdminsCount: ")
		sb.WriteString(fmt.Sprint(c.AdminsCount))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(2) {
		sb.WriteString("\tKickedCount: ")
		sb.WriteString(fmt.Sprint(c.KickedCount))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(2) {
		sb.WriteString("\tBannedCount: ")
		sb.WriteString(fmt.Sprint(c.BannedCount))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(13) {
		sb.WriteString("\tOnlineCount: ")
		sb.WriteString(fmt.Sprint(c.OnlineCount))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tReadInboxMaxID: ")
	sb.WriteString(fmt.Sprint(c.ReadInboxMaxID))
	sb.WriteString(",\n")
	sb.WriteString("\tReadOutboxMaxID: ")
	sb.WriteString(fmt.Sprint(c.ReadOutboxMaxID))
	sb.WriteString(",\n")
	sb.WriteString("\tUnreadCount: ")
	sb.WriteString(fmt.Sprint(c.UnreadCount))
	sb.WriteString(",\n")
	sb.WriteString("\tChatPhoto: ")
	sb.WriteString(c.ChatPhoto.String())
	sb.WriteString(",\n")
	sb.WriteString("\tNotifySettings: ")
	sb.WriteString(c.NotifySettings.String())
	sb.WriteString(",\n")
	sb.WriteString("\tExportedInvite: ")
	sb.WriteString(c.ExportedInvite.String())
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range c.BotInfo {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	if c.Flags.Has(4) {
		sb.WriteString("\tMigratedFromChatID: ")
		sb.WriteString(fmt.Sprint(c.MigratedFromChatID))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(4) {
		sb.WriteString("\tMigratedFromMaxID: ")
		sb.WriteString(fmt.Sprint(c.MigratedFromMaxID))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(5) {
		sb.WriteString("\tPinnedMsgID: ")
		sb.WriteString(fmt.Sprint(c.PinnedMsgID))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(8) {
		sb.WriteString("\tStickerset: ")
		sb.WriteString(c.Stickerset.String())
		sb.WriteString(",\n")
	}
	if c.Flags.Has(9) {
		sb.WriteString("\tAvailableMinID: ")
		sb.WriteString(fmt.Sprint(c.AvailableMinID))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(11) {
		sb.WriteString("\tFolderID: ")
		sb.WriteString(fmt.Sprint(c.FolderID))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(14) {
		sb.WriteString("\tLinkedChatID: ")
		sb.WriteString(fmt.Sprint(c.LinkedChatID))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(15) {
		sb.WriteString("\tLocation: ")
		sb.WriteString(c.Location.String())
		sb.WriteString(",\n")
	}
	if c.Flags.Has(17) {
		sb.WriteString("\tSlowmodeSeconds: ")
		sb.WriteString(fmt.Sprint(c.SlowmodeSeconds))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(18) {
		sb.WriteString("\tSlowmodeNextSendDate: ")
		sb.WriteString(fmt.Sprint(c.SlowmodeNextSendDate))
		sb.WriteString(",\n")
	}
	if c.Flags.Has(12) {
		sb.WriteString("\tStatsDC: ")
		sb.WriteString(fmt.Sprint(c.StatsDC))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tPts: ")
	sb.WriteString(fmt.Sprint(c.Pts))
	sb.WriteString(",\n")
	if c.Flags.Has(21) {
		sb.WriteString("\tCall: ")
		sb.WriteString(c.Call.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChannelFull) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelFull#ef3a6acd as nil")
	}
	b.PutID(ChannelFullTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelFull#ef3a6acd: field flags: %w", err)
	}
	b.PutInt(c.ID)
	b.PutString(c.About)
	if c.Flags.Has(0) {
		b.PutInt(c.ParticipantsCount)
	}
	if c.Flags.Has(1) {
		b.PutInt(c.AdminsCount)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.KickedCount)
	}
	if c.Flags.Has(2) {
		b.PutInt(c.BannedCount)
	}
	if c.Flags.Has(13) {
		b.PutInt(c.OnlineCount)
	}
	b.PutInt(c.ReadInboxMaxID)
	b.PutInt(c.ReadOutboxMaxID)
	b.PutInt(c.UnreadCount)
	if c.ChatPhoto == nil {
		return fmt.Errorf("unable to encode channelFull#ef3a6acd: field chat_photo is nil")
	}
	if err := c.ChatPhoto.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelFull#ef3a6acd: field chat_photo: %w", err)
	}
	if err := c.NotifySettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelFull#ef3a6acd: field notify_settings: %w", err)
	}
	if c.ExportedInvite == nil {
		return fmt.Errorf("unable to encode channelFull#ef3a6acd: field exported_invite is nil")
	}
	if err := c.ExportedInvite.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelFull#ef3a6acd: field exported_invite: %w", err)
	}
	b.PutVectorHeader(len(c.BotInfo))
	for idx, v := range c.BotInfo {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channelFull#ef3a6acd: field bot_info element with index %d: %w", idx, err)
		}
	}
	if c.Flags.Has(4) {
		b.PutInt(c.MigratedFromChatID)
	}
	if c.Flags.Has(4) {
		b.PutInt(c.MigratedFromMaxID)
	}
	if c.Flags.Has(5) {
		b.PutInt(c.PinnedMsgID)
	}
	if c.Flags.Has(8) {
		if err := c.Stickerset.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channelFull#ef3a6acd: field stickerset: %w", err)
		}
	}
	if c.Flags.Has(9) {
		b.PutInt(c.AvailableMinID)
	}
	if c.Flags.Has(11) {
		b.PutInt(c.FolderID)
	}
	if c.Flags.Has(14) {
		b.PutInt(c.LinkedChatID)
	}
	if c.Flags.Has(15) {
		if c.Location == nil {
			return fmt.Errorf("unable to encode channelFull#ef3a6acd: field location is nil")
		}
		if err := c.Location.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channelFull#ef3a6acd: field location: %w", err)
		}
	}
	if c.Flags.Has(17) {
		b.PutInt(c.SlowmodeSeconds)
	}
	if c.Flags.Has(18) {
		b.PutInt(c.SlowmodeNextSendDate)
	}
	if c.Flags.Has(12) {
		b.PutInt(c.StatsDC)
	}
	b.PutInt(c.Pts)
	if c.Flags.Has(21) {
		if err := c.Call.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channelFull#ef3a6acd: field call: %w", err)
		}
	}
	return nil
}

// SetCanViewParticipants sets value of CanViewParticipants conditional field.
func (c *ChannelFull) SetCanViewParticipants(value bool) {
	if value {
		c.Flags.Set(3)
		c.CanViewParticipants = true
	} else {
		c.Flags.Unset(3)
		c.CanViewParticipants = false
	}
}

// SetCanSetUsername sets value of CanSetUsername conditional field.
func (c *ChannelFull) SetCanSetUsername(value bool) {
	if value {
		c.Flags.Set(6)
		c.CanSetUsername = true
	} else {
		c.Flags.Unset(6)
		c.CanSetUsername = false
	}
}

// SetCanSetStickers sets value of CanSetStickers conditional field.
func (c *ChannelFull) SetCanSetStickers(value bool) {
	if value {
		c.Flags.Set(7)
		c.CanSetStickers = true
	} else {
		c.Flags.Unset(7)
		c.CanSetStickers = false
	}
}

// SetHiddenPrehistory sets value of HiddenPrehistory conditional field.
func (c *ChannelFull) SetHiddenPrehistory(value bool) {
	if value {
		c.Flags.Set(10)
		c.HiddenPrehistory = true
	} else {
		c.Flags.Unset(10)
		c.HiddenPrehistory = false
	}
}

// SetCanSetLocation sets value of CanSetLocation conditional field.
func (c *ChannelFull) SetCanSetLocation(value bool) {
	if value {
		c.Flags.Set(16)
		c.CanSetLocation = true
	} else {
		c.Flags.Unset(16)
		c.CanSetLocation = false
	}
}

// SetHasScheduled sets value of HasScheduled conditional field.
func (c *ChannelFull) SetHasScheduled(value bool) {
	if value {
		c.Flags.Set(19)
		c.HasScheduled = true
	} else {
		c.Flags.Unset(19)
		c.HasScheduled = false
	}
}

// SetCanViewStats sets value of CanViewStats conditional field.
func (c *ChannelFull) SetCanViewStats(value bool) {
	if value {
		c.Flags.Set(20)
		c.CanViewStats = true
	} else {
		c.Flags.Unset(20)
		c.CanViewStats = false
	}
}

// SetBlocked sets value of Blocked conditional field.
func (c *ChannelFull) SetBlocked(value bool) {
	if value {
		c.Flags.Set(22)
		c.Blocked = true
	} else {
		c.Flags.Unset(22)
		c.Blocked = false
	}
}

// SetParticipantsCount sets value of ParticipantsCount conditional field.
func (c *ChannelFull) SetParticipantsCount(value int) {
	c.Flags.Set(0)
	c.ParticipantsCount = value
}

// GetParticipantsCount returns value of ParticipantsCount conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetParticipantsCount() (value int, ok bool) {
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.ParticipantsCount, true
}

// SetAdminsCount sets value of AdminsCount conditional field.
func (c *ChannelFull) SetAdminsCount(value int) {
	c.Flags.Set(1)
	c.AdminsCount = value
}

// GetAdminsCount returns value of AdminsCount conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetAdminsCount() (value int, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.AdminsCount, true
}

// SetKickedCount sets value of KickedCount conditional field.
func (c *ChannelFull) SetKickedCount(value int) {
	c.Flags.Set(2)
	c.KickedCount = value
}

// GetKickedCount returns value of KickedCount conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetKickedCount() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.KickedCount, true
}

// SetBannedCount sets value of BannedCount conditional field.
func (c *ChannelFull) SetBannedCount(value int) {
	c.Flags.Set(2)
	c.BannedCount = value
}

// GetBannedCount returns value of BannedCount conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetBannedCount() (value int, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.BannedCount, true
}

// SetOnlineCount sets value of OnlineCount conditional field.
func (c *ChannelFull) SetOnlineCount(value int) {
	c.Flags.Set(13)
	c.OnlineCount = value
}

// GetOnlineCount returns value of OnlineCount conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetOnlineCount() (value int, ok bool) {
	if !c.Flags.Has(13) {
		return value, false
	}
	return c.OnlineCount, true
}

// SetMigratedFromChatID sets value of MigratedFromChatID conditional field.
func (c *ChannelFull) SetMigratedFromChatID(value int) {
	c.Flags.Set(4)
	c.MigratedFromChatID = value
}

// GetMigratedFromChatID returns value of MigratedFromChatID conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetMigratedFromChatID() (value int, ok bool) {
	if !c.Flags.Has(4) {
		return value, false
	}
	return c.MigratedFromChatID, true
}

// SetMigratedFromMaxID sets value of MigratedFromMaxID conditional field.
func (c *ChannelFull) SetMigratedFromMaxID(value int) {
	c.Flags.Set(4)
	c.MigratedFromMaxID = value
}

// GetMigratedFromMaxID returns value of MigratedFromMaxID conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetMigratedFromMaxID() (value int, ok bool) {
	if !c.Flags.Has(4) {
		return value, false
	}
	return c.MigratedFromMaxID, true
}

// SetPinnedMsgID sets value of PinnedMsgID conditional field.
func (c *ChannelFull) SetPinnedMsgID(value int) {
	c.Flags.Set(5)
	c.PinnedMsgID = value
}

// GetPinnedMsgID returns value of PinnedMsgID conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetPinnedMsgID() (value int, ok bool) {
	if !c.Flags.Has(5) {
		return value, false
	}
	return c.PinnedMsgID, true
}

// SetStickerset sets value of Stickerset conditional field.
func (c *ChannelFull) SetStickerset(value StickerSet) {
	c.Flags.Set(8)
	c.Stickerset = value
}

// GetStickerset returns value of Stickerset conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetStickerset() (value StickerSet, ok bool) {
	if !c.Flags.Has(8) {
		return value, false
	}
	return c.Stickerset, true
}

// SetAvailableMinID sets value of AvailableMinID conditional field.
func (c *ChannelFull) SetAvailableMinID(value int) {
	c.Flags.Set(9)
	c.AvailableMinID = value
}

// GetAvailableMinID returns value of AvailableMinID conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetAvailableMinID() (value int, ok bool) {
	if !c.Flags.Has(9) {
		return value, false
	}
	return c.AvailableMinID, true
}

// SetFolderID sets value of FolderID conditional field.
func (c *ChannelFull) SetFolderID(value int) {
	c.Flags.Set(11)
	c.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetFolderID() (value int, ok bool) {
	if !c.Flags.Has(11) {
		return value, false
	}
	return c.FolderID, true
}

// SetLinkedChatID sets value of LinkedChatID conditional field.
func (c *ChannelFull) SetLinkedChatID(value int) {
	c.Flags.Set(14)
	c.LinkedChatID = value
}

// GetLinkedChatID returns value of LinkedChatID conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetLinkedChatID() (value int, ok bool) {
	if !c.Flags.Has(14) {
		return value, false
	}
	return c.LinkedChatID, true
}

// SetLocation sets value of Location conditional field.
func (c *ChannelFull) SetLocation(value ChannelLocationClass) {
	c.Flags.Set(15)
	c.Location = value
}

// GetLocation returns value of Location conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetLocation() (value ChannelLocationClass, ok bool) {
	if !c.Flags.Has(15) {
		return value, false
	}
	return c.Location, true
}

// SetSlowmodeSeconds sets value of SlowmodeSeconds conditional field.
func (c *ChannelFull) SetSlowmodeSeconds(value int) {
	c.Flags.Set(17)
	c.SlowmodeSeconds = value
}

// GetSlowmodeSeconds returns value of SlowmodeSeconds conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetSlowmodeSeconds() (value int, ok bool) {
	if !c.Flags.Has(17) {
		return value, false
	}
	return c.SlowmodeSeconds, true
}

// SetSlowmodeNextSendDate sets value of SlowmodeNextSendDate conditional field.
func (c *ChannelFull) SetSlowmodeNextSendDate(value int) {
	c.Flags.Set(18)
	c.SlowmodeNextSendDate = value
}

// GetSlowmodeNextSendDate returns value of SlowmodeNextSendDate conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetSlowmodeNextSendDate() (value int, ok bool) {
	if !c.Flags.Has(18) {
		return value, false
	}
	return c.SlowmodeNextSendDate, true
}

// SetStatsDC sets value of StatsDC conditional field.
func (c *ChannelFull) SetStatsDC(value int) {
	c.Flags.Set(12)
	c.StatsDC = value
}

// GetStatsDC returns value of StatsDC conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetStatsDC() (value int, ok bool) {
	if !c.Flags.Has(12) {
		return value, false
	}
	return c.StatsDC, true
}

// SetCall sets value of Call conditional field.
func (c *ChannelFull) SetCall(value InputGroupCall) {
	c.Flags.Set(21)
	c.Call = value
}

// GetCall returns value of Call conditional field and
// boolean which is true if field was set.
func (c *ChannelFull) GetCall() (value InputGroupCall, ok bool) {
	if !c.Flags.Has(21) {
		return value, false
	}
	return c.Call, true
}

// Decode implements bin.Decoder.
func (c *ChannelFull) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelFull#ef3a6acd to nil")
	}
	if err := b.ConsumeID(ChannelFullTypeID); err != nil {
		return fmt.Errorf("unable to decode channelFull#ef3a6acd: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field flags: %w", err)
		}
	}
	c.CanViewParticipants = c.Flags.Has(3)
	c.CanSetUsername = c.Flags.Has(6)
	c.CanSetStickers = c.Flags.Has(7)
	c.HiddenPrehistory = c.Flags.Has(10)
	c.CanSetLocation = c.Flags.Has(16)
	c.HasScheduled = c.Flags.Has(19)
	c.CanViewStats = c.Flags.Has(20)
	c.Blocked = c.Flags.Has(22)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field about: %w", err)
		}
		c.About = value
	}
	if c.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field participants_count: %w", err)
		}
		c.ParticipantsCount = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field admins_count: %w", err)
		}
		c.AdminsCount = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field kicked_count: %w", err)
		}
		c.KickedCount = value
	}
	if c.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field banned_count: %w", err)
		}
		c.BannedCount = value
	}
	if c.Flags.Has(13) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field online_count: %w", err)
		}
		c.OnlineCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field read_inbox_max_id: %w", err)
		}
		c.ReadInboxMaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field read_outbox_max_id: %w", err)
		}
		c.ReadOutboxMaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field unread_count: %w", err)
		}
		c.UnreadCount = value
	}
	{
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field chat_photo: %w", err)
		}
		c.ChatPhoto = value
	}
	{
		if err := c.NotifySettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field notify_settings: %w", err)
		}
	}
	{
		value, err := DecodeExportedChatInvite(b)
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field exported_invite: %w", err)
		}
		c.ExportedInvite = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field bot_info: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotInfo
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode channelFull#ef3a6acd: field bot_info: %w", err)
			}
			c.BotInfo = append(c.BotInfo, value)
		}
	}
	if c.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field migrated_from_chat_id: %w", err)
		}
		c.MigratedFromChatID = value
	}
	if c.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field migrated_from_max_id: %w", err)
		}
		c.MigratedFromMaxID = value
	}
	if c.Flags.Has(5) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field pinned_msg_id: %w", err)
		}
		c.PinnedMsgID = value
	}
	if c.Flags.Has(8) {
		if err := c.Stickerset.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field stickerset: %w", err)
		}
	}
	if c.Flags.Has(9) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field available_min_id: %w", err)
		}
		c.AvailableMinID = value
	}
	if c.Flags.Has(11) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field folder_id: %w", err)
		}
		c.FolderID = value
	}
	if c.Flags.Has(14) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field linked_chat_id: %w", err)
		}
		c.LinkedChatID = value
	}
	if c.Flags.Has(15) {
		value, err := DecodeChannelLocation(b)
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field location: %w", err)
		}
		c.Location = value
	}
	if c.Flags.Has(17) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field slowmode_seconds: %w", err)
		}
		c.SlowmodeSeconds = value
	}
	if c.Flags.Has(18) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field slowmode_next_send_date: %w", err)
		}
		c.SlowmodeNextSendDate = value
	}
	if c.Flags.Has(12) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field stats_dc: %w", err)
		}
		c.StatsDC = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field pts: %w", err)
		}
		c.Pts = value
	}
	if c.Flags.Has(21) {
		if err := c.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channelFull#ef3a6acd: field call: %w", err)
		}
	}
	return nil
}

// construct implements constructor of ChatFullClass.
func (c ChannelFull) construct() ChatFullClass { return &c }

// Ensuring interfaces in compile-time for ChannelFull.
var (
	_ bin.Encoder = &ChannelFull{}
	_ bin.Decoder = &ChannelFull{}

	_ ChatFullClass = &ChannelFull{}
)

// ChatFullClass represents ChatFull generic type.
//
// See https://core.telegram.org/type/ChatFull for reference.
//
// Example:
//  g, err := DecodeChatFull(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ChatFull: // chatFull#dc8c181
//  case *ChannelFull: // channelFull#ef3a6acd
//  default: panic(v)
//  }
type ChatFullClass interface {
	bin.Encoder
	bin.Decoder
	construct() ChatFullClass
	fmt.Stringer
}

// DecodeChatFull implements binary de-serialization for ChatFullClass.
func DecodeChatFull(buf *bin.Buffer) (ChatFullClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatFullTypeID:
		// Decoding chatFull#dc8c181.
		v := ChatFull{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatFullClass: %w", err)
		}
		return &v, nil
	case ChannelFullTypeID:
		// Decoding channelFull#ef3a6acd.
		v := ChannelFull{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatFullClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatFullClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatFull boxes the ChatFullClass providing a helper.
type ChatFullBox struct {
	ChatFull ChatFullClass
}

// Decode implements bin.Decoder for ChatFullBox.
func (b *ChatFullBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatFullBox to nil")
	}
	v, err := DecodeChatFull(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatFull = v
	return nil
}

// Encode implements bin.Encode for ChatFullBox.
func (b *ChatFullBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatFull == nil {
		return fmt.Errorf("unable to encode ChatFullClass as nil")
	}
	return b.ChatFull.Encode(buf)
}
