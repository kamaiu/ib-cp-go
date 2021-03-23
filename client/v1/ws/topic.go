package ws

type TopicType string

const (
	TopicSystem        TopicType = "system"
	TopicStatus        TopicType = "sts"
	TopicNotifications TopicType = "ntf"
	TopicBulletins     TopicType = "blt"
	TopicProfitLoss    TopicType = "spl"
	TopicLiveOrders    TopicType = "sor"
)
