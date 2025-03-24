package datastructs

type ChatMessage struct {
	Room      string
	Username  string
	Message   string
	Timestamp int64
	Time      string
}

type ChatMessagesSortable []ChatMessage

func (m ChatMessagesSortable) Len() int           { return len(m) }
func (m ChatMessagesSortable) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ChatMessagesSortable) Less(i, j int) bool { return m[i].Timestamp < m[j].Timestamp }
