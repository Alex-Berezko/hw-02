package player

type User struct {
	name       string
	pieceColor string
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetColor(color string) {
	u.pieceColor = color
}

func (u User) SayName() string {
	return u.name

}

func (u *User) RenameUser(newName string) {
	if newName != "" {
		u.name = newName
	}
}

func (u User) SayPieceColor() string {
	return u.pieceColor
}
