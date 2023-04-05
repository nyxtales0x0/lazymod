package stat

type Stat struct {
	statName string
	statType string
	minValue int
	maxValue int
}

func (stat Stat) GetName() string {
	return stat.statName
}

func (stat Stat) GetType() string {
	return stat.statType
}

func (stat Stat) GetMinValue() int {
	return stat.minValue
}

func (stat Stat) GetMaxValue() int {
	return stat.maxValue
}

func (stat Stat) SetName(name string) {
	stat.statName = name
}

func (stat Stat) SetMinValue(value int) {
	stat.minValue = value
}

func (stat Stat) SetMaxValue(value int) {
	stat.maxValue = value
}

// write logic to identify stat type
// create stat accordingly
func New() Stat {
	return Stat{}
}
