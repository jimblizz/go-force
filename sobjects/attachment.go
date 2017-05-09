package sobjects

type Attachment struct {
	BaseSObject

	BodyLength 	float64 `force:"BodyLength"`
	Body       	string `force:"Body"`
}

func (a Attachment) ApiName() string {
	return "Attachment"
}
