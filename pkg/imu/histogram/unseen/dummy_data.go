package unseen

import (
	"encoding/json"
	"fmt"
)

const addBins = 10

func GenerateDummyData() {
	v2 := NewEmptyHistogramV2(-500, 1500, 0.1)
	fmt.Printf("%v", v2)
}

func GenerateDummyData_() {
	var data AllGainsHistogramResponseV2
	err := json.Unmarshal([]byte(json_), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for i, d := range data.Data {
		data.Data[i].Unseen = d.Seen.Clone()
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(marshal))
}

const json_ = `
{
  "data": [
    {
      "seen": {
        "bins": [
          {
            "max": -35669.31,
            "min": -36614.33,
            "count": 0
          },
          {
            "max": -34724.29,
            "min": -35669.31,
            "count": 5
          },
          {
            "max": -33779.27,
            "min": -34724.29,
            "count": 0
          },
          {
            "max": -32834.254,
            "min": -33779.27,
            "count": 0
          },
          {
            "max": -31889.234,
            "min": -32834.254,
            "count": 0
          },
          {
            "max": -30944.215,
            "min": -31889.234,
            "count": 0
          },
          {
            "max": -29999.195,
            "min": -30944.215,
            "count": 0
          },
          {
            "max": -29054.176,
            "min": -29999.195,
            "count": 0
          },
          {
            "max": -28109.158,
            "min": -29054.176,
            "count": 0
          },
          {
            "max": -27164.139,
            "min": -28109.158,
            "count": 0
          },
          {
            "max": -26219.12,
            "min": -27164.139,
            "count": 0
          },
          {
            "max": -25274.102,
            "min": -26219.12,
            "count": 0
          },
          {
            "max": -24329.082,
            "min": -25274.102,
            "count": 0
          },
          {
            "max": -23384.062,
            "min": -24329.082,
            "count": 0
          },
          {
            "max": -22439.045,
            "min": -23384.062,
            "count": 0
          },
          {
            "max": -21494.025,
            "min": -22439.045,
            "count": 0
          },
          {
            "max": -20549.006,
            "min": -21494.025,
            "count": 0
          },
          {
            "max": -19603.988,
            "min": -20549.006,
            "count": 0
          },
          {
            "max": -18658.969,
            "min": -19603.988,
            "count": 0
          },
          {
            "max": -17713.95,
            "min": -18658.969,
            "count": 0
          },
          {
            "max": -16768.932,
            "min": -17713.95,
            "count": 0
          },
          {
            "max": -15823.912,
            "min": -16768.932,
            "count": 0
          },
          {
            "max": -14878.893,
            "min": -15823.912,
            "count": 0
          },
          {
            "max": -13933.874,
            "min": -14878.893,
            "count": 0
          },
          {
            "max": -12988.855,
            "min": -13933.874,
            "count": 0
          },
          {
            "max": -12043.836,
            "min": -12988.855,
            "count": 0
          },
          {
            "max": -11098.817,
            "min": -12043.836,
            "count": 0
          },
          {
            "max": -10153.799,
            "min": -11098.817,
            "count": 0
          },
          {
            "max": -9208.779,
            "min": -10153.799,
            "count": 0
          },
          {
            "max": -8263.761,
            "min": -9208.779,
            "count": 0
          },
          {
            "max": -7318.7417,
            "min": -8263.761,
            "count": 0
          },
          {
            "max": -6373.7227,
            "min": -7318.7417,
            "count": 0
          },
          {
            "max": -5428.7036,
            "min": -6373.7227,
            "count": 0
          },
          {
            "max": -4483.6846,
            "min": -5428.7036,
            "count": 0
          },
          {
            "max": -3538.666,
            "min": -4483.6846,
            "count": 0
          },
          {
            "max": -2593.647,
            "min": -3538.666,
            "count": 0
          },
          {
            "max": -1648.628,
            "min": -2593.647,
            "count": 0
          },
          {
            "max": -703.60913,
            "min": -1648.628,
            "count": 0
          },
          {
            "max": 241.40979,
            "min": -703.60913,
            "count": 0
          },
          {
            "max": 1186.4287,
            "min": 241.40979,
            "count": 0
          },
          {
            "max": 2131.4478,
            "min": 1186.4287,
            "count": 0
          },
          {
            "max": 3076.4666,
            "min": 2131.4478,
            "count": 0
          },
          {
            "max": 4021.4854,
            "min": 3076.4666,
            "count": 0
          },
          {
            "max": 4966.5044,
            "min": 4021.4854,
            "count": 0
          },
          {
            "max": 5911.5234,
            "min": 4966.5044,
            "count": 0
          },
          {
            "max": 6856.542,
            "min": 5911.5234,
            "count": 0
          },
          {
            "max": 7801.561,
            "min": 6856.542,
            "count": 0
          },
          {
            "max": 8746.58,
            "min": 7801.561,
            "count": 0
          },
          {
            "max": 9691.599,
            "min": 8746.58,
            "count": 0
          },
          {
            "max": 10636.618,
            "min": 9691.599,
            "count": 0
          }
        ],
        "mean": -14598.303,
        "median": -15056.243,
        "std_dev": 4682.7827,
        "bin_scale": 945.0189,
        "max_value": 10636.616,
        "min_value": -36614.33,
        "max_bin_size": 0,
        "min_bin_size": 0,
        "values_count": 0
      },
      "opacity": 1,
      "dependent_name": "extra_plot_var",
      "dependent_type": "",
      "confidence_rate": 1,
      "independent_name": "ti8910",
      "independent_type": "MV"
    }
  ]
}
`
