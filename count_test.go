package tuna

import "testing"

func TestCount(t *testing.T) {
	AggTestCases{
		{
			stream: NewStream(
				Row{"ice-cream": "1", "flavor": "chocolate"},
				Row{"ice-cream": "1", "flavor": "coffee"},
				Row{"ice-cream": "1", "flavor": "vanilla"},
				Row{"ice-cream": "2", "flavor": "mango"},
				Row{"ice-cream": "2", "flavor": "yoghurt"},
			),
			agg:    NewGroupBy("ice-cream", func() Agg { return NewCount() }),
			output: "count,ice-cream\n3,1\n2,2\n",
			size:   2,
		},
		{
			stream: NewStream(
				Row{"ice-cream": "1", "flavor": "chocolate"},
				Row{"ice-cream": "1", "flavor": "coffee"},
				Row{"ice-cream": "2", "flavor": "mango"},
				Row{"ice-cream": "1", "flavor": "vanilla"},
				Row{"ice-cream": "2", "flavor": "yoghurt"},
			),
			agg:    NewGroupBy("ice-cream", func() Agg { return NewCount() }),
			output: "count,ice-cream\n3,1\n2,2\n",
			size:   2,
		},
	}.Run(t)
}
