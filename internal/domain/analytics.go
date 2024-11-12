package domain

type Analytics struct {
	UserID    int64
	ProductID int64
	Activity  string
}

type AnalyticsEvent struct {
	UserID    int64
	ProductID int64
	Activity  string
}

func (e AnalyticsEvent) EventType() string {
	return "AnalyticsLogged"
}
