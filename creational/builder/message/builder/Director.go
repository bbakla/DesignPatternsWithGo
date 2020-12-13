package builder

type Director struct {
}

func (d *Director) ConstructMessage(builder MessageSender) (*Message, error) {

	return builder.AddVote().Send().Build()

}
