package category

import "lazymod/src/structs/stat"

type Category struct {
	categoryName string
	stats        []stat.Stat
}

func (category Category) GetName() string {
	return category.categoryName
}

func (category Category) GetStats() []stat.Stat {
	return category.stats
}

func (category Category) SetName(name string) {
	category.categoryName = name
}

func (category Category) AddStat(stat stat.Stat) {
	category.stats = append(category.stats, stat)
}

func (category Category) RemoveStat(name string) {
	newStats := []stat.Stat{}
	for _, stat := range category.stats {
		if stat.GetName() != name {
			newStats = append(newStats, stat)
		}
	}
	category.stats = newStats
}
