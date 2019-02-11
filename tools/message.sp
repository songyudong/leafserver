message Header{
	type int32
    session int32
}
	
message Hello{
	Name string
}

message CSLogin{
	UserName string
	Password string
}

message CSChat {
	Content string
}

message CSMatch {
	Mode int32
}

message CSEnterGame {
	Room int32
}

message CSMove {
	Left bool
}

message CSStop {
}

message CSFloat {
}

message CSDrop {
}

message SCLogin {
	ErrorCode int32
	UserId    int32
}

message SCChat {
	UserId   int32
	UserName string
	Content  string
}

message SCMatch {
	Result int32
	Room   int32
}

message SCEnterGame {
	Result int32
}

message SCGameStart {
	TimeStamp float64
}

message SCSpawnUnit {
	Iid      int32
	UType    int32
	Pos      Vector2D
	FaceLeft bool
	UFaction int32
	UserId   int32
}

message SCGameState {
	CurTime     float64
	FrameNumber int32
	UnitStates  []UnitState
}

message SCMove {
	Iid  int32
	Left bool
}

message SCStop {
	Iid int32
}

message SCFloat {
	Iid int32
}

message SCDrop {
	Iid int32
}

message SCFire {
	Iid int32
}

message SCBurst {
	Iid int32
}

message SCBlowStart {
	Iid int32
}

message SCBlowCancel {
	Iid int32
}

message SCBlowSuccess {
	Iid int32
}

message UserData {
	UserId   int32
	UserName string
	level    int32
	money    int32
}

message Vector2D{
	X float64
	Y float64
}

message UnitState {
	Iid      int32
	Pos      Vector2D
	FaceLeft bool
	Moving   bool
	Floating bool
	Blowing  bool
}

