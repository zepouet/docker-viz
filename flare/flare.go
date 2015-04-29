package flare

type(
	Flare interface {
		DendrogamAndBubbleImages() string;
		BubbleContainers() string;
	}
)

const(
	BeginJson = "{\"name\": \"Docker\", \"children\": ["
	EndJson = "]}"
)