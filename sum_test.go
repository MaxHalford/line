package tuna

import "testing"

func TestSum(t *testing.T) {
	ExtractorTestCases{
		{
			stream: ZipStreams(
				NewStream(
					Row{"flux": "1.0"},
					Row{"flux": "4.0"},
					Row{"flux": "-2.0"},
				),
				NewStream(
					Row{"flux": "1.0"},
					Row{"flux": "2.0"},
					Row{"flux": "3.0"},
				),
			),
			extractor: NewSum("flux"),
			output:    "flux_sum\n9\n",
		},
	}.Run(t)
}
