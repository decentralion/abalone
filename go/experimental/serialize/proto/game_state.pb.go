// Code generated by protoc-gen-gogo.
// source: game_state.proto
// DO NOT EDIT!

/*
Package game_state_pb is a generated protocol buffer package.

It is generated from these files:
	game_state.proto

It has these top-level messages:
	Hex
	Board
	State
*/
package game_state_pb

import proto "code.google.com/p/gogoprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Player int32

const (
	Player_WHITE Player = 0
	Player_BLACK Player = 1
)

var Player_name = map[int32]string{
	0: "WHITE",
	1: "BLACK",
}
var Player_value = map[string]int32{
	"WHITE": 0,
	"BLACK": 1,
}

func (x Player) Enum() *Player {
	p := new(Player)
	*p = x
	return p
}
func (x Player) String() string {
	return proto.EnumName(Player_name, int32(x))
}
func (x *Player) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Player_value, data, "Player")
	if err != nil {
		return err
	}
	*x = Player(value)
	return nil
}

type Hex struct {
	Q                *int32 `protobuf:"zigzag32,1,opt,name=q" json:"q,omitempty"`
	R                *int32 `protobuf:"zigzag32,2,opt,name=r" json:"r,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Hex) Reset()         { *m = Hex{} }
func (m *Hex) String() string { return proto.CompactTextString(m) }
func (*Hex) ProtoMessage()    {}

func (m *Hex) GetQ() int32 {
	if m != nil && m.Q != nil {
		return *m.Q
	}
	return 0
}

func (m *Hex) GetR() int32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

type Board struct {
	EdgeLength       *uint32 `protobuf:"varint,1,opt,name=edge_length" json:"edge_length,omitempty"`
	BlackPositions   []*Hex  `protobuf:"bytes,2,rep,name=black_positions" json:"black_positions,omitempty"`
	WhitePositions   []*Hex  `protobuf:"bytes,3,rep,name=white_positions" json:"white_positions,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Board) Reset()         { *m = Board{} }
func (m *Board) String() string { return proto.CompactTextString(m) }
func (*Board) ProtoMessage()    {}

func (m *Board) GetEdgeLength() uint32 {
	if m != nil && m.EdgeLength != nil {
		return *m.EdgeLength
	}
	return 0
}

func (m *Board) GetBlackPositions() []*Hex {
	if m != nil {
		return m.BlackPositions
	}
	return nil
}

func (m *Board) GetWhitePositions() []*Hex {
	if m != nil {
		return m.WhitePositions
	}
	return nil
}

type State struct {
	Board            *Board  `protobuf:"bytes,1,opt,name=board" json:"board,omitempty"`
	NextPlayer       *Player `protobuf:"varint,2,opt,name=next_player,enum=game.state.pb.Player" json:"next_player,omitempty"`
	MovesRemaining   *uint32 `protobuf:"varint,3,opt,name=moves_remaining" json:"moves_remaining,omitempty"`
	MarblesPerMove   *uint32 `protobuf:"varint,4,opt,name=marbles_per_move" json:"marbles_per_move,omitempty"`
	LossThreshold    *uint32 `protobuf:"varint,5,opt,name=loss_threshold" json:"loss_threshold,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}

func (m *State) GetBoard() *Board {
	if m != nil {
		return m.Board
	}
	return nil
}

func (m *State) GetNextPlayer() Player {
	if m != nil && m.NextPlayer != nil {
		return *m.NextPlayer
	}
	return Player_WHITE
}

func (m *State) GetMovesRemaining() uint32 {
	if m != nil && m.MovesRemaining != nil {
		return *m.MovesRemaining
	}
	return 0
}

func (m *State) GetMarblesPerMove() uint32 {
	if m != nil && m.MarblesPerMove != nil {
		return *m.MarblesPerMove
	}
	return 0
}

func (m *State) GetLossThreshold() uint32 {
	if m != nil && m.LossThreshold != nil {
		return *m.LossThreshold
	}
	return 0
}

func init() {
	proto.RegisterEnum("game.state.pb.Player", Player_name, Player_value)
}